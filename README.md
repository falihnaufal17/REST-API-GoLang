# REST API with Golang and MySql
![](https://img.shields.io/badge/Languange-GO-blue)
![](https://img.shields.io/badge/Database-MySql-orange)


## Prerequiste
- GoLang - Download and Install [Golang](https://golang.org/doc/install)
- MySQL - Download and Install [MySQL](https://www.mysql.com/downloads/) - Make sure it's running on the default port.

## Installation
### Clone
```
$ git clone https://github.com/falihnaufal17/REST-API-GoLang.git
$ go get github.com/go-sql-driver/mysql (for driver mysql)
$ go get github.com/gorilla/mux (for Router)
$ cd REST-API-GoLang
$ go build
```
### Setup Database
1. Open mysql.go
2. Change "root:password@tcp(127.0.0.1:3306)/simple_rest" to your mysql configuration
3. If not using password, just remove "password" example: "root:@tcp(127.0.0.1:3306)/simple_rest"

### Start Development Server
```
$ .\REST-API-GoLang.exe
```

### License
----

© [Falih Naufal](https://github.com/falihnaufal17 "Falih")
