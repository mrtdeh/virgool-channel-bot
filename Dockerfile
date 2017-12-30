# build stage
FROM golang AS build-env
ADD . /go/src/telebot
RUN cd /go \
	&& go get gopkg.in/telegram-bot-api.v4 \
	&& go get github.com/PuerkitoBio/goquery \
	&& CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o test telebot 


# final stage
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /go/test /app
ENTRYPOINT ./test
