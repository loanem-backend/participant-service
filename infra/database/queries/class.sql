-- name: InsertClass :one
INSERT INTO classes (name, course_id)
VALUES ($1, $2)
RETURNING id;

-- name: DeleteClass :exec
DELETE FROM classes
WHERE id = $1;

-- name: FindClassesByCourseID :many
SELECT
    c.id,
    c.name,
    c.course_id,
    rc.name AS course_name,
    rc.year AS course_year,
    c.created_at,
    c.updated_at
FROM classes c
INNER JOIN repl_courses rc ON c.course_id = rc.id
WHERE c.course_id = $1
ORDER BY c.name;
