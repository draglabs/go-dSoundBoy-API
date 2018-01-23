FROM golang:1.9.0




RUN mkdir -p /go/src/dsound
WORKDIR /go/src/dsound

COPY . /go/src/dsound

# download any thirdparty packages
RUN go-wrapper download
RUN go-wrapper install


expose 8080

CMD ["go-wrapper", "run"]