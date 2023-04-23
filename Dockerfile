ARG BASE_IMAGE=golang:1.16-alpine 

# build stage
FROM ${BASE_IMAGE} AS build 
WORKDIR /app 
COPY . .
RUN go mod download 
RUN go mod vendor
RUN go build -o bin/main main.go

# main stage
FROM alpine  
WORKDIR /app 
COPY --from=build /app /app/

RUN apk add --no-cache bash

EXPOSE 5045

CMD ["/app/bin/main"]
