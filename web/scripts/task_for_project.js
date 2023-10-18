document.addEventListener("DOMContentLoaded", function () {
    // Ваш существующий код для получения задач

    // Предположим, что у вас есть переменная projectName, содержащая название проекта
    // После успешного выполнения запроса задач, обновляем заголовок
    let projectsHeading = document.querySelector('.projects-heading');
    const id = window.location.pathname.split('/').pop();
    projectsHeading.textContent = `Задачи проекта "` + id + `"`;

});
document.getElementById("dashboardLink").addEventListener("click", function(event) {
    event.preventDefault();
    let accessToken = localStorage.getItem('token');

    if (accessToken) {
        makeRequest("../dashboard", accessToken); // Выполняем переход с токеном в заголовке Authorization
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
        makeRequest("../createNewTask", accessToken); // Выполняем переход с токеном в заголовке Authorization
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
        makeRequest("../", accessToken); // Выполняем переход с токеном в заголовке Authorization
        // window.location.href = "./dashboard";
    } else {
        // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
        window.location.href = "../sing-in"; // Измени в соответствии с реальным путем
    }
    // window.location.href = "./";
});

document.getElementById("Projects").addEventListener("click", function(event) {
    event.preventDefault();
    let accessToken = localStorage.getItem('token');

    if (accessToken) {
        makeRequest("../Projects", accessToken); // Выполняем переход с токеном в заголовке Authorization
        // window.location.href = "./dashboard";
    } else {
        // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
        window.location.href = "../sing-in"; // Измени в соответствии с реальным путем
    }
    // window.location.href = "./Projects";
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
document.addEventListener("DOMContentLoaded", function () {

    let createProjectBtn = document.getElementById("createProjectBtn");

    // Добавляем обработчик события клика
    createProjectBtn.addEventListener("click", function () {
        let accessToken = localStorage.getItem('token');

        if (accessToken) {
            makeRequest("../createNewTask", accessToken); // Выполняем переход с токеном в заголовке Authorization
            // window.location.href = "./dashboard";
        } else {
            // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
            window.location.href = "./sing-in"; // Измени в соответствии с реальным путем
        }
        // Выполняем редирект на "./createNewProject"
        // window.location.href = "../createNewTask";
    });


    // Создаем объект XMLHttpRequest
    let xhr = new XMLHttpRequest();

    const urlParams = new URLSearchParams(window.location.search);
    let accessToken = localStorage.getItem('token');
    // const id = urlParams.get("id");
    const id = window.location.pathname.split('/').pop();
    const decodedString = decodeURIComponent(id);
    // Устанавливаем метод и URL для запроса
    xhr.open("POST", `../api/dashboard/tasks/${id}`, true);

    xhr.setRequestHeader('Authorization', 'Bearer ' + accessToken);

    // Устанавливаем обработчик события при успешном завершении запроса
    xhr.onload = function () {
        if (xhr.status === 200) {
            // Попытаемся преобразовать ответ в объект JSON
            try {
                let tasks = JSON.parse(xhr.responseText);
                // Здесь tasks представляет массив задач
                console.log("Успешный запрос!");
                console.log(tasks);

                // Теперь вы можете использовать данные для заполнения вашей страницы
                // Например, добавить задачи в соответствующие списки на вашей странице

                // Пример:
                let freeTasksList = document.getElementById("column-free");
                let inProgressTasksList = document.getElementById("column-in-progress");
                let inReviewTasksList = document.getElementById("column-in-review");
                let completedTasksList = document.getElementById("column-completed");

                tasks.forEach(function (task) {
                    let listItem = document.createElement("li");
                    listItem.textContent = task.taskTitle;

                    // В зависимости от значения task.Colum, добавьте задачу в соответствующий список
                    switch (task.Colum) {
                        case "Свободные":
                            freeTasksList.appendChild(listItem);
                            break;
                        case "В работе":
                            inProgressTasksList.appendChild(listItem);
                            break;
                        case "В ревью":
                            inReviewTasksList.appendChild(listItem);
                            break;
                        case "Завершенные":
                            completedTasksList.appendChild(listItem);
                            break;
                        default:
                            // Обработка неизвестных значений Colum
                            break;
                    }
                });
            } catch (e) {
                console.error("Ошибка при разборе JSON: " + e);
            }
        } else {
            // Обработка ошибок
            console.error("Произошла ошибка: " + xhr.status);
        }
    };
    xhr.send();
});