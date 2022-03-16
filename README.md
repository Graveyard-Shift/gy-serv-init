<div align="center">
    <image src="./assets/logo.png" alt="">
    <h1>gy-serv-init</h1>
</div>

> `gy-serv-init` initiates a new Golang webserver that uses the `gin-gonic` package

### Usage

```
gy-serv-init MyAppName
cd MyAppName
```

> The `gy-serv-init` command takes in one argument; the project's name. It takes the project name and creates a directory named ${projectName} inside the cwd. This directory is the directory the boilerplate will be written in.

```
MyAppName
├── src
│   ├── controllers
│   |   ├── main.controllers.go
│   |   └── utils.go
|   |
│   └── router
|       └── router.go
├── .editorconfig
├── .gitignore
├── go.mod
└── main.go
```
