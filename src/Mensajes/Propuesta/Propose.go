package Propose

import (
	//	"log"
	"golang.org/x/net/context"
	//"time"
	//"google.golang.org/grpc"
	//"io"
	//"fmt"
	//"io/ioutil"
	//"os"
	//"strconv"
)

type Server struct{
}

func (s *Server) Proponer(ctx context.Context, prop *Propuesta) (*PropuestaFinal, error) {
	
	mv1 := InfoMaquina {
		Puerto: ":8001",
		IsAvailable: int64(1),
	}
	mv2 := InfoMaquina {
		Puerto: ":7000",
		IsAvailable: int64(1),
	}
	mv3 := InfoMaquina {
		Puerto: ":6000",
		IsAvailable: int64(1),
	}
	
	/*
	var conn *grpc.ClientConn
	var t time.Duration = 5000000000

	conn, err := grpc.Dial(":8001",grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
	if err != nil{
		fmt.Println("Servidor :8001 no disponible ",err)
		mv1.IsAvailable = int64(0)
	}
	defer conn.Close()

	
	conn, err = grpc.Dial(":7000",grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
	if err != nil{
		fmt.Println("Servidor :7000 no disponible",err)
		mv2.IsAvailable = int64(0)
	}
	defer conn.Close()

	conn, err = grpc.Dial(":6000",grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
	if err != nil{
		fmt.Println("Servidor :6000 no disponible ",err)
		mv1.IsAvailable = int64(0)
	}
	defer conn.Close()

	*/
	propose := PropuestaFinal {
		Vm1: &mv1,
		Vm2: &mv2,
		Vm3: &mv3,
	}

	return &propose, nil
}
