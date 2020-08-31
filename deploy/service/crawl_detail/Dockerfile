FROM centos:7

COPY bin/crawl_detail /app/
COPY config.json /app/config/

RUN chmod 777 /app/crawl_detail

WORKDIR /app

ENTRYPOINT ["./crawl_detail"]