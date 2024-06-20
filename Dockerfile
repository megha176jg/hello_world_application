FROM alpine as builder
WORKDIR /app
COPY hello_world ./
RUN chmod +x hello_world
EXPOSE 8080
CMD ["./hello_world"]