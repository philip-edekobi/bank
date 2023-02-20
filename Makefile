postgres:
	docker run --name bankdb -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=168912 -d postgres:14

createdb:
	docker exec -it bankdb createdb --username=root --owner=root bank

dropdb:
	docker exec -it bankdb dropdb bank

migrateup:
	migrate -path db/migration -database "postgres://root:168912@localhost:5430/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:168912@localhost:5430/bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server