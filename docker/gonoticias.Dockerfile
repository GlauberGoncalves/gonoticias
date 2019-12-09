FROM golang

LABEL maintainer="Glauber <sistemas.glauber@gmail.com>"

WORKDIR /app/src/gonoticias
ENV GOPATH=/app
COPY . /app/src/gonoticias
RUN go build main.go
ENTRYPOINT ["./main"]
EXPOSE 8080