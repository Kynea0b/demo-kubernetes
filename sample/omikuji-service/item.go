package main

import (
	"context"
	"fmt"
	"log"
	omikujipb "omikuji-service/pkg/grpc/proto"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client omikujipb.ItemServiceClient
	conn   *grpc.ClientConn
)

func init() {
	if client != nil {
		return
	}

	host := os.Getenv("ITEM_SERVICE_HOST")
	if host == "" {
		log.Fatal("ITEM_SERVICE_HOST is not set")
	}

	var err error
	conn, err = grpc.Dial(
		fmt.Sprintf("%s:%s", host, "8080"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	client = omikujipb.NewItemServiceClient(conn)
}

func GetItem(ctx context.Context, userId, itemId int64, itemName string, rarity omikujipb.Rarity) (*omikujipb.GetItemResponse, error) {
	req := &omikujipb.GetItemRequest{
		UserId:   userId,
		ItemId:   itemId,
		ItemName: itemName,
		Rarity:   rarity,
	}
	return client.GetItem(ctx, req)
}

func Close() error {
	if conn != nil {
		return conn.Close()
	}
	return nil
}
