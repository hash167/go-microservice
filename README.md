## Introduction

Building a REST API with Golang. Learning about the Golang standard library and the `http` package. Current learnings
- Go package management
- JSON encoding/decoding
- Gorrilla MUX
- Swagger


Basic react frontend
## Services

### Product API [./product-api](./product-api)
Simple Go based JSON API built using the Gorilla framework. The API allows CRUD based operations on a product list.

### Frontend website [./frontend](./frontend)
ReactJS website for presenting the Product API information

## Deployment

### Backend

```
cd product-api
go run main.go
```

### Frontend

`yarn run`