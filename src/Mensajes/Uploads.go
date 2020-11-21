package Uploads

import (
	"log"
  //"golang.org/x/net/context"
	"io"

)

type Server1 struct{

	 x  map[int][]Chunk

}

func (s *Server1) Upload(stream GuploadService_UploadServer) error {
	x := make(map[int][]Chunk)
	for i:=0;;i++ {
		str, err := stream.Recv()
		log.Printf("Ha llegado el chunk",str.GetPart())
		x[i] = append(x[i],*str)
		if err == io.EOF {
			return SendAndClose(&UploadStatus{
				Message:   "Su estado es OK",
			})
		}
		if err != nil {
			return SendAndClose(&UploadStatus{
				Message:   "Su estado es ERROR",
			})
		}
	}
}
