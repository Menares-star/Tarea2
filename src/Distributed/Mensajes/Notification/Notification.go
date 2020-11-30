package Notification

import (
	//	"log"
	"golang.org/x/net/context"
	//"time"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/connectivity"
	//"io"
	//"fmt"
	//"io/ioutil"
	"os"
	"strconv"
)

type ServerN struct{
}

func (s *ServerN) Notificar(ctx context.Context, noti *Notificacion) (*Respuesta, error) {

	/* GUARDAR EN LOG */
	fi, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	fi.WriteString(noti.BookName + " " + strconv.FormatUint(uint64(noti.NChunks1+noti.NChunks2+noti.NChunks3), 10) + "\n")
	
	var i int64 = 0
	for ;i<noti.NChunks1 ;i++{
		fi.WriteString(noti.BookName+"_parte_"+ strconv.FormatUint(uint64(i), 10)+" "+noti.Puerto1+"\n")
	}
	for ;i<(noti.NChunks1 +noti.NChunks2);i++{
	 	fi.WriteString(noti.BookName+"_parte_"+ strconv.FormatUint(uint64(i), 10)+" "+noti.Puerto2+"\n")
	}
	for ;i<(noti.NChunks1 +noti.NChunks2 +noti.NChunks3);i++{
	   	fi.WriteString(noti.BookName+"_parte_"+ strconv.FormatUint(uint64(i), 10)+" "+noti.Puerto3+"\n")
	}
	fi.Close()
	
	resp := Respuesta {
		Resultado: "GUARDADO EXITOSAMENTE EN LOG",
	}

	return &resp, nil
}
