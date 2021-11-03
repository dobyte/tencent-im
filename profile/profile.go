/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/28 11:23 上午
 * @Desc: TODO
 */

package profile

import (
	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/internal/entity"
	"github.com/dobyte/tencent-im/internal/enum"
)

type Profile struct {
	entity.User
}

func NewProfile(userId ...string) *Profile {
	p := &Profile{}
	if len(userId) > 0 {
		p.SetUserId(userId[0])
	}
	return p
}

// CheckError 检测错误
func (p *Profile) CheckError() (err error) {
	if userId := p.GetUserId(); userId == "" {
		return core.NewError(enum.InvalidParamsCode, "the userid is not set")
	}

	if err = p.GetError(); err != nil {
		return
	}

	return
}
