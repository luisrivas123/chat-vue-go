package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	
	// Inicializar el servidor 
	server := socketio.NewServer(nil)
	
	// SOCKETS

	// Cuando ocurra un evento de conexión
	// Gauarda los datos del cliente conectado en la variable s

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		s.Join("chat_room")
		fmt.Println("connected:", s.ID())

		return nil
	})

	// Escucha el evento "chat message"
	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		fmt.Println("chat message:", msg)
		// s.Emit("reply", msg, "chat_room")
		server.BroadcastToRoom("", "chat_room", "reply", msg)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	go server.Serve()
	defer server.Close()

	// conexión http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
