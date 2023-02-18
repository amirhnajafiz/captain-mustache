<p align="center">
    <img src=".github/readme/logo.png" width="300" alt="logo" />
</p>

<br />

<p align="center">
    <img src="https://img.shields.io/badge/Docker-3.9-66ADD8?style=for-the-badge&logo=docker" alt="go version" />
    <img src="https://img.shields.io/badge/Version-1.2.1-informational?style=for-the-badge&logo=github" alt="version" />
</p>

<br />

Dockerize your Golang applications under Captain *Mustache*. 
A fast, safe, dynamic way to create dockerfiles for you golang applications.
With this tool you don't need to know docker to dockerize your applications, 
all you need is to have docker installed on your system.

<br />

## Table of contents

- [Install](#how-to-use)
- [Commands](#commands)
- [Stubs](#stubs)

<br />

## How to use?

Clone into the repository:

```shell
git clone https://github.com/amirhnajafiz/mustache.git
```

Execute the mustache script:

```shell
chmod +x ./bin/mustache
./bin/mustache build
```
<br />

After that you will get a _docker_out/_ directory. Move this directory next to your golang main file where
you start your application.

Now use the following command to start the container:

```shell
./bin/mustache up
```

For stopping the container, use:

```shell
./bin/mustache down
```

<br />

## Commands

|   Command   | Description                                                      |
|:-----------:|------------------------------------------------------------------|
| ```build``` | Build dockerfile and docker compose for your application         |
|  ```up```   | Start building docker containers based on ```docker_out``` files |
| ```down```  | Stop docker containers                                           |

<br />

## Stubs

|   Stub   | Description                                     | Values                                  |
|:--------:|-------------------------------------------------|-----------------------------------------|
| Database | Create database containers for your application | ```redis, mongodb, mysql, postgresql``` |
| Network  | Create docker network for your containers       | ```mustache```                          |
|  Volume  | Create docker volume for your containers        | ```mustache-volume```                   |

<br />

## Supported OS

- MacOS
- Linux
- Windows
