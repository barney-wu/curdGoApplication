# curdGoApplication

This is a simple CURD application written in Golang.


Create a local database with mysql.

  CREATE DATABASE TST;

  USE DATABASE TST;

  CREATE TABLE users (

    user_id INT(11) AUTO_INCREMENT PRIMARY KEY,

    username VARCHAR(40),

    password VARCHAR(255),

    email VARCHAR(255),
  )


To build it

go run main.go
