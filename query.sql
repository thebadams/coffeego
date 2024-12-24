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
SELECT name from coffees
ORDER BY name;
