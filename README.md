# consent-enforcement

Consent management and enforcement PoC

PoC for consent management / enforcement using Go + graphql + couchdb

Install all dependencies.

```
dep ensure
```

Re-building the code from schema (if needed or if has any update).

```
go run scripts/gqlgen.go
```

Start the server.

```
go run server/server.go
```

Open http://localhost:8080/ for GraphQL Playground