# Go REST API with JSON / MySQL / PostgreSQL

Sales Program using JSON or Database (PT. Phincon Bootcamp Exercise)

## Prerequisite

- Go 1.19.7 or higher version
- MySQL or PosgtreSQL

## Running :rocket:

```bash
go install
```

```bash
go run main.go
```

## List Endpoints

- GET `localhost:5000/product`

- POST `localhost:5000/product`

- GET `localhost:5000/voucher`

- POST `localhost:5000/voucher`

- GET `localhost:5000/transaction`

- POST `localhost:5000/transaction`

## Running Swagger UI

- Install OpenAPI (Swagger) Editor via VSCode Extensions

- Open swagger.yml in ./docs

- Press keyboard F1

- Choose OpenAPI: show preview using Swagger UI

## Running Unit Test

```bash
go test ./usecase/product -v -run TestUsecaseProduct
```

```bash
go test ./usecase/transaction -v -run TestUsecaseTransaction
```

```bash
go test ./usecase/voucher -v -run TestUsecaseTransaction
```

```bash
go test ./repository/product -v -run TestRepoProduct
```

```bash
go test ./repository/transaction -v -run TestRepoTransaction
```

```bash
go test ./repository/voucher -v -run TestRepoTransaction
```