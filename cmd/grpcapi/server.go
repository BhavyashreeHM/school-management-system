package main

import (
	"employee-management-system/internal/api/handlers"
	"employee-management-system/internal/reposioriy/mongodb"
	pb "employee-management-system/proto/gen"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	//load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	port := os.Getenv("APP_PORT")
	fmt.Println("Port:", port)

	// Connect to MongoDB
	_, err = mongodb.CreateMongoClient()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	//create a new gRPC server instance
	grpcServer := grpc.NewServer()

	//register the gRPC server with the handlers
	pb.RegisterExecsServiceServer(grpcServer, &handlers.Server{})
	pb.RegisterTeachersServiceServer(grpcServer, &handlers.Server{})
	pb.RegisterStudentsServiceServer(grpcServer, &handlers.Server{})

	log.Println("gRPC API SERVER IS UP AND RUNNING ON PORT", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
