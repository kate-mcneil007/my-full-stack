# my-full-stack

go build -ldflags "-s -w" -o api-backend.bin ./cmd/main.go
^ use when changes are made 

separate terminal cd into deployments
docker-compose up --build --remove-orphans --force-recreate