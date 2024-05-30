const searchParams = new URLSearchParams(window.location.search);

let urlManagerName = searchParams.get('managerName')
let managerName = urlManagerName.replace('_', ' ')

let h1 = '<h1>Hello, </h1>' + managerName

document.write(h1)