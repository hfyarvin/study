var e = require("events");

var emitter = new e.EventEmitter();

var connectHandler = function connected() {
    console.log("connect success.");

    emitter.emit("data_received");
}
emitter.on("connection", connectHandler);

var receicedHandler = function receiced() {
    console.log("received success");
}
emitter.on("data_received", receicedHandler);

emitter.emit("connection");

console.log("over");