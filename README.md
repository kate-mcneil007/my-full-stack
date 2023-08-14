# my-full-stack

use when changes are made 
go build -ldflags "-s -w" -o api-backend.bin ./cmd/main.go


- separate terminal 
- cd into deployments
- Uses dockerfile to spin up postgres db and runs api backend.bin 
docker-compose up --build --remove-orphans --force-recreate