# rpi-dev
My development fun for my Pi

Exemple: `env GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "-X main.buildTime=$(date +"%Y.%m.%d.%H%M%S")" -o bin main.go`