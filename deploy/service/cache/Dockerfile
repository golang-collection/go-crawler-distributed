FROM centos:7

ADD bin/cache /app/
ADD config.json /app/config/

RUN chmod 777 /app/cache

WORKDIR /app

ENTRYPOINT ["./cache"]