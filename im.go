/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:29
 * @Desc: 腾讯云IM
 */

package im

import (
	"sync"
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
	"github.com/dobyte/tencent-im/recentcontact"
	"github.com/dobyte/tencent-im/sns"
)

type Error = core.Error

type (
	IM interface {
		// GetUserSig 获取UserSig签名
		GetUserSig(userId string, expiration ...int) UserSig
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
		// Operation 获取运营管理接口
		Operation() operation.API
		// RecentContact 获取最近联系人接口
		RecentContact() recentcontact.API
		// Callback 获取回调接口
		Callback() callback.Callback
	}

	Options struct {
		AppId      int    // 应用SDKAppID，可在即时通信 IM 控制台 的应用卡片中获取。
		AppSecret  string // 密钥信息，可在即时通信 IM 控制台 的应用详情页面中获取，具体操作请参见 获取密钥
		UserId     string // 用户ID
		Expiration int    // UserSig过期时间
	}

	UserSig struct {
		UserSig  string // 用户签名
		ExpireAt int64  // 签名过期时间
	}

	im struct {
		opt    *Options
		client core.Client
		sns    struct {
			once     sync.Once
			instance sns.API
		}
		mute struct {
			once     sync.Once
			instance mute.API
		}
		push struct {
			once     sync.Once
			instance push.API
		}
		group struct {
			once     sync.Once
			instance group.API
		}
		account struct {
			once     sync.Once
			instance account.API
		}
		profile struct {
			once     sync.Once
			instance profile.API
		}
		private struct {
			once     sync.Once
			instance private.API
		}
		operation struct {
			once     sync.Once
			instance operation.API
		}
		recentcontact struct {
			once     sync.Once
			instance recentcontact.API
		}
		callback struct {
			once     sync.Once
			instance callback.Callback
		}
	}
)

func NewIM(opt *Options) IM {
	return &im{opt: opt, client: core.NewClient(&core.Options{
		AppId:      opt.AppId,
		AppSecret:  opt.AppSecret,
		UserId:     opt.UserId,
		Expiration: opt.Expiration,
	})}
}

// GetUserSig 获取UserSig签名
func (i *im) GetUserSig(userId string, expiration ...int) UserSig {
	if len(expiration) == 0 {
		expiration = append(expiration, i.opt.Expiration)
	}

	userSig, _ := sign.GenUserSig(i.opt.AppId, i.opt.AppSecret, userId, expiration[0])
	expireAt := time.Now().Add(time.Duration(i.opt.Expiration) * time.Second).Unix()
	return UserSig{UserSig: userSig, ExpireAt: expireAt}
}

// SNS 获取关系链管理接口ok
func (i *im) SNS() sns.API {
	i.sns.once.Do(func() {
		i.sns.instance = sns.NewAPI(i.client)
	})
	return i.sns.instance
}

// Mute 获取全局禁言管理接口ok
func (i *im) Mute() mute.API {
	i.mute.once.Do(func() {
		i.mute.instance = mute.NewAPI(i.client)
	})
	return i.mute.instance
}

// Push 获取全员推送接口
func (i *im) Push() push.API {
	i.push.once.Do(func() {
		i.push.instance = push.NewAPI(i.client)
	})
	return i.push.instance
}

// Group 获取群组管理接口
func (i *im) Group() group.API {
	i.group.once.Do(func() {
		i.group.instance = group.NewAPI(i.client)
	})
	return i.group.instance
}

// Account 获取账号管理接口ok
func (i *im) Account() account.API {
	i.account.once.Do(func() {
		i.account.instance = account.NewAPI(i.client)
	})
	return i.account.instance
}

// Profile 获取资料管理接口ok
func (i *im) Profile() profile.API {
	i.profile.once.Do(func() {
		i.profile.instance = profile.NewAPI(i.client)
	})
	return i.profile.instance
}

// Private 获取私聊消息接口ok
func (i *im) Private() private.API {
	i.private.once.Do(func() {
		i.private.instance = private.NewAPI(i.client)
	})
	return i.private.instance
}

// Operation 获取运营管理接口ok
func (i *im) Operation() operation.API {
	i.operation.once.Do(func() {
		i.operation.instance = operation.NewAPI(i.client)
	})
	return i.operation.instance
}

// RecentContact 获取最近联系人接口ok
func (i *im) RecentContact() recentcontact.API {
	i.recentcontact.once.Do(func() {
		i.recentcontact.instance = recentcontact.NewAPI(i.client)
	})
	return i.recentcontact.instance
}

// Callback 获取回调接口
func (i *im) Callback() callback.Callback {
	i.callback.once.Do(func() {
		i.callback.instance = callback.NewCallback(i.opt.AppId)
	})
	return i.callback.instance
}
