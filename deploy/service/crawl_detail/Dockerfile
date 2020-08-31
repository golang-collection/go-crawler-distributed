FROM centos:7

COPY bin/crawlDetail /app/
COPY config.json /app/config/

RUN chmod 777 /app/crawlDetail

WORKDIR /app

ENTRYPOINT ["./crawlDetail"]