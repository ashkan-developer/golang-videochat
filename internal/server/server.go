package server

import (
	"flag"
	"os"
)

var (
	addr = flag.String("addr", "", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()
	if *addr == ":" {
		*addr = ":8080"
	}

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", handlers.Welcome)
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.new(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.new(handlers.RoomViewerWebsocket))
	app.Get("/strean/:ssuid", handlers.Stream)
	app.Get("/strean/:ssuid/websocket", handlers.Stream)
	app.Get("/strean/:ssuid/chat/websocket", handlers.Stream)
	app.Get("/strean/:ssuid/viewer/websocket", handlers.Stream)

}
