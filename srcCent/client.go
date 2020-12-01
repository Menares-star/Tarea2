package main

import (
	"bufio"
	"log"
	//"errors"
	"fmt"
	//"io/ioutil"
	"context"
	"github.com/Menares-star/Tarea2/src/Mensajes/Propuesta"
	"github.com/Menares-star/Tarea2/srcCent/Mensajes/Upload"
	"google.golang.org/grpc"
	"io"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func folderReader() string {

	var files []string

	root := "./Uploads"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, info.Name())
		return nil
	})
	if err != nil {
		panic(err)
	}
	for num, fil := range files {
		fmt.Println(strconv.Itoa(num+1) + ".- " + fil)
	}

	var upload int
	fmt.Println("Ingrese el numero del libro que desea subir: ")
	fmt.Scan(&upload)

	file := files[upload-1]

	reader := bufio.NewReader(os.Stdin) //habilita un lector
	fmt.Println("Quieres subir '" + file + "'? [s/n]")
	text, _ := reader.ReadString('\n') //el lector lee un string

	if text == "s\n" {
		fmt.Println("subien12")
		return file
	} else if text == "n\n" {
		folderReader()
	}

	return file

}

func main() {
	var flag bool = true
	servers := make([]string, 3)
	servers[0] = "dist74:8001"
	servers[1] = "dist75:7000"
	servers[2] = "dist76:6000"
	for flag == true {

		var cliente int
		fmt.Println("Que tipo de cliente es?: ")
		fmt.Println("1. Uploader")
		fmt.Println("2. Downloader")
		fmt.Println("3. Exit")
		fmt.Println("Ingrese el número: ")
		fmt.Scan(&cliente)

		if cliente == 1 {

			var t time.Duration = 5000000000
			var p string = servers[0]
			/* CONEXION*/
			//conn, err := grpc.Dial(":8000", grpc.WithInsecure(), grpc.WithBlock())
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(servers[0], grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(t))
			if err != nil {
				fmt.Println("Servidor "+servers[0]+" no disponible: ", err)
				conn, err = grpc.Dial(servers[1], grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(t))
				if err != nil {
					fmt.Println("Servidor "+servers[1]+" no disponible: ", err)
					conn, err = grpc.Dial(servers[2], grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(t))
					if err != nil {
						fmt.Println("Servidor "+servers[2]+" no disponible: ", err)
						fmt.Println("TODOS LOS SERVIDORES ESTÁN CAIDOS")
						os.Exit(1)
					} else {
						p = servers[2]
						defer conn.Close()
					}
				} else {
					p = servers[1]
					defer conn.Close()
				}
			}
			defer conn.Close()

			//REGISTRANDO SERVICIO POR PARTE DE CLIENTE
			streaming := Uploads.NewGuploadServiceClient(conn)
			stream, err := streaming.Upload(context.Background())
			if err != nil {
				log.Fatalf("%v.Upload(_) = _, %v", streaming, err)
			}

			/*ESCOGIENDO LIBRO*/
			fileName := folderReader()

			fileToBeChunked := "./Uploads/" + fileName

			file, err := os.Open(fileToBeChunked)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			defer file.Close()

			fileInfo, _ := file.Stat()

			var fileSize int64 = fileInfo.Size()

			const fileChunk = 256000

			// calculate total number of parts the file will be chunked into

			totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

			fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

			for i := uint64(0); i < totalPartsNum; i++ {

				partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
				partBuffer := make([]byte, partSize) //inicializa un arreglo de tamaño partSize

				file.Read(partBuffer)

				// write to disk
				chunk := Uploads.Chunk{
					Content: partBuffer,
					Name:    fileName[0:(len(fileName) - 4)],
					Part:    int32(i),
					Puerto:  p,
				}
				/*fileName := fileName[0:(len(fileName)-4)] +"_part_" + strconv.FormatUint(i, 10) + ".pdf"
				_, err := os.Create(fileName)

				if err != nil {
						fmt.Println(err)
						os.Exit(1)
				}*/
				///STREAM DE CHUNKS
				if err := stream.Send(&chunk); err != nil {
					log.Fatalf("%v.Send(%v) = %v", stream, chunk, err)
				}
				// write/save buffer to disk
				//ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
				fileName := fileName[0:(len(fileName)-4)] + "_part_" + strconv.FormatUint(i, 10) + ".pdf"
				fmt.Println("Split to : ", fileName)
			}

			reply, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
			}
			fmt.Println(reply.Message)
		}
		if cliente == 2 {

			fmt.Println("Soy downloader!")
			//fmt.Println("SORRY, NOT WORKING")

			var t time.Duration = 5000000000
			var conn *grpc.ClientConn
			conn, err := grpc.Dial("dist73:9002", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(t))
			defer conn.Close()
			if err != nil {
				fmt.Println("Servidor dist73:9002 no disponible: ", err)
			}
			want := Propose.ReceptStatus{
				Message: "ENTREGAME LA LISTA DE LIBROS",
			}
			streaming := Propose.NewListarServiceClient(conn)
			stream, err := streaming.Listar(context.Background(), &want)
			if err != nil {
				log.Fatalf("%v.Listar(_) = _, %v", streaming, err)
			}
			var save []string
			for i := 1; ; i++ {
				Libro, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("%v.Listar(_) = _, %v", streaming, err)
				}
				fmt.Printf("%d. %s\n", i, Libro.Name)
				save = append(save, Libro.Name)
			}
			var namelibro int
			fmt.Println("Ingrese el número: ")
			fmt.Scan(&namelibro)

			var book string
			book = save[namelibro-1]
			fmt.Println(book)

			sendbook := Propose.Libro{
				Name: book,
			}

			var conn1 *grpc.ClientConn
			conn1, err1 := grpc.Dial("dist73:9002", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(t))
			defer conn1.Close()
			if err1 != nil {
				fmt.Println("Servidor dist73:9002 no disponible: ", err1)
			}
			streaming1 := Propose.NewListarChunksServiceClient(conn1)
			stream1, err := streaming1.ListarChunks(context.Background(), &sendbook)
			if err != nil {
				log.Fatalf("%v.Listar(_) = _, %v", streaming1, err)
			}
			var save1 []string
			var save2 []string
			var save3 []string
			for {
				Ichunk, err := stream1.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("%v.Listar(_) = _, %v", streaming, err)
				}
				fmt.Printf("%s\n", Ichunk.Part)
				if Ichunk.Puerto == servers[0] {
					save1 = append(save1, Ichunk.Part)
				}
				if Ichunk.Puerto == servers[1] {
					save2 = append(save2, Ichunk.Part)
				}
				if Ichunk.Puerto == servers[2] {
					save3 = append(save3, Ichunk.Part)
				}
			}

			var conn11 *grpc.ClientConn
			conn11, err11 := grpc.Dial(servers[0], grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(t))
			defer conn11.Close()
			if err11 != nil {
				fmt.Println("Servidor"+servers[0]+"no disponible: ", err11)
			}
			streaming11 := Uploads.NewDownloadServiceClient(conn11)
			stream11, err := streaming11.Download(context.Background())
			if err != nil {
				log.Fatalf("%v.Download(_) = _, %v", streaming11, err)
			}
			waitc := make(chan struct{})
			go func(){
				for i:=0;i<len(save1);i++ {
					info:=Uploads.InfoChunk{
						Port: servers[0],
						Namepart: save1[i],
					}
					if err := stream11.Send(&info); err != nil {
						log.Fatalf("Failed to send a info: %v", err)
					}
				}
				if err := stream11.CloseSend(); err != nil {
					log.Println(err)
				}
			}()
			go func() {
				for {
					in, err := stream11.Recv()
					if err == io.EOF {
						// read done.
						close(waitc)
						return
					}
					if err != nil {
						log.Fatalf("Failed to receive a note : %v", err)
					}
					fmt.Println(in.Nametrozo)
				}
			}()
			/*var conn2 *grpc.ClientConn
			conn2, err2 := grpc.Dial(servers[1],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
			defer conn2.Close()
			if err2 != nil{
				fmt.Println("Servidor"+servers[1]+"no disponible: ",err2)
			}
			streaming2 := Uploads.NewDownloadServiceClient(conn2)
			stream2, err := streaming2.Download(context.Background())
			if err != nil {
				log.Fatalf("%v.Download(_) = _, %v", streaming2, err)
			}


			var conn3 *grpc.ClientConn
			conn3, err3 := grpc.Dial(servers[1],grpc.WithInsecure(),grpc.WithBlock(),grpc.WithTimeout(t))
			defer conn3.Close()
			if err3 != nil{
				fmt.Println("Servidor"+servers[2]+"no disponible: ",err3)
			}
			streaming3 := Uploads.NewDownloadServiceClient(conn3)
			stream3, err := streaming3.Download(context.Background())
			if err != nil {
				log.Fatalf("%v.Download(_) = _, %v", streaming3, err)
			}
			var totalParts int = len(save1)+len(save2)+len(save3)*/

		}
		if cliente == 3 {
			flag = false
		}
	}
}
