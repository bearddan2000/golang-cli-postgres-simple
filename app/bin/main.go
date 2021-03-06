package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

const (
  host     = "db"
  port     = 5432
  user     = "maria"
  password = "pass"
  dbname   = "beverage"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)

  if err != nil {
    panic(err)
  }

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")

  db.Close()
}
