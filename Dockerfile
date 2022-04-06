FROM golang:buster

RUN useradd appUser

WORKDIR /code

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY go.mod ./
COPY go.sum ./

RUN go build ./cmd/twitter-countdown/main.go
RUN mkdir /twitter-countdown
RUN mv main /twitter-countdown
RUN chown -R appUser:appUser /twitter-countdown
RUN chmod -R 750 /twitter-countdown

USER appUser
WORKDIR /twitter-countdown

ENTRYPOINT ["./main"]
