# echo-chamber

A tiny example server that has some useful routes, and is supposed to be
well-instrumented.

## What's it do?

```
/echo     - Echoes the request that comes in
/latency  - Returns within up to 200ms, randomly
/404      - Always returns 404

/_metrics - Exposes metrics for prometheus
```

And there's `cmd/load`, which is a tiny tool that generates "load":

```
# Run 50 req/s against /latency on localhost:8080
$ go run cmd/load/load.go -url http://localhost:8080/latency -per-second 50 -verbose
```

## How to use it

Run `make build && ./echo-chamber`.  This starts the HTTP server on
<http://localhost:12345>.

If you want, you can run a [prometheus](https://prometheus.io) instance,
and configure it to fetch metrics from
<http://localhost:12345/_metrics>.

And if you want to be even fancier, use
[grafana](https://github.com/grafana/grafana) to get a fancy dashboard.
