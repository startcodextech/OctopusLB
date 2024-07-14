FROM golang:latest AS base

ENV GOPATH /go
ENV GO111MODULE=on

WORKDIR /app/src

COPY . .

FROM base AS dev

RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

EXPOSE 80
EXPOSE 443
EXPOSE 3000

CMD ["air"]

FROM base AS build

RUN go mod download

RUN go build -o /app/bin/main cmd/main.go

RUN chwon -R root:root /app/bin/main | chmod u+s /app/bin/main

FROM almalinux:latest AS rhel

COPY --from=build /app/bin/main /app/bin/main

CMD ["/app/bin/main"]