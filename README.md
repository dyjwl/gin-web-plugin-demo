# Gin Web Application with plugins

## layout

This gin web application layout use the layout of [project layout](`https://github.com/golang-standards/project-layout`)


## plugins

The third-party libraries used in this project are open source solutions, which are only used for learning. If you use them in commercial projects, pay attention to the open source protocols used by their open source libraries to prevent commercial disputes.
- gin
- mysql
- postgresql
- sqlite3
- reids
- kafka
- rbac
- redis lock
- go-playground validator
- zap log
- token  refresh token
- swagger
- viper
- docker sdk api
- kubernetes client sdk api
- smtp
- middlware
- shutdown
- zip / unzip
- breakpoint
- uuid


## documentation

### Distributed Locks with Redis

A Distributed Lock Pattern with Redis
Distributed locks are a very useful primitive in many environments where different processes must operate with shared resources in a mutually exclusive way.

There are a number of libraries and blog posts describing how to implement a DLM (Distributed Lock Manager) with Redis, but every library uses a different approach, and many use a simple approach with lower guarantees compared to what can be achieved with slightly more complex designs.

This page describes a more canonical algorithm to implement distributed locks with Redis. We propose an algorithm, called Redlock, which implements a DLM which we believe to be safer than the vanilla single instance approach. We hope that the community will analyze it, provide feedback, and use it as a starting point for the implementations or more complex or alternative designs.

More Information you can read here: redis.io/docs/reference/patterns/distributed-locks/