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

document.querySelector(".buttons-container .btn-primary").addEventListener("click", function(event) {
    event.preventDefault();
    window.location.href = "./sing-up";
});

// Добавляем действие для кнопки "Вход"
document.querySelector(".buttons-container .btn-success").addEventListener("click", function(event) {
    event.preventDefault();
    window.location.href = "./sing-in";
});