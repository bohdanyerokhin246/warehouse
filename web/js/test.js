document.write("<h1>Hello world</h1>")

fetch("/getTestInfo",{method:'GET'}
).then(
    response => response.text()
).then(
    (data) => {console.log(data);document.getElementById("serverMessageBox").innerHTML=data}
)