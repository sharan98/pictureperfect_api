FROM golang as build

WORKDIR /go/src/app
COPY . .

RUN go get -d github.com/gorilla/mux

RUN CGO_ENABLED=0 go build -o /go/src/app/movies .

FROM alpine

WORKDIR /home/app

COPY --from=build /go/src/app/movies .
EXPOSE 8000

CMD [ "/home/app/movies" ]