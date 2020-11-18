package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"path/filepath"
)

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

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Quieres subir '" + file + "'? [s/n]")
    text, _ := reader.ReadString('\n')
	
	if text == "s\n" {
		fmt.Println("subien12")
		return file
	} else if text == "n\n" {
		folderReader()
	}

	return file
	
}

func main() {

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
		partBuffer := make([]byte, partSize)

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

}