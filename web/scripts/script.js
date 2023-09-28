document.addEventListener('DOMContentLoaded', function () {
    // Инициализация Sortable.js для каждой колонки
    const columnFree = new Sortable(document.getElementById('column-free'), {
        group: 'task-list',
        animation: 150,
        onEnd: function (evt) {
            handleTaskMove(evt);
        }
    });

    const columnInProgress = new Sortable(document.getElementById('column-in-progress'), {
        group: 'task-list',
        animation: 150,
        onEnd: function (evt) {
            handleTaskMove(evt);
        }
    });

    const columnInReview = new Sortable(document.getElementById('column-in-review'), {
        group: 'task-list',
        animation: 150,
        onEnd: function (evt) {
            handleTaskMove(evt);
        }
    });

    const columnCompleted = new Sortable(document.getElementById('column-completed'), {
        group: 'task-list',
        animation: 150,
        onEnd: function (evt) {
            handleTaskMove(evt);
        }
    });

    // Обработчик события при отправке формы
    const taskForm = document.getElementById('task-form');
    taskForm.addEventListener('submit', function (e) {
        e.preventDefault();

        // Получаем значения из формы
        const title = document.getElementById('task-title').value;
        const description = document.getElementById('task-description').value;

        // Создаем новую задачу
        const taskItem = document.createElement('li');
        taskItem.className = 'list-group-item'; // Базовый класс для элемента списка
        taskItem.innerHTML = `
            <strong>${title}</strong>
            <p>${description}</p>
        `;

        // Добавляем задачу в колонку "Свободные"
        columnFree.appendChild(taskItem);

        // Очищаем форму
        taskForm.reset();
    });

    // Обработчик перемещения задачи между колонками
    function handleTaskMove(evt) {
        const item = evt.item;
        const fromColumn = evt.from.id;
        const toColumn = evt.to.id;

        if (fromColumn !== toColumn) {
            // Если задача перемещается между колонками, измените статус задачи соответственно
            if (toColumn === 'column-in-progress') {
                // Задача перемещается в "В работе"
                item.classList.add('list-group-item-primary');
            } else if (toColumn === 'column-in-review') {
                // Задача перемещается в "В ревью"
                item.classList.add('list-group-item-warning');
            } else if (toColumn === 'column-completed') {
                // Задача перемещается в "Завершенные"
                item.classList.add('list-group-item-success');
            } else if (toColumn === 'column-free') {
                // Задача перемещается обратно в "Свободные"
                item.classList.remove('list-group-item-primary', 'list-group-item-warning', 'list-group-item-success');
            }
        }
    }
});

