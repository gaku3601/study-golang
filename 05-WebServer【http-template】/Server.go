package main

import(
  "fmt"
  "net/http"
  "log"
  "html/template"
  "time"
)

type TemplateData struct{
  Title string
  Datetime string
  Unixtime string
}

func HelloServer(w http.ResponseWriter, r *http.Request){
  tmpl := template.Must(template.ParseFiles("views/index.html","views/body.html"))
  title := "現在の時刻"
  datetime := fmt.Sprint(time.Now())
  unixtime := fmt.Sprint(time.Now().Unix())

  templatedata := TemplateData{title,datetime,unixtime}
  if err := tmpl.ExecuteTemplate(w, "base", templatedata); err != nil {
    fmt.Println(err)
  }
}

func main(){
  http.HandleFunc("/",HelloServer)
  log.Printf("Start Go HTTP Server")
  err := http.ListenAndServe(":8080",nil)
  if err != nil {
    log.Fatal("ListenAndServer: ", err)
  }
}
