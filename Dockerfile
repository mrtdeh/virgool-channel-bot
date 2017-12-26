FROM alpine
RUN mkdir /telebot
COPY ./telebot /telebot/
WORKDIR /telebot
CMD ["telebot"]