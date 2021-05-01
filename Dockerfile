FROM golang:1.16 AS build

WORKDIR         /go/src/app
ADD .           /go/src/app

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app

FROM gcr.io/distroless/base AS runtime

COPY --from=build /go/bin/app /

CMD ["/app"]
