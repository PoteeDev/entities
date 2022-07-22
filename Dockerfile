FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
# COPY auth /auth
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /teams .

##
## Deploy
##
FROM alpine
WORKDIR /
COPY --from=build /teams .
ENV PORT=8080
ENTRYPOINT [ "./teams"]
