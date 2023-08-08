DB_URL=postgresql://root:secret@localhost:5432/test_db?sslmode=disable

.PHONY: docker-network postgres createdb dropdb cleanup test-suite

docker-network:
	docker network create test-network

# Start a new PostgreSQL container named "postgres14" running on network "test-network"
# Expose the PostgreSQL service at port 5432 on the host machine
# Set the PostgreSQL username to "root" and password to "secret"
postgres:
	docker run --name postgres14 --network test-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

# Execute the `createdb` command inside the "postgres14" container
# The `--username=root --owner=root` options specify the PostgreSQL username and the owner of the new database
# "test_db" is the name of the database to be created
createdb:
	sleep 5
	docker exec -it postgres14 createdb --username=root --owner=root test_db 

# Execute the `dropdb` command inside the "postgres14" container
# "test_db" is the name of the database to be dropped
dropdb:
	docker exec -it postgres14 dropdb test_db

cleanup:
	docker-compose down --volumes
# Stop and remove the PostgreSQL container
	docker stop postgres14 || true
	docker rm postgres14 || true
	docker volume rm postgres14-data || true
	
# Remove the test network
	docker network rm test-network || true
	
# Remove all unused Docker objects (images, containers, volumes, networks)
# docker system prune --volumes -a -f

migrateup:
	migrate -path internal/db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path internal/db/migrations -database "postgresql://root:secret@localhost:5432/test_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

test-suite: 
	@make docker-network ;\
	make postgres ;\
	make createdb ;\
	make migrateup ;\
	trap 'make cleanup' ERR ;\
	make test ;\
	make cleanup

dev:
	docker-compose build --no-cache && docker-compose up
