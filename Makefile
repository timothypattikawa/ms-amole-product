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
      --go_out=paths=import:./api/grpc/generated \
      --go-grpc_out=paths=import:./api/grpc/generated \
      --proto_path=./api/grpc/protos/ \
      ./api/grpc/protos/*.proto
.PHONY: createdb dropdb migrate migrateup migratedown proto