package main

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	omikujipb "omikuji-service/pkg/grpc/proto"
	"time"

	timestamp "google.golang.org/protobuf/types/known/timestamppb"
)

type Item struct {
	Id     int64
	Name   string
	Rarity omikujipb.Rarity
	Weight int
}

type History struct {
	ItemId    int64
	ItemName  string
	Rarity    string
	CreatedAt time.Time
}

type omikujiServiceServer struct {
	omikujipb.UnimplementedOmikujiServiceServer
	db    *sql.DB
	items []Item
}

func NewOmikujiServiceServer(db *sql.DB, items []Item) omikujipb.OmikujiServiceServer {
	return &omikujiServiceServer{
		db:    db,
		items: items,
	}
}

func (g *omikujiServiceServer) Draw(ctx context.Context, req *omikujipb.DrawRequest) (*omikujipb.DrawResponse, error) {
	// itemsからitemを重み付抽選する
	weights := make([]int, len(g.items))
	for i, item := range g.items {
		weights[i] = item.Weight
	}

	seed := time.Now().UnixNano()
	i := linearSearchLottery(weights, seed)
	item := g.items[i]

	// DBに保存する
	if err := save(ctx, g.db, req.UserId, item); err != nil {
		return nil, err
	}

	// item所持情報も更新する
	res, err := GetItem(ctx, req.UserId, item.Id, item.Name, item.Rarity)
	if err != nil {
		return nil, err
	}

	fmt.Printf("get_item_response: %+v\n", res)

	return &omikujipb.DrawResponse{
		ItemId:   item.Id,
		ItemName: item.Name,
		Rarity:   item.Rarity,
	}, nil
}

func save(ctx context.Context, db *sql.DB, userId int64, item Item) error {
	_, err := db.ExecContext(
		ctx,
		"INSERT INTO histories (user_id, item_id, item_name, rarity, created_at) VALUES (?, ?, ?, ?, ?)",
		userId,
		item.Id,
		item.Name,
		item.Rarity.String(),
		time.Now().Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return err
	}
	return nil
}

/*
線形探索で重み付抽選する
@return 当選した要素のインデックス
*/
func linearSearchLottery(weights []int, seed int64) int {
	//  重みの総和を取得する
	var total int
	for _, weight := range weights {
		total += weight
	}

	// 乱数取得
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	r := rand.New(rand.NewSource(seed))
	rnd := r.Intn(total)

	var currentWeight int
	for i, w := range weights {
		// 現在要素までの重みの総和
		currentWeight += w

		if rnd < currentWeight {
			return i
		}
	}

	return len(weights) - 1
}

func (g *omikujiServiceServer) GetHistories(ctx context.Context, req *omikujipb.GetHistoriesRequest) (*omikujipb.GetHistoriesResponse, error) {
	rows, err := g.db.QueryContext(
		ctx,
		"SELECT item_id, item_name, rarity, created_at FROM histories WHERE user_id = ? ORDER BY created_at DESC",
		req.UserId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []*omikujipb.History
	for rows.Next() {
		var history History
		if err := rows.Scan(
			&history.ItemId,
			&history.ItemName,
			&history.Rarity,
			&history.CreatedAt,
		); err != nil {
			return nil, err
		}
		histories = append(histories, toMessage(history))
	}

	return &omikujipb.GetHistoriesResponse{
		Histories: histories,
	}, nil
}

func toMessage(h History) *omikujipb.History {
	return &omikujipb.History{
		ItemId:    h.ItemId,
		ItemName:  h.ItemName,
		Rarity:    parseRarity(h.Rarity),
		CreatedAt: timestamp.New(h.CreatedAt),
	}
}

func parseRarity(s string) omikujipb.Rarity {
	value, ok := omikujipb.Rarity_value[s]
	if !ok {
		fmt.Printf("invalid Rarity: %s\n", s)
		return omikujipb.Rarity_RARITY_UNKNOWN
	}
	return omikujipb.Rarity(value)
}
