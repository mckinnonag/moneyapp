postgres:
	docker run --name moneydb --rm -e POSTGRES_USER=root -e POSTGRES_PASSWORD=4y7sV96vA9wv46VR -e PGDATA=/var/lib/postgresql/data/pgdata -v /tmp:/var/lib/postgresql/data -p 5432:5432 -it postgres:alpine3.15

createdb:
	docker exec -it moneydb createdb --username=root --owner=root moneyapp

dropdb:
	docker exec -it moneydb dropdb moneyapp

migrateup:
	cmd/migrate -path db/migration -database postgresql://root:4y7sV96vA9wv46VR@localhost:5234/moneydb -verbose up


migratedown:
	migrate -path db/migration -database "postgressql://root:4y7sV96vA9wv46VR@localhost:5234/moneydb?sslmode=disable" -verbose down


.PHONY: postgres createdb dropdb migrateup migratedown
