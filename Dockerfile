FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
# COPY auth /auth
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /entities .

##
## Deploy
##
FROM alpine
WORKDIR /
COPY --from=build /entities .
ENV PORT=8080
ENTRYPOINT [ "./entities"]
