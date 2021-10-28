/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:29
 * @Desc: 腾讯云IM
 */

package im

import (
	"time"

	"github.com/dobyte/tencent-im/account"
	"github.com/dobyte/tencent-im/callback"
	"github.com/dobyte/tencent-im/group"
	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/internal/sign"
	"github.com/dobyte/tencent-im/mute"
	"github.com/dobyte/tencent-im/operation"
	"github.com/dobyte/tencent-im/private"
	"github.com/dobyte/tencent-im/profile"
	"github.com/dobyte/tencent-im/push"
	"github.com/dobyte/tencent-im/session"
	"github.com/dobyte/tencent-im/sns"
)

type Error = core.Error

type (
	IM interface {
		// GetUserSig 获取UserSig签名
		GetUserSig() UserSig
		// SNS 获取关系链管理接口
		SNS() sns.API
		// Mute 获取全局禁言管理接口
		Mute() mute.API
		// Push 获取全员推送接口
		Push() push.API
		// Group 获取群组管理接口
		Group() group.API
		// Account 获取账号管理接口
		Account() account.API
		// Profile 获取资料管理接口
		Profile() profile.API
		// Private 获取私聊消息接口
		Private() private.API
		// Session 获取最近联系人接口
		Session() session.API
		// Operation 获取运营管理接口
		Operation() operation.API
		// Callback 获取回调接口
		Callback() callback.Callback
	}

	Options struct {
		AppId     int    // 应用SDKAppID，可在即时通信 IM 控制台 的应用卡片中获取。
		AppSecret string // 密钥信息，可在即时通信 IM 控制台 的应用详情页面中获取，具体操作请参见 获取密钥
		UserId    string // 用户ID
		Expire    int    // UserSig过期时间
	}

	UserSig struct {
		UserSig  string // 用户签名
		ExpireAt int64  // 签名过期时间
	}

	im struct {
		opt    *Options
		client core.Client
	}
)

func NewIM(opt *Options) IM {
	return &im{opt: opt, client: core.NewClient(&core.Options{
		AppId:     opt.AppId,
		AppSecret: opt.AppSecret,
		UserId:    opt.UserId,
		Expire:    opt.Expire,
	})}
}

// GetUserSig 获取UserSig签名
func (i *im) GetUserSig() UserSig {
	userSig, _ := sign.GenUserSig(i.opt.AppId, i.opt.AppSecret, i.opt.UserId, i.opt.Expire)
	expireAt := time.Now().Add(time.Duration(i.opt.Expire) * time.Second).Unix()
	return UserSig{UserSig: userSig, ExpireAt: expireAt}
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
func (i *im) Push() push.API {
	return push.NewAPI(i.client)
}

// Group 获取群组管理接口
func (i *im) Group() group.API {
	return group.NewAPI(i.client)
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

// Session 获取最近联系人接口
func (i *im) Session() session.API {
	return session.NewAPI(i.client)
}

// Operation 获取运营管理接口
func (i *im) Operation() operation.API {
	return operation.NewAPI(i.client)
}

// Callback 获取回调接口
func (i *im) Callback() callback.Callback {
	return callback.NewCallback(i.opt.AppId)
}
