存放proto文件夹生成的pb.go文件

首先cd到proto文件目录，执行下面一行，生成pb.go到pb目录下

cd proto
protoc --gofast_out=plugins=grpc:../pb helloworld.proto
