# Inventory Service

This is the Inventory service

Generated with

```
micro new microservice-in-micro/part1/inventory-srv --namespace=mu.micro.book --alias=inventory --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.srv.inventory
- Type: srv
- Alias: inventory

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./inventory-srv
```

Build a docker image
```
make docker
```