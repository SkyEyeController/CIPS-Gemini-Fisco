FROM docker.1panel.live/library/golang:1.23

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app


CMD ["bash", "-c", "go mod tidy && go mod download && tail -f /dev/null"]

