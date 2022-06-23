FROM golang:latest

WORKDIR /usr/src/server

RUN apt-get update && apt-get -y install mariadb-client && apt-get -y install git

#RUN go mod tidy && go install github.com/cosmtrek/air@latest

#CMD ["air", "-c", ".air.toml"]
CMD ["go", "run", "main.go"]