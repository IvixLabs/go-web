ARG GO_VERSION=1.22
FROM golang:${GO_VERSION} as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -C /usr/src/app/cmd/web -v -o /goapp main.go
RUN chmod a+x /goapp

CMD ["/goapp"]
