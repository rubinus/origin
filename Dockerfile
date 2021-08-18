#使用朱大仙儿 build的带有curl的apline
FROM rubinus/alpine-curl
LABEL maintainer="rubinus.chu@mail.com"

#接受参数,把源代码git commitid放到镜像内
ARG BUILD
LABEL VERSION=$BUILD

#设置工作目录
WORKDIR /opt/origin/

#添加可执行文件
ADD _output/origin /opt/origin/
ADD entrypoint.sh /opt/origin/

#添加配置及html等
ADD config /opt/origin/config
ADD views /opt/origin/views
ADD public /opt/origin/public
ADD deploy /opt/origin/deploy

RUN ["chmod", "+x", "origin"]
RUN ["chmod", "+x", "entrypoint.sh"]

#设置Web端口，一般不用更改
EXPOSE 80

#设置GRPC端口，一般不用更改
EXPOSE 50051

ENTRYPOINT ["./entrypoint.sh"]