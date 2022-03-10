## crud-api-go: simple CRUD operation using Mux, GoLang, Postgres 

* to use Postgres cmd commands to handle data

    ```bash
  # to create database college
    CREATE DATABASE college;
    
  # to select database college  
    \connect college;
  
  # to create table students
    CREATE TABLE students (
    id SERIAL,
    name varchar(50) NOT NULL,
    department varchar(50) NOT NULL,
    address varchar(50) NOT NULL,
    PRIMARY KEY (id)
    );
  
  # to insert data into students table
    INSERT INTO students (
    name,
    department,
    address
    )
    VALUES
    ('varun', 'cse', 'delhi'),
    ('vicky ', 'it', 'mumbai'),
    ('rohan', 'electronics', 'bangalore');
    ```

* First, run these commands to get third party packages and go mod initialisation before code implementation

    ```bash
    go get github.com/gorilla/mux
    go get github.com/lib/pq
    go mod init
    ```

* import packages

    ```bash
    # for logging errors and printing messaging
    fmt and log 

    # Go core package for handling JSON data.
    encoding/json

    # Go core package for handling SQL-based database communication.
    net/http 

    # a Go HTTP package for handling HTTP requesting when creating GO APIs
    net/http 
  
    # for URL matcher and routing. It helps in implementing request routers and match each incoming request with its matching handler.
    github.com/gorilla/mux 

    # a Go PostgreSQL driver for handling the database/SQL package.
    github.com/lib/pq
    ```

* endpoints of crud-api

    ```bash
    # GET all records  
    localhost:8000/student
  
    # POST insert a record
    localhost:8000/student?name=vimal&department=cse&address=delhi
  
    # UPDATE a record
    localhost:8000/student?id=1&department=it&address=mumbai
  
    # DELETE a record 
    localhost:8000/student/1
  
   # DELETE all records 
    localhost:8000/student
  ```