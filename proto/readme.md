proto文件夹生成的proto源文件

生成pb.go到pb目录下

确保在当前项目的目录下，比如 git.zhugefang.com/gobase/base-to-base-wait-copy/

protoc --proto_path=proto --gofast_out=plugins=grpc:pb helloworld/helloworld.proto





##使用 gogo插件  any.proto 或更多 替换下面的 jgpush/jgpush.proto
protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=grpc:pb --proto_path=proto base-to-base-wait-copy/base-to-base-wait-copy.proto