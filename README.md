   # Assignmnet Backend : 
   ![c02a76c4-ec77-4426-bb61-0ea8d98a3ebe](https://user-images.githubusercontent.com/116354441/231779799-9d7ab8da-5e58-4f4a-9c2e-2a63e4f71c26.jpg)

 # Output Samples : 
## 1) Condition  : JSON Formate Data 
## Methode : POST 
   -Input
   1) Name 
   2) Email
   3) DOB
   4) Mobile
 
 ![Screenshot (73)](https://user-images.githubusercontent.com/116354441/231775509-2691b27d-b8c4-4fa6-8b19-b0fec6a6c934.png)
 ## 2) Condition : Code : 200 , record saved sucessfully message with entered value name Parameter
 ![Screenshot (71)](https://user-images.githubusercontent.com/116354441/231776515-32ffa4d3-a4c6-4669-bb67-eccefbc262ff.png)
 ## 3) Condition : Optional Step : Save the data to the database MySQL
![Screenshot (74)](https://user-images.githubusercontent.com/116354441/231775943-01a65a8c-96d0-4970-b965-9a3f35bc7018.png)


 


# Rsense_Backend_Assignment
Before we begin, make sure that you have Go installed on your machine and the Go-Zero framework installed.

Step 1: Install Go-Zero Framework
To get started with the Go-Zero framework, you need to install it first. You can use the following command to install the framework:
# go get -u github.com/tal-tech/go-zero

Step 2: Create a New Project
Create a new directory for your project and initialize it as a Go module:
mkdir myproject
cd myproject
go mod init myproject

Step 2: Install Dependencies
We need to install the following dependencies:

Go-Zero Framework
MySQL driver for Go
To install them, run the following commands:


# Copy code
go get -u github.com/tal-tech/go-zero
go get -u github.com/go-sql-driver/mysql

# Step 3: Create a Database
Create a new MySQL database and table to store the user data. You can use any MySQL client like MySQL Workbench to create the database.

CREATE DATABASE mydb;
USE mydb;
CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    dob DATE NOT NULL,
    mobile VARCHAR(20) NOT NULL,
    PRIMARY KEY (id)
);

# Step 4: Create a Go-Zero API Project
Create a new Go-Zero API project using the following command:
type (
    Request struct {
        Name   string `path:"name,options=Tom|Jerry"`
        Email  string `json:"email"`
        DOB    string `json:"dob"`
        Mobile string `json:"mobile"`
    }

    Response struct {
        Code    int    `json:"code"`
        Message string `json:"message"`
    }

    Service interface {
        AddUser(ctx context.Context, req Request) (*Response, error)
    }
    )

// add user
@post /user/add
func (svc *Service) AddUser(ctx context.Context, req Request) (*Response, error)

# Step 6: Generate the API Code
Generate the Go code for the API using the following command:


Copy code
goctl api generate -api myapi.api -dir .
This will generate the necessary Go code for the API in the current directory.

# Step 7: Implement the Service
Open the internal/service/user.go file and implement the AddUser function as follows:


Copy code
func (s *Service) AddUser(ctx context.Context, req types.Request) (*types.Response, error) {
    _, err := s.model.Insert(&model.User{
        Name:   req.Name,
        Email:  req.Email,
        DOB:    req.DOB,
        Mobile: req.Mobile,
    })
    if err != nil {
        return nil, err
    }

    return &types.Response{
        Code:    200,
        Message: fmt.Sprintf("record saved successfully for %s", req.Name),
    }, nil

# Step 8: Configure the Database Connection
Open the etc/user-api.yaml file and configure the database connection as follows:


Copy code
DataSource:
  DriverName: mysql
  DataSource: root:password@tcp(127.0.0.1:3306)/mydb
  MaxIdle: 50
  MaxOpen: 150
  MaxLifetime: 600
  MaxConnLifetime: 600
Replace root and password with your MySQL username and password respectively.

# Step 9: Run the API
Run the API using the following command:


Copy code
# go run user.go -
