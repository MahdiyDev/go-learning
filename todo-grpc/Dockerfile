FROM golang:1.21.4
FROM postgres:16.1

ENV POSTGRES_DB=todo
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=1407

EXPOSE 5432

WORKDIR /todo-grpc

COPY . .

RUN go mod download

RUN make compile

EXPOSE 4040

CMD ["./server.out"]