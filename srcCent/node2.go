package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea2/srcCent/Mensajes/Upload"

)

func main() {

	lis, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Uploads.Server1{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	Uploads.RegisterGuploadServiceServer(grpcServer, &s)
	Uploads.RegisterRepartirServiceServer(grpcServer, &s)
	Uploads.RegisterDownloadServiceServer(grpcServer,&s)
	//ordenes.RegisterSeguimientoServiceServer(grpcServer, &s)
	//FIN REGISTRO DE SERVICIOS
	fmt.Println("NODO2 FUNCIONANDO")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
