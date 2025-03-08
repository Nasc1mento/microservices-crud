package main

import (
	"context"
	"fmt"
	"log"
	"microservices-crud/user-service/config"
	"microservices-crud/user-service/internal/db/repo"
	user "microservices-crud/user-service/internal/pb"
	"microservices-crud/user-service/internal/service"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	grpcAddr := fmt.Sprintf("%s:%d", cfg.Server.GRpcHost, cfg.Server.GRpcPort)

	gwAddr := fmt.Sprintf("%s:%d", cfg.Server.ResHost, cfg.Server.RestPort)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Host,
		cfg.Database.Port,
	)

	dbConn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dbConn.Close(ctx)

	q := repo.New(dbConn)

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	us := service.NewServer(q)
	user.RegisterUserServiceServer(s, us)

	go func() {
		log.Fatal(s.Serve(lis))
	}()

	conn, err := grpc.NewClient(
		grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	gmux := runtime.NewServeMux()
	err = user.RegisterUserServiceHandler(ctx, gmux, conn)
	if err != nil {
		log.Fatal(err)
	}

	gwServer := &http.Server{
		Addr:    gwAddr,
		Handler: gmux,
	}

	log.Printf("Serving gRPC-Gateway on http://%s", gwAddr)
	log.Fatal(gwServer.ListenAndServe())
}
