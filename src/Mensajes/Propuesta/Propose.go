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
	"os"
	//"bufio"
	"strconv"
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

	var cont int = 1
	var available1 string = servers[0]
	var available2 string = servers[1]
	var available3 string = servers[2]

	/* CONEXION*/
	if available1!=prop.Puerto{
		var conn1 *grpc.ClientConn
		conn1, err1 := grpc.Dial(servers[0],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
		if err1 != nil{
			fmt.Println("Servidor "+servers[0]+" no disponible: ",err1)
			available1="No disponible"
				}else{
					fmt.Println("Servidor "+servers[0]+" disponible: ")
					cont = cont +1
					defer conn1.Close()
				}
	}


	if available2!=prop.Puerto{
		var conn2 *grpc.ClientConn
		conn2, err2 := grpc.Dial(servers[1],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
		if err2 != nil{
			fmt.Println("Servidor "+servers[1]+" no disponible: ",err2)
			available2="No disponible"
				}else{
					fmt.Println("Servidor "+servers[1]+" disponible: ")
					cont = cont+1
					defer conn2.Close()
				}
	}

	if available3!=prop.Puerto{
		fmt.Println("ver3")
		var conn3 *grpc.ClientConn
		conn3, err3 := grpc.Dial(servers[2],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
		if err3 != nil{
			fmt.Println("Servidor "+servers[2]+" no disponible: ",err3)
			available3="No disponible"
				}else{
					fmt.Println("Servidor "+servers[2]+" disponible: ")
					cont = cont + 1
					defer conn3.Close()
				}
	}

	var c1 int =0
	var c2 int =0
	var c3 int =0

	var nc int = int(prop.Nchunks)

	if nc<cont{
		if servers[0]==prop.Puerto{
			c1=nc
		}
		if servers[1]==prop.Puerto{
			c2=nc
		}
		if servers[2]==prop.Puerto{
			c3=nc
		}
	}

	fmt.Println(float64(float64(nc)/float64(cont)))

	if nc>=cont{
		if (nc%cont)==0{
			if available1==servers[0]{
				c1=(nc/cont)
			}
			if available2==servers[1]{
				c2=(nc/cont)
			}
			if available3==servers[2]{
				c3=(nc/cont)
			}

		}else{
			if available1==servers[0]{
				c1=(nc/cont)
				if servers[0]==prop.Puerto{
					c1=c1 + (nc%cont)
				}
			}
			if available2==servers[1]{
				c2=(nc/cont)
				if servers[1]==prop.Puerto{
					c2=c2 + (nc%cont)
				}
			}
			if available3==servers[2]{
				c3=(nc/cont)
				if servers[2]==prop.Puerto{
					c3=c3 + (nc%cont)
				}
			}
		}
	}

	// open input file
  fi, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  if err != nil {
  panic(err)
	}
	fi.WriteString(prop.Name+" "+ strconv.Itoa(nc)+"\n")
	var i int = 0
  for ;i<c1;i++{
	 fi.WriteString(prop.Name+"_parte_"+ strconv.Itoa(i)+" "+servers[0]+"\n")
  }
	for ;i<(c1+c2);i++{
	 fi.WriteString(prop.Name+"_parte_"+ strconv.Itoa(i)+" "+servers[1]+"\n")
  }
	for ;i<(c1+c2+c3);i++{
	 fi.WriteString(prop.Name+"_parte_"+ strconv.Itoa(i)+" "+servers[2]+"\n")
  }
  fi.Close()

	propose := Propuesta{
	Vm1: available1,
	Vm2: available2,
	Vm3: available3,
	Cant1: int64(c1),
	Cant2: int64(c2),
	Cant3: int64(c3),
	}
	fmt.Println(propose)
	return &propose, nil
}
