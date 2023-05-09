FROM golang:alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /mini-project

EXPOSE 8080

CMD [ "/mini-project" ]