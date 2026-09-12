-- name: InsertParticipant :exec
INSERT INTO participants (id, nim, name, class_id)
VALUES ($1, $2, $3, $4);

-- name: FindParticipantByID :one
SELECT *
FROM participants
WHERE id = $1;

-- name: FindParticipantsByClassID :many
WITH sort_param AS (
    SELECT COALESCE(sqlc.narg('sort_by'), 'nim')::varchar AS sort_by
)
SELECT
    p.id,
    p.nim,
    p.name,
    p.class_id,
    p.created_at,
    p.updated_at,
    tp.team_id,
    t.number AS team_number
FROM participants p
LEFT JOIN team_participants tp ON tp.participant_id = p.id
LEFT JOIN teams t ON tp.team_id = t.id
CROSS JOIN sort_param sp
WHERE p.class_id = $1
ORDER BY
    CASE WHEN sp.sort_by = 'nim' THEN p.nim END NULLS LAST,
    CASE WHEN sp.sort_by = 'team' THEN team_number END NULLS LAST;

-- name: FindParticipantsByTeamID :many
SELECT
    p.id,
    p.nim,
    p.name,
    p.class_id,
    p.created_at,
    p.updated_at,
    tp.team_id,
    t.number AS team_number
FROM participants p
LEFT JOIN team_participants tp ON tp.participant_id = p.id
LEFT JOIN teams t ON tp.team_id = t.id
WHERE p.class_id = $1
ORDER BY p.nim;

-- name: SetParticipantTeam :exec
INSERT INTO team_participants (team_id, participant_id)
VALUES ($1, $2);

-- name: ResetParticipantTeam :exec
DELETE FROM team_participants
WHERE
    team_id = $1 AND
    participant_id = $2;
