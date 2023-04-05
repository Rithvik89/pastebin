dev:
	go run http/*.go

remove_mysql:
	docker rmi -f mysql_pb

create_mysql:
	docker run --name mysql_pb -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=mypassword -v mysql:/var/lib/mysql mysql:8