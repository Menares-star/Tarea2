package Propose

import (
	"golang.org/x/net/context"
	"time"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/connectivity"
	//"io"
	"fmt"
	//"io/ioutil"
	"log"
  "strings"
	"bufio"
	"os"
	"strconv"
)

type Server struct{
}

var t time.Duration = 2500000000

func (s *Server) Proponer(ctx context.Context, prop *InfoMaquina) (*Propuesta, error) {

	fmt.Println("Verificando Propuesta")
	start := time.Now()

	servers:=make([]string,3)
	servers[0]="dist74:8001"
	servers[1]="dist75:7000"
	servers[2]="dist76:6000"

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
					fmt.Println("Servidor "+servers[0]+" disponible")
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
					fmt.Println("Servidor "+servers[1]+" disponible")
					cont = cont+1
					defer conn2.Close()
				}
	}

	if available3!=prop.Puerto{
		var conn3 *grpc.ClientConn
		conn3, err3 := grpc.Dial(servers[2],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
		if err3 != nil{
			fmt.Println("Servidor "+servers[2]+" no disponible: ",err3)
			available3="No disponible"
				}else{
					fmt.Println("Servidor "+servers[2]+" disponible")
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

  elapsed := time.Since(start)
  fmt.Println("Archivo LOG creado satisfactoriamente")
  fmt.Println("EXECUTION TIME: ", elapsed)

	propose := Propuesta{
	Vm1: available1,
	Vm2: available2,
	Vm3: available3,
	Cant1: int64(c1),
	Cant2: int64(c2),
	Cant3: int64(c3),
	}
	fmt.Println("Enviando distribucion a " + prop.Puerto)
	return &propose, nil
}


func (s *Server) Listar(recept *ReceptStatus, stream ListarService_ListarServer) error {
	var save[] string
  file, err := os.Open( "log.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    librop:= strings.Split(scanner.Text(),"_parte")
    libro:= strings.Split(librop[0]," ")
    //fmt.Println(libro[0])
    var flag bool = false
    for i:=0;flag==false && i<len(save);i++ {
      if save[i]==libro[0]{
        flag=true
        break
      }
    }
    if flag==false{
      save = append(save,libro[0])
      fmt.Println(libro[0])
			book:= Libro{
				Name: libro[0],
			}
			if err := stream.Send(&book); err != nil {
				return err
			}
    }
  }
 return nil
}

func (s *Server) ListarChunks(book *Libro, stream ListarChunksService_ListarChunksServer) error {
	fmt.Println(book.Name)
	//var save[] string
  file, err := os.Open( "log.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
	var i int=-1
	var Npartes int =-1
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    librop:= strings.Split(scanner.Text()," ")
		fmt.Println(librop[0])
		fmt.Println(i,Npartes)
		if i<Npartes{
			info:=InfoChunk{
				Puerto: librop[1],
				Part: librop[0],
			}
			if err := stream.Send(&info); err != nil {
				return err
			}
			i=i+1
		}
		if i==Npartes && i>0{
			break
		}
    if librop[0]==book.Name{
			i = 0
			Npartes,err = strconv.Atoi(librop[1])
			if err!=nil{
				fmt.Println(Npartes,err)
				return err
			}
		}
  }

 return nil
}
