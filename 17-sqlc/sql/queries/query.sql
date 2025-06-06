-- name: ListCategories :many
SELECT * FROM category;

-- name: GetCategory :one
SELECT * FROM category WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO category (name, description) VALUES (?, ?);

-- name: UpdateCategory :exec
UPDATE category SET name = ?, description = ? WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM category WHERE id = ?;

-- name: CreateCourse :exec
INSERT INTO course (name, description, category_id) VALUES (?, ?, ?);