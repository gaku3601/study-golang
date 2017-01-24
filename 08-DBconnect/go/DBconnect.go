package main

import( 
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)

func main() {
	db, _ := sql.Open("postgres", "user=postgres host=localhost dbname=godbtest port=5432 sslmode=disable")
	defer db.Close()

  // SELECT（JSONのフィールドを直接指定）
  rows, _ := db.Query("SELECT data->'character'->>'hero' as hero FROM jojo;")

  //ポインタが先頭より前なので1つ目でも`Next()`が必要
  for rows.Next() {
    var hero string // 普通にstringとして読める
    rows.Scan(&hero)
    fmt.Println(hero)
  }
}
