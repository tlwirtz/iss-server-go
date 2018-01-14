FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY iss-server-go /app/iss-server-go
CMD ["/app/iss-server-go"]