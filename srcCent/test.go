package main
import(
  "fmt"
  "log"
  //"bufio"
  "os"

  //"strings"
)

func main() {
  //var save[] string
  file, err := os.OpenFile("Node1/Alicia_en_el_pais_de_las_maravillas-Carroll_Lewis_parte_1.pdf", os.O_RDONLY, 0600)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  fmt.Println("leido")
  /*
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    librop:= strings.Split(scanner.Text(),"_parte")
    libro:= strings.Split(librop[0]," ")
    //fmt.Println(libro[0])
    var flag bool = false
    for i:=0;flag==false && i<len(save);i++ {
      if save[i]==libro[0]{
        flag=true
        break
      }
    }
    if flag==false{
      save = append(save,libro[0])
      fmt.Println(libro[0])
    }
  }
  */
}
