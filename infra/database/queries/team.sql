-- name: InsertTeam :exec
INSERT INTO teams (id, number, class_id)
VALUES ($1, $2, $3);

-- name: UpdateTeamNumber :exec
UPDATE teams
SET
    number = $1,
    updated_at = $2
WHERE id = $3;

-- name: DeleteTeam :exec
DELETE FROM teams
WHERE id = $1;

-- name: FindTeamsByClassID :many
SELECT
    t.id,
    t.number,
    t.class_id,
    c.name AS class_name,
    c.course_id,
    rc.name AS course_name,
    rc.year AS course_year,
    t.created_at,
    t.updated_at
FROM teams t
INNER JOIN classes c ON t.class_id = c.id
INNER JOIN repl_courses rc ON c.course_id = rc.id
WHERE t.class_id = $1
ORDER BY t.number;

-- name: FindTeamsByCourseID :many
SELECT
    t.id,
    t.number,
    t.class_id,
    c.name AS class_name,
    c.course_id,
    rc.name AS course_name,
    rc.year AS course_year,
    t.created_at,
    t.updated_at
FROM teams t
INNER JOIN classes c ON t.class_id = c.id
INNER JOIN repl_courses rc ON c.course_id = rc.id
WHERE c.course_id = $1
ORDER BY c.name, t.number;
