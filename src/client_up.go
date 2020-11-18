package main

import (
	"bufio"
	"log"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"path/filepath"
	"math/rand"
	"time"
	//"golang.org/x/net/context"
  "google.golang.org/grpc"
	//"github.com/Menares-star/Tarea1/src/Mensajes"
)

func random(min int, max int)int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min)+min
}

func folderReader() string{

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
        fmt.Println(strconv.Itoa(num + 1) + ".- " + fil)
	}

	var upload int
	fmt.Println("Ingrese el numero del libro que desea subir: ")
	fmt.Scan(&upload)

	file := files[upload - 1]

	reader := bufio.NewReader(os.Stdin) //habilita un lector
	fmt.Println("Quieres subir '" + file + "'? [s/n]")
    text, _ := reader.ReadString('\n')//el lector lee un string

	if text == "s\n" {
		fmt.Println("subien12")
		return file
	} else if text == "n\n" {
		folderReader()
	}

	return file

}

func main() {

	var cliente int
	fmt.Println("Que tipo de cliente es?: ")
	fmt.Println("1. Uploader")
	fmt.Println("2. Downloader")
	fmt.Println("Ingrese el número: ")
	fmt.Scan(&cliente)

	if cliente==1 {

		servers:=make([]string,3)
		servers[0]="8000"
		servers[1]="7000"
		servers[2]="6000"
		number:=random(0,3)
		fmt.Println(servers[number])

		/* CONEXION*/
	   /*var conn *grpc.ClientConn
	   conn, err := grpc.Dial(":8000", grpc.WithInsecure(), grpc.WithBlock())
	   if err != nil{
	     log.Fatalf("could not connect: %s",err)
	   }
	   defer conn.Close()*/


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

		const fileChunk = 250000 // 1 MB, change this to your requirement

		// calculate total number of parts the file will be chunked into

		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

		for i := uint64(0); i < totalPartsNum; i++ {

			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize) //inicializa un arreglo de tamaño partSize

			file.Read(partBuffer)

			// write to disk
			fileName := "bigfile_" + strconv.FormatUint(i, 10) + ".pdf"
			_, err := os.Create(fileName)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}


			// write/save buffer to disk
			ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

			fmt.Println("Split to : ", fileName)
		}
	}else{
		fmt.Println("Soy downloader!")
	}

}
