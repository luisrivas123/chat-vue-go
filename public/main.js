// const socket = io();

var socket = io();

new Vue({
  el: '#chat-app',
    created() {
      socket.on("chat message", (message) => {
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
      sendMessage() {
        socket.emit("chat message", this.message)
        this.message = "";
      }
    }
})

// let msg = {
//     name: "jhjhjk",
//     lastName: "kjkjlk"
// }

// socket.on('some:event', function (msg, sendAckCb) {
//     //Sending ACK with data to server after receiving some:event from server
//     sendAckCb(JSON.stringify(data)); // for example used serializing to JSON
// })