package main

import (
	"context"
	"fmt"
	"microservices-crud/user-service/config"
	"microservices-crud/user-service/db/user"
	"microservices-crud/user-service/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Host,
		cfg.Database.Port,
	)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close(ctx)

	q := user.New(conn)
	uh := handlers.NewUserHandler(q)

	r := chi.NewRouter()
	r.Mount("/users", uh.Routes())

	http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), r)
}
