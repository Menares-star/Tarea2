package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea2/src/Mensajes/Propuesta"

)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Propose.Server{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	Propose.RegisterProponerServiceServer(grpcServer, &s)
	//ordenes.RegisterSeguimientoServiceServer(grpcServer, &s)
	//FIN REGISTRO DE SERVICIOS
	fmt.Println("NODO1 FUNCIONANDO")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
