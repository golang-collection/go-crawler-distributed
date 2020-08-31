FROM centos:7

COPY bin/crawlTags /app/
COPY config.json /app/config/

RUN chmod 777 /app/crawlTags

WORKDIR /app

ENTRYPOINT ["./crawlTags"]