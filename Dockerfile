FROM alpine:edge
LABEL maintainer="zhuhonglei@zhuge.com"

#设置国内镜像
RUN echo 'https://mirrors.ustc.edu.cn/alpine/edge/community' > /etc/apk/repositories
RUN echo 'https://mirrors.ustc.edu.cn/alpine/edge/main' >> /etc/apk/repositories
RUN echo 'https://mirrors.ustc.edu.cn/alpine/edge/testing' >> /etc/apk/repositories
RUN apk update

#访问外部的https及安装curl
ENV CURL_VERSION 7.67.0

RUN apk add --update --no-cache openssl openssl-dev nghttp2-dev ca-certificates
RUN apk add --update --no-cache --virtual curldeps g++ make perl && \
wget https://curl.haxx.se/download/curl-$CURL_VERSION.tar.bz2 && \
tar xjvf curl-$CURL_VERSION.tar.bz2 && \
rm curl-$CURL_VERSION.tar.bz2 && \
cd curl-$CURL_VERSION && \
./configure \
    --with-nghttp2=/usr \
    --prefix=/usr \
    --with-ssl \
    --enable-ipv6 \
    --enable-unix-sockets \
    --without-libidn \
    --disable-static \
    --disable-ldap \
    --with-pic && \
make && \
make install && \
cd / && \
rm -r curl-$CURL_VERSION && \
rm -r /var/cache/apk && \
rm -r /usr/share/man && \
apk del curldeps

#设置东八区，北京时间
ENV TZ=Asia/Shanghai
RUN apk add --no-cache tzdata && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


#设置工作目录
WORKDIR $GOPATH/src/origin/

#添加可执行文件
ADD origin $GOPATH/src/origin/
ADD entrypoint.sh $GOPATH/src/origin/

#添加配置及html等
ADD config $GOPATH/src/origin/config
ADD views $GOPATH/src/origin/views
ADD public $GOPATH/src/origin/public
ADD deploy $GOPATH/src/origin/deploy


RUN ["chmod", "+x", "origin"]
RUN ["chmod", "+x", "entrypoint.sh"]

#设置Web端口，一般不用更改
EXPOSE 80

#设置GRPC端口，一般不用更改
EXPOSE 50051

ENTRYPOINT ["./entrypoint.sh"]