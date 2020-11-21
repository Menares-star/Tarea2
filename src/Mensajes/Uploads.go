package Uploads

import (
	"log"
  "golang.org/x/net/context"
	"io/ioutil"
	"time"

)

type Server1 struct{

	x:= make(map[int][]Chunk)

}

func (s *Server1) Upload(stream GuploadService_UploadServer) error {

	startTime := time.Now()
	for i:=0;;i++ {
		chunk, err := stream.Recv()
		log.Printf("Ha llegado el chunk",chunk.Chunk.GetPart())
		x[i] = append(x[i],chunk)
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(UploadStatus{
				Message:   "Su estado es",
				Code: enum.OK,
			})
		}
		if err != nil {
			return err
		}
	}
}
