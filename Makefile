build:
	go build -o payment_service main.go

run:
	go run main.go -config config.yaml

docker-build:
	docker-compose build

docker-up:
	docker-compose up
