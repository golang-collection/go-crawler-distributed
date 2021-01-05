FROM centos:7

COPY bin/crawl_urllist /app/
COPY config.json /app/config/

RUN chmod 777 /app/crawl_urllist

WORKDIR /app

ENTRYPOINT ["./crawl_urllist"]