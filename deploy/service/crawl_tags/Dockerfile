FROM centos:7

COPY bin/crawl_tags /app/
COPY config.json /app/config/

RUN chmod 777 /app/crawl_tags

WORKDIR /app

ENTRYPOINT ["./crawl_tags"]