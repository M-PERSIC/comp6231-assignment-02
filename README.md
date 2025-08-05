# COMP6231 Assignment 02 (Michael Persico)

- **Project Name:** comp6231-assignment-02
- **Programming Languages:** Go, SQL, Java (fallback)
- **Frameworks / Libraries:** go-kit, go-duckdb, Spring Boot Dev Tools (spring), Spring Boot Actuator (spring), Spring Web (spring), Config Client (spring)
- **Database:** DuckDB, H2 (spring)
- **Source Control:** Git (GitHub)
- **Development Tools:** Visual Studio Code, GitHub Codespaces (devcontainer)
- **Architecture:** Dual microservice
- **Deployment:** locally (single executable) or Docker or Kubernetes
- **CI/CD:** GitHub Actions
- **Repository Link:** [github.com/M-PERSIC/comp6231-assignment-02](https://github.com/M-PERSIC/comp6231-assignment-02.git)                                                                     

The main dual microservice program is written in Go, with the Spring Boot fallback program (under the`spring-boot-fallback` directory) written in case the Go program does not meet the assignment requirements.

# Installation

The latest linux binary is available at [github.com/M-PERSIC/comp6231-assignment-02/releases](https://github.com/M-PERSIC/comp6231-assignment-02/releases). To grab the latest binary:

```bash
curl -sL $(curl -sL https://api.github.com/repos/M-PERSIC/comp6231-assignment-02/releases/latest | \
jq -r '.assets[] | select(.name? | match(".*\\.tar\\.gz$")) | .browser_download_url') -o latest-release.tar.gz
tar -xvf latest-release.tar.gz 
```

The binary can be built either manually:

```bash
go build cmd/main.go #./main 
```

or via [GoReleaser](https://goreleaser.com/):

```bash
goreleaser release --snapshot --clean
cd dist/comp6231-assignment-02_linux_amd64_v1/
```

A [Docker image](https://github.com/M-PERSIC/comp6231-assignment-02/pkgs/container/comp6231-assignment-02) is also available via the GitHub Container Registry. 

> NOTE: The Dockerfile present at the top-level directory does not build the image on its own.
> It is used by GoReleaser, which also builds the docker image alongside a local executable.

# Usage

## Docker (recommended)

```bash
docker pull ghcr.io/m-persic/comp6231-assignment-02:latest
# FruitMonthPrice microservice (can omit the fmpPort, starts the FruitMonthPrice at port 8000 by default)
docker run -d \
  --name fmp-service \
  --network host \
  ghcr.io/m-persic/comp6231-assignment-02:latest

# FruitTotalPrice microservice (port 8100)
docker run -d \
  --name ftp-service \
  --network host \
  ghcr.io/m-persic/comp6231-assignment-02:latest --ftp
```

For custom ports:

```bash
# FruitMonthPrice microservice (port 8001)
docker run -d \
  --name fmp-service-custom \
  -e FMP_PORT=8001 \
  --network host \
  ghcr.io/m-persic/comp6231-assignment-02:latest

# FruitTotalPrice microservice (port 8100)
docker run -d \
  --name ftp-service-custom \
  -e FMP_PORT=8001 \
  -e FTP_PORT=8101 \
  --network host \
  ghcr.io/m-persic/comp6231-assignment-02:latest --ftp
```

## Binary (GitHub Releases)

```bash
curl -sL $(curl -sL https://api.github.com/repos/M-PERSIC/comp6231-assignment-02/releases/latest | \
jq -r '.assets[] | select(.name? | match(".*\\.tar\\.gz$")) | .browser_download_url') -o latest-release.tar.gz
tar -xvf latest-release.tar.gz 

# FruitMonthPrice microservice (port 8000)
./comp6231-assignment-02 

# FruitTotalPrice microservice (port 8100)
# requires separate terminal
./comp6231-assignment-02 --ftp
```

For custom ports:

```bash
# FruitMonthPrice microservice (port 8001)
FMP_PORT=8001 ./comp6231-assignment-02

# FruitTotalPrice microservice (port 8101)
# requires separate terminal
FMP_PORT=8001 FTP_PORT=8101 ./comp6231-assignment-02 --ftp
```

## Binary (manual build)

```bash
go build cmd/main.go

# FruitMonthPrice microservice (port 8000)
./main

# FruitTotalPrice microservice (port 8100)
./main --ftp
```

For custom ports:

```bash
# FruitMonthPrice microservice (port 8001)
FMP_PORT=8001 ./main

# FruitTotalPrice microservice (port 8101)
# requires separate terminal
FMP_PORT=8001 FTP_PORT=8101 ./main --ftp
```

# Usage

```bash
# Valid requests
curl -X GET "http://localhost:8000/fruit-price/fruit/banana/month/jul" \
  -H "Accept: application/json"

curl -X GET "http://localhost:8100/fruit-total/fruit/banana/month/jul/quantity/10" \
  -H "Accept: application/json"

# Invalid requests
curl -X GET "http://localhost:8000/fruit-price/fruit/mango/month/jan" \
  -H "Accept: application/json"

curl -X GET "http://localhost:8100/fruit-total/fruit/banana/month/jul/quantity/-5" \
  -H "Accept: application/json"
```
