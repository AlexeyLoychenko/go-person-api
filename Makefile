#-- Env Variables --#
include .env
DOCKER_COMPOSE_FILE?=docker-compose.yml

db-connect:
	docker compose -f ${DOCKER_COMPOSE_FILE} exec db psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}
db-build:
	docker compose -f ${DOCKER_COMPOSE_FILE} build db
db-run:
	docker compose up db -d

containers-down:
	@powershell $$running=docker ps -q; \
	if ($$running -eq $$null) { \
		echo "no running containers" \
	} else { \
		docker stop $$running \
	}

run:
	@go run ./cmd/server/.

#-- Migrations --#

migrate-up:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up

migrate-down:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down

migrate-create: ##`make migrate-create name=migration-name`
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migrations $(name)
migrate-force:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate force $(version)

