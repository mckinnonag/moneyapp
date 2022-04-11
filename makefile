postgres:
	docker run --name moneydb -p 5001:5001 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it moneydb createdb --username=root --owner=root moneyapp

dropdb:
	docker exec -it moneydb dropdb moneyapp

migrateup:
	migrate -path db/migration -database postgresql://root:secret@localhost:5001/moneydb -verbose up

migratedown:
	migrate -path db/migration -database "postgressql://root:secret@localhost:5001/moneydb?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown