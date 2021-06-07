# Fetch-app
This API serves to get commodity data as well as provincial and weekly reports. Some of the features available are commodities and commodity reports by province and weekly. You can find complete feature documentation regarding this application at the following [link](https://github.com/hudabikhoir/efishery-monorepo/blob/master/API.md)

Built using a hexagonal architecture so that the module can be structured properly. This design architecture supports several database drivers such as sql, sqlite, and mongodb. And is equipped with a cache if one day it is used. For now I'm implementing caching in http requests when calling the eksternal API.

## Technology
- Golang
- Echo Framework
- Linq (language integrated query)
- Docker
- Support mysql, sqlite and mongodb driver

## Installation

there are 2 ways to run fetch-app manually and docker. To run the fetch-app manually we use reflex so when there is a change in our code, the program will automatically compile it. You can use it for the development process. The installation is as follows:
```
    $ cp config/config.yaml.example config/config.yaml
    $ ./run.sh
```
or you can build with docker file
```
    $ docker build -t efishery-go .
    $ docker run -p 5001:5001 -it efishery-go
```
- you can access `http://127.0.0.1:5001/` on your browser or postman

## Module Structure
The code implementation was inspired by port and adapter pattern or known as [hexagonal](blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example):

- **Business**<br/>Contains all the logic in domain business. Also called this as a service. All the interface of all the repository needed and the service itself will be put here.
- **Core**<br/>Contains model/entity and all pure function helper relate to this will
- **Modules**<br/>Contains implementation of interfaces that defined at the business (also called as adapters in hexagonal's term) and also http handler (controller).


```
├── Dockerfile
├── README.md
├── api
│   ├── common
│   │   └── response.go
│   ├── middleware
│   │   └── middleware.go
│   ├── router.go
│   └── v1
│       └── commodity
│           ├── controller.go
│           └── response
│               ├── get_commodity.go
│               └── get_commodity_report.go
├── app
│   ├── main.go
│   └── modules
│       └── modules.go
├── business
│   ├── auth
│   │   ├── service.go
│   │   └── user.go
│   ├── commodity
│   │   ├── comodity.go
│   │   └── service.go
│   └── error.go
├── config
│   ├── config.go
│   ├── config.yaml
│   └── config.yaml.example
├── go.mod
├── go.sum
├── modules
│   └── repository
│       ├── auth
│       │   ├── factory.go
│       │   └── rest_api.go
│       └── commodity
│           ├── factory.go
│           └── rest_api.go
├── run.sh
└── util
    ├── dbdriver.go
    └── rediscon.go
```
