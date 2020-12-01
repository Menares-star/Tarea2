package Uploads

import (
	"log"
  "golang.org/x/net/context"
	"io"
	"fmt"
	"google.golang.org/grpc"
	"github.com/Menares-star/Tarea2/srcCent/Mensajes/Propuesta"
	"io/ioutil"
	"os"
	"strconv"
)

type Server1 struct{
}


func (s *Server1) Upload(stream GuploadService_UploadServer) error {
	//x := make(map[int][]Chunk)
	var save [] Chunk
	cont := 0
	for i:=0;;i++ {
		str, err := stream.Recv()
		if err == io.EOF {
			break
		}
		cont = cont + 1

		save = append(save,*str)

		if err != nil {
			return stream.SendAndClose(&UploadStatus{
				Message:   "Su estado es CHUNKS NO ENVIADOS",
			})
		}
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9002",grpc.WithInsecure(),grpc.WithBlock())
	if err != nil{
		fmt.Println("Servidor :9002 no disponible: ", err)
	}
	defer conn.Close()

	propose := Propose.Propuesta {
		Vm1: ":8001",
		Vm2: ":7000",
		Vm3: ":6000",
	}

  var port string = save[len(save)-1].Puerto
  var name string = save[len(save)-1].Name

	Info := Propose.InfoMaquina {
  	Puerto: port,
  	Propuesta: &propose,
  	Nchunks: int64(cont),
    Name: name,
	}

	prop := Propose.NewProponerServiceClient(conn)
	propFinal, err := prop.Proponer(context.Background(), &Info)

	fmt.Println("Resp: vm1(" + propFinal.Vm1 + ")")
	fmt.Println("Resp: vm2(" + propFinal.Vm2 + ")")
	fmt.Println("Resp: vm3(" + propFinal.Vm3 + ")")

  var i int = 0

  if propFinal.Vm1==":8001"{
    if propFinal.Vm1==port{
      for ;i<int(propFinal.Cant1);i++{
        fileName := "Node1/"+name +"_part_" + strconv.FormatUint(uint64(i), 10) + ".pdf"
    		_, err := os.Create(fileName)

    		if err != nil {
    				fmt.Println(err)
    				os.Exit(1)
    		}
        //write/save buffer to disk
  			ioutil.WriteFile(fileName, save[i].Content, os.ModeAppend)
      }
    }else{

      var conn1 *grpc.ClientConn
      conn1, err1 := grpc.Dial(":8001",grpc.WithInsecure(),grpc.WithBlock())
      if err1 != nil{
        fmt.Println("Servidor :8001 no disponible: ", err1)
      }
      defer conn1.Close()
      streaming:= NewRepartirServiceClient(conn1)
  		streaml, err := streaming.Repartir(context.Background())
      if err != nil {
  			log.Fatalf("%v.Repartir(_) = _, %v", streaming, err)
		  }
		  fmt.Println("ENVIANDO A NODO 1")
      for ;i<int(propFinal.Cant1);i++{
        save[i].Puerto = ":8001"
        chunk:= Chunk{
  				Content: save[i].Content,
  				Name: save[i].Name,
  				Part: save[i].Part,
  				Puerto: save[i].Puerto,
  	    	}
        if err := streaml.Send(&chunk); err != nil {
  				log.Fatalf("%v.Send(%v) = %v", streaml, chunk, err)
  			}
      }
      reply, err := streaml.CloseAndRecv()
  		if err != nil {
  			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", streaml, err, nil)
  		}
  		fmt.Println(reply.Message)
    }
  }

  if propFinal.Vm2==":7000"{
    if propFinal.Vm2==port{
      for ;i<int(propFinal.Cant1+propFinal.Cant2);i++{
        fileName := "Node2/"+name +"_part_" + strconv.FormatUint(uint64(i), 10) + ".pdf"
    		_, err := os.Create(fileName)

    		if err != nil {
    				fmt.Println(err)
    				os.Exit(1)
    		}
        //write/save buffer to disk
  			ioutil.WriteFile(fileName, save[i].Content, os.ModeAppend)
      }
    }else{
      var conn2 *grpc.ClientConn
      conn2, err2 := grpc.Dial(":7000",grpc.WithInsecure(),grpc.WithBlock())
      if err2 != nil{
        fmt.Println("Servidor :7000 no disponible: ", err2)
      }
      defer conn2.Close()
      streaming:=NewRepartirServiceClient(conn2)
  		streaml, err := streaming.Repartir(context.Background())
      if err != nil {
  			log.Fatalf("%v.Repartir(_) = _, %v", streaming, err)
  		}
      fmt.Println("ENVIANDO A NODO 2")
      for ;i<int(propFinal.Cant1+propFinal.Cant2);i++{
        save[i].Puerto = ":7000"
        chunk:= Chunk{
          Content: save[i].Content,
          Name: save[i].Name,
          Part: save[i].Part,
          Puerto: save[i].Puerto,
          }
        if err := streaml.Send(&chunk); err != nil {
  				log.Fatalf("%v.Send(%v) = %v", streaml, chunk, err)
  			}
      }
      reply, err4 := streaml.CloseAndRecv()
  		if err4 != nil {
  			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", streaml, err4, nil)
  		}
  		fmt.Println(reply.Message)

    }
  }

  if propFinal.Vm3==":6000"{
    if propFinal.Vm3==port{
      for ;i<int(propFinal.Cant1+propFinal.Cant2+propFinal.Cant3);i++{
        fileName := "Node3/"+name +"_part_" + strconv.FormatUint(uint64(i), 10) + ".pdf"
    		_, err := os.Create(fileName)

    		if err != nil {
    				fmt.Println(err)
    				os.Exit(1)
    		}
        //write/save buffer to disk
  			ioutil.WriteFile(fileName, save[i].Content, os.ModeAppend)
      }
    }else{
      var conn3 *grpc.ClientConn
      conn3, err3 := grpc.Dial(":6000",grpc.WithInsecure(),grpc.WithBlock())
      if err3 != nil{
        fmt.Println("Servidor :6000 no disponible: ", err3)
      }
      defer conn3.Close()
      streaming:= NewRepartirServiceClient(conn3)
  		streaml, err := streaming.Repartir(context.Background())
      if err != nil {
  			log.Fatalf("%v.Repartir(_) = _, %v", streaming, err)
		  }
		  fmt.Println("ENVIANDO A NODO 3")
      for ;i<int(propFinal.Cant1+propFinal.Cant2+propFinal.Cant3);i++{
        save[i].Puerto = ":6000"
        chunk:= Chunk{
          Content: save[i].Content,
          Name: save[i].Name,
          Part: save[i].Part,
          Puerto: save[i].Puerto,
          }
        if err := streaml.Send(&chunk); err != nil {
  				log.Fatalf("%v.Send(%v) = %v", streaml, chunk, err)
  			}
      }
      reply, err := streaml.CloseAndRecv()
  		if err != nil {
  			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", streaml, err, nil)
  		}
  		fmt.Println(reply.Message)

    }
  }

	return stream.SendAndClose(&UploadStatus{
		Message:   "Su estado es CHUNKS ENVIADOS",
	})

}

func (s *Server1) Repartir(stream1 RepartirService_RepartirServer) error {
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

    if str.Puerto==":7000"{
      fileName := "Node2/"+str.Name +"_part_" + strconv.FormatUint(uint64(str.Part), 10) + ".pdf"
      _, err := os.Create(fileName)

      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }
      //write/save buffer to disk
      ioutil.WriteFile(fileName, str.Content, os.ModeAppend)
    }

    if str.Puerto==":6000"{
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
		Message: "CHUNKS ALMACENADOS",
	}
  fmt.Println("Se han recibido los chunks")
  return stream1.SendAndClose(&msg)
}

func (s *Server1) Download(stream DownloadService_DownloadServer) error {
		var trozo string
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			/*var Dir string = ""
			if in.Port == ":8001"{
				Dir="Node1"
			}
			if in.Port == ":7000"{
				Dir="Node2"
			}
			if in.Port == ":6000"{
				Dir="Node3"
			}*/
			/*fileToBeSend:=in.Namepart+".pdf"
			//outputDirRead, _ := os.Open("/"+Dir+"/")
			file, err := os.Open(fileToBeSend)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			//defer file.Close()
			fileInfo, _ := file.Stat()
			var fileSize int64 = fileInfo.Size()*/
			partBuffer := make([]byte, 0) //inicializa un arreglo de tamaño partSize

			//file.Read(partBuffer)

			trozo = in.Namepart+".pdf"

			// write to disk
			content:= Content{
				Content: partBuffer,
				Nametrozo: in.Namepart,
	    	}
			if err := stream.Send(&content); err != nil {
			 return err
			}

		}
		//outputDirRead, _ := os.Open("/"+Dir+"/")
		//fmt.Println("Voy a abrir el trozo: ")
		trozo = "Uploads/Alicia_en_el_pais_de_las_maravillas-Carroll_Lewis.pdf"
		fmt.Println(trozo)
		file, err := os.Open(trozo)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//defer file.Close()
		fileInfo, _ := file.Stat()
		var fileSize int64 = fileInfo.Size()
		fmt.Println(fileSize)
		partBuffer := make([]byte, fileSize) //inicializa un arreglo de tamaño partSize

		file.Read(partBuffer)
	return nil
}