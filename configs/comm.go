package configs

import (
  "fmt"
  "github.com/gitcpu-io/zgo"
)

/*
@Time : 2019-09-04 19:07
@Author : zhuhonglei
@File : comm
@project: origin
*/

//从这里定义整个项目使用的公共变量

var (
  ComMap = zgo.Map.New()
)

func Goodbye() {
  goodbye := `
                ##### | #####
Oh we finish ? # _ _ #|# _ _ #
               #      |      #
         |       ############
                     # #
  |                  # #
                    #   #
         |     |    #   #      |        |
  |  |             #     #               |
         | |   |   # .-. #         |
                   #( O )#    |    |     |
  |  ################. .###############  |
   ##  _ _|____|     ###     |_ __| _  ##
  #  |                                |  #
  #  |    |    |    |   |    |    |   |  #
   ######################################
                   #     #
                    #####
`

  fmt.Println(goodbye)
}
