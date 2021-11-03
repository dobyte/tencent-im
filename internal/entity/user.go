/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/30 4:23 下午
 * @Desc: 用户
 */

package entity

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/internal/enum"
	"github.com/dobyte/tencent-im/internal/types"
)

type User struct {
	userId string
	attrs  map[string]interface{}
	err    error
}

// SetUserId 设置用户账号
func (u *User) SetUserId(userId string) {
	u.userId = userId
}

// GetUserId 获取用户账号
func (u *User) GetUserId() string {
	return u.userId
}

// SetNickname 设置昵称
func (u *User) SetNickname(nickname string) {
	u.SetAttr(enum.StandardAttrNickname, nickname)
}

// GetNickname 获取昵称
func (u *User) GetNickname() (nickname string, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrNickname); exist {
		nickname = v.(string)
	}

	return
}

// SetGender 设置性别
func (u *User) SetGender(gender types.GenderType) {
	u.SetAttr(enum.StandardAttrGender, gender)
}

// GetGender 获取性别
func (u *User) GetGender() (gender types.GenderType, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrGender); exist {
		gender = types.GenderType(v.(string))
	}

	return
}

// SetBirthday 设置生日
func (u *User) SetBirthday(birthday time.Time) {
	b, _ := strconv.Atoi(birthday.Format("20060102"))
	u.SetAttr(enum.StandardAttrBirthday, b)
}

// GetBirthday 获取昵称
func (u *User) GetBirthday() (birthday time.Time, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrBirthday); exist {
		if val := v.(string); val != "" {
			birthday, _ = time.Parse("20060102", val)
		}
	}

	return
}

// SetLocation 设置所在地
func (u *User) SetLocation(country uint32, province uint32, city uint32, region uint32) {
	var (
		str      string
		location = []uint32{country, province, city, region}
		builder  strings.Builder
	)

	builder.Grow(16)

	for _, v := range location {
		str = strconv.Itoa(int(v))

		if len(str) > 4 {
			u.SetError(enum.InvalidParamsCode, "invalid location params")
			break
		}

		builder.WriteString(strings.Repeat("0", 4-len(str)))
		builder.WriteString(str)
	}

	u.SetAttr(enum.StandardAttrLocation, builder.String())
}

// GetLocation 获取所在地
func (u *User) GetLocation() (country uint32, province uint32, city uint32, region uint32, exist bool) {
	var v interface{}

	if v, exist = u.attrs[enum.StandardAttrLocation]; exist {
		str := v.(string)

		if len(str) != 16 {
			exist = false
			return
		}

		if c, err := strconv.Atoi(str[0:4]); err != nil || c < 0 {
			exist = false
			return
		} else {
			country = uint32(c)
		}

		if c, err := strconv.Atoi(str[4:8]); err != nil || c < 0 {
			exist = false
			return
		} else {
			province = uint32(c)
		}

		if c, err := strconv.Atoi(str[8:12]); err != nil || c < 0 {
			exist = false
			return
		} else {
			city = uint32(c)
		}

		if c, err := strconv.Atoi(str[12:16]); err != nil || c < 0 {
			exist = false
			return
		} else {
			region = uint32(c)
		}
	}

	return
}

// SetSignature 设置个性签名
func (u *User) SetSignature(signature string) {
	u.SetAttr(enum.StandardAttrSignature, signature)
}

// GetSignature 获取个性签名
func (u *User) GetSignature() (signature string, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrSignature); exist {
		signature = v.(string)
	}

	return
}

// SetAllowType 设置加好友验证方式
func (u *User) SetAllowType(allowType types.AllowType) {
	u.SetAttr(enum.StandardAttrAllowType, allowType)
}

// GetAllowType 获取加好友验证方式
func (u *User) GetAllowType() (allowType types.AllowType, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrAllowType); exist {
		allowType = types.AllowType(v.(string))
	}

	return
}

// SetLanguage 设置语言
func (u *User) SetLanguage(language uint) {
	u.SetAttr(enum.StandardAttrLanguage, language)
}

// GetLanguage 获取语言
func (u *User) GetLanguage() (language uint, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrLanguage); exist {
		language = uint(v.(float64))
	}

	return
}

// SetAvatar 设置头像URL
func (u *User) SetAvatar(avatar string) {
	u.SetAttr(enum.StandardAttrAvatar, avatar)
}

// GetAvatar 获取头像URL
func (u *User) GetAvatar() (avatar string, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrAvatar); exist {
		avatar = v.(string)
	}

	return
}

// SetMsgSettings 设置消息设置
func (u *User) SetMsgSettings(settings uint) {
	u.SetAttr(enum.StandardAttrMsgSettings, settings)
}

// GetMsgSettings 获取消息设置
func (u *User) GetMsgSettings() (settings uint, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrMsgSettings); exist {
		settings = uint(v.(float64))
	}

	return
}

// SetAdminForbidType 设置管理员禁止加好友标识
func (u *User) SetAdminForbidType(forbidType types.AdminForbidType) {
	u.SetAttr(enum.StandardAttrAdminForbidType, forbidType)
}

// GetAdminForbidType 获取管理员禁止加好友标识
func (u *User) GetAdminForbidType() (forbidType types.AdminForbidType, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrAdminForbidType); exist {
		forbidType = types.AdminForbidType(v.(string))
	}

	return
}

// SetLevel 设置等级
func (u *User) SetLevel(level uint) {
	u.SetAttr(enum.StandardAttrLevel, level)
}

// GetLevel 获取等级
func (u *User) GetLevel() (level uint, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrLevel); exist {
		level = uint(v.(float64))
	}

	return
}

// SetRole 设置角色
func (u *User) SetRole(role uint) {
	u.SetAttr(enum.StandardAttrRole, role)
}

// GetRole 获取角色
func (u *User) GetRole() (role uint, exist bool) {
	var v interface{}

	if v, exist = u.GetAttr(enum.StandardAttrRole); exist {
		role = uint(v.(float64))
	}

	return
}

// SetCustomAttr 设置自定义属性
func (u *User) SetCustomAttr(name string, value interface{}) {
	u.SetAttr(fmt.Sprintf("%s_%s", enum.CustomAttrPrefix, name), value)
}

// GetCustomAttr 获取自定义属性
func (u *User) GetCustomAttr(name string) (val interface{}, exist bool) {
	val, exist = u.GetAttr(fmt.Sprintf("%s_%s", enum.CustomAttrPrefix, name))
	return
}

// IsValid 检测用户是否有效
func (u *User) IsValid() bool {
	return u.err == nil
}

// SetError 设置异常错误
func (u *User) SetError(code int, message string) {
	if code != enum.SuccessCode {
		u.err = core.NewError(code, message)
	}
}

// GetError 获取异常错误
func (u *User) GetError() error {
	return u.err
}

// SetAttr 设置属性
func (u *User) SetAttr(name string, value interface{}) {
	if u.attrs == nil {
		u.attrs = make(map[string]interface{})
	}
	u.attrs[name] = value
}

// GetAttr 获取属性
func (u *User) GetAttr(name string) (value interface{}, exist bool) {
	value, exist = u.attrs[name]
	return
}

// GetAttrs 获取所有属性
func (u *User) GetAttrs() map[string]interface{} {
	return u.attrs
}
