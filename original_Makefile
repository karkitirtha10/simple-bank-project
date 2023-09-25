step ?= 1

server:
	go run cmd/main.go

migration:
	go run cmd/main.go -migrate_flag=create -migration_name=$(name)
	
# migrate create -ext sql -dir pkg/common/db/migration -seq $(name)

migrate-up:
	go run cmd/main.go -migrate_flag=up -migrate_step=$(step)

migrate-down:
	go run cmd/main.go -migrate_flag=down -migrate_step=$(step)

# migrate -path pkg/common/db/migration -database "postgresql://pgsuperuser:Admin@1@localhost:5432/gotodo?sslmode=disable" -verbose up 	

docker-up:
	docker-compose up -d

.PHONY: server migration migrate-up migrate-down docker-up