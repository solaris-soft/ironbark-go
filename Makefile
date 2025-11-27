migrate:
	migrate -database postgres://postgres:postgres@127.0.0.1:5432/ironbark?sslmode=disable -path schema/migrations up

generate-ui:
	go tool templ generate

generate-db:
	sqlc generate

dev:
	air