package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea2/srcCent/Mensajes/Propuesta"

)

func main() {

	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Propose.Server{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	Propose.RegisterProponerServiceServer(grpcServer, &s)
	Propose.RegisterListarServiceServer(grpcServer, &s)
	Propose.RegisterListarChunksServiceServer(grpcServer,&s)
	//ordenes.RegisterSeguimientoServiceServer(grpcServer, &s)
	//FIN REGISTRO DE SERVICIOS
	fmt.Println("nameNode FUNCIONANDO")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
