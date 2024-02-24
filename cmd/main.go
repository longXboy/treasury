package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "github.com/longxboy/treasure/api/server"
)

var (
	Name = "treasury"
)

type server struct {
	pb.UnimplementedTreasuryServer
}

func (s *server) CreateWithdraw(ctx context.Context, in *pb.CreateWithdrawRequest) (*pb.CreateWithdrawReply, error) {
	return &pb.CreateWithdrawReply{}, nil
}

func (s *server) GetWithdrawStatus(ctx context.Context, in *pb.GetWithdrawStatusRequest) (*pb.GetWithdrawStatusReply, error) {
	return &pb.GetWithdrawStatusReply{}, nil
}

func (s *server) ConfirmWithdraw(ctx context.Context, in *pb.ConfirmWithdrawRequest) (*pb.ConfirmWithdrawReply, error) {
	return &pb.ConfirmWithdrawReply{}, nil
}

func main() {
	s := &server{}
	httpSrv := http.NewServer(
		http.Address(":8000"),
		http.Middleware(
			recovery.Recovery(),
		),
	)

	pb.RegisterTreasuryHTTPServer(httpSrv, s)

	app := kratos.New(
		kratos.Name(Name),
		kratos.Server(
			httpSrv,
		),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
