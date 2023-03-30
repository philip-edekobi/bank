postgres:
	docker run --name bankdb --network bank-network -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=168912 -d postgres:14

createdb:
	docker exec -it bankdb createdb --username=root --owner=root bank

dropdb:
	docker exec -it bankdb dropdb bank

migrateup:
	migrate -path db/migration -database "postgres://root:168912@localhost:5430/bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgres://root:168912@localhost:5430/bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://root:168912@localhost:5430/bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgres://root:168912@localhost:5430/bank?sslmode=disable" -verbose down 1

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/philip-edekobi/bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migratedown1 migrateup1
