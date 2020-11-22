package Uploads

import (
	"log"
  //"golang.org/x/net/context"
	"io"
	"fmt"
)

var save [] Chunk

type Server1 struct{
}

func (s *Server1) Upload(stream GuploadService_UploadServer) error {
	//x := make(map[int][]Chunk)
	for i:=0;;i++ {
		str, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&UploadStatus{
				Message:   "Su estado es CHUNKS ENVIADOS",
			})
		}
		log.Printf("Ha llegado el chunk %d",str.GetPart())
		save = append(save,*str)
		//node1.save1 = append(node1.save1,*str)
		fmt.Println(save[i].Part)
		//fmt.Println("kaka")
		if err != nil {
			return stream.SendAndClose(&UploadStatus{
				Message:   "Su estado es CHUNKS NO ENVIADOS",
			})
		}
	}
}
