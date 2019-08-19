# Auction

Api with the standard lib for do manage a auction oferted prices

## Requirements

- docker or go

## Description
```sh
├── api
│   └── server.go (save persistent data) 
├── db
│   ├── mem.go (functions for manage data in men and transform to json for dump and load)
│   └── mem_test.go
├── handlers (handlers used for api, communicate by interface)
│   ├── bid.go
│   ├── bid_test.go
│   ├── common.go (common functions in handlers)
│   ├── stats.go
│   └── stats_test.go
├── main.go (try load file with data dump and if can't do this, start empty, when you press ctr+c this save the backup) 
├── Makefile 
├── models (common structs in program) 
│   ├── bid.go 
│   ├── item.go
│   └── stats.go 
└── start.sh (start server)
```

## Usage

The focus is for usage of Makefile, you can see it for get more info, but some infos here

### Run
For run the project you can use:
```sh
go run main.go -h
# or 
make run
# or
make run args="-port 3000 -filepath auction.dump"
# or
./start -port 3000 -filepath auction.dump
```

## Test

For run all tests you can use
```sh
make test
````

### Test with args

For run tests with args you can use some like this
```sh
make test args="-v db/*.go -run TestDb"
```

or for specify benchmark tests
```sh
make test args="./... -bench Db"
```

### Test while run

first run the server
```sh
make run
```

so seed this server
```sh
curl -X POST -s localhost:3000/bid --data '{"item_id": "23", "price": 1.94, "client_id": "1"}'
curl -X POST -s localhost:3000/bid --data '{"item_id": "23", "price": 1.94, "client_id": "2"}'
curl -X POST -s localhost:3000/bid --data '{"item_id": "2", "price": 1.94, "client_id": "2"}'
curl -X POST -s localhost:3000/bid --data '{"item_id": "2", "price": 1.98, "client_id": "2"}'
curl -X POST -s localhost:3000/bid --data '{"item_id": "23", "price": 1.93, "client_id": "3"}'
curl -X POST -s localhost:3000/bid --data '{"item_id": "2", "price": 3.10, "client_id": "3"}'
```

then see the stats
```sh
curl -X GET -s localhost:3000/stats | jq "."
```
