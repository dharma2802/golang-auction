
# Auction System

Golang-based auction system coupled with Docker


## Deployment

To deploy this project run

```bash
  docker-compose build --no-cache && docker-compose up -d
```


## Features

- Basic Golang REST APIs
- Mysql Integration
- Golang & Mysql Images
- Schema creation druing deployment


## Usage

Use below link to export postman collection

[POSTMAN Collection](https://www.getpostman.com/collections/8fe34d8209d3b6b5ce74)


## Run Locally

Clone the project

```bash
  git clone git@github.com:dharma2802/golang-auction.git
```

Go to the project directory

```bash
  cd golang-auction
```

DB Connection changes

```bash
  db/connection.go
```

DB Schema

```bash
  db/migration.sql
```

Start the server

```bash
  go run main.go
```

