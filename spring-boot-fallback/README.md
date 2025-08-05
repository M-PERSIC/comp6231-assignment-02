# Installation

# Docker build (recommended)

```bash
# Build images
cd fruit_month_price_service/
docker build --no-cache -t fruit-month-price-service:latest .

cd ../fruit_total_pricing_service/
docker build --no-cache -t fruit-total-price-service:latest .

docker run -d \
  --name fmp-service \
  --network host \
  -e SERVER_PORT=8000 \
  fruit-month-price-service:latest

docker run -d \
  --name ftp-service \
  --network host \
  -e SERVER_PORT=8100 \
  -e FMP_PORT=8000 \
  fruit-total-price-service:latest

# Custom ports
docker run -d \
  --name fmp-service-custom \
  -p 8001:8001 \
  --network host \
  -e SERVER_PORT=8001 \
  fruit-month-price-service:latest

docker run -d \
  --name ftp-service-custom \
  -p 8101:8101 \
  --network host \
  -e SERVER_PORT=8101 \
  -e FMP_PORT=8001 \
  fruit-total-price-service:latest
```

## Maven Build

```bash
cd fruit_month_price_service/
mvn clean compile
mvn spring-boot:run
# separate terminal
cd fruit_total_pricing_service/
mvn clean compile
mvn spring-boot:run
```

# Usage

```bash
# Valid requests
curl -X GET "http://localhost:8000/fruit-price/fruit/banana/month/jul" \
  -H "Accept: application/json"

curl -X GET "http://localhost:8000/fruit-price/fruit/apple/month/jan" \
  -H "Accept: application/json"

curl -X GET "http://localhost:8100/fruit-total/fruit/banana/month/jul/quantity/10" \
  -H "Accept: application/json"

curl -X GET "http://localhost:8100/fruit-total/fruit/apple/month/jan/quantity/5" \
  -H "Accept: application/json"

# Invalid requests
curl -X GET "http://localhost:8000/fruit-price/fruit/mango/month/jan" \
  -H "Accept: application/json"

curl -X GET "http://localhost:8000/fruit-price/fruit/banana/month/invalid" \
  -H "Accept: application/json"

curl -X GET "http://localhost:8100/fruit-total/fruit/banana/month/jul/quantity/-5" \
  -H "Accept: application/json"

curl -X GET "http://localhost:8100/fruit-total/fruit/banana/month/jul/quantity/0" \
  -H "Accept: application/json"
```