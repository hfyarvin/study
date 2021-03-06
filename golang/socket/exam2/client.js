var wsServer = "ws://localhost:1234";
var websocket = new WebSocket(wsServer);

websocket.onopen = function(evt) {
    console.log("Connected to WebSocket server.");
};

websocket.onclose = function(evt) {
    console.log("Disconnected");
};

websocket.onmessage = function(evt) {
    console.log('Retrieved data from server: ' + evt.data);
};

websocket.onerror = function(evt) {
    console.log('Error occured: ' + evt.data);
};