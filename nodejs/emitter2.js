var e = require("events");
var emitter = new e.EventEmitter();

var listener1 = function listener1() {
    console.log("listener1");
}

var listener2 = function listener2() {
    console.log("listener2");
}

emitter.addListener("connection", listener1);

emitter.on("connection", listener2);

var listeners = e.EventEmitter.listenerCount(emitter,"connection");
console.log(listeners, "个");

emitter.emit("connection");

// 移除
emitter.removeListener("connection", listener1);
console.log("remove listener1");

emitter.emit("connection");

var listeners = e.EventEmitter.listenerCount(emitter,"connection");
console.log(listeners, "个");