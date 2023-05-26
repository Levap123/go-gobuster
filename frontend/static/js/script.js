
    const socket = new WebSocket("ws://localhost:8080/bust"); // Замените на актуальный URL вашего WebSocket сервера

    socket.addEventListener("open", (event) => {
    console.log("Connected to WebSocket server");
});

    socket.addEventListener("message", (event) => {
    const message = event.data;
    addMessageToDOM(message);
});

    socket.addEventListener("close", (event) => {
    console.log("Disconnected from WebSocket server");
});

    function sendMessage() {
        const inputElement = document.getElementById("inputMessage");
        let message = inputElement.value;

        message = sanitizeURL(message);

        socket.send(JSON.stringify({ url: message }));
        inputElement.value = "";
    }

    function addMessageToDOM(message) {
    const messagesElement = document.getElementById("messages");
    const messageElement = document.createElement("div");
    messageElement.textContent = message;
    messageElement.className = "typing";
    messagesElement.appendChild(messageElement);
}
    function stopWebSocket() {
        socket.close();
    }
    //сheck input
    function sanitizeURL(inputURL) {

        if (!inputURL.startsWith('http://') && !inputURL.startsWith('https://')) {
            inputURL = 'http://' + inputURL;
        }

        if (inputURL.endsWith('/')) {
            inputURL = inputURL.slice(0, -1);
        }

        return inputURL;
    }