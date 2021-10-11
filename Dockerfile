#使用朱大仙儿 build的带有curl的apline
FROM rubinus/alpine-nb:v1.0
LABEL maintainer="rubinus.chu@mail.com"

#接受参数,把源代码git commitid放到镜像内
ARG BUILD
LABEL VERSION=$BUILD

#设置工作目录
WORKDIR /opt/origin/

#添加可执行文件
COPY _output/origin-linux-amd64 /opt/origin/
COPY entrypoint.sh /opt/origin/

#添加配置及html等
COPY config /opt/origin/config
COPY views /opt/origin/views
COPY public /opt/origin/public
COPY deploy /opt/origin/deploy

RUN ["chmod", "+x", "origin-linux-amd64"]
RUN ["chmod", "+x", "entrypoint.sh"]

#设置Web端口，一般不用更改
EXPOSE 80

#设置GRPC端口，一般不用更改
EXPOSE 50051

ENTRYPOINT ["/opt/origin/origin-linux-amd64","--env=container"]
