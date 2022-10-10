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

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// sockets
	// Eventos del lado del client

	// Sever cuando ocurra un evento de conexión
	// Gauarda los datos delc liente conectado en la variable so
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	// conexión http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
