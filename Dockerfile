ARG GO_VERSION=1.22
FROM node:20-bookworm as nodejs

COPY internal /usr/src/app/internal
COPY webapp /usr/src/app/webapp

WORKDIR /usr/src/app/webapp
RUN npm install
RUN npm run tailwindcss_build
RUN npm run build

WORKDIR /usr/src/app
RUN rm -rf webapp

RUN ls /usr/src/app

FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY ./cmd /usr/src/app/cmd/
COPY --from=nodejs /usr/src/app/internal ./internal
RUN go build -C /usr/src/app/cmd/web -v -o /goapp main.go
RUN chmod a+x /goapp

FROM debian:bookworm

COPY --from=builder /goapp /goapp

CMD ["/goapp"]
