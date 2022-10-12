# syntax=docker/dockerfile:1

FROM golang:1.19-bullseye as build
WORKDIR /app
COPY backend/. /app
COPY moneyapp-3615b-firebase-adminsdk-nnqws-3e0b8d838e.json /app
 
RUN go mod download
# RUN go build -o /docker-moneyapp
RUN go build -o /docker-moneyapp /app/cmd/server/main.go

CMD [ "/docker-moneyapp" ]