# Etapa de construcción: usar una imagen de Go
FROM golang:1.16 as builder

WORKDIR /app

# Copiar los archivos go.mod y go.sum e instalar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación para Linux (compilación cruzada)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./

# Etapa de ejecución: usar la imagen base de Lambda
FROM public.ecr.aws/lambda/go:1

COPY --from=builder /app/main /var/task/main

# Definir el comando de ejecución
CMD ["./main"]
