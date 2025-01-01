-- name: CreateCoffee :one
INSERT INTO coffees(
name, roaster_id

) VALUES (
?, ?
	)
RETURNING *;


-- name: CreateRoaster :one
INSERT INTO roasters(
name
) VALUES (
?
	)
RETURNING *;


-- name: ListCoffees :many
SELECT c.name as coffee, r.name as roaster from coffees as c
INNER JOIN roasters as r
ON c.roaster_id=r.id
ORDER BY c.name;


-- name: FindRoasterByName :one
SELECT * from roasters
WHERE name = ? LIMIT 1;
