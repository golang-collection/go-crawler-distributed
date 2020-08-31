FROM centos:7

COPY bin/crawlList /app/
COPY config.json /app/config/

RUN chmod 777 /app/crawlList

WORKDIR /app

ENTRYPOINT ["./crawlList"]