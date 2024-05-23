# Challenges

Create projeto to stress test

```bash
go build -o loadtest main.go
```

```bash
./loadtest --url=http://google.com --requests=1000 --concurrency=10
```

```bash
sudo docker build -t loadtest .
```

```bash
sudo docker run loadtest --url=http://google.com --requests=1000 --concurrency=10
```