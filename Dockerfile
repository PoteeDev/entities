FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
# COPY auth /auth
RUN go mod download
COPY registration registration
COPY vpn vpn
COPY handlers handlers
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /teams .

##
## Deploy
##
FROM alpine
WORKDIR /
COPY --from=build /teams .
ENV PORT=8080
ENTRYPOINT [ "./teams"]
