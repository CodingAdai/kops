FROM alpine:latest
LABEL authors="daixiaodong"

WORKDIR /app

COPY server-go/server-go .

COPY dist dist

ENTRYPOINT ["./server-go"]
