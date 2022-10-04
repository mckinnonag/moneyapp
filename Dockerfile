# syntax=docker/dockerfile:1

FROM golang:1.19-bullseye as build
WORKDIR /app
COPY server/. /app
COPY moneyapp-3615b-firebase-adminsdk-nnqws-3e0b8d838e.json /app
 
RUN go mod download
RUN go build -o /docker-moneyapp

# CMD [ "/docker-moneyapp" ]
ENTRYPOINT ["go", "test", "-v", "./...", "-coverprofile", "cover.out"]