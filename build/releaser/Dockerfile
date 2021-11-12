FROM debian:9-slim
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Taipei

RUN apt-get update \
    && apt-get install -y --no-install-recommends tzdata
    
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata 

ADD . /go/src/app/
WORKDIR /go/src/app

EXPOSE 5000
ENTRYPOINT ["app-server", "daemon"]