/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/8/31 15:14
 * @Desc: TODO
 */

package example

import (
	"fmt"

	"github.com/dobyte/tencent-im"
	"github.com/dobyte/tencent-im/account"
)

func main() {
	tim := im.NewIM(&im.Options{
		AppId:     1400579830,                                                         // 无效的AppId,请勿直接使用
		AppSecret: "0d2a321b087fdb8fd5ed5ea14fe0489139086eb1b03541283fc9feeab8f2bfd3", // 无效的AppSecret,请勿直接使用
		UserId:    "administrator",                                                    // 管理员用户账号，请在腾讯云IM后台设置管理账号
	})

	// 导入账号
	if err := tim.Account().ImportAccount(&account.Account{
		UserId:   "test1",
		Nickname: "测试账号1",
		FaceUrl:  "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
	}); err != nil {
		if e, ok := err.(im.Error); ok {
			fmt.Println(fmt.Sprintf("import account failed, code:%d, message:%s.", e.Code(), e.Message()))
		} else {
			fmt.Println(fmt.Sprintf("import account failed:%s.", err.Error()))
		}
	}

	fmt.Println("import account success.")
}
