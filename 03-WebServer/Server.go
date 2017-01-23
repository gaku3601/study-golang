package main

import (
  "fmt"
  "log"
  "net/http"
)

func helloWorld(w http.ResponseWriter,r *http.Request){
  fmt.Fprintf(w,"HelloWorld")
}

func main(){
  http.HandleFunc("/",helloWorld)
  if err := http.ListenAndServe(":8080",nil); err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
