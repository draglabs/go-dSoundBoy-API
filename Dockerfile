FROM golang:1.9.0




RUN mkdir -p /go/src/dsound
WORKDIR /go/src/dsound

COPY . /go/src/dsound

 ENV AWS_ACCESS_KEY_ID="AKIAIHIPQHLGCPDR44HQ"
 ENV AWS_SECRET_ACCESS_KEY="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"
 ENV AWS_REGION="us-west-1"
# download any thirdparty packages
RUN go-wrapper download
RUN go-wrapper install


expose 8080

CMD ["go-wrapper", "run"]