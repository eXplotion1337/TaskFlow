document.getElementById("dashboardLink").addEventListener("click", function(event) {
    event.preventDefault();
    let accessToken = localStorage.getItem('token');

    if (accessToken) {
        makeRequest("./dashboard", accessToken); // Выполняем переход с токеном в заголовке Authorization
        // window.location.href = "./dashboard";
    } else {
        // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
        window.location.href = "./sing-in"; // Измени в соответствии с реальным путем
    }
});

document.getElementById("createNewTask").addEventListener("click", function(event) {
    event.preventDefault();
    let accessToken = localStorage.getItem('token');

    if (accessToken) {
        makeRequest("./createNewTask", accessToken); // Выполняем переход с токеном в заголовке Authorization
        // window.location.href = "./dashboard";
    } else {
        // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
        window.location.href = "./sing-in"; // Измени в соответствии с реальным путем
    }
    // window.location.href = "./createNewTask";
});

document.getElementById("TaskFlowLink").addEventListener("click", function(event) {
    event.preventDefault();
    let accessToken = localStorage.getItem('token');

    if (accessToken) {
        makeRequest("./", accessToken); // Выполняем переход с токеном в заголовке Authorization
        // window.location.href = "./dashboard";
    } else {
        // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
        window.location.href = "./sing-in"; // Измени в соответствии с реальным путем
    }
    // window.location.href = "./";
});

document.getElementById("Projects").addEventListener("click", function(event) {
    event.preventDefault();
    let accessToken = localStorage.getItem('token');

    if (accessToken) {
        makeRequest("./Projects", accessToken); // Выполняем переход с токеном в заголовке Authorization
        // window.location.href = "./dashboard";
    } else {
        // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
        window.location.href = "./sing-in"; // Измени в соответствии с реальным путем
    }
    window.location.href = "./Projects";
});

document.getElementById("loginButton").addEventListener("click", function() {
    // Получаем значения полей ввода
    let username = document.getElementById("username").value
    let password = document.getElementById("password").value

    // Создаем объект с данными для отправки на сервер в формате JSON
    let data = {
        Username: username,
        Password: password
    };

    console.log(data);
    let accessToken = localStorage.getItem('token');
    // Отправляем POST-запрос на сервер
    fetch('/auth/sing-in', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + accessToken
        },

        body: JSON.stringify(data)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            // Получаем токен из ответа сервера
            console.log(data);
            let token = data.token;
            //
            // // Теперь можно сохранить токен, например, в localStorage
            localStorage.setItem('token', token);
            //
            // // Перенаправляем пользователя на нужную страницу
            let accessToken = localStorage.getItem('token');

            if (accessToken) {
                makeRequest("./dashboard", accessToken); // Выполняем переход с токеном в заголовке Authorization
                // window.location.href = "./dashboard";
            } else {
                // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
                window.location.href = "./sing-in"; // Измени в соответствии с реальным путем
            }
            // window.location.href = '/dashboard'; // Измени в соответствии с реальным путем
        })
        .catch(error => {
            console.error('Ошибка:', error);
            // Обработка ошибок, например, показ сообщения об ошибке пользователю
        });
});

function makeRequest(url, token) {
    let xhr = new XMLHttpRequest();

    xhr.open('GET', url, true);
    // Устанавливаем заголовок Authorization
    xhr.setRequestHeader('Authorization', 'Bearer ' + token);

    xhr.setRequestHeader('Cache-Control', 'no-cache');

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                // Обработка успешного ответа
                console.log(xhr.responseText);

                // Редирект на другую страницу
                window.location.href = url;
            } else {
                // Обработка ошибки
                console.error('Ошибка запроса:', xhr.status, xhr.statusText);
            }
        }
    }
    // Отправляем запрос
    // xhr.open('GET', url, true);
    xhr.send();
}