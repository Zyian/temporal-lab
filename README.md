# Temporal Lab - Food Delivery

A simple combination of CLI + Worker process to replicate a simple food delivery system

## Usage

### Required Env Variables
Copy `sample.env` to `.env` and set the appropriate values to your Temporal instance and namespace

### CLI
```shell
go build -o cli cmd/order-cli/main.go
```
Then use:

```shell
source .env
./cli order -l
./cli order 2ce0
./cli order-status <uuid>
./cli order-list
./cli pickup <uuid>
./cli deliver <uuid>
```

### Worker
```shell
go build -o food-worker cmd/delivery-worker/main.go
```

```shell
source .env
./worker
```