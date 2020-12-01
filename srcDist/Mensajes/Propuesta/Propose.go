package Propose

import (
	"golang.org/x/net/context"
)

type Server struct{
}

func (s *Server) Proponer(ctx context.Context, prop *Propuesta) (*Respuesta, error) {

	resp := Respuesta {
		Status: "ACEPTO",
	}

	return &resp, nil
}
