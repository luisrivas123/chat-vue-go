const socket = io();

new Vue({
  el: '#chat-app',
    created() {
      // Escucha el envento del servidor
      socket.on("reply", (message) => {
        console.log(message)
        this.messages.push({
          text: message,
          date: new Date().toLocaleString()
        })
      })
    },
    data: {
      message: '',
      messages: []
    },
    methods: {
      // Envía el mensaje del cliente
      sendMessage() {
        socket.emit("chat message", this.message)
        this.message = "";
      }
    }
})