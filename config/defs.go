package config

import (
	"bytes"
	"errors"
	"fmt"
	"git.zhugefang.com/gocore/zgo"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

//从.json文件中加载配置项

var Conf *allConfig

type allConfig struct {
	Env           string `json:"env"`
	Version       string `json:"version"`
	Project       string `json:"project"`
	EtcdHosts     string `json:"etcdHosts"`
	Loglevel      string `json:"loglevel"`
	RpcHost       string `json:"rpcHost"`
	RpcPort       string `json:"rpcPort"`
	ServerPort    int    `json:"serverPort"`
	UsePreAbsPath int    `json:"usePreAbsPath"`
}

func InitConfig(e, project, etcdHosts, port, rpcPort string) {
	initConfig(e, project, etcdHosts, port, rpcPort)
}

func initConfig(e, project, etcdHosts, port, rpcPort string) {
	var cf string
	if e == "local" {
		_, f, _, ok := runtime.Caller(1)
		if !ok {
			panic(errors.New("Can not get current file info"))
		}
		cf = fmt.Sprintf("%s/%s.json", filepath.Dir(f), e)

	} else {
		cf = fmt.Sprintf("./config/%s.json", e)
	}

	bf, err := ioutil.ReadFile(cf)
	if err != nil {
		panic(err)
	}

	//Conf = LoadConfig(cf)

	//使用zgo.Utils中的反序列化
	err = zgo.Utils.Unmarshal(bf, &Conf)
	if err != nil {
		panic(err)
	}

	if project != "" {
		Conf.Project = project
	}
	if etcdHosts != "" {
		Conf.EtcdHosts = etcdHosts
	}
	if port != "" {
		portInt, err := strconv.Atoi(port)
		if err != nil {
			zgo.Log.Error(err)
		} else {
			Conf.ServerPort = portInt
		}

	}
	if rpcPort != "" {
		Conf.RpcPort = rpcPort
	}

	fmt.Printf("origin %s is started on the ... %s\n", Conf.Version, Conf.Env)
}

// LoadConfigByFile暂时不用
func LoadConfigByFile(path string) *allConfig {
	var config allConfig
	configFile, err := os.Open(path)
	if err != nil {
		emit("Failed to open config file '%s': %s\n", path, err)
		return &config
	}

	fi, _ := configFile.Stat()
	if size := fi.Size(); size > (10 << 20) {
		emit("config file (%q) size exceeds reasonable limit (%d) - aborting", path, size)
		return &config // REVU: shouldn't this return an error, then?
	}

	if fi.Size() == 0 {
		emit("config file (%q) is empty, skipping", path)
		return &config
	}

	buffer := make([]byte, fi.Size())
	_, err = configFile.Read(buffer)
	//emit("\n %s\n", buffer)

	buffer, err = StripComments(buffer) //去掉注释
	if err != nil {
		emit("Failed to strip comments from json: %s\n", err)
		return &config
	}

	buffer = []byte(os.ExpandEnv(string(buffer))) //特殊

	err = zgo.Utils.Unmarshal(buffer, &config) //解析json格式数据
	if err != nil {
		emit("Failed unmarshalling json: %s\n", err)
		return &config
	}
	return &config
}

func StripComments(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0) // Windows
	lines := bytes.Split(data, []byte("\n"))                //split to muli lines
	filtered := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			filtered = append(filtered, line)
		}
	}

	return bytes.Join(filtered, []byte("\n")), nil
}

func emit(msgfmt string, args ...interface{}) {
	fmt.Printf(msgfmt, args...)
}
