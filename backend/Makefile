include ../.devcontainer/.env
export

.PHONY: test db migration

test:
	go test ./... -v

init-db:
	psql -U postgres --host db --port 5432 -c "CREATE DATABASE sensorly;"

db:
	psql -U postgres --host db --port 5432 -d sensorly

migration:
	migrate create -seq -ext=.sql -dir=./migrations $(name)

migrate:
	migrate -path=./migrations -database="$(DB_DSN)?sslmode=disable" up

db-cmd:
	migrate -path=./migrations -database="$(DB_DSN)?sslmode=disable" $(cmd)

run-sims:
	go run ./tools/sims/sensor_sims.go -host=$(HOST) -token=$(token)