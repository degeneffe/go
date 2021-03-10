package main

import (
  "fmt"
  "net/http"
  "log"
)

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hola mundo desde mi servidor web con go")
  })

  server := http.ListenAndServe(":8080", nil)

  log.Fatal(server)

  fmt.Println("El servidor está corriendo")
}
