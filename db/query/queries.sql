-- name: CreateProject :one
INSERT INTO projects (name, description, start_date, end_date, active)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetProjects :many
SELECT * FROM projects;

-- name: GetProjectById :one
SELECT * FROM projects
WHERE id = $1 LIMIT 1;

-- name: UpdateProjectById :one
UPDATE  projects
SET active = $2
WHERE id = $1 RETURNING *;

-- name: DeleteProjectById :exec
DELETE FROM projects
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email ,password)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetEmail :one
SELECT email, password FROM users
WHERE email = $1 LIMIT 1;
