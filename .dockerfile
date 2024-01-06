FROM golang:latest

WORKDIR /restAPi
COPY . .

RUN go mod download
RUN go build -o yourappname .

EXPOSE 8080

CMD ["./restAPi"]
