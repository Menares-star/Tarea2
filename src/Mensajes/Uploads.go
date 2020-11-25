package Uploads

import (
	//"log"
  //"golang.org/x/net/context"
	"io"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

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
		part := str.GetPart()
		name := str.GetName()
		content := str.GetContent()

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
		vm1 = ":8001"
		vm2 = ":7000"
		vm3 = ":6000"
	}

	prop := Propose.NewProponerServiceClient(conn)
	propFinal, err := prop.Propose(context.Background(), &propose)

	fmt.Println("pF: " + propFinal.vm2.puerto)

	return stream.SendAndClose(&UploadStatus{
		Message:   "Su estado es CHUNKS ENVIADOS",
	})

}
