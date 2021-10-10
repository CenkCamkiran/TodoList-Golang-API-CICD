FROM golang:latest

RUN apt update
RUN apt-get install build-essential -y

WORKDIR /todoapi

COPY . .

#RUN go get -d github.com/gorilla/mux
#RUN go get -d github.com/stretchr/testify
#RUN go get -d go.mongodb.org/mongo-driver/mongo

RUN go mod vendor

CMD ["go","run","todo_api.go"]
