FROM golang:1.19.5 AS build

WORKDIR /workdir
COPY go.mod ./
RUN go mod download -x
COPY . .
RUN go build -o bin/loop

FROM golang:1.19.5-alpine3.17
RUN apk add --no-cache libc6-compat
COPY --from=build /workdir/bin/loop /loop
EXPOSE 8090
CMD ["/loop"]