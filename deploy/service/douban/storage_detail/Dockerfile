FROM centos:7

COPY bin/storage_detail /app/
COPY config.json /app/config/

RUN chmod 777 /app/storage_detail

WORKDIR /app

ENTRYPOINT ["./storage_detail"]