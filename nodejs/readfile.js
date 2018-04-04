// 阻塞
var fs = require("fs");

var data = fs.readFileSync("input.txt");

console.log(data.toString());//读取文件信息
console.log("over!");