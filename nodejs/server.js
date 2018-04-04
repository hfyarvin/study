var http = require('http');//自带http模块

var c =http.createServer(
    function (request, response){
        response.writeHead(200, {'Conten-Type': 'text/plain'});

        response.end("Hello World\n");
    }
);

c.listen(8888);

console.log("server 8888no");