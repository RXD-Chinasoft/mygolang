
FROM scratch

#RUN mkdir -p /go/src/mygolang

ENV PORT 8000
EXPOSE $PORT

COPY . .
CMD ["/models"]