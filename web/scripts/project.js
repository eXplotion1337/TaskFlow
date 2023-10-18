document.addEventListener("DOMContentLoaded", function () {
    // Создаем объект XMLHttpRequest
    let xhr = new XMLHttpRequest();
    let accessToken = localStorage.getItem('token');
    // Устанавливаем метод и URL для запроса
    xhr.open("POST", "./api/allprojects", true);

    xhr.setRequestHeader('Authorization', 'Bearer ' + accessToken);

    // Устанавливаем обработчик события при успешном завершении запроса
    xhr.onload = function () {
        if (xhr.status === 200) {
            // Попытаемся преобразовать ответ в объект JSON
            try {
                let projects = JSON.parse(xhr.responseText);
                // Здесь projects представляет массив проектов
                console.log("Успешный запрос!");
                console.log(projects);

                // Теперь вы можете использовать данные для заполнения вашей страницы
                // Например, добавить проекты в контейнер на вашей странице
                // Получаем кнопку "Создать новый проект"
                let createProjectBtn = document.getElementById("createProjectBtn");

                // Добавляем обработчик события клика
                createProjectBtn.addEventListener("click", function () {
                    let accessToken = localStorage.getItem('token');

                    if (accessToken) {
                        makeRequest("./createNewProject", accessToken); // Выполняем переход с токеном в заголовке Authorization
                        // window.location.href = "./dashboard";
                    } else {
                        // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
                        window.location.href = "./sing-in"; // Измени в соответствии с реальным путем
                    }

                    // Выполняем редирект на "./createNewProject"
                    // window.location.href = "./createNewProject";
                });


                // Пример:
                let projectsContainer = document.getElementById("projectsContainer");



                projects.forEach(function (project) {
                    let projectBox = document.createElement("div");
                    projectBox.classList.add("col-md-3", "mb-4");
                    projectBox.innerHTML = `
                                <div class="project-box">
                                    <h2>${project.NameProject}</h2>
                                    <p>${project.Description}</p>
                                    <p>Участники: ${project.Collaborators}</p>
                                </div>
                            `;

                    projectBox.addEventListener("click", function () {
                        let accessToken = localStorage.getItem('token');
                        let url_to_redirect = `./dashboard/${project.NameProject}`
                        if (accessToken) {
                            makeRequest(url_to_redirect, accessToken); // Выполняем переход с токеном в заголовке Authorization
                            // window.location.href = "./dashboard";
                        } else {
                            // Если токен отсутствует, можешь перенаправить пользователя на страницу входа
                            window.location.href = "./sing-in"; // Измени в соответствии с реальным путем
                        }


                        // window.location.href = `./dashboard/${project.NameProject}`;
                    });

                    projectsContainer.appendChild(projectBox); // добавляем проект
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