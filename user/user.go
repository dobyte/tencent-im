/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/30 4:23 下午
 * @Desc: 用户
 */

package user

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	
	"github.com/dobyte/tencent-im/internal/core"
)

type (
	// 性别类型
	GenderType string
	
	// 加好友验证方式
	AllowType string
	
	// 管理员禁止加好友标识类型
	AdminForbidType string
)

const (
	// 标准资料字段
	StandardAttrNickname        = "Tag_Profile_IM_Nick"            // 昵称
	StandardAttrGender          = "Tag_Profile_IM_Gender"          // 性别
	StandardAttrBirthday        = "Tag_Profile_IM_BirthDay"        // 生日
	StandardAttrLocation        = "Tag_Profile_IM_Location"        // 所在地
	StandardAttrSignature       = "Tag_Profile_IM_SelfSignature"   // 个性签名
	StandardAttrAllowType       = "Tag_Profile_IM_AllowType"       // 加好友验证方式
	StandardAttrLanguage        = "Tag_Profile_IM_Language"        // 语言
	StandardAttrAvatar          = "Tag_Profile_IM_Image"           // 头像URL
	StandardAttrMsgSettings     = "Tag_Profile_IM_MsgSettings"     // 消息设置
	StandardAttrAdminForbidType = "Tag_Profile_IM_AdminForbidType" // 管理员禁止加好友标识
	StandardAttrLevel           = "Tag_Profile_IM_Level"           // 等级
	StandardAttrRole            = "Tag_Profile_IM_Role"            // 角色
	
	// 自定义属性前缀
	CustomAttrPrefix = "Tag_Profile_Custom" // 自定义属性前缀
	
	// 性别类型
	GenderTypeUnknown GenderType = "Gender_Type_Unknown" // 没设置性别
	GenderTypeFemale  GenderType = "Gender_Type_Female"  // 女性
	GenderTypeMale    GenderType = "Gender_Type_Male"    // 男性
	
	// 加好友验证方式
	AllowTypeNeedConfirm AllowType = "AllowType_Type_NeedConfirm" // 需要经过自己确认对方才能添加自己为好友
	AllowTypeAllowAny    AllowType = "AllowType_Type_AllowAny"    // 允许任何人添加自己为好友
	AllowTypeDenyAny     AllowType = "AllowType_Type_DenyAny"     // 不允许任何人添加自己为好友
	
	// 管理员禁止加好友标识类型
	AdminForbidTypeNone    AdminForbidType = "AdminForbid_Type_None"    // 默认值，允许加好友
	AdminForbidTypeSendOut AdminForbidType = "AdminForbid_Type_SendOut" // 禁止该用户发起加好友请求
)

type User struct {
	userId string
	attrs  map[string]interface{}
	err    error
}

// InitAttrs 初始化属性
func (u *User) InitAttrs() {
	u.attrs = make(map[string]interface{})
}

// SetUserId 设置用户ID
func (u *User) SetUserId(userId string) {
	u.userId = userId
}

// GetUserId 获取用户ID
func (u *User) GetUserId() string {
	return u.userId
}

// SetNickname 设置昵称
func (u *User) SetNickname(nickname string) {
	u.SetAttr(StandardAttrNickname, nickname)
}

// GetNickname 获取昵称
func (u *User) GetNickname() (nickname string, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrNickname); exist {
		nickname = v.(string)
	}
	
	return
}

// SetGender 设置性别
func (u *User) SetGender(gender GenderType) {
	u.SetAttr(StandardAttrGender, gender)
}

// GetGender 获取性别
func (u *User) GetGender() (gender GenderType, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrGender); exist {
		gender = GenderType(v.(string))
	}
	
	return
}

// SetBirthday 设置生日
func (u *User) SetBirthday(birthday time.Time) {
	b, _ := strconv.Atoi(birthday.Format("20060102"))
	u.SetAttr(StandardAttrBirthday, b)
}

// GetBirthday 获取昵称
func (u *User) GetBirthday() (birthday time.Time, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrBirthday); exist {
		if val := v.(string); val != "" {
			birthday, _ = time.Parse("20060102", val)
		}
	}
	
	return
}

// SetLocation 所在地
func (u *User) SetLocation(country uint32, province uint32, city uint32, region uint32) {
	location := []uint32{country, province, city, region}
	buff := make([]byte, 4*len(location))
	for i, v := range location {
		buff[i*4+0] = byte(v / 256 / 256 / 256 % 256)
		buff[i*4+1] = byte(v / 256 / 256 % 256)
		buff[i*4+2] = byte(v / 256 % 256)
		buff[i*4+3] = byte(v % 256)
	}
	
	fmt.Println("国家，省，市，地区分别为：", country, province, city, region)
	fmt.Println("最终字节：", buff)
	fmt.Println("最终字符串：", strings.Split(string(buff), "\\0"))
	
	u.SetAttr(StandardAttrLocation, "abcdefghijklmlopq")
}

func (u *User) GetLocation() (country uint32, province uint32, city uint32, region uint32, exist bool) {
	fmt.Println(u.GetAttr(StandardAttrLocation))
	
	//var v interface{}
	//if v, exist = p.attrs[AttrLocation]; exist {
	//	fmt.Println(p.attrs)
	//	b := []byte(v.(string))
	//	fmt.Println(b)
	//
	//}
	
	return
}

// SetSignature 设置个性签名
func (u *User) SetSignature(signature string) {
	u.SetAttr(StandardAttrSignature, signature)
}

// GetSignature 获取个性签名
func (u *User) GetSignature() (signature string, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrSignature); exist {
		signature = v.(string)
	}
	
	return
}

// SetAllowType 设置加好友验证方式
func (u *User) SetAllowType(allowType AllowType) {
	u.SetAttr(StandardAttrAllowType, allowType)
}

// GetSignature 获取个性签名
func (u *User) GetAllowType() (allowType AllowType, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrAllowType); exist {
		allowType = AllowType(v.(string))
	}
	
	return
}

// SetLanguage 设置语言
func (u *User) SetLanguage(language uint) {
	u.SetAttr(StandardAttrLanguage, language)
}

// GetLanguage 获取语言
func (u *User) GetLanguage() (language uint, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrLanguage); exist {
		language = uint(v.(float64))
	}
	
	return
}

// SetAvatar 设置头像URL
func (u *User) SetAvatar(avatar string) {
	u.SetAttr(StandardAttrAvatar, avatar)
}

// GetAvatar 获取头像URL
func (u *User) GetAvatar() (avatar string, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrAvatar); exist {
		avatar = v.(string)
	}
	
	return
}

// SetMsgSettings 设置消息设置
func (u *User) SetMsgSettings(settings uint) {
	u.SetAttr(StandardAttrMsgSettings, settings)
}

// GetMsgSettings 获取消息设置
func (u *User) GetMsgSettings() (settings uint, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrMsgSettings); exist {
		settings = uint(v.(float64))
	}
	
	return
}

// SetAdminForbidType 设置管理员禁止加好友标识
func (u *User) SetAdminForbidType(forbidType AdminForbidType) {
	u.SetAttr(StandardAttrAdminForbidType, forbidType)
}

// GetAdminForbidType 获取管理员禁止加好友标识
func (u *User) GetAdminForbidType() (forbidType AdminForbidType, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrAdminForbidType); exist {
		forbidType = AdminForbidType(v.(string))
	}
	
	return
}

// SetLevel 设置等级
func (u *User) SetLevel(level uint) {
	u.SetAttr(StandardAttrLevel, level)
}

// GetLevel 获取等级
func (u *User) GetLevel() (level uint, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrLevel); exist {
		level = uint(v.(float64))
	}
	
	return
}

// SetRole 设置角色
func (u *User) SetRole(role uint) {
	u.SetAttr(StandardAttrRole, role)
}

// GetRole 获取角色
func (u *User) GetRole() (role uint, exist bool) {
	var v interface{}
	if v, exist = u.GetAttr(StandardAttrRole); exist {
		role = uint(v.(float64))
	}
	
	return
}

// SetCustomAttr 设置自定义属性
func (u *User) SetCustomAttr(name string, value interface{}) {
	u.SetAttr(fmt.Sprintf("%s_%s", CustomAttrPrefix, name), value)
}

// GetCustomAttr 获取自定义属性
func (u *User) GetCustomAttr(name string) (val interface{}, exist bool) {
	val, exist = u.GetAttr(fmt.Sprintf("%s_%s", CustomAttrPrefix, name))
	return
}

// IsValid 检测用户是否有效
func (u *User) IsValid() bool {
	return u.err == nil
}

// SetError 设置异常错误
func (u *User) SetError(code int, message string) {
	if code != core.SuccessCode {
		u.err = core.NewError(code, message)
	}
}

// GetError 获取异常错误
func (u *User) GetError() error {
	return u.err
}

// SetAttr 设置属性
func (u *User) SetAttr(name string, value interface{}) {
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
