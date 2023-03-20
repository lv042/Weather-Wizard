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



You can decide if you want 
to build your backend completely from scratch or if you want to build it with a specific framework. This is often referred to as a Backend as a Service (BaaS) solution.


