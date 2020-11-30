package Uploads

import (
	"log"
  	"golang.org/x/net/context"
	"io"
	"fmt"
	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea2/src/Mensajes/Propuesta"
	"github.com/Menares-star/Tarea2/src/Mensajes/Notification"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Server1 struct{
}

func (s *Server1) Upload(stream GuploadService_UploadServer) error {

	var save [] Chunk
	cont := 0
	for i:=0;;i++ {
		str, err := stream.Recv()
		if err == io.EOF {
			break
		}
		cont++

		save = append(save,*str)

		if err != nil {
			return stream.SendAndClose(&UploadStatus{
				Message:   "Su estado es CHUNKS NO ENVIADOS",
			})
		}
	}

	/* CUANTOS CHUNKS ENVIO A CADA UNO SI ES QUE TODOS ESTAN VIVOS */
	var cociente int = cont / 3
	var resto int = cont % 3
	var cociente2 int = cont / 2
	var resto2 int = cont % 2

	var chunks_a_1 int = cociente
	var chunks_a_2 int = cociente
	var chunks_a_3 int = cociente

	if resto > 0 {
		chunks_a_1++
		if resto == 2{
			chunks_a_2++
		}
	}

	/* QUIENES SON MIS COMPAÑEROS DATANODES */
	servers:=make([]string,3)
	servers[0]=":8001"
	servers[1]=":7001"
	servers[2]=":6001"

	s1 := "NOT SET"
	s2 := "NOT SET"
	s0 := "DUNNO"

	if save[0].Puerto == servers[0]{
		s1 = ":7002"
		s2 = ":6002"
		s0 = "Node1"
	}
	if save[0].Puerto == servers[1]{
		s1 = ":8002"
		s2 = ":6002"
		s0 = "Node2"
	}
	if save[0].Puerto == servers[2]{
		s1 = ":8002"
		s2 = ":7002"
		s0 = "Node3"
	}

	/* CREO CONEXIONES A DATANODES */
	var t time.Duration = 5000000000
	status1 := "NOT SET"
	status2 := "NOT SET"

	var conn1 *grpc.ClientConn
	conn1, err := grpc.Dial(s1,grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
	if err != nil{
		fmt.Println("Servidor " + s1 + " no disponible: ", err)
		s1 = "NOT AVAILABLE"
		status1 = "No me mandes cositas rey"
	}else{
		fmt.Println("Servidor " + s1 + " esta disponible")
		defer conn1.Close()
	}

	var conn2 *grpc.ClientConn
	conn2, err2 := grpc.Dial(s2,grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
	if err2 != nil{
		fmt.Println("Servidor " + s2 + " no disponible: ", err2)
		s2 = "NOT AVAILABLE"
		status2 = "No me mandes cositas rey"
	}else{
		fmt.Println("Servidor " + s2 + " esta disponible")
		defer conn2.Close()
	}


	/* PROPUESTAS */
	propose1 := Propose.Propuesta {
		Puerto: save[0].Puerto,
		NChunks: int64(chunks_a_2),
	}

	propose2 := Propose.Propuesta {
		Puerto: save[0].Puerto,
		NChunks: int64(chunks_a_3),
	}

	/* ENVIO Y RECIBO RESPUESTAS */
	if s1 != "NOT AVAILABLE"{
		prop1 := Propose.NewProponerServiceClient(conn1)
		resp1, err3 := prop1.Proponer(context.Background(), &propose1)
		if err3 != nil {
			fmt.Println(s1 + " no acepto la propuesta ",err3)
		}else {
			status1 = resp1.Status
		}
	} else {
		chunks_a_2 = 0

		chunks_a_1 = cociente2
		chunks_a_3 = cociente2

		if resto2 > 0 {
			chunks_a_1++
		}
	}

	if s2 != "NOT AVAILABLE"{
		prop2 := Propose.NewProponerServiceClient(conn2)
		resp2, err4 := prop2.Proponer(context.Background(), &propose2)
		if err4 != nil {
			fmt.Println(s2 + " no acepto la propuesta ",err4)
		}else {
			status2 = resp2.Status
		}
	} else {
		chunks_a_3 = 0

		chunks_a_1 = cociente2
		chunks_a_2 = cociente2

		if resto2 > 0 {
			chunks_a_1++
		}
	}

	if (s1 == "NOT AVAILABLE") && (s2 == "NOT AVAILABLE") {
		chunks_a_1 = cont
		chunks_a_2 = 0
		chunks_a_3 = 0
	}

	fmt.Println("s1: " + s1 + ", resp: " + status1)
	fmt.Println("s2: " + s2 + ", resp: " + status2)
	fmt.Println("Chunks 1: " + strconv.Itoa(chunks_a_1) + ", Chunks 2: " + strconv.Itoa(chunks_a_2) + ", Chunks 3: " + strconv.Itoa(chunks_a_3))

	/* GUARDO CHUNKS Y ENVIO A NODOS CORRESPONDIENTES */
	name := save[0].Name
	var i int = 0
	for ; i<chunks_a_1; i++ {
	
		fileName := s0 + "/" + name +"_part_" + strconv.Itoa(i) + ".pdf"
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//write/save buffer to disk
		ioutil.WriteFile(fileName, save[i].Content, os.ModeAppend)
	}

		//SOLO NODO S1 VIVO
	if s1 != "NOT AVAILABLE" && s2 == "NOT AVAILABLE"{
		var conn11 *grpc.ClientConn
		conn11, err11 := grpc.Dial(s1[0:4] + "1",grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
		if err11 != nil{
			fmt.Println("Servidor " + s1[0:4] + "1" + " no disponible: ", err11)
		}else{
			defer conn11.Close()
		}

		streaming1 := NewRepartirServiceClient(conn11)
		stream1, err111 := streaming1.Repartir(context.Background())
		if err111 != nil {
			log.Fatalf("%v.Repartir(_) = _, %v", streaming1, err111)
		}

		fmt.Println(chunks_a_1, chunks_a_2)
		for ; i<chunks_a_1+chunks_a_2; i++ {
			
			fmt.Println(i)

			chunk:= Chunk{
				Content: save[i].Content,
				Name: name,
				Part: int32(i),
				Puerto: s1[0:4] + "1",
	    	}

			if err := stream1.Send(&chunk); err != nil {
				log.Fatalf("%v.Send(%v) = %v", stream1, chunk, err)
			} else {
				fmt.Println("envia3")
			}
		}
		reply1, errr := stream1.CloseAndRecv()
		if errr != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream1, errr, nil)
		}
		fmt.Println(reply1.Message)
	} 

		//SOLO NODO S2 VIVO
	if s2 != "NOT AVAILABLE" && s1 == "NOT AVAILABLE"{
		var conn21 *grpc.ClientConn
		conn21, err21 := grpc.Dial(s2[0:4] + "1",grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
		if err21 != nil{
			fmt.Println("Servidor " + s2[0:4] + "1" + " no disponible: ", err21)
		}else{
			defer conn21.Close()
		}

		streaming2 := NewRepartirServiceClient(conn21)
		stream2, err211 := streaming2.Repartir(context.Background())
		if err211 != nil {
			log.Fatalf("%v.Repartir(_) = _, %v", streaming2, err211)
		}

		for i:=chunks_a_1; i<chunks_a_1+chunks_a_3; i++ {

			chunk:= Chunk{
				Content: save[i].Content,
				Name: name,
				Part: int32(i),
				Puerto: s2[0:4] + "1",
			}

			if err := stream2.Send(&chunk); err != nil {
				log.Fatalf("%v.Send(%v) = %v", stream2, chunk, err)
			}
		}
		reply2, errr := stream2.CloseAndRecv()
		if errr != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream2, errr, nil)
		}
		fmt.Println(reply2.Message)
	}

		//AMBOS NODOS VIVOS
	if s1 != "NOT AVAILABLE" && s2 != "NOT AVAILABLE"{
		//S1
		var conn11 *grpc.ClientConn
		conn11, err11 := grpc.Dial(s1[0:4] + "1",grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
		if err11 != nil{
			fmt.Println("Servidor " + s1[0:4] + "1" + " no disponible: ", err11)
		}else{
			defer conn11.Close()
		}

		streaming1 := NewRepartirServiceClient(conn11)
		stream11, err111 := streaming1.Repartir(context.Background())
		if err111 != nil {
			log.Fatalf("%v.Repartir(_) = _, %v", streaming1, err111)
		}

		for i=chunks_a_1; i<chunks_a_1+chunks_a_2; i++ {

			chunk:= Chunk{
				Content: save[i].Content,
				Name: name,
				Part: int32(i),
				Puerto: s1[0:4] + "1",
	    	}

			if err := stream11.Send(&chunk); err != nil {
				log.Fatalf("%v.Send(%v) = %v", stream11, chunk, err)
			}
		}
		reply11, errr := stream11.CloseAndRecv()
		if errr != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream11, errr, nil)
		}
		fmt.Println(reply11.Message)

		//S2
		var conn21 *grpc.ClientConn
		conn21, err21 := grpc.Dial(s2[0:4] + "1",grpc.WithInsecure(),grpc.WithBlock(), grpc.WithTimeout(t))
		if err21 != nil{
			fmt.Println("Servidor " + s2[0:4] + "1" + " no disponible: ", err21)
		}else{
			defer conn21.Close()
		}

		streaming2 := NewRepartirServiceClient(conn21)
		stream22, err211 := streaming2.Repartir(context.Background())
		if err211 != nil {
			log.Fatalf("%v.Repartir(_) = _, %v", streaming2, err211)
		}

		for i=chunks_a_1+chunks_a_2; i<chunks_a_1+chunks_a_2+chunks_a_3; i++ {

			chunk:= Chunk{
				Content: save[i].Content,
				Name: name,
				Part: int32(i),
				Puerto: s2[0:4] + "1",
			}

			if err := stream22.Send(&chunk); err != nil {
				log.Fatalf("%v.Send(%v) = %v", stream22, chunk, err)
			}
		}
		reply22, errr := stream22.CloseAndRecv()
		if errr != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream22, errr, nil)
		}
		fmt.Println(reply22.Message)

	}
	
	/* NOTIFICO A NAMENODE */
	
	notification := Notification.Notificacion {
		Puerto1: save[0].Puerto,
		Puerto2: s1[0:4] + "1",
		Puerto3: s2[0:4] + "1",
		NChunks1: int64(chunks_a_1),
		NChunks2: int64(chunks_a_2),
		NChunks3: int64(chunks_a_3),
		BookName: name,
	}

	var connN *grpc.ClientConn
	connN, errN := grpc.Dial(":9000",grpc.WithInsecure(),grpc.WithBlock())
	fmt.Println("Enviando distribucion a NameNode")
	if errN != nil{
		fmt.Println("Servidor :9000 no disponible: ", errN)
	}else{
		fmt.Println("Servidor :9000 esta disponible")
		defer connN.Close()
	}

	noti := Notification.NewNotificationServiceClient(connN)
	resultN, errn := noti.Notificar(context.Background(), &notification)
	if errn != nil {
		fmt.Println(":9000 no pudo recibir la notificacion: ",errn)
	}else {
		fmt.Println(":9000 Guardo la distribucion en Log exitosamente: ", resultN.Resultado)
	}
	

	return stream.SendAndClose(&UploadStatus{
		Message:   "Su estado es CHUNKS ENVIADOS",
	})

}

func (s *Server1) Repartir(stream1 RepartirService_RepartirServer) error {

	var port string = "NOT SET"
	for i:=0;;i++ {
		str, err := stream1.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			msg:=UploadStatus{
				Message: "CHUNKS NO ALMACENADOS",
			}
			return stream1.SendAndClose(&msg)
		}

		port = str.Puerto
			
		if str.Puerto==":8001"{
			fileName := "Node1/"+str.Name +"_part_" + strconv.FormatUint(uint64(str.Part), 10) + ".pdf"
			_, err := os.Create(fileName)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			//write/save buffer to disk
			ioutil.WriteFile(fileName, str.Content, os.ModeAppend)
		}

		if str.Puerto==":7001"{
			fileName := "Node2/"+str.Name +"_part_" + strconv.FormatUint(uint64(str.Part), 10) + ".pdf"
			_, err := os.Create(fileName)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			//write/save buffer to disk
			ioutil.WriteFile(fileName, str.Content, os.ModeAppend)
		}

		if str.Puerto==":6001"{
			fileName := "Node3/"+str.Name +"_part_" + strconv.FormatUint(uint64(str.Part), 10) + ".pdf"
			_, err := os.Create(fileName)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			//write/save buffer to disk
			ioutil.WriteFile(fileName, str.Content, os.ModeAppend)
		}
		
	}
	msg:=UploadStatus{
		  Message: "CHUNKS ALMACENADOS EN " + port,
	}
	return stream1.SendAndClose(&msg)
}