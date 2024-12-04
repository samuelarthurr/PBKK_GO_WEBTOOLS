# GOWT
Sample crud web application project using Golang(http, templates, os, sql), Bootstrap 4, DataTables, MySQL, Docker.

### Prerequisites

* Golang, preferably the latest version (1.16).
* MySQL Database
* Docker (optional)

Run below command and install dependencies

```
go mod download
```

3. Create database on MySQL

```
CREATE DATABASE gowtdb CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE gowtdb;

CREATE TABLE tools (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  category varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL,
  url varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  rating int(11) DEFAULT NULL,
  notes text COLLATE utf8_unicode_ci,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
```

4. Create a .env file with the variables listed bellow and change values as needed

```
DATABASE_NAME="gowtdb"
DATABASE_USERNAME="user"
DATABASE_PASSWORD="pass"
DATABASE_SERVER="localhost"
DATABASE_PORT="3306"
```

6. Run the application

```
make run
```

## Deployment

1. Create an executable

```
make build
```

2. Run the application

```
./out/bin/gowt
```
## Create Docker image

1. To build and tag your image locally

```
make docker-build
```

2. To push your image to registry

```
make docker-release
```

## Run Docker image locally

```
make docker-run
```
