package main

import (
  "github.com/gabrielgmusskopf/crud_go/server"
  "github.com/gabrielgmusskopf/crud_go/db"
)

func main() {
    db.Connect()
    server.Init()
}
