#!/bin/bash
# Script para ejecutar mi aplicacion y pueda ser probada localmente
cd ./go-app


# GIN_MODE a produccion
#export GIN_MODE=release

# Mongo Configs
export ENV=production
export MONGO_URI=mongodb://localhost:27017
export PORT=8000
export HOST=http://localhost

go mod tidy

go run main.go
