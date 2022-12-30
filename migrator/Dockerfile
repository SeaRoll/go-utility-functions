FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /target/main.out ./auth-service/cmd/main.go

#------------------#

FROM golang:alpine

WORKDIR /

COPY ./auth-service/sql /sql
COPY --from=build /target/main.out .

RUN file="$(ls -l)" && echo $file

EXPOSE 8080
ENTRYPOINT ["/main.out"]