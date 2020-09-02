FROM centos:7

COPY bin/crawl_list /app/
COPY config.json /app/config/

RUN chmod 777 /app/crawl_list

WORKDIR /app

ENTRYPOINT ["./crawl_list"]