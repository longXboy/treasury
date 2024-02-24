package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "github.com/longxboy/treasure/api/server"
	"github.com/longxboy/treasure/ent"
)

var (
	Name     = "treasury"
	dbClient *ent.Client
)

func main() {
	connectDb()
	defer dbClient.Close()
	go runJob()
	runServer()
}

func runServer() {
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

func connectDb() {
	var err error
	dbClient, err = ent.Open("mysql", "root:root@tcp(127.0.0.1:3306)/treasury?timeout=500ms&readTimeout=500ms&writeTimeout=600ms&parseTime=true&loc=Local&charset=utf8mb4,utf8")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// Run the auto migration tool.
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
