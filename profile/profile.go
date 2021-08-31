/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/28 11:23 上午
 * @Desc: TODO
 */

package profile

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
	AttrNickname        = "Tag_Profile_IM_Nick"            // 昵称
	AttrGender          = "Tag_Profile_IM_Gender"          // 性别
	AttrBirthday        = "Tag_Profile_IM_BirthDay"        // 生日
	AttrLocation        = "Tag_Profile_IM_Location"        // 所在地
	AttrSignature       = "Tag_Profile_IM_SelfSignature"   // 个性签名
	AttrAllowType       = "Tag_Profile_IM_AllowType"       // 加好友验证方式
	AttrLanguage        = "Tag_Profile_IM_Language"        // 语言
	AttrAvatar          = "Tag_Profile_IM_Image"           // 头像URL
	AttrMsgSettings     = "Tag_Profile_IM_MsgSettings"     // 消息设置
	AttrAdminForbidType = "Tag_Profile_IM_AdminForbidType" // 管理员禁止加好友标识
	AttrLevel           = "Tag_Profile_IM_Level"           // 等级
	AttrRole            = "Tag_Profile_IM_Role"            // 角色
	CustomAttrPrefix    = "Tag_Profile_Custom"             // 自定义属性前缀
	
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

type profile struct {
	userId string
	attrs  map[string]interface{}
	err    error
}

func NewProfile(userId string) *profile {
	return &profile{userId: userId, attrs: make(map[string]interface{})}
}

// SetUserId 设置用户ID
func (p *profile) SetUserId(userId string) *profile {
	p.userId = userId
	return p
}

// GetNickname 获取昵称
func (p *profile) GetUserId() string {
	return p.userId
}

// SetNickname 设置昵称
func (p *profile) SetNickname(nickname string) *profile {
	return p.setAttr(AttrNickname, nickname)
}

// GetNickname 获取昵称
func (p *profile) GetNickname() (nickname string, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrNickname]; exist {
		nickname = v.(string)
	}
	
	return
}

// SetGender 设置性别
func (p *profile) SetGender(gender GenderType) *profile {
	return p.setAttr(AttrGender, gender)
}

// GetGender 获取性别
func (p *profile) GetGender() (gender GenderType, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrGender]; exist {
		gender = GenderType(v.(string))
	}
	
	return
}

// SetBirthday 设置生日
func (p *profile) SetBirthday(birthday time.Time) *profile {
	b, _ := strconv.Atoi(birthday.Format("20060102"))
	return p.setAttr(AttrBirthday, b)
}

// GetBirthday 获取昵称
func (p *profile) GetBirthday() (birthday time.Time, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrBirthday]; exist {
		if val := v.(string); val != "" {
			birthday, _ = time.Parse("20060102", val)
		}
	}
	
	return
}

// SetLocation 所在地
func (p *profile) SetLocation(country uint32, province uint32, city uint32, region uint32) *profile {
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
	fmt.Println("最终字符串：", strings.Split(string(buff) , "\\0"))
	
	
	return p.setAttr(AttrLocation, "abcdefghijklmlopq")
}

func (p *profile) GetLocation() (country uint32, province uint32, city uint32, region uint32, exist bool) {
	fmt.Println(p.attrs[AttrLocation])
	
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
func (p *profile) SetSignature(signature string) *profile {
	return p.setAttr(AttrSignature, signature)
}

// GetSignature 获取个性签名
func (p *profile) GetSignature() (signature string, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrSignature]; exist {
		signature = v.(string)
	}
	
	return
}

// SetAllowType 设置加好友验证方式
func (p *profile) SetAllowType(allowType AllowType) *profile {
	return p.setAttr(AttrAllowType, allowType)
}

// GetSignature 获取个性签名
func (p *profile) GetAllowType() (allowType AllowType, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrAllowType]; exist {
		allowType = AllowType(v.(string))
	}
	
	return
}

// SetLanguage 设置语言
func (p *profile) SetLanguage(language uint) *profile {
	return p.setAttr(AttrLanguage, language)
}

// GetLanguage 获取语言
func (p *profile) GetLanguage() (language uint, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrLanguage]; exist {
		language = uint(v.(float64))
	}
	
	return
}

// SetAvatar 设置头像URL
func (p *profile) SetAvatar(avatar string) *profile {
	return p.setAttr(AttrAvatar, avatar)
}

// GetAvatar 获取头像URL
func (p *profile) GetAvatar() (avatar string, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrAvatar]; exist {
		avatar = v.(string)
	}
	
	return
}

// SetMsgSettings 设置消息设置
func (p *profile) SetMsgSettings(settings uint) *profile {
	return p.setAttr(AttrMsgSettings, settings)
}

// GetMsgSettings 获取消息设置
func (p *profile) GetMsgSettings() (settings uint, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrMsgSettings]; exist {
		settings = uint(v.(float64))
	}
	
	return
}

// SetAdminForbidType 设置管理员禁止加好友标识
func (p *profile) SetAdminForbidType(forbidType AdminForbidType) *profile {
	return p.setAttr(AttrAdminForbidType, forbidType)
}

// GetAdminForbidType 获取管理员禁止加好友标识
func (p *profile) GetAdminForbidType() (forbidType AdminForbidType, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrAdminForbidType]; exist {
		forbidType = AdminForbidType(v.(string))
	}
	
	return
}

// SetLevel 设置等级
func (p *profile) SetLevel(level uint) *profile {
	return p.setAttr(AttrLevel, level)
}

// GetLevel 获取等级
func (p *profile) GetLevel() (level uint, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrLevel]; exist {
		level = uint(v.(float64))
	}
	
	return
}

// SetRole 设置角色
func (p *profile) SetRole(role uint) *profile {
	return p.setAttr(AttrRole, role)
}

// GetRole 获取角色
func (p *profile) GetRole() (role uint, exist bool) {
	var v interface{}
	if v, exist = p.attrs[AttrRole]; exist {
		role = uint(v.(float64))
	}
	
	return
}

// SetCustomAttr 设置自定义属性
func (p *profile) SetCustomAttr(name string, value interface{}) *profile {
	return p.setAttr(fmt.Sprintf("%s_%s", CustomAttrPrefix, name), value)
}

// GetCustomAttr 获取自定义属性
func (p *profile) GetCustomAttr(name string) (val interface{}, exist bool) {
	val, exist = p.attrs[fmt.Sprintf("%s_%s", CustomAttrPrefix, name)]
	return
}

// IsValid 检测用户是否有效
func (p *profile) IsValid() bool {
	return p.err == nil
}

// 设置异常错误
func (p *profile) setAbnormal(code int, message string) *profile {
	if code != core.SuccessCode {
		p.err = core.NewError(code, message)
	}
	
	return p
}

// 设置属性
func (p *profile) setAttr(name string, value interface{}) *profile {
	p.attrs[name] = value
	return p
}
