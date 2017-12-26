FROM alpine
RUN mkdir /telebot
ADD ./telebot /telebot/
CMD ["/telebot"]