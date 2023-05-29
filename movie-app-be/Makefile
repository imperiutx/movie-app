run:
	go run ./cmd/api

startdb:
	docker run --name db_movies -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.3-alpine

createdb:
	docker exec -it db_movies createdb --username=root --owner=root db_movies