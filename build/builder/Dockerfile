FROM golang:1.16.7-stretch AS build-env 
ENV GO111MODULE=on
ADD . /go/src/app/
WORKDIR /go/src/app
RUN go get -d -v cmd/basicapp
RUN go install -v cmd/basicapp
RUN GOOS=linux GOARCH=amd64 go build -v -o /go/src/app/app-server cmd/basicapp

FROM debian:9-slim
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Taipei

RUN apt-get update \
    && apt-get install -y --no-install-recommends tzdata
    
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata 
    
COPY --from=build-env /go/src/app/app-server /usr/local/bin/app-server
EXPOSE 5000
ENTRYPOINT ["app-server", "daemon"]