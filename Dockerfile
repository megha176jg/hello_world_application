FROM golang:1.22.3-alpine
WORKDIR /app
COPY helloworld /app
RUN chmod +x helloworld
EXPOSE 8080
CMD ["./hello_world"]