-- name: GetTodos :many
SELECT todo FROM todos;

-- name: CreateTodo :exec
INSERT INTO todos (todo)
VALUES ($1);