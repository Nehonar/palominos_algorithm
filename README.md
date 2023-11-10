# palominos_algorithm
Extract big data information

## Build docker
```bash
docker build -t palominos_algorithm .
```
## Run docker
```bash
docker run -p 8080:8080 palominos_algorithm
```

## Compile
```bash
go build main.go
del main.zip
tar.exe -a -cf main.zip main
```