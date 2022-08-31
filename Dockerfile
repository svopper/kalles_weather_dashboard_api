FROM golang

ENV GO111MODULE=on
ENV DMI_MET_OBS_API_KEY=xxx
ENV GIN_MODE=release

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build app/main.go

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]