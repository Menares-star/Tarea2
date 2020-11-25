package Propose

import (
	"log"
  	"golang.org/x/net/context"
	//"io"
	//"fmt"
	//"io/ioutil"
	//"os"
	//"strconv"
)

type Server struct{
}

func (s *Server) Proponer(ctx context.Context, prop *Propuesta) (*PropuestaFinal, error) {
	
	mv1 := Propose.InfoMaquina {
		puerto:":8001",
		isAvailable: int32(1),
	}
	mv2 := Propose.InfoMaquina {
		puerto: ":7000",
		isAvailable: int32(0),
	}
	mv3 := Propose.InfoMaquina {
		puerto: ":6000",
		isAvailable: int32(1),
	}
	propose := Propose.PropuestaFinal {
		vm1: mv1,
		vm2: mv2,
		vm3: mv3,
	}

	return &propose, nil
}
