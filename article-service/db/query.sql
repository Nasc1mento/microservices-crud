-- name: CreateArticle :one
INSERT INTO articles (
    title,
    content,
    author_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetArticleById :one
SELECT * FROM articles 
WHERE id = $1 LIMIT 1;

-- name: ListArticles :many
SELECT * FROM articles
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateArticle :exec
UPDATE articles SET
    title = $1,
    content = $2
WHERE id = $3
RETURNING *;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1;