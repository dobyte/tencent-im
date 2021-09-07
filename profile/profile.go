/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/28 11:23 上午
 * @Desc: TODO
 */

package profile

import (
    "github.com/dobyte/tencent-im/internal/entity"
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
