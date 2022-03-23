# blockchain-query-api

Service to query blockchain data from `cardano-db-sync`'s database

## Building

Build the binary.

```
$ make
```

Build the docker image.

```
$ make image
```

## Testing locally

Setup a port forward to the `cardano-db-sync` Postgres instance.

```
$ kubectl port-forward svc/cardano-kstack-cardano-postgres-testnet 5432
```

Run the service.

```
$ DATASOURCE_CARDANODBSYNC_USERNAME=postgres DATASOURCE_CARDANODBSYNC_PASSWORD=xxxxxxxx ./blockchain-query-api
```

Perform a request.

```
$ curl -s http://localhost:8080/api/v0/block/xxxxxxxx | jq .
```
