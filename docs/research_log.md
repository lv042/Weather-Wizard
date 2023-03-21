# Research Log

In this document the current state of the research phase of the project.


## Starting out with the project

I'm currently planning to get a good general overview of Golang. I want to know how to use the language, which libraries and what architectures to use.

I found plenty of info material I want to read and watch through:

-https://www.youtube.com/watch?v=YS4e4q9oBaU
-https://madappgang.com/blog/backend-development-with-golang/
-https://www.youtube.com/watch?v=bj77B59nkTQ

In addition to that I plan to build the new backend in a separate git repository. This way I can easily switch between the old and the new backend. 
The git repository will be added as submodule to the main project, so it can be easily integrated into the project.


## Update on the submodule approach

After doing some research on this topic during the weekend, it turns out that all commits of the included submodule will not show up in the commits of the main project. That would be okay,
if the commits to our projects weren't important for our grade. So I decided to just import the subrepo as a normal folder and push it in the end to the other repository.

## Choosing a backend framework

During the last weekend, I also did some research on what backend frameworks to use with Golang. I found out that there are a lot of interesting solutions out there:

#### Web-Frameworks

To build a API with Golang, you can use the following frameworks:

Gin - A lightweight and fast HTTP web framework with a focus on high performance.
Echo - A high-performance, minimalist Go web framework with a focus on RESTful APIs.
Fiber - A fast and easy-to-use web framework with a focus on performance and minimalism.
Chi - A lightweight, flexible and fast HTTP router and micro-framework for Go.
Revel - A full-stack web framework for building scalable and maintainable web applications in Go.
Gorilla - A toolkit of packages for building web applications and services in Go.

I have to admit that I am not very familiar with any of these frameworks, but have tried Gin and Fiber during the last weekend.
Gin and Fiber a very similar structure to each other, but Fiber is a bit more lightweight and has a more modern approach to the code structure, but both are very easy to learn and look
very similar to the popular Express framework in NodeJS. 

#### Databases

For databases, I think the best choice would be one you are already familiar with or want to learn. SQL and NO-SQL databases are both very popular and have their own advantages and disadvantages:

__Use SQL databases when:__

- Your data has a well-defined schema and requires complex querying.
- Your application requires transactional integrity and ACID compliance.
- Your application requires high levels of data normalization.

__Use NoSQL databases when:__

- Your data is unstructured or semi-structured.
- Your application requires high scalability and high write throughput.
- Your application requires fast data access and real-time analytics.

I think for the Weather Wizard the selection of the database is not that important, because the data is not that complex and the application does not require high scalability or high write throughput. All the mentioned 
arguments are not very important for such a small application.


#### ORM

ORM stands for Object Relational Mapping. It is a technique that lets you query and manipulate data from a database using an object-oriented paradigm.
It might be not that important for the Weather Wizard, but it is very useful technique to learn. I will try to use an ORM for the backend, because it makes the code more readable and easier to maintain.

The number one framework which is recommended to use with Golang is GORM. It is a very popular ORM framework for Golang and has a lot of features. It is also very easy to use and has a great documentation.

#### BaaS 

Backend as a Service (BaaS) is a cloud-based service that provides a backend for mobile and web applications. I have used Firebase in on of my last group projects for the backend, but I worked on a Rust client and not the backend. Since these BaaS became very popular in the last years
and I want to get familiar with one of these tools. I think it is a good idea to use one of them for the Weather Wizard.

There are many BaaS providers out there, these are the three most popular ones:

1. Firebase - Firebase is a comprehensive mobile and web development platform that provides real-time databases, authentication, cloud messaging, and hosting. It's backed by Google and has a large community of developers, making it one of the most popular BaaS solutions.

2. AWS Amplify - Amplify is a set of tools and services from Amazon Web Services (AWS) that simplifies the process of building scalable and secure cloud-powered mobile and web applications. It provides a wide range of features, including data storage, user authentication, APIs, and analytics.

3. Parse - Parse is an open-source BaaS platform that provides APIs for data storage, user management, and push notifications. It's easy to set up and provides a flexible and customizable backend for mobile and web applications.

Since our research project is a bit more small scale and I also want to focus on working with Golang, I think the following BaaS provider is the best choice for the Weather Wizard:

##### PocketBase

PocketBase is an open source solution which was founded in 2020 by a team of developers from the United States.

These are the features of PocketBase, most of them are not that important for the Weather Wizard:

- Compact size: PocketBase has a small size of approximately 15 MB, making it lightweight and easy to deploy [1].

- API and SDK support: It offers complete API and SDK support for Dart and JavaScript, enabling developers to create database collections with ease.

- Real-time operations: PocketBase supports real-time operations within the SDK and through a web API, which is beneficial for applications that require immediate data updates.

- One-file backend: It allows developers to create a one-file backend with all the required functionalities, making it an efficient framework for various applications.

- Security features: PocketBase provides SSL encryption to secure data transmissions, access controls to restrict data access, and backup and restore features to protect against data loss.

- Integration with popular libraries and frameworks: It can be easily integrated with React, a popular JavaScript library for building user interfaces, and React Context for managing state across components.

- Authentication and user management: PocketBase offers authentication features and user management capabilities, making it easier to secure and manage user access to applications.

- Flexible hosting options: PocketBase can be hosted on various platforms, such as fly.io, providing flexibility in deployment and scaling.

If you want to check out PocketBase, you can find some useful links about it here: 

- PocketBase Website: https://pocketbase.io/
- PocketBase Live Demo: https://pocketbase.io/demo
- PocketBase Github: https://github.com/pocketbase/pocketbase
- PocketBase Documentation: https://pocketbase.io/docs

## How to deploy the backend with PocketBase

Since the Frontend is only a static website, I am planning to provide docker container for the backend which runs both the database and the rest of the backend. This way I can integrate nicely with the existing containers or just host it on a public linux server.
At the moment I am considering the following cloud providers:

- DigitalOcean -> A very popular cloud provider which also offers 
- Fly.io -> A new cloud provider with a very interesting pricing model
- AWS -> Amazon Web Services
- BWCloud -> A german cloud provider which is free for students

After doing some research yesterday, I found out that Pocketbase compiles to only one binary file which runs backend and database at once. This makes it very easy to deploy the backend on your own device without the need of a cloud provider or a docker container.

## Is Pocketbase usable for a three-week project?

At moment, I am not sure if Pocketbase is usable for a three-week project. I think it is a very interesting project and I would like to use it for the Weather Wizard, but I am not sure if it is stable enough for a three-week project. It seems to have too many features, I won't need. I will try to use it for the backend, but if it does not work out, I will switch to 
the conventional development of a backend.




