# Workshop performance testing with K6 (Distributed testing)
* K6
* Prometheus
* Grafana

## Step to run

### 1. Start API for testing

```
$docker compose build api
$docker compose up -d api
```

### 2. Run and Store result in Prometheus and Grafana

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

Access to Grafana with url = http://localhost:3000
* user=admin
* password=admin

### 3. Run K6 script with Docker compose
```
$docker compose build k6
$docker compose up k6 --remove-orphans
```
See result in Grafana dashboard

### 4. Delete all
```
$docker compose down
```