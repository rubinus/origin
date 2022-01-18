package demo_crypto

import (
  "encoding/base64"
  "encoding/hex"
  "fmt"
  "github.com/gitcpu-io/origin/config"
  "github.com/gitcpu-io/zgo"
)

/*
@Time : 2019-03-15 11:49
@Author : rubinus.chu
@File : demo
@project: origin
*/

var (
  bt  = []byte("http://www.baidu.com")
  key = []byte("ZhugeAd")
  str = "zgo engine is niubility"
)

func Demo() {
  config.InitConfig("local", "", "", "", "")

  err := zgo.Engine(&zgo.Options{
    Env:      config.Conf.Env,
    Project:  config.Conf.Project,
    Loglevel: config.Conf.Loglevel,
  })

  if err != nil {
    panic(err)
  }

  md5 := zgo.Crypto.Md5(str)
  fmt.Printf("val:%s Md5:%s\n", str, md5)

  sha1 := zgo.Crypto.SHA1(str)
  fmt.Printf("val:%s SHA1:%s\n", str, sha1)

  sha256 := zgo.Crypto.SHA256String(str)
  fmt.Printf("val:%s SHA256String:%s\n", str, sha256)

  b, _ := zgo.Crypto.AESCFBEncrypt(bt, key)
  fmt.Printf("val:%s AESCFBEncrypt:%s\n", string(bt), hex.EncodeToString(b))

  b, _ = zgo.Crypto.AESCBCEncrypt(bt, key)
  fmt.Printf("val:%s AESCBCEncrypt:%s\n\n", string(bt), hex.EncodeToString(b))

  bs := zgo.Crypto.AesEncrypt(string(bt), "9871267812345mn812345xyz")
  fmt.Printf("val:%s AesEncrypt加密:%s\n", string(bt), bs)

  decryptCode := zgo.Crypto.AesDecrypt(bs, "9871267812345mn812345xyz")
  fmt.Println("AesEncrypt解密结果：", decryptCode)

  //====================
  var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDfw1/P15GQzGGYvNwVmXIGGxea
8Pb2wJcF7ZW7tmFdLSjOItn9kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQol
DOkEzNP0B8XKm+Lxy4giwwR5LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TU
GQD6QzKY1Y8xS+FoQQIDAQAB
-----END PUBLIC KEY-----
`)
  var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDfw1/P15GQzGGYvNwVmXIGGxea8Pb2wJcF7ZW7tmFdLSjOItn9
kvUsbQgS5yxx+f2sAv1ocxbPTsFdRc6yUTJdeQolDOkEzNP0B8XKm+Lxy4giwwR5
LJQTANkqe4w/d9u129bRhTu/SUzSUIr65zZ/s6TUGQD6QzKY1Y8xS+FoQQIDAQAB
AoGAbSNg7wHomORm0dWDzvEpwTqjl8nh2tZyksyf1I+PC6BEH8613k04UfPYFUg1
0F2rUaOfr7s6q+BwxaqPtz+NPUotMjeVrEmmYM4rrYkrnd0lRiAxmkQUBlLrCBiF
u+bluDkHXF7+TUfJm4AZAvbtR2wO5DUAOZ244FfJueYyZHECQQD+V5/WrgKkBlYy
XhioQBXff7TLCrmMlUziJcQ295kIn8n1GaKzunJkhreoMbiRe0hpIIgPYb9E57tT
/mP/MoYtAkEA4Ti6XiOXgxzV5gcB+fhJyb8PJCVkgP2wg0OQp2DKPp+5xsmRuUXv
720oExv92jv6X65x631VGjDmfJNb99wq5QJBAMSHUKrBqqizfMdOjh7z5fLc6wY5
M0a91rqoFAWlLErNrXAGbwIRf3LN5fvA76z6ZelViczY6sKDjOxKFVqL38ECQG0S
pxdOT2M9BM45GJjxyPJ+qBuOTGU391Mq1pRpCKlZe4QtPHioyTGAAMd4Z/FX2MKb
3in48c0UX5t3VjPsmY0CQQCc1jmEoB83JmTHYByvDpc8kzsD8+GmiPVrausrjj4p
y2DQpGmUic2zqCxl6qXMpBGtFEhrUbKhOiVOJbRNGvWW
-----END RSA PRIVATE KEY-----
`)
  data, _ := zgo.Crypto.RsaEncrypt([]byte(str), publicKey)
  fmt.Println()
  fmt.Println()
  fmt.Println("RsaEncrypt加密:", base64.StdEncoding.EncodeToString(data))
  origData, _ := zgo.Crypto.RsaDecrypt(data, privateKey)
  fmt.Println("RsaDecrypt解密:", string(origData))

}
