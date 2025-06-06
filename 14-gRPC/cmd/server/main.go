package main

import (
	"database/sql"
	"net"

	"log"

	"github.com/felipeazsantos/pos-goexpert/14-gRPC/internal/database"
	"github.com/felipeazsantos/pos-goexpert/14-gRPC/internal/pb"
	"github.com/felipeazsantos/pos-goexpert/14-gRPC/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(categoryDB)

	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, categoryService)
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("gRPC server is listening on %s", lis.Addr().String())

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
