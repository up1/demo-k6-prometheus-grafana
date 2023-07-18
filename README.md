# Workshop performance testing with K6 (Distributed testing)
* K6
* InfluxDB 2
* Grafana

## Step to run

1. Install and [build K6 + InfluxDB 2](https://k6.io/docs/results-output/real-time/influxdb/)
* Go
* Git
```
$go install go.k6.io/xk6/cmd/xk6@latest
$xk6 build --with github.com/grafana/xk6-output-influxdb
```

Run K6
```
$k6 version

k6 v0.43.1 ((devel), go1.21rc2, darwin/arm64)
Extensions:
  github.com/grafana/xk6-output-influxdb v0.4.0, xk6-influxdb [output]
```

Run script with K6
```
$k6 run ./k6-script/demo.js
```

2. Run and Store result in InfluxDB 2

Start InfluxDB 2
```
$docker compose up -d influxdb
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
*password=admin

Run K6 script
```
$export K6_INFLUXDB_ORGANIZATION="influxdata-org"
$export K6_INFLUXDB_BUCKET="k6_results"
$export K6_INFLUXDB_TOKEN="5up3r-S3cr3t-auth-t0k3n"
$export K6_INFLUXDB_ADDR="http://localhost:8086"
$k6 run ./k6-script/demo.js -o xk6-influxdb
```
