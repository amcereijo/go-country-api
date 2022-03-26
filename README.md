
# env file
A `.env` file is required with
```
PORT=3001

MYSQL_ROOT_PASSWORD = rootpass
MYSQL_DATABASE = gocountriesapi
MYSQL_LOCAL_PORT = 3306
MYSQL_DOCKER_PORT = 3306
```

# Run Mysql
* build: docker build -t mysql-gocountryapi -f ./Dockerfile-mysql .
* run: docker run --name=mysql-go -p3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=gocountriesapi mysql-gocountryapi

# generate mocks for interfaces with gomock
`> /Users/amcereijo/go/bin/mockgen -source core/ports/ports.go -destination mocks/ports_mock.go`

# Run:
```
> docker compose build

> docker compose up
```

# Run test
```
> go test ./...
```

