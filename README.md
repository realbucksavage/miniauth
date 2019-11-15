# MiniAuth - A small UAA server implemented in go.

---

## Installation

For now, this project depends on a static PSQL config.

### PSQL

Pull from Docker

```bash
$ docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
$ docker exec -it postgres psql -U postgres
```

Set up user and DB.

```postgresql
CREATE DATABASE miniauth;
CREATE USER miniauth WITH ENCRYPTED PASSWORD 'miniauth';
GRANT ALL PRIVILEGES ON DATABASE miniauth TO miniauth; 
```

### Clone and run

```bash
$ git clone git@github.com:realbucksavage/miniauth
$ cd miniauth
$ dep ensure
$ go run main.go
```