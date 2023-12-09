gen-swag:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go

mg-up:
	migrate -path migrations/sql_model/ -database "postgresql://nodirbek:0021@localhost:5432/easy_travel" -verbose up


mg-down:
	migrate -path migrations/sql_model/ -database "postgresql://nodirbek:0021@localhost:5432/easy_travel" -verbose down