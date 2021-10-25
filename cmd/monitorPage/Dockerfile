FROM golang:alpine AS builder
WORKDIR /builder

COPY go.mod ./
COPY go.sum ./
COPY cmd/monitorPage/main.go ./
COPY mirroring /builder/mirroring
COPY internal /builder/internal

RUN go build -o monitorPage

FROM alpine
WORKDIR /monitor-page
COPY cmd/monitorPage/index.html ./
COPY --from=builder /builder/monitorPage ./
ENTRYPOINT ["/monitor-page/monitorPage"]
EXPOSE 8000
