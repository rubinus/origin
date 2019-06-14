proto文件夹生成的proto源文件

生成pb.go到pb目录下

确保在当前项目的目录下，比如 git.zhugefang.com/goymd/visource/

protoc --proto_path=proto --gofast_out=plugins=grpc:pb helloworld.proto

