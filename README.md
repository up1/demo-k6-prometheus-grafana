# Workshop performance testing with K6 (Distributed testing)
* K6
* Prometheus
* Grafana

## Step to run

1. Install and [build K6 + Prometheus remote write](https://grafana.com/docs/k6/latest/results-output/real-time/prometheus-remote-write/)
* Go
* Git
```
$go install go.k6.io/xk6/cmd/xk6@latest
$xk6 build --with github.com/grafana/xk6-output-influxdb
```

Run K6
```
$k6 version
```

Run script with K6
```
$k6 run ./k6-script/demo.js
```

2. Run and Store result in Prometheus and Grafana

Start Prometheus
```
$docker compose up -d prometheus
$docker compose ps
$docker compose logs --follow
```

Start Grafana
```
$docker compose up -d grafana
$docker compose ps
$docker compose logs --follow
```

Access to Grafana eith url = http://localhost:3000
* user=admin
* password=admin

Run K6 script
```
$export K6_PROMETHEUS_RW_SERVER_URL=http://localhost:9090/api/v1/write
$export K6_PROMETHEUS_RW_TREND_AS_NATIVE_HISTOGRAM=true
$export K6_OUT=experimental-prometheus-rw
$k6 run ./k6-script/demo.js -o experimental-prometheus-rw
```

or Run K6 script with Docker compose
```
$docker compose build k6
$docker compose up k6 --remove-orphans
```
