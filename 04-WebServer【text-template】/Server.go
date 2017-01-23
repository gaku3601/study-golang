package main

import (
  "net/http"
  "text/template"
)

type Page struct {
  Title string
  Count int
}

func viewHandler(w http.ResponseWriter,r *http.Request) {
  page := Page{"HelloWorld",1}
  tmpl, err := template.New("new").Parse("{{.Title}}{{.Count}} count")
  if err != nil {
    panic(err)
  }

  err = tmpl.Execute(w, page)
  if err != nil {
    panic(err)
  }
}

func main() {
  http.HandleFunc("/", viewHandler)
  http.ListenAndServe(":8080", nil)
}
