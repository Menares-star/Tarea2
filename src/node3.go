package node3

import (
	"log"
	"net"
	//"fmt"

	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea2/src/Mensajes"
)


func main() {

	lis, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//s := ordenes.Server{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	//ordenes.RegisterOrdenServiceServer(grpcServer, &s)
	//ordenes.RegisterSeguimientoServiceServer(grpcServer, &s)
	//FIN REGISTRO DE SERVICIOS


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	//SEGUIMIENTOS
	//FIN SEGUIMIENTOS
	//COMUNICACION CON CAMIONES

}
