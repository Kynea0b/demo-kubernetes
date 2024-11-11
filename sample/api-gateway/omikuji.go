package main

import (
	omikujipb "api-gateway/pkg/grpc/proto"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	omikujiClient omikujipb.OmikujiServiceClient
	omikujiConn   *grpc.ClientConn
)

func init() {
	if omikujiClient != nil {
		return
	}

	host := os.Getenv("OMIKUJI_SERVICE_HOST")
	if host == "" {
		log.Fatal("OMIKUJI_SERVICE_HOST is not set")
	}

	var err error
	omikujiConn, err = grpc.Dial(
		fmt.Sprintf("%s:%s", host, "8080"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	omikujiClient = omikujipb.NewOmikujiServiceClient(omikujiConn)
}

type DrawRequest struct {
	UserId int64 `json:"user_id"`
}

// /draw
func Draw(c echo.Context) error {
	dr := new(DrawRequest)
	if err := c.Bind(dr); err != nil {
		return err
	}

	req := &omikujipb.DrawRequest{
		UserId: dr.UserId,
	}
	res, err := omikujiClient.Draw(context.Background(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func GetHistories(c echo.Context) error {
	p := c.Param("user_id")
	userId, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return err
	}

	req := &omikujipb.GetHistoriesRequest{
		UserId: userId,
	}
	res, err := omikujiClient.GetHistories(context.Background(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func CloseOmikujiConnection() error {
	if omikujiConn != nil {
		return omikujiConn.Close()
	}
	return nil
}
