go get -u github.com/go-sql-driver/mysql

su

 docker pull mysql:5.7

!!!first use mysql docker , goto 1 , or goto 2
1.
    docker run -p 3306:3306 --name mysql2 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7

2. 
    docker ps -a
    docker restart ******* 

go run main.go
