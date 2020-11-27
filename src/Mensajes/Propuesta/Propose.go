package Propose

import (
	//	"log"
	"golang.org/x/net/context"
	"time"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/connectivity"
	//"io"
	"fmt"
	//"io/ioutil"
	//"os"
	//"strconv"
)

type Server struct{
}

var t time.Duration = 2500000000

func (s *Server) Proponer(ctx context.Context, prop *InfoMaquina) (*Propuesta, error) {

	fmt.Println(prop)
	fmt.Println("Verificando Propuesta")


	servers:=make([]string,3)
	servers[0]=":8001"
	servers[1]=":7000"
	servers[2]=":6000"

	var available1 string = servers[0]
	var available2 string = servers[1]
	var available3 string = servers[2]

	var t time.Duration = 2500000000
	/* CONEXION*/
	//conn, err := grpc.Dial(":8000", grpc.WithInsecure(), grpc.WithBlock())
	var conn1 *grpc.ClientConn
	conn1, err1 := grpc.Dial(servers[0],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
	if err1 != nil{
		fmt.Println("Servidor "+servers[0]+" no disponible: ",err1)
		available1="No disponible"
			}else{
				fmt.Println("Servidor "+servers[0]+" disponible: ")
				defer conn1.Close()
			}

	var conn2 *grpc.ClientConn
	conn2, err2 := grpc.Dial(servers[1],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
	if err2 != nil{
		fmt.Println("Servidor "+servers[1]+" no disponible: ",err2)
		available2="No disponible"
			}else{
				fmt.Println("Servidor "+servers[1]+" disponible: ")
				defer conn2.Close()
			}

	var conn3 *grpc.ClientConn
	conn3, err3 := grpc.Dial(servers[2],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))

	if err3 != nil{
		fmt.Println("Servidor "+servers[2]+" no disponible: ",err3)
		available3="No disponible"
			}else{
				fmt.Println("Servidor "+servers[2]+" disponible: ")
				defer conn3.Close()
			}


		propose := Propuesta{
		Vm1: available1,
		Vm2: available2,
		Vm3: available3,
	}
	fmt.Println("AQUI")
	return &propose, nil
}
