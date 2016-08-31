# Golang Boilerplate
> A starter project with Golang, Gin and MongoDB

### Boilerplate structure

```
.
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
├── server
│   ├── router.go
│   └── server.go
├── utils
├── test   
└── vendor
    ├── your_vendor
    └── vendor.json

```

## Installation

__Install Govendor__

`go get -u github.com/kardianos/govendor`

then run:

```
govendor sync
```

## Release History

* 0.0.1
    * First commit :)

Inspired from https://github.com/vsouza/go-gin-boilerplate
