package Uploads

import (
	//"log"
  	"golang.org/x/net/context"
	"io"
	"fmt"
	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea2/src/Mensajes/Propuesta"
	//"io/ioutil"
	//"os"
	//"strconv"
)

var memory1 int64 = 250000
var memory2 int64 = 125000
var memory3 int64 = 75000

var save [] Chunk

type Server1 struct{
}

func (s *Server1) Upload(stream GuploadService_UploadServer) error {
	//x := make(map[int][]Chunk)
	for i:=0;;i++ {
		str, err := stream.Recv()
		if err == io.EOF {
			break
		}
		//DECOMPOSING CHUNK
		//part := str.GetPart()
		//name := str.GetName()
		//content := str.GetContent()

		save = append(save,*str)
		//node1.save1 = append(node1.save1,*str)
		fmt.Println(save[i].Part)
		//fmt.Println("kaka")

		//SAVE IN DISK
		/*
		partName := name + "_part_" + strconv.Itoa(i) + ".pdf"
		fmt.Println("chunk " + strconv.Itoa(i) + " name: " + partName, part)
		ioutil.WriteFile("chunks/" + partName, content, os.ModeAppend)
		*/

		if err != nil {
			return stream.SendAndClose(&UploadStatus{
				Message:   "Su estado es CHUNKS NO ENVIADOS",
			})
		}
	}
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000",grpc.WithInsecure(),grpc.WithBlock())
	if err != nil{
		fmt.Println("Servidor :9000 no disponible: ", err)
	}
	defer conn.Close()

  propose := Propose.Propuesta {
		Vm1: ":8001",
		Vm2: ":7000",
		Vm3: ":6000",
	}
  var mem int64 = 0
  if save[0].Puerto==":8001"{
    mem=memory1
  }
  if save[0].Puerto==":7000"{
    mem=memory2
  }
  if save[0].Puerto==":6000"{
    mem=memory3
  }


  Info := Propose.InfoMaquina {
		Puerto: save[0].Puerto,
		Memory: mem,
		Propuesta: &propose,
	}

	prop := Propose.NewProponerServiceClient(conn)
	propFinal, err := prop.Proponer(context.Background(), &Info)

	fmt.Println("R: vm1(" + propFinal.Vm1 + ")")
	fmt.Println("R: vm2(" + propFinal.Vm2 + ")")
	fmt.Println("R: vm3(" + propFinal.Vm3 + ")")

	return stream.SendAndClose(&UploadStatus{
		Message:   "Su estado es CHUNKS ENVIADOS",
	})

}
