package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea2/src/Distributed/Mensajes/Uploads"
	"github.com/Menares-star/Tarea2/src/Distributed/Mensajes/Propuesta"

)

func uploadsServer(lis net.Listener) {
	s := Uploads.Server1{}
	grpcServer := grpc.NewServer()
	Uploads.RegisterGuploadServiceServer(grpcServer, &s)
	Uploads.RegisterRepartirServiceServer(grpcServer, &s)
	fmt.Println("NODO3 Escuchando a nuevos clientes")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func proposesServer(lis net.Listener) {
	s := Propose.Server{}
	grpcServer := grpc.NewServer()
	Propose.RegisterProponerServiceServer(grpcServer, &s)
	fmt.Println("NODO3 Escuchando a las propuestas")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}



func main() {

	lis1, err1 := net.Listen("tcp", ":6001")
	if err1 != nil {
		log.Fatalf("failed to listen: %v", err1)
	}

	lis2, err2 := net.Listen("tcp", ":6002")
	if err2 != nil {
		log.Fatalf("failed to listen: %v", err2)
	}

	go func() {
		uploadsServer(lis1)
	}()
	proposesServer(lis2)
}
