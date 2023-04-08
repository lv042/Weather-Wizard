# How can Golang be used to implement a better backend for Weather Wizard?

In the first version of Weather Wizard, we used PHP for the backend. PHP is very popular, has great community support
is used in a lot of projects and is easy to learn. However, there are more modern languages which are coming up and
offer
plenty of benefits over old languages. One of those languages is Go.

Go, often referred to as Golang is an open-source programming language that was developed by Google in 2009.
Go is gaining a lot of popularity in the last years and has become one of the most popular languages for backend
development. The languages philosophy is to be simple, fast and efficient.

## For whom is this research relevant?

The target audience for this research are developers and engineers who are interested in using Golang for backend
development in small scale or personal projects.

## What are methods and sources to research this question?

To narrow down the research question, I will focus on the following topics:

- What are the benefits of developing in Golang?
- What libraries and technologies can be used for backend development in Golang?
- What is the best architecture for a Weather Wizard backend in Golang?

These topics will be researched by reading articles and developer blogs, watching videos and tutorials and by analyzing
surveys and benchmarks.

## What are the benefits of developing in Golang?

Many features of Go's make it a great choice for backend development. Go is a compiled language, which means that it is
actually
converted to machine code before it is executed.
This makes it a lot faster than interpreted languages like PHP. Another great feature of Golang is its Garbage
Collector. This makes it a lot easier to
manage memory in Go. Most other compiled languages require you to manually manage your memory, which can be a lot of
work and is very error-prone. Microsoft estimated that 70% of all security bugs are caused by bad memory management. https://msrc.microsoft.com/blog/2019/07/a-proactive-approach-to-more-secure-code/
Furthermore, Go is a statically typed language, which means that you have to define the type of variable when you
declare it, which makes it a lot easier to
find errors in your code before runtime.https://instil.co/blog/static-vs-dynamic-types/#:~:text=Protection%20from%20Runtime%20Errors,a%20much%20smoother%20development%20experience. The easy syntax of Go also makes it a lot easier to learn than other languages.
The syntax of Go
is very similar to C, which makes it easy to switch between those languages. The last feature that I want to
mention is the concurrency of Go.
Golang has very good support for concurrency. That means that you run multiple tasks at the same time. This is very
useful for backend development,
since it allows you to handle many requests and perform multiple tasks simultaneously,what improves performance and
responsiveness.
https://medium.com/@julienetienne/why-go-the-benefits-of-golang-6c39ea6cff7e

For all those reasons many companies have switched to Golang for their backend development. Here are some examples:

- Uber: Uber heavily relies on Golang for backend development, notably its dispatch system and map services. Uber can
  handle a huge number of requests concurrently thanks to Golang's concurrency characteristics, making it simpler to
  deliver a smooth experience for its consumers.
  https://www.uber.com/en-NL/blog/aresdb/
- Dropbox: Dropbox's desktop client, which manages file syncing between devices, is written in Golang. The desktop
  client can sync data rapidly and effectively thanks to Golang's fast execution speeds and concurrency characteristics.
  https://dropbox.tech/infrastructure/open-sourcing-our-go-libraries
- Docker: Docker uses Golang to implement its core container runtime, as well as various other components such as the
  Docker CLI and Docker Compose.
  https://news.ycombinator.com/item?id=6709517

And there are many more examples, like Netflix, Meta, Cloudflare and of course Google itself. Most developer teams have
said that they switched to Golang because of its performance and concurrency characteristics.
https://go.dev/solutions/case-studies

### Drawbacks of Golang

Even though Golang has many great features, it also has some drawbacks.

- Lack of Generics: Because Golang doesn't have a generic type system, it is more difficult to write code that requires
  generic programming capabilities.

- A smaller standard library: Golang has a standard library, however it is small in contrast to other programming
  languages, which could build up a big ecosystem of libraries over the last years. This means that developers must
  frequently rely on third-party libraries to do particular tasks.

- Different Syntax: Golang has a particular syntax and approach to programming that may be different for certain
  developers to learn and understand. Also, because it is a younger language, there aren't as many learning resources as
  there are for more established languages like Java or Python.

## What libraries and technologies can be used for backend development in Golang?

Since Golang is very popular for backend development, there are also a lot of libraries available for it.

To get familiar with good libraries which are offered in the Goland ecosystem, I can recommend the following Github
page:
https://github.com/avelino/awesome-go

It contains a list of the most popular libraries and frameworks for Golang. It is structured by category, so it is easy
to find the right library for your use cases. It is also very well maintained and updated frequently.

https://madappgang.com/blog/backend-development-with-golang/

For weather wizard we need the following types of libraries and technologies to build a modern backend:

https://kinsta.com/blog/postgresql-vs-mysql/#:~:text=MySQL%20is%20a%20purely%20relational,%2C%20ACID%2Dcompliant%20storage%20engine.

#### Database:

https://survey.stackoverflow.co/2022/
Since we already worked with MySQL it makes sense to use another sql database. I decided to pick PostgreSQL, because
it is a very popular and modern database. It has a lot of features that make it subjectively better than MySQL. Of
course, it is
also open-source and free to use. The stackoverflow developer survey 2022 shows that PostgreSQL is the most popular
database under professional developers. It is also the most
popular database for backend development.

![image](./images/stack.png)
Nevertheless, there is not a big one decisive reason why I should pick one of the popular sql databases over the others
Almost all popular sql databases, which are not deprecated, are very good and have a lot of features.

NoSQL databases on the other hand would require us to change our whole data model, which is not worth the effort.

https://github.com/mingrammer/go-web-framework-stars

#### Web-framework:

The most popular and most used web-framework for Golang is currently Gin, but after doing some research I decided to
pick Fiber. Fiber is a better choice than Gin due to its better performance, nicer syntax, and more comprehensive
documentation. 
https://www.youtube.com/watch?v=10miByMOGfY&t=735s
The strong performance is also shown in the TechEmpower benchmarks.The benchmark consists of three parts.

1. Database access: The test demands the web application to connect to a database, get records, and execute specified
   operations on the data.

2. Server-side templating: After retrieving records from the database, the web application must generate an HTML view
   using server-side templates. This tests the framework's ability to handle dynamic HTML creation.

3. JSON serialization: The test requires the web application to serialize data into JSON format, which tests the
   framework's JSON handling skills.

![benchmark](./images/bench.png)

Fiber is the third fastest web-framework for Golang, with only rather small and less popular frameworks being faster.
It is also the 24. fastest web-framework for all
languages. Gin is only the 162. fastest web-framework.
https://www.techempower.com/benchmarks/#section=data-r21&test=fortune
For me, it was also easier to get started with Fiber, because I was already familiar Express.js, which is a similar
web-framework for Node.js.

#### ORM (Object Relational Mapping): https://github.com/go-gorm/gorm

https://blog.bitsrc.io/what-is-an-orm-and-why-you-should-use-it-b2b6f75f5e2a
ORM is a technique that allows you to query and manipulate data from a database using an object-oriented paradigm.
Using an ORM allows developers to leverage their fluency in a programming language, simplifying database interactions
and abstracting away SQL complexities. It offers easy database system switching, advanced built-in features, and often
leads to better-performing queries than handwritten SQL. It is typically recommended for small projects to make rapid
development easier.

#### Monitoring and Logging and Notification:
For monitoring the most popular frameworks are Prometheus and Grafana. Both are open-source and free to use. You can
deploy them in a docker container and connect them to your application.
After setting them up you can see the performance of your application in real-time like here:
https://prometheus.io/docs/introduction/overview/
https://grafana.com/docs/

![grafana](images/graf.jpg)
https://grafana.com/

You can also set up alerts, so you get notified if something goes wrong.

For logging, a logging framework doesnt have to be used for a small scale project like Weather Wizard. The standard
library of Golang already provides a good logging framework.
https://betterstack.com/community/guides/logging/logging-framework/
For bigger projects it is recommended to use a proper logging framework. The most popular and widely used logging
framework is Logrus. https://github.com/sirupsen/logrus
https://betterprogramming.pub/awesome-logging-in-go-with-logrus-70606a49f285
Logrus supports a variety of output formats, log levels, and hooks, as well as structured logging. Its adaptability and
extensibility make it an excellent choice for a wide range of applications.

Notifications can be implemented by Grafana or Prometheus, but also by a separate notification framework. The
notifications can be
implemented by email, Slack, Telegram or other messaging services.

I decided to use an email service for the notifications since it doesn't require setting up a mail server. The most
popular email service which also supports
Golang development is SendGrid. It also offers a free tier with 100 free emails per day and is very easy to set up.
https://rapidapi.com/blog/email-apis-which-one-is-right-for-you/

### Authentication

https://dev.to/kcdchennai/how-jwt-json-web-token-authentication-works-21e7
https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication
For authentication, the most popular token based authentication is JWT. JSON Web Tokens (JWTs) are a widely used modern
standard for securely exchanging data between parties. JSON Web Tokens are often used to authenticate users and convey
authorization information and consist of a header, payload, and signature. JWTs are adaptable and simple to implement,
and they may be
used in a broad range of applications and use situations. Although for a small scale project like Weather Wizard, a more
simplistic authentication method can also be used. Basic authentication is a simple authentication scheme built into the
HTTP protocol. The client sends HTTP requests with
the Authorization header that contains the word Basic word followed by a space and a base64-encoded string username:
password. The server decodes the base64 string, splits the string into the username and password components, and
verifies that the given credentials are valid. Basic authentication may not be very feature rich, but it is very easy to
implement. It also doesn't require any additional libraries. It's only draw is that it doesn't offer confidentiality, so
the password is sent in plain text, but this can be mitigated by using HTTPS instead of HTTP.

The implementation of Basic authentication looks like this:

![basic auth](images/auth.png)

### Testing:

The most popular testing framework for Golang is Testify, but most Golang developers stated in the JetBrains survey,
that they prefer the standard library of Golang for testing.
https://www.jetbrains.com/research/devecosystem-2018/go/
https://www.jetbrains.com/lp/devecosystem-2022/go/
Altogether, the standard library of Golang is very good and provides all the necessary tools for testing. It is also
very easy to use and has a good documentation.
Testify can be used if more advanced features are needed, but for a small scale project, the
standard library is sufficient.
Testify has a lot of features like mocking, assertions, and test suites.

## What is the best architecture for a Weather Wizard backend in Golang?

https://www.atlassian.com/microservices/microservices-architecture/microservices-vs-monolith#:~:text=A%20monolithic%20architecture%20is%20a,monolith%20architecture%20for%20software%20design

Before you can implement a new backend, you should decide what architecture pattern you want to use. There are two main
architecture patterns that are relevant for this project. The first one is the monolithic architecture and the second
one is the microservices' architecture.

### Microservices' architecture:

Microservice architecture is an architectural approach in which an application is divided into smaller, autonomous
services that communicate with one another using APIs. This method enhances scalability, flexibility, and fault
separation while also making deployment and maintenance easier. Nevertheless, it can result in additional complexity and
administrative overhead, higher infrastructure and operating expenses, and more challenging end-to-end functionality
testing.

### Monolithic architecture:

Monolithic architecture is an architectural style in which an application is constructed as a single, cohesive system.
This technique encourages simplicity, quicker development and testing, and improved speed owing to fewer network calls.
Nevertheless, it can limit scalability and flexibility, have weaker fault isolation and resilience, and be more complex
to deploy and manage. It also provides limited support for polyglot development, which allows various programming
languages and technologies to be utilized for different areas of the application.

### Which architecture pattern should be used for what project?

A monolithic architecture may be a better choice for a small-scale project than a microservice design. A monolithic
design is quicker to construct and administer, needs less infrastructure resources, and is easier to verify end-to-end
functionality.

Microservice architecture, on the other hand, is more complicated and necessitates more infrastructure resources. It is
often better suited for bigger and more complicated applications that require high scalability, fault tolerance, and
autonomous service deployment.

Nevertheless, the particular objectives and constraints of the project ultimately decide between monolithic and
microservice design. If the project is projected to grow significantly in the future or if it requires independent
scaling and deployment of various services, microservice architecture may be a better option. On the other hand, if the
project is anticipated to have a lower scope and scale or if simplicity and convenience of development and management
are the primary criteria, a monolithic design could be more suited.

## How could such an implementation look like?

https://www.ibm.com/topics/rest-apis
The most widely used standard for APIs in a monolith architecture is REST.

The REST API (Representational State Transfer) is a web-based software architecture approach for developing scalable and
adaptable distributed systems. A RESTful API enables customers to access and alter resources over HTTP using established
methods.

Clients should be able to conduct conventional CRUD (Create, Read, Update, Delete) activities on resources defined by
URIs via a REST API (Uniform Resource Identifiers). The API should use a stateless communication paradigm, in which the
server does not save the client's state and each request provides all the information needed to process it.

#### Why should be a REST API be used?

https://www.integrate.io/blog/why-you-need-a-rest-api/

- Simplifies development by defining a standard language via which diverse systems may communicate.
- Offers an abstraction layer over technical information, making it easier for developers to deal with.
- REST is the most common API design, with a plethora of tools and courses accessible for developers.
- The API provides technical advantages such as interoperability with multiple message formats and being more
  lightweight than other API designs.
- REST APIs may be used for a variety of applications, including automated ETL and data integration.

## Conclusion

To summarize, Golang is an excellent alternative for creating high-performance, scalable, and concurrent applications.
Because of its compiled nature, garbage collector, static typing, easy syntax, and high concurrency support, it is a
popular choice for backend development. Organizations like Uber, Dropbox, and Docker have turned to Golang for backend
development, citing its high speed and parallelism.

I propose PostgreSQL as the database, Fiber as the web framework, and GORM as the ORM for Weather Wizard. We suggest
Prometheus and Grafana for monitoring and logging. JWT may be used for authentication, or Basic authentication for lower
scale projects. The basic Golang library is adequate for testing, while Testify can be used for more sophisticated
capabilities.

In terms of architecture, whether to utilize a monolithic or microservice architecture ultimately relies on the
project's unique aims and restrictions. A monolithic design may be more suitable for a small-scale project, but a
microservice architecture may be better suited for larger and more complex systems that require great scalability and
fault tolerance.

Lastly, a REST API should be utilized for the backend because it is a widely used and standardized way for constructing
scalable and adaptive distributed systems. It simplifies development, offers technological advantages, and may be
utilized for a wide range of applications.

## Sources
