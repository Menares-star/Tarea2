package main

import (
  "time"
  "io"
  "os"
  "fmt"
  "log"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "encoding/csv"
  "strconv"
  "math/rand"
  "github.com/Menares-star/Tarea1/src/Mensajes"
)


type Orden struct{
  Id string
  Producto string
  Valor int32
  Tienda string
  Destino string
  Prioridad int32
  Codigo int32
  Tipo string
}

func RemoveIndex(s []Orden, index int) []Orden {
    return append(s[:index], s[index+1:]...)
}

func main() {
  var age int
  fmt.Println("Ingresa un tiempo para ordenes mayor que 10: ")
  fmt.Scan(&age)
 /* CONEXION*/
  var conn *grpc.ClientConn
  conn, err := grpc.Dial(":7000", grpc.WithInsecure(), grpc.WithBlock())
  if err != nil{
    log.Fatalf("could not connect: %s",err)
  }
  defer conn.Close()

  c := ordenes.NewOrdenServiceClient(conn)
  s := ordenes.NewSeguimientoServiceClient(conn)
  /*slice y map*/
  var siguiendo []Orden
  var recientes []Orden

  /* Leyendo CSV Pymes*/
  f,err := os.Open("pymes.csv")
  if err != nil{
    log.Printf("error abriendo el archivo: %v",err)
  }
  defer f.Close()

  r := csv.NewReader(f)
  r.Comma = ','
  r.Comment = '#'
  r.FieldsPerRecord = -1

  for{
    record, err := r.Read()
    if err ==io.EOF{
      break
    }
    if err!= nil{
      log.Printf("error leyendo linea: %v",err)
    }
    v,e:=strconv.Atoi(record[2])
    if e!=nil{
      log.Printf("error tratando de procesar la valor: %v",e)
      continue
    }
    pr,h:=strconv.Atoi(record[5])
    if h!=nil{
      log.Printf("error tratando de procesar la prioridad: %v",h)
      continue
    }
    p:= Orden{
      Id: record[0],
      Producto: record[1],
      Valor:int32(v),
      Tienda: record[3] ,
      Destino: record[4] ,
      Prioridad: int32(pr),
      Codigo:0,
      Tipo:"pymes" ,
    }
    recientes=append(recientes,p)
  }
  /* Leyendo CSV Retail*/
  a,err := os.Open("retail.csv")
  if err != nil{
    log.Printf("error abriendo el archivo: %v",err)
  }
  defer a.Close()

  y := csv.NewReader(a)
  y.Comma = ','
  y.Comment = '#'
  y.FieldsPerRecord = -1

  for{
    record, err := y.Read()
    if err ==io.EOF{
      break
    }
    if err!= nil{
      log.Printf("error leyendo linea: %v",err)
    }
    g,e:=strconv.Atoi(record[2])
    if e!=nil{
      log.Printf("error tratando de procesar la valor: %v",e)
      continue
    }

    re:= Orden{
      Id: record[0],
      Producto: record[1],
      Valor:int32(g),
      Tienda: record[3] ,
      Destino: record[4] ,
      Prioridad: -1,
      Codigo:0,
      Tipo:"retail" ,
    }
    recientes=append(recientes,re)
  }

  fmt.Println(recientes)
  fmt.Println(len(recientes))
  for i:=0;len(recientes)>0;i++{
    rand.Seed(time.Now().UTC().UnixNano())
  	oal:= rand.Intn(len(recientes))
    // Remove the element at index i from a.
    message := ordenes.Orden{
      Id :recientes[oal].Id,
      Producto:recientes[oal].Producto,
      Valor:recientes[oal].Valor,
      Tienda:recientes[oal].Tienda,
      Destino:recientes[oal].Destino,
      Prioridad:recientes[oal].Prioridad,
      Codigo:recientes[oal].Codigo,
      Tipo:recientes[oal].Tipo,
    }

    response,err := c.ReceivedOrden(context.Background(),&message)
    if err!= nil{
      log.Fatalf("Error when calling SayHello: %s",err)
    }

    log.Printf("Response from Server: %d",response.Codigo)

    recientes[oal].Codigo=response.Codigo
    //Println(recientes[oal])
    siguiendo=append(siguiendo,recientes[oal])
    //Eliminando de la lista de ordenes
    recientes= RemoveIndex(recientes,oal)
    fmt.Println(len(recientes))
    rand.Seed(time.Now().UTC().UnixNano())
  	rseg:= rand.Intn(age/2)+1 //HACER RSEG CONSULTAS DE SEGUIMIENTO
    for j:=0;j<rseg && len(siguiendo)>3;j++{
      rand.Seed(time.Now().UTC().UnixNano())
    	sal:= rand.Intn(len(siguiendo))
      //CONSULTAR SEGUIMIENTO CON CODIGO DEL ELEMENTO EN POSICION SAL
      consulta := ordenes.Seguimiento{
        Codigo:siguiendo[sal].Codigo,
        Estado:"",
      }

      respuesta,err := s.ReceivedSeguimiento(context.Background(),&consulta)
      if err!= nil{
        log.Fatalf("Error when calling SayHello: %s",err)
      }

      log.Printf("Response from Server: %s",respuesta.Estado)
      ///FIN CONSULTA SEGUIMEINTO
      time.Sleep(2* time.Second)
    }
    fmt.Println(rseg)
    time.Sleep(time.Duration(age-(2*rseg))* time.Second)
  }

}
