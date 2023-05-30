# my-full-stack

go build -ldflags "-s -w" -o api-backend.bin ./cmd/main.go
docker-compose up --build --remove-orphans --force-recreate