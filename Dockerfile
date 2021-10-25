FROM alpine:3.7
LABEL Maintainer="AlicFeng <a@samego.com>" \
      Description="gin_api based on alpine"

ENV PATH "$PATH:/usr/local/samego/bin"

WORKDIR /usr/local/samego/

COPY bin /usr/local/samego/bin

RUN chmod a+x /usr/local/samego/bin/* && \
    mkdir -p /usr/local/samego/log

CMD ["server"]