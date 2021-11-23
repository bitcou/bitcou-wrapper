FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV PORT 8000
# ENV BITCOU_KEY 
# ENV BLOCKCHAIN_KEY passwordWithLove

RUN go build
CMD [ "./bitcou-wrapper" ]

