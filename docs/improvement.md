# Improvements

My research focused on the backend of Weather Wizard. I wanted to implement the best possible backend in the given timeframe. In the
following I will explain the improvements I made.

## DOCUMENTATION

Im using Golang language for the backend of Weather Wizard. before talking about the implementation of the backend, I will explain the language itself.
Golang is a compiled language, which means that the code is compiled to machine code before it is executed. This makes it much faster than interpreted languages like PHP.
Golang is also a statically typed language, which means that you have to define the type of variable when you declare it. This makes it easier to find bugs in your code.
It is also garbage collected language like Java, which means that you don't have to worry about memory management like in C or C++. The garbage collector will clean up after you.
Furthermore, Golang is a concurrent language, which means that you can run multiple threads at the same time. This is very useful for web development, since you can run multiple
requests at the same time. It is language that is very easy to learn and has a lot of useful features. This makes it a very popular language for backend development and also the reason
why I wanted to learn it.

Because Golang is a very popular language for backend development, there are a lot of libraries available for backend development. I used the following libraries for this project:

These are the rather regular libraries that I used for this project, which dont need much explanation:

- encoding/json: Functions for encoding and decoding JSON data.
-  fmt: Formatted I/O functions for I/O and string-based operations.
-  log: A straightforward logging package with basic logging features.
-  net/url: Methods for modifying URL components and parsing URLs.
-  encoding/base64: Helps to convert data into base64 format and vice versa.
-  io/ioutil: Functions for basic I/O utility activities like reading and writing files.
-  os and path/filepath: Tools for working with the file system and operating system, including file manipulation.
-  Regexp: Regular expression working functions.
-  time: Time, durations, and date-related functions.


These are the libraries that are a bit more important for this project:

-  Fiber: A fast, simple, and minimalist web framework for Go. 
- Gorm: An ORM library for Golang. It helps to work with databases, specifically PostgreSQL.
- fatih/color: Makes it easy to add colors to text in the console.
- sync: Helps to coordinate tasks and protect shared resources. 
- Sendgrid: A library for sending emails. 
- fatih/color: Makes it easy to add colors to text in the console. 

Why I picked specifically these libraries over others is compared in the research assignment.

Next to these libraries, I also used the following other technologies:

- PostgreSQL: A relational database management system setup in a docker container.
- Docker: A container platform.
- Docker-compose: A tool for defining and running multi-container Docker applications.
- Sendgrid: A service for sending emails.
- Air: A live reload tool for Go applications. 
- Grafana: A tool for visualizing data.
- Insomnia: A tool for testing tool for http, websockets and graphql requests.


Why I picked them specifically is again compared in the research assignment.

Since I added around 1000 lines of code, I will only focus on the most important parts of the source code and not too much on language specific
details. The starting point of the backend is the main file and its main function as you might know it from
Java or C#. 

First there are two global variables declared, which are the dbManager and the fiberApp. Thet get initialized in the initBackend function and
are created with the New functions The new functions is similar to the constructor in Java or C#. It returns a new object of the given type and also initializes it.
The DBManager in this case is initialized with the NewDBManager function. This function takes 2 parameters, the database type and the connection string. 
The fiberApp is initialized with a seperate function. This is because the fiberApp now listens for requests and is blocking the main thread.

```go
package main

var dbManager *DBManager
var fiberApp *FiberApp

func main() {
	initBackend()
}

func initBackend() {
	//new db manager object
	dbManager = NewDBManager("Postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	dbManager.GetInfo()
	dbManager.setupDb()

	//new fiber app object
	fiberApp = NewFiberApp()
	fiberApp.InitFiber() //This must be the last thing since it starts the server
}
```

I will continue next with the DBManager. This is a important part of the backend, since it is responsible for all database related tasks. 
This Function is creating a new DBManager object and also initializes it. It takes 2 parameters, the database type and the connection string.
It sets up the postgres database connection with the gorm library. 

```go
func NewDBManager(name string, dsn string) *DBManager {
	var d = DBManager{db: nil, name: name, insertSql: "insert.sql", rebuildSql: "rebuild.sql"}

	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		d.LogError("Failed to connect database")
	}

	d.Log("Connected to database")

	return &d
}
```


This the struct of the object. It contains the database connection, the name of the database, the insert sql file and the rebuild sql file.

```go
type DBManager struct {
	db         *gorm.DB
	name       string
	insertSql  string
	rebuildSql string
}
```

The function called afterwards is the GetInfo function. This function returns information about the DBManager. This function is only used to ensure that the DBManager is initialised correctly.

```go
func (d *DBManager) GetInfo() {
d.Log(fmt.Sprintf("%+v", d))
}
```

It is one of many utility functions that I added to the DBManager and the fiberAPP to make it easier to develop and debug the backend:

```go

func (d *DBManager) SetDBManager(db *gorm.DB, name string) {
	d.SetName(name)
	d.SetDB(db)
}

func (d *DBManager) SetDB(db *gorm.DB) {
	d.db = db
}

func (d *DBManager) SetName(name string) {
	d.name = name
}

func (d *DBManager) GetDB() *gorm.DB {
	return d.db
}

func (d *DBManager) GetName() string {
	return d.name
}

func (d *DBManager) ToString() string {
	return fmt.Sprintf("Running %s ", d.name)
}
```

But also more complex functions which I cant show complete in the scope of this assignment. 

These functions include:



```go


## Architecture

During writing this assignment it is much easier to follow if you have a look at the architecture first.

Main architecture:
3

SUBSTANTIATION.
The improvement proposals are substantiated with explicit links to your research outcomes.

REALISATION.
The realised improvements are documented with code explanations, pictures, videos, etc. on your markdown site.
