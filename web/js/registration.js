let button = document.querySelector(".form > button");
if (button) {
    button.onclick = function (e) {
        let inputs = document.querySelectorAll(".form > input");
        let data = {};
        data.login = inputs[0].value;
        data.email = inputs[1].value;
        data.password = inputs[2].value;
        data.role = inputs[3].value;
        data.name = inputs[4].value;


        let xhr = new XMLHttpRequest();
        xhr.open("POST", "/user/create");
        xhr.onload = function (e) {
            let response = JSON.parse(e.currentTarget.response);
            if ("Error" in response) {
                if (response.Error == null) {
                    alert("Користувач успішно зареєстрований");
                    document.location.href = "/"
                } else {
                    console.log(response.Error);
                }
            } else {
                    alert("Введено некоректні дані");
            }
        };
        xhr.send(JSON.stringify(data));
    }
}