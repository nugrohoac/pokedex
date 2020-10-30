# Build
build:
	go build -o pokedexapp cmd/main.go

# Run app has been build
run-app:
	./pokedexapp

run:
	go run cmd/main.go

#Swagger
swag-docs:
	swag init -g cmd/main.go