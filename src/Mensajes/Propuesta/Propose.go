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

func (s *Server) Proponer(ctx context.Context, prop *Propuesta) (*PropuestaFinal, error) {

	fmt.Println("Verificando Propuesta")


	servers:=make([]string,3)
	servers[0]=":8001"
	servers[1]=":7000"
	servers[2]=":6000"

	var available1 = 1
	var available2 = 1
	var available3 = 1

	var t time.Duration = 2500000000
	/* CONEXION*/
	//conn, err := grpc.Dial(":8000", grpc.WithInsecure(), grpc.WithBlock())
	var conn1 *grpc.ClientConn
	conn1, err1 := grpc.Dial(servers[0],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
	if err1 != nil{
		fmt.Println("Servidor "+servers[0]+" no disponible: ",err1)
		available1=0
			}else{
				fmt.Println("Servidor "+servers[0]+" disponible: ")
			}

	var conn2 *grpc.ClientConn
	conn2, err2 := grpc.Dial(servers[1],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
	if err2 != nil{
		fmt.Println("Servidor "+servers[1]+" no disponible: ",err2)
		available2=0
			}else{
				fmt.Println("Servidor "+servers[1]+" disponible: ")
			}

	var conn3 *grpc.ClientConn
	conn3, err3 := grpc.Dial(servers[2],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
	if err3 != nil{
		fmt.Println("Servidor "+servers[2]+" no disponible: ",err3)
		available3=0
			}else{
				fmt.Println("Servidor "+servers[2]+" disponible: ")
			}
	fmt.Println(conn1,conn2,conn3)
	/*defer conn1.Close()
	defer conn2.Close()
	defer conn3.Close()*/
	mv1 := InfoMaquina {
		Puerto: ":8001",
		IsAvailable: int64(available1),
	}
	mv2 := InfoMaquina {
		Puerto: ":7000",
		IsAvailable: int64(available2),
	}
	mv3 := InfoMaquina {
		Puerto: ":6000",
		IsAvailable: int64(available3),
	}

	propose := PropuestaFinal {
		Vm1: &mv1,
		Vm2: &mv2,
		Vm3: &mv3,
	}
	fmt.Println("AQUI")
	return &propose, nil
}
