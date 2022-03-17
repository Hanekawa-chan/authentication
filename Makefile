db_build:
	sudo docker-compose up --build -d mongodb
db_run:
	sudo docker-compose up
go_run:
	go run main.go
run:
	db_build db_run go_run
