FROM centos:7

ADD bin/elastic /app/
ADD config.json /app/config/

RUN chmod 777 /app/elastic

WORKDIR /app

ENTRYPOINT ["./elastic"]