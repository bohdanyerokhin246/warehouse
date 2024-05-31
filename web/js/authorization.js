let button = document.querySelector(".form > button");
if (button) {
    button.onclick = function (e) {
        let user = document.querySelectorAll(".form > input");

        let data = {};
        data.login = user[0].value;
        data.password = user[1].value;


        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/user/auth");
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response);
            if ("Error" in response) {
                if (response.Error == null) {
                   document.location.href = "/adminPage?managerName=" + response.Name
                } else {
                    console.log(response.Error);
                }
            } else {
                    alert("Введено некоректні дані або такого користувача не існує");
            }
        };
        xhr.send(JSON.stringify(data));
    }
}