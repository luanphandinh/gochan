var client = new WebSocket("ws://localhost:8080/status");
var textArea = document.getElementById("text");
client.onmessage = function (event) {
    textArea.innerHTML += event.data + '<br />'
};
