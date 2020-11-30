package Propose

import (
	//	"log"
	"golang.org/x/net/context"
	//"time"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/connectivity"
	//"io"
	//"fmt"
	//"io/ioutil"
	//"os"
	//"strconv"
)

type Server struct{
}

func (s *Server) Proponer(ctx context.Context, prop *Propuesta) (*Respuesta, error) {

	resp := Respuesta {
		Status: "ACECTO",
	}

	return &resp, nil
}
