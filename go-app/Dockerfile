# definir la imagen base de Golang
FROM golang:alpine

# Instalar git
RUN apk update && apk add --no-cache git

# localizacion de nuestros archivos en el docker container
WORKDIR /opt/go-app

# Copiar la fuente desde nuestro directorio actual hasta el directorio de trabajo dentro del contenedor
# La fuente también contiene go.mod y go.sum, que son archivos de dependencia
COPY . .

# Obtener dependencias
RUN go mod download

# Instalar Air para hot reload
RUN go get -u github.com/cosmtrek/air

# 
# El ENTRYPOINT define el comando que se ejecutará cuando se inicie el contenedor.
# En este caso, air command para hot reload, vaya a aplicaciones en cambios de archivo
ENTRYPOINT air