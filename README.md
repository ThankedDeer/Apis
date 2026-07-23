# Apis

Small Go/Gin project built to practice Go fundamentals that don't show up in CRUD tutorials: concurrency with goroutines, fan-in/fan-out with channels, `sync.WaitGroup`, context propagation, generics, and clean error handling — without frameworks that hide the mechanics.

## What it does

A single `/dashboard` endpoint fans out 5 concurrent requests to independent public APIs (news, music, canonical entities, breweries, weather METARs), aggregates whichever ones succeed, and returns partial results if some fail — instead of failing the whole request.

```
GET /dashboard
        │
        ├── goroutine → news API
        ├── goroutine → music API
        ├── goroutine → canonical API
        ├── goroutine → breweries API
        └── goroutine → weather API
                │
        (fan-in via buffered channel)
                │
         aggregate + respond (partial or full)
```

## Endpoints

| Method | Path        | Description                                   |
|--------|-------------|------------------------------------------------|
| GET    | `/health`   | Liveness check                                  |
| GET    | `/dashboard`| Aggregates 5 external APIs concurrently         |

## Concepts exercised

- **Goroutines + `sync.WaitGroup`**: each external call runs in its own goroutine; the handler waits for all of them without blocking on any single one.
- **Fan-in with channels**: a buffered `chan models.TaskResult` collects results as they arrive; a separate goroutine closes the channel once `wg.Wait()` returns, so the consumer loop can safely `range` over it.
- **Context propagation**: `context.Context` flows from the incoming `*gin.Context` down to every outbound `http.Request`, so cancellation/timeouts propagate to in-flight external calls.
- **Partial failure as a first-class outcome**: one slow/broken upstream doesn't take down the whole response — errors are collected per source and reported alongside whatever data did come back.
- **Generics**: `Result[T any]` is a single response envelope reused across `HealthStatus` and `DashboardData`.
- **Error wrapping**: `http_client` wraps errors with `%w` and treats any non-2xx response as an error, keeping failure handling centralized instead of repeated per call site.

## Project layout

```
.
├── main.go              # router setup, route registration
├── handlers/            # request handlers + concurrent orchestration (Dashboard)
├── http_client/          # thin HTTP client wrapper (context, headers, status validation)
└── models/               # DTOs for each upstream API + shared response types
```

## Running it

```bash
go run main.go
```

Default Gin port (`:8080`) unless `PORT`/`GIN_MODE` env vars say otherwise.

```bash
curl localhost:8080/health
curl localhost:8080/dashboard
```

## Roadmap / ideas to keep practicing

- [ ] Add timeouts per upstream call (`context.WithTimeout`) instead of relying only on request cancellation.
- [ ] Replace unbuffered error handling with `errgroup` for comparison against the manual `WaitGroup` approach.
- [ ] Add unit tests with a fake `http.RoundTripper` to test partial-failure paths deterministically.
- [ ] Add structured logging (`slog`) instead of `log.Printf`.
- [ ] Add a request-level timeout/circuit breaker per external source.
