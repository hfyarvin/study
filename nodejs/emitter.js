var events = require("events");

var emitter = new events.EventEmitter();

var eHandler = function eh(arg1, arg2) {
    console.log("触发:", arg1, arg2);
}

emitter.on("some", eHandler);

var timeoutHandler = function tH() {
    emitter.emit("some", "arg1", "arg2");
}

setTimeout(timeoutHandler, 1000);//函数延时