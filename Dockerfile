FROM alpine:latest
WORKDIR /app
RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*
COPY ./dfns_linux ./dfns
RUN chmod +x dfns
EXPOSE 9000
ENTRYPOINT [ "./dfns", "--addr=:9000" ]