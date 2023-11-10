# bukalawak-web
Web for Bukalawak

HTMX + Go

Cross Compile
```
GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux main.go
```

Build Docker image
```
docker build -t bukalawak-web .
```

Docker run
```
docker run -dp 8000:8000 bukalawak-web
```