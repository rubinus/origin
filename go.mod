module github.com/gitcpu-io/origin

go 1.16

require (
	github.com/0x19/goesl v0.0.0-20191107044804-3efcc2f41ccb
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/fiorix/go-eventsocket v0.0.0-20180331081222-a4a0ee7bd315
	github.com/gitcpu-io/zgo v1.0.2
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/kataras/iris/v12 v12.1.8
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/ryanuber/columnize v2.1.2+incompatible // indirect
	github.com/spf13/pflag v1.0.5
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
	golang.org/x/sys v0.0.0-20210817190340-bfb29a6856f2 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.40.0
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.2
	google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.26.0
//github.com/gitcpu-io/zgo => ../zgo
)
