FROM centos:7

COPY bin/storageDetail /app/
COPY config.json /app/config/

RUN chmod 777 /app/storageDetail

WORKDIR /app

ENTRYPOINT ["./storageDetail"]