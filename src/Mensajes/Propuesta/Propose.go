package Propose

import (
	//	"log"
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
	
	mv1 := InfoMaquina {
		Puerto:":8001",
		IsAvailable: int32(1),
	}
	mv2 := InfoMaquina {
		Puerto: ":7000",
		IsAvailable: int32(0),
	}
	mv3 := InfoMaquina {
		Puerto: ":6000",
		IsAvailable: int32(1),
	}

	propose := PropuestaFinal {
		Vm1: &mv1,
		Vm2: &mv2,
		Vm3: &mv3,
	}

	return &propose, nil
}
