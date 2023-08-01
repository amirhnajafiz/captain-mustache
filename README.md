# Captain Mustache

![GitHub release (with filter)](https://img.shields.io/github/v/release/amirhnajafiz/captain-mustache)
![GitHub language count](https://img.shields.io/github/languages/count/amirhnajafiz/captain-mustache)
![GitHub top language](https://img.shields.io/github/languages/top/amirhnajafiz/captain-mustache)
![GitHub Repo stars](https://img.shields.io/github/stars/amirhnajafiz/captain-mustache)



Dockerize your Golang applications under Captain *Mustache*. 
A fast, safe, dynamic way to create dockerfiles for you golang applications.
With this tool you don't need to know docker to dockerize your applications, 
all you need is to have docker installed on your system.

## Table of contents

- [Install](#how-to-use)
- [Commands](#commands)
- [Stubs](#stubs)

## How to use?

Clone into the project repository and build the executable file using the following commands:

```shell
git clone https://github.com/amirhnajafiz/captain-mustache.git
make build
```

## Commands

In here you can find a list of cli commands.

|   Command    | Description                                              |
|:------------:|----------------------------------------------------------|
| ```build```  | Build dockerfile and docker compose for your application |
|   ```up```   | Start building docker containers                         |
|  ```down```  | Stop docker containers                                   |
| ```status``` | Displays a status of your containers                     |

## Stubs

Available stubs on captain mustache.

|   Stub   | Description                                     | Values                                  |
|:--------:|-------------------------------------------------|-----------------------------------------|
| Database | Create database containers for your application | ```redis, mongodb, mysql, postgresql``` |

## Supported OS

- MacOS
- Linux
- Windows
