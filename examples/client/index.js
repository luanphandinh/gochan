var client = new WebSocket("ws://localhost:3000/status");
var textArea = document.getElementById("text");
client.onmessage = function (event) {
    textArea.innerHTML = textArea.innerHTML + event.data + '<br />'
};
