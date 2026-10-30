# Gateway Service

Fiber API gateway for ShopCore. It proxies authentication and user requests to the User Service, catalog requests to the Product Service, and order requests to the Order Service. Redis is used for IP-based rate limiting.

## Run

```sh
go mod download
go run ./cmd/gateway
```

Default port: `8080`.
