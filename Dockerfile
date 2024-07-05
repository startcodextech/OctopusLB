FROM golang:latest as base

ENV GOPATH /go
ENV GO111MODULE=on

WORKDIR /app/src

COPY . .

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]

FROM base as build

RUN go mod download

RUN go build -o /app/bin/main cmd/main.go

RUN chwon -R root:root /app/bin/main | chmod u+s /app/bin/main

FROM almalinux:latest as rhel

COPY --from=build /app/bin/main /app/bin/main

CMD ["/app/bin/main"]