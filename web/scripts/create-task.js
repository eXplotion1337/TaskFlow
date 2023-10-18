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
    // window.location.href = "./Projects";
});

document.getElementById("createTaskForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Предотвратить стандартное действие отправки формы

    // Получите значение NameProjects из формы
    let projectName = document.getElementById("NameProjects").value;
    let accessToken = localStorage.getItem('token');
    // Отправьте POST-запрос для проверки наличия проекта
    fetch(`/api/checkProdject/${projectName}`, {
        method: "POST",
        headers: {
            'Authorization': 'Bearer ' + accessToken

        },
    })
        .then(function(response) {
            if (response.status === 200) {
                // Если проект существует, создайте задачу
                createTask();
                let accessToken = localStorage.getItem('token');

                if (accessToken) {
                    makeRequest("./dashboard", accessToken); // Выполняем переход с токеном в заголовке Authorization
                    // window.location.href = "./dashboard";
                } else {
                    // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
                    window.location.href = "./sing-in"; // Измени в соответствии с реальным путем
                }
            } else {
                // Если проект не существует, предупредите пользователя
                alert("Проект не найден. Пожалуйста, проверьте название проекта.");
            }
        })
        .catch(function(error) {
            console.error("Произошла ошибка при проверке проекта:", error);
        });
});

// Функция для создания задачи
function createTask() {
    // Соберите данные из формы в объект JSON
    let formData = {
        taskTitle: document.getElementById("taskTitle").value,
        taskDescription: document.getElementById("taskDescription").value,
        TimeStart: document.getElementById("TimeStart").value,
        TimeEnd: document.getElementById("TimeEnd").value,
        NameProjects: document.getElementById("NameProjects").value,
        taskAssignee: document.getElementById("taskAssignee").value

    };

    let accessToken = localStorage.getItem('token');

    // Отправьте POST-запрос на сервер для создания задачи
    fetch("./createNewTask", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": 'Bearer ' + accessToken
        },
        body: JSON.stringify(formData)
    })
        .then(function(response) {
            if (response.status === 201) {
                window.location.href = './dashboard';
            } else {
                // Обработайте другие статусы, если необходимо
            }
            // Обработайте ответ от сервера, если необходимо
        })
        .catch(function(error) {
            // Обработайте ошибку, если она произошла
            console.error("Произошла ошибка при отправке запроса:", error);
        });
}