package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea2/src/Mensajes"
)


func main() {

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Uploads.Server1{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	Uploads.RegisterGuploadServiceServer(grpcServer, &s)
	//ordenes.RegisterSeguimientoServiceServer(grpcServer, &s)
	//FIN REGISTRO DE SERVICIOS
	fmt.Println("NODO1 FUNCIONANDO")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	fmt.Println("NODO1 FUNCIONANDO")
	//SEGUIMIENTOS
	//FIN SEGUIMIENTOS
	//COMUNICACION CON CAMIONES

}
