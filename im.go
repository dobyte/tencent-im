/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:29
 * @Desc: 腾讯云IM
 */

package im

import (
	"github.com/dobyte/tencent-im/account"
	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/mute"
	"github.com/dobyte/tencent-im/operation"
	"github.com/dobyte/tencent-im/private"
	"github.com/dobyte/tencent-im/profile"
	"github.com/dobyte/tencent-im/push"
	"github.com/dobyte/tencent-im/sns"
)

type Error = core.Error

type (
	IM interface {
		// SNS 获取关系链管理接口
		SNS() sns.API
		// Mute 获取全局禁言管理接口
		Mute() mute.API
		// Push 获取全员推送接口
		Push() push.Push
		// Account 获取账号管理接口
		Account() account.API
		// Profile 获取资料管理接口
		Profile() profile.API
		// Private 获取私聊消息接口
		Private() private.API
		// Operation 获取运营管理接口
		Operation() operation.API
	}

	Options struct {
		AppId     int    // 应用 SDKAppID，可在即时通信 IM 控制台 的应用卡片中获取。
		AppSecret string // 密钥信息，可在即时通信 IM 控制台 的应用详情页面中获取，具体操作请参见 获取密钥
		UserId    string // 用户ID
	}

	im struct {
		client    core.Client
		appId     int
		appSecret string
	}
)

// NewIM create a im instance.
func NewIM(opt Options) IM {
	return &im{
		appId: opt.AppId,
		client: core.NewClient(core.Options{
			AppId:     opt.AppId,
			AppSecret: opt.AppSecret,
			UserId:    opt.UserId,
		}),
	}
}

// SNS 获取关系链管理接口
func (i *im) SNS() sns.API {
	return sns.NewAPI(i.client)
}

// Mute 获取全局禁言管理接口
func (i *im) Mute() mute.API {
	return mute.NewAPI(i.client)
}

// Push 获取全员推送接口
func (i *im) Push() push.Push {
	return push.NewPush(i.client)
}

// Account 获取账号管理接口
func (i *im) Account() account.API {
	return account.NewAPI(i.client)
}

// Profile 获取资料管理接口
func (i *im) Profile() profile.API {
	return profile.NewAPI(i.client)
}

// Private 获取私聊消息接口
func (i *im) Private() private.API {
	return private.NewAPI(i.client)
}

// Operation 获取运营管理接口
func (i *im) Operation() operation.API {
	return operation.NewAPI(i.client)
}

func (i *im) Callback() {

}
