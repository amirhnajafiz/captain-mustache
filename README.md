<p align="center">
    <img src=".github/readme/logo.png" width="192" alt="logo" />
</p>

<h1 align="center">
    Mustache
</h1>

<p align="center">
    <img src="https://img.shields.io/badge/Docker-3.9-66ADD8?style=for-the-badge&logo=docker" alt="go version" />
    <img src="https://img.shields.io/badge/Version-1.1.2-informational?style=for-the-badge&logo=none" alt="version" />
</p>

Dockerize your Golang applications under Captain *Mustache*.

A fast, safe, dynamic way to create dockerfiles for you golang applications.<br />
With this tool you don't need to know docker to dockerize your applications, all you need is to have docker installed
on your system.

## How to use?
Clone the repository:
```shell
git clone https://github.com/amirhnajafiz/mustache.git
```

Start script:
```shell
./bin/mustache
```

If you got error for executing the script, enter the following command and then try again:
```shell
chmod +x ./bin/mustache
```

## Next?
Afet that you will get a _docker_out/_ directory. Move this directory next to your golang main file where
you start your application.

Now use the following command to start the container:
```shell
./bin/mustache up
```

For stopping the container, use:
```shell
./bin/mustache down
```

## Supports
- MacOS
- Linux
