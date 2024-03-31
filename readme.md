

protoc --go_out=../gen/go --go_opt=paths=source_relative --go-grpc_out=../gen/go --go-grpc_opt=paths=source_relative orders/orders.proto


migrate -database postgres://postgres:secret@localhost:5433/postgres?sslmode=disable -path migrations up/down

migrate migrate -database postgres://postgres:secret@localhost:5433/postgres?sslmode=disable -path migrations force V

migrate create -ext sql -dir migrations -seq create_user_table

docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres