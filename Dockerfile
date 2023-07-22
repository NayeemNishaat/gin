FROM golang:latest

LABEL maintainer="Nayeem <NayeemNishaat@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=5000

RUN go build
RUN find . -name "*.go" -type f -delete

EXPOSE $PORT

CMD [ "./gin" ]