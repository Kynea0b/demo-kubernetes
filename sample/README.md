# demo
Go, kubernetes, gRPC...


## grpc

```
go get -u google.golang.org/grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### generate grpc code

```
mkdir -p proto/grpc
make proto
```

## build

```
make build
```

## request

### draw omikuji

```
make draw USER_ID=1 PORT=50942

{"item_id":3,"item_name":"item3","rarity":3}
```

### get omikuji histories

```
make history USER_ID=1 PORT=50942 | jq

{
  "histories": [
    {
      "item_id": 3,
      "item_name": "item3",
      "rarity": 3,
      "created_at": {
        "seconds": 1687778806
      }
    },
    {
      "item_id": 2,
      "item_name": "item2",
      "rarity": 2,
      "created_at": {
        "seconds": 1687778249
      }
    },
    ...
}
```

### get item inventory

```
make inventry USER_ID=1 PORT=50942 | jq

{
  "items": [
    {
      "item_id": 1,
      "item_name": "item1",
      "rarity": 1,
      "count": 6
    },
    {
      "item_id": 2,
      "item_name": "item2",
      "rarity": 2,
      "count": 2
    },
    {
      "item_id": 3,
      "item_name": "item3",
      "rarity": 3,
      "count": 2
    },
    {
      "item_id": 4,
      "item_name": "item4",
      "rarity": 4,
      "count": 2
    }
  ]
}
```
