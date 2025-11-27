-- name: GetContact :one
SELECT * FROM contacts
WHERE id = $1 LIMIT 1;

-- name: ListContacts :many
SELECT * FROM contacts
ORDER BY first_name;

-- name: CreateContact :one
INSERT INTO contacts (
  first_name, last_name, email, phone, address, city, state, zip, country
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateContact :exec
UPDATE contacts
  set first_name = $2, last_name = $3, email = $4, phone = $5, address = $6, city = $7, state = $8, zip = $9, country = $10
WHERE id = $1;

-- name: DeleteContact :exec
DELETE FROM contacts
WHERE id = $1;
