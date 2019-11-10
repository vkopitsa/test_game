# builder
FROM golang:alpine as builder

ENV CGO_ENABLED=1

WORKDIR /go/src/github.com/v.kopitsa/test_game
#WORKDIR /app

RUN apk --update upgrade && \
   apk add git && \
   rm -rf /var/cache/apk/*

COPY . .

ENV GO111MODULE=on

RUN sh -c 'go build \ 
    -ldflags "-s -w" \ 
    -o /app cmd/server/*.go' 

# runner
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /
COPY --from=builder /app .
COPY --from=builder /go/src/github.com/v.kopitsa/test_game/front /front

# run it!
CMD ["./app", "server"]
