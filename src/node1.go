package node1

import (
	"log"
	"net"
	//"fmt"
	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea1/src/Mensajes"
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

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	//SEGUIMIENTOS
	//FIN SEGUIMIENTOS
	//COMUNICACION CON CAMIONES

}
