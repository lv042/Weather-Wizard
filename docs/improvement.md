# Improvements

My research focused on the backend of Weather Wizard. I wanted to implement the best possible backend in the given timeframe. In the
following I will explain the improvements I made.

## DOCUMENTATION
Since I added around 1000 lines of code, I will only focus on the most important parts of the source code and not too much on language specific
details. In the documentation I will explain the architecture of the backend and the most important parts.


To explain the strucutre of the new backend, I will start with the most important part: the `main.go` file. This file is the entry point of every Golang application.

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
Im using Golang language for the backend of Weather Wizard. before talking about the implementation of the backend, I will explain the language itself.
Golang is a compiled language, which means that the code is compiled to machine code before it is executed. This makes it much faster than interpreted languages like PHP.
Golang is also a statically typed language, which means that you have to define the type of variable when you declare it. This makes it easier to find bugs in your code.
It is also garbage collected language like Java, which means that you don't have to worry about memory management like in C or C++. The garbage collector will clean up after you.
Furthermore, Golang is a concurrent language, which means that you can run multiple threads at the same time. This is very useful for web development, since you can run multiple
requests at the same time. It is language that is very easy to learn and has a lot of useful features. This makes it a very popular language for backend development.


First of all I should mention that programming in Golang is rather functional than object-oriented. This means that there are no classes, but instead
you can use structs. A struct is a collection of fields. In the example above you can see that I created a new `DBManager` and `FiberApp` struct. These objects
are managing the api and the database. I will go into more detail about these objects later. 


## Architecture

During writing this assignment it is much easier to follow if you have a look at the architecture first.

Main architecture:
3

SUBSTANTIATION.
The improvement proposals are substantiated with explicit links to your research outcomes.

REALISATION.
The realised improvements are documented with code explanations, pictures, videos, etc. on your markdown site.
