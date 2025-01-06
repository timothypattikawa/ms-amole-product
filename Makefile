postgres:
	docker run -d --name postgres-amole -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -v ~/docker/postgres-amole:/var/lib/postgresql/data -p 5432:5432 postgres:17-alpine
createdb:
	docker exec -it postgres-amole createdb --username=root --owner=root amole_db
dropdb:
	docker exec -it postgres-amole dropdb --username=root --owner=root amole_db
migrate:
	migrate create -ext sql -dir ./script/migrations -seq tb_amole_product
migrateup:
	migrate -path script/migrations -database "postgresql://root:secret@localhost:5432/amole_db?sslmode=disable" -verbose up
migratedown:
	migrate -path script/migrations -database "postgresql://root:secret@localhost:5432/amole_db?sslmode=disable" -verbose down
proto:
	protoc \
      --go_out=paths=import:./api/grpc/protos/ \
      --go-grpc_out=paths=import:./api/grpc/protos/ \
      --proto_path=./api/grpc/protos/ \
      ./api/grpc/protos/*.proto
redis-product:
	docker run -d \
		--name redis-amole \
		-v ~/docker/redis-amole:/data \
		-p 6379:6379 \
		redis:7.4-alpine
		
.PHONY: createdb dropdb migrate migrateup migratedown proto