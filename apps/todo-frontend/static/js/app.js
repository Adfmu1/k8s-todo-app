const apiLink = "/api"

function getCurrentImageIndex() {
    const now = new Date();
    return now.getHours() % 10;
}

function updateImage() {
    const imgElement = document.getElementById("image");
    const i = getCurrentImageIndex();
    imgElement.src = `/cache/images/image${i}.jpg`;
}

async function getTodos() {
    const link = apiLink + "/todos"
    const todosContainer = document.getElementById('todos');

    // delete all children
    while (todosContainer.firstChild) todosContainer.removeChild(todosContainer.firstChild);
    
    // get requests from remote
    try {
        const response = await fetch(link);
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }

        const result = await response.json();

        if (result == null) {
            console.error("No values retrieved from backend");
        }

        result.forEach(element => {
            const todoDiv = document.createElement('div');
            todoDiv.classList.add('todo');

            const p = document.createElement('p');
            p.textContent = element;

            todoDiv.appendChild(p);
            todosContainer.appendChild(todoDiv);
        });

    } catch (error) {
        console.error(error.message);
    }
}

getTodos()
updateImage();

setInterval(updateImage, 60 * 1000);

document.getElementById('add_todo_btn').addEventListener('click', addTodo);

document.getElementById('todo_input').addEventListener('keydown', function (e) {
    if (e.key === 'Enter') {
        addTodo();
    }
});

async function addTodo() {
    const input = document.getElementById('todo_input');
    const text = input.value.trim();
    const link = apiLink + "/todos"

    if (text === '') {
        return;
    }

    response = await fetch(link, {
        method: "POST",
        body: JSON.stringify({ newTodo: text })
    })

    getTodos()

    input.value = '';
    input.focus();
}