## Local development

We use [goose][goose] for database migrations. The test database is
`rickover_test` and the development database is `rickover`. The authenticating
user is `rickover`.

### Dependency setup

Run compose-up to set up necessary dependency
```
make compose-up
```

If you work in host OS, please make sure you have go1.17 setup and run migrate to setup database
```
make migrate
```

### Start the server


```
make server
```

Will start the example server on port 9090.

### Start the dequeuer

```
make worker
```

Will try to pull jobs out of the database and send them to the downstream
worker. Note you will need to set `DOWNSTREAM_WORKER_AUTH` as the basic auth
password for the downstream service (the user is hardcoded to "jobs"), and
`DOWNSTREAM_URL` as the URL to hit when you have a job to dequeue.

## Debugging variables

- `DEBUG_HTTP_TRAFFIC` - Dump all incoming and outgoing http traffic to stdout
