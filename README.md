onlinejudge-server
==================

Backend component for the codeIIEST online judge.
Written in Golang, it is meant to be fast and handle
concurrency well.

The project is currently in its development stage but
there are several guidelines which should be followed.

What it should support:

- Parallel execution of test programs in a sandboxed environment (no system("shutdown"))

- Administrator mode for adding/modifying problems, adding test cases etc.

- User mode for contests and free practice

- Fair job scheduler (FIFO).


## Table of Contents
1. [ Installation ](#install)
2. [ Dependencies ](#depend)
3. [ Contributing ](#contrib)

<a name="install"></a>
## 1. Installation

>Make a .env file and config it accordingly. A sample is provided

### Starting the database

Make sure you have docker installed. 

To verify the docker config, check the output of the following command. 

```bash
$ docker-compose config
```

Once verified run

```bash
$ docker-compose up
```
This will fire up a docker container with postgress and pg-admin.

### Run Migrations

Make sure your postgres container is running. Now run the following commands

```bash
$ go run migrations/*.go  init
```
This will print something like `version is 0`. If no error then continue

```bash
$ go run migrations/*.go
```

This will create the required tables, and migrate the version. 

### Start the go server

``` go build -o build/server```


<a name="depend"></a>
## 2. Dependencies

- [ test-runner ](https://github.com/raydwaipayan/test-runner)
- [ Gin Gonic ](https://github.com/gin-gonic/gin)

<a name="contrib"></a>
## 3. Contributing

Please use the issue tracker.
All contributions are more than welcome :)
