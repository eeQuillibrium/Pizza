.SILENT:
services := runapi rundelivery
runapi:
	go run pizza-api/cmd/main.go
rundelivery:
	go run pizza-delivery/cmd/main.go
run: $(services)
	go run pizza-kitchen/cmd/main.go
	