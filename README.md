# curdGoApplication

This is a simple CURD application written in Golang.


Create a local database with mysql.

    CREATE DATABASE TST;

    USE DATABASE TST;

    CREATE TABLE users (

      user_id INT(11) AUTO_INCREMENT PRIMARY KEY,

      username VARCHAR(40),

      password VARCHAR(255),

      email VARCHAR(255)
    );
    
   Create test cases: 
   
    INSERT INTO users (username, password, email) VALUES ('TOM', 'TOM123', 'TOM123@G.COM');
   
    UPDATE users SET username = 'Jerry', password = 'Jerry456', email = 'Jerry456@gg.com' WHERE user_id % 5 = 1;

    UPDATE users SET username = 'Bill', password = 'gates456', email = 'micro@m.com' WHERE user_id % 5 = 2;

    UPDATE users SET username = 'Jeff', password = 'Bezos6', email = 'aws@a.com' WHERE user_id % 5 = 3;

    UPDATE users SET username = 'JACK', password = 'alibaba', email = '955@alibaba.com' WHERE user_id % 5 = 4;


To build it

  go run main.go
  
 Visit
    localhost:8080
  
