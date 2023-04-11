# Prometheus exporter

Prometheus exporter for MariaDB users statistics (using USER_STATISTICS info <https://mariadb.com/kb/en/information-schema-user_statistics-table/>).

## Usage

Setup or export all needed environment variables:

- IP (listening IP of the exporter, default: 0.0.0.0)
- PORT (listening port of the exporter, default: 8080)
- HTTP_AUTH (enable http authentification on the exporter, default: false)
- HTTP_USER (http user of the exporter, default: none)
- HTTP_PASSWORD_HASH (http password of the exporter, default: none)
- SQL_HOST (MariaDB host, default: localhost)
- SQL_PORT (MariaDB port, default: 3306)
- SQL_USER (MariaDB user, default: root)
- SQL_PASSWORD (MariaDB password, default: none)

Launch the exporter:

```bash
./mariadb_stats_exporter
```

or with some environment variables:

```bash
PORT='8082' HTTP_AUTH='true' HTTP_USER='prometheus' HTTP_PASSWORD_HASH='$2a$10$VccJVw2Cn2NWjEwS0./lmObb7JHrGvOzCz4tsE7yumxkwPf2pGZMi' ./mariadb_stats_exporter
```

(you could now do a request with `curl -u "prometheus:secret" 127.0.0.1:8082/metrics`)

To get password hash

```bash
echo $(htpasswd -bnBC 10 "" "<PASSWORD>" | tr -d ':\n' | sed 's/$2y/$2a/' | sed 's/://')
```

### set environment variable from a file

For development it could be usefull to set environment variable of the current shell from a file

```bash
set -o allexport
source .env
set +o allexport
```

## Build

```bash
git clone git@github.com:ethersys/ethersys-mariadb-stats-exporter.git
go install
CGO_ENABLED=0 go build
```
