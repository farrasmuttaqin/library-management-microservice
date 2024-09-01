build:
	go build -o main .

run:
	./main

docker:
	docker build -t microservice .
	docker run -p 8080:8080 microservice
