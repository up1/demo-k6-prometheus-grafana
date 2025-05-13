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

### 3. Convert Postman collection file to K6 Script
* [Postman to K6](https://github.com/apideck-libraries/postman-to-k6)

Install
```
$npm install -g @apideck/postman-to-k6
```

Convert
```
$cd postman
$postman-to-k6 demo.postman_collection.json -o demo-api.js
```

### 4. Run K6 script with Docker compose
* [Types of load testing](https://grafana.com/load-testing/types-of-load-testing/)
```
$docker compose build k6
$docker compose up k6
```
See result in Grafana dashboard


### 5. Delete all
```
$docker compose down
```