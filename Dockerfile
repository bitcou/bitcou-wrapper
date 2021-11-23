FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV PORT 8000
# ENV BITCOU_KEY f54fef0b-eb61-40e9-93f5-ed8c1260b4ea
ENV BLOCKCHAIN_KEY passwordWithLove

RUN go build
CMD [ "./bitcou-wrapper" ]

