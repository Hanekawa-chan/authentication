db_build:
	sudo docker-compose up --build -d mongodb
db_stop:
	sudo docker-compose down
go_run:
	go run main.go
run: db_build go_run
