# Cards Board (Go)

API made with Go and CockroachDB to create a board with anonymous and non anonymous messages.

## Run

First migrate the database running:

```bash
cat ./models/initdb.sql | cockroach sql --url your_url
```

Then you can run the project:

```bash
go run .
```
