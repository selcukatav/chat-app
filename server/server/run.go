package server

import (
	"github.com/selcukatav/chat-app/api/route"
	"github.com/selcukatav/chat-app/database"
)


func Run() {
	e := route.New()
	database.Postgres()
	database.Redis()
	e.Logger.Fatal(e.Start(":3000"))
}
