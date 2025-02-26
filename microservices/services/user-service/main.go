package main

import (
    "log"
    "net"
    "google.golang.org/grpc"
    pb "user-service/proto"
    "user-service/handlers"
    "user-service/db"
)

func main() {
    db.InitDB()

    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen on port 50051: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterUserServiceServer(grpcServer, &handlers.UserServiceServer{})

    log.Println("User Service running on port 50051")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to start gRPC server: %v", err)
    }
}
