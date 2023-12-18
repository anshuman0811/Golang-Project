GoFr Project

Open the terminal in project location and run the following command

go mod init github.com/example
go get gofr.dev


run the command in terminal

go mod tidy



command for connection of mySQL run the following commands

docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3307:3306 -d mysql:8.0.30
docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE Student_details (id INT AUTO_INCREMENT PRIMARY KEY NOT NULL, name VARCHAR(255) NOT NULL ,age INT ,class INT ,teacher_id INT NOT NULL ,fees INT NOT NULL ,address VARCHAR(255) NOT NULL ,mobile_no VARCHAR(255) NOT NULL);"
docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE Teacher_details (teach_id INT PRIMARY KEY, teach_name VARCHAR(255) NOT NULL ,teach_age INT ,teach_salary INT NOT NULL,teach_num VARCHAR(255) NOT NULL ,teach_add VARCHAR(255) NOT NULL);"
go run main.go 



to access data

localhost:9000/students



to insert data

localhost:9000/students/Anshuman/Singh/22/9/M/Near tilak nagar, bikaner/828845



to update data

localhost:9000/2/students/Anshuman/Singh/22/9/M/Near tilak nagar, bikaner/828845



to delete data

localhost:9000/students/1
