# Challange - Create CRUD  using Golang

## Tecnologies

- MySQL
- Golang
  
Requirements

Run command in the terminal for create container mysql:

```bash
docker run --name some-mysql \
-e MYSQL_PASSWORD={YOUR_PASSWORD} \
-e MYSQL_USER={YOUR_USER} \
-e MYSQL_DATABASE={YOUR_DATABASE} \
-d mysql:{VERSION}
```

Example:

```bash
docker run --name golang \
-e MYSQL_ROOT_PASSWORD=root_golang \
-e MYSQL_PASSWORD=golang \
-e MYSQL_USER=golang \
-e MYSQL_DATABASE=golang \
-d mysql:8.0.31
```

Use [documentation](https://hub.docker.com/_/mysql) to modify variabels:

- `YOUR_PASSWORD`;
- `YOUR_USER`;
- `YOUR_DATABASE`;
- `VERSION`.

> Create table

Access container and execute command:

```bash
mysql -u golang -p
```

```bash
use golang
```

```mysql
CREATE TABLE user(
    id int auto_increment primary key, 
    name varchar(50) not null,
    email varchar(50) not null) ENGINE=INNODB;
```

In the file `db/db.go` it's necessary modify variable `stringConection`
with your credentials and IP.

> Initialize modules

```bash
go mod init crud
go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql
```

> Methods

Methods utilization in project:

- [GET](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET)
- [POST](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST)
- [PUT](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/PUT)
- [DELETE](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/DELETE)
