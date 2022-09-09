# syntax=docker/dockerfile:1

FROM golang:1.19-bullseye as build
WORKDIR /app
COPY server/. /app
COPY moneyapp-3615b-firebase-adminsdk-g2mab-af1c8228b3.json /app
 
RUN go mod download
RUN go build -o /docker-moneyapp

CMD [ "/docker-moneyapp" ]