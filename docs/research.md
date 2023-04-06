# How can Golang be used to implement a better backend for Weather Wizard?


Go, often referred to as Golang is an open-source programming language that was developed by Google in 2009.
Go is gaining a lot of popularity in the last years and has become one of the most popular languages for backend development.
Many features of Go's make it a great choice for that. Go is a compiled language, which means that it is actually converted to machine code before it is executed.
This makes it a lot faster than interpreted languages like PHP. Another great feature of Golang is its Garbage Collector. This makes it a lot easier to
manage memory in Go. Most other compiled languages require you to manually manage your memory, which can be a lot of work and is very error-prone.
Furthermore, Go is a statically typed language, which means that you have to define the type of variable when you declare it, which makes it a lot easier to
find errors in your code before runtime. The easy syntax of Go also makes it a lot easier to learn than other languages. The syntax of Go
is very similar to C and Python, which makes it easy to switch between those languages. The last feature that I want to mention is the concurrency of Go.
Golang has very good support for concurrency. That means that you run multiple tasks at the same time. This is very useful for backend development,
since it allows you to handle many requests and perform multiple tasks simultaneously,what improves performance and responsiveness.

For all those reasons many companies have switched to Golang for their backend development. Here are some examples:

- Uber: Uber heavily relies on Golang for backend development, notably its dispatch system and map services. Uber can handle a huge number of requests concurrently thanks to Golang's concurrency characteristics, making it simpler to deliver a smooth experience for its consumers.

- Dropbox: Dropbox's desktop client, which manages file syncing between devices, is written in Golang. The desktop client can sync data rapidly and effectively thanks to Golang's fast execution speeds and concurrency characteristics.

- Docker: Docker is a popular containerization program that heavily relies on Golang.


## For whom is this research relevant?

This researchs target audience are developers who wish to create high-performance, scalable, and concurrent applications. Golang is especially well-suited for developing microservices, APIs, and web servers. Because of its simplicity and speed, it is also a popular choice for building DevOps tools like as continuous integration and deployment (CI/CD) systems.


## What are methods to research this question?

Since this research question can be interpreted very broadly, I will have to narrow it down. I will focus on the following topics:

- What libraries could be used?
- What are modern features of Golang?
- What is the best architecture for a backend in Golang?

These topics will be researched by reading literature, watching videos and trying out the language myself. Other sources of information I want to use are the official documentation of Golang, case studies and performance benchmarks.

## What are libraries that could be used?

Since Golang is very popular for backend development, there are also a lot of libraries available for it. 

To get familiar with good libraries which are offered in the Goland ecosystem, I can recommend the following Github page:
https://github.com/mingrammer/go-web-framework-stars

It contains a list of the most popular libaries and frameworks for Golang. It is structured by category, so it is easy to find the right library for your use cases.

https://madappgang.com/blog/backend-development-with-golang/

For weather wizard we need the following types libraries to build a modern backend:

https://kinsta.com/blog/postgresql-vs-mysql/#:~:text=MySQL%20is%20a%20purely%20relational,%2C%20ACID%2Dcompliant%20storage%20engine.
Database: What database should we use?

Since we already have a sql database it makes sense to use another sql database. I decided to pick PostgreSQL, because it is a very popular and modern database. It has a lot of features that make it subjectively better than MySQL. It is also open-source and free to use. 
The stackoverflow developer survey 2022 shows that PostgreSQL is one of the most popular databases for backend development.








