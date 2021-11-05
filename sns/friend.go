/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/30 2:55 下午
 * @Desc: 好友关系
 */

package sns

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dobyte/tencent-im/internal/entity"
)

var (
	errNotSetAccount   = errors.New("the friend's account is not set")
	errNotSetAddSource = errors.New("the friend's add source is not set")
)

type Friend struct {
	entity.User
	customAttrs map[string]interface{}
}

func NewFriend(userId ...string) *Friend {
	f := new(Friend)
	f.customAttrs = make(map[string]interface{})

	if len(userId) > 0 {
		f.SetUserId(userId[0])
	}

	return f
}

// SetAddSource 设置添加来源
func (f *Friend) SetAddSource(addSource string) {
	f.SetAttr(FriendAttrAddSource, "AddSource_Type_"+addSource)
}

// GetAddSource 获取添加来源
func (f *Friend) GetAddSource() (addSource string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(FriendAttrAddSource); exist {
		addSource = strings.TrimLeft(v.(string), "AddSource_Type_")
	}

	return
}

// GetSrcAddSource 获取添加来源（原始的）
func (f *Friend) GetSrcAddSource() (addSource string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(FriendAttrAddSource); exist {
		addSource = v.(string)
	}

	return
}

// SetRemark 设置备注
func (f *Friend) SetRemark(remark string) {
	f.SetAttr(FriendAttrRemark, remark)
}

// GetRemark 获取备注
func (f *Friend) GetRemark() (remark string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(FriendAttrRemark); exist {
		remark = v.(string)
	}

	return
}

// SetGroup 设置分组
func (f *Friend) SetGroup(groupName ...string) {
	f.SetAttr(FriendAttrGroup, groupName)
}

// GetGroup 获取分组
func (f *Friend) GetGroup() (groups []string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(FriendAttrGroup); exist && v != nil {
		if vv, ok := v.([]interface{}); ok {
			for _, group := range vv {
				groups = append(groups, group.(string))
			}
		}
	}

	return
}

// SetAddWording 设置形成好友关系时的附言信息
func (f *Friend) SetAddWording(addWording string) {
	f.SetAttr(FriendAttrAddWording, addWording)
}

// GetAddWording 获取形成好友关系时的附言信息
func (f *Friend) GetAddWording() (addWording string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(FriendAttrAddWording); exist {
		addWording = v.(string)
	}

	return
}

// SetAddTime 设置添加时间（忽略）
func (f *Friend) SetAddTime(addTime int64) {
	f.SetAttr(FriendAttrAddTime, addTime)
}

// GetAddTime 获取添加时间
func (f *Friend) GetAddTime() (addTime int64, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(FriendAttrAddTime); exist {
		addTime = v.(int64)
	}

	return
}

// SetRemarkTime 设置备注时间
func (f *Friend) SetRemarkTime(remarkTime int64) {
	f.SetAttr(FriendAttrRemarkTime, remarkTime)
}

// GetRemarkTime 获取备注时间
func (f *Friend) GetRemarkTime() (remarkTime int64, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(FriendAttrRemarkTime); exist {
		remarkTime = v.(int64)
	}

	return
}

// SetSNSCustomAttr 设置SNS自定义关系数据（自定义字段需要单独申请，请在 IM 控制台 >应用配置>功能配置申请自定义好友字段，申请提交后，自定义好友字段将在5分钟内生效）
func (f *Friend) SetSNSCustomAttr(name string, value interface{}) {
	if f.customAttrs == nil {
		f.customAttrs = make(map[string]interface{})
	}

	f.customAttrs[fmt.Sprintf("%s_%s", "Tag_SNS_Custom", name)] = value
}

// GetSNSCustomAttr 设置SNS自定义关系数据 （自定义字段需要单独申请，请在 IM 控制台 >应用配置>功能配置申请自定义好友字段，申请提交后，自定义好友字段将在5分钟内生效）
func (f *Friend) GetSNSCustomAttr(name string) (value interface{}, exist bool) {
	if f.customAttrs == nil {
		return
	}

	value, exist = f.customAttrs[fmt.Sprintf("%s_%s", "Tag_SNS_Custom", name)]
	return
}

// GetSNSAttrs 获取SNS标准关系数据
func (f *Friend) GetSNSAttrs() (attrs map[string]interface{}) {
	attrs = make(map[string]interface{})

	for k, v := range f.GetAttrs() {
		switch k {
		case FriendAttrAddSource, FriendAttrRemark, FriendAttrGroup, FriendAttrAddWording, FriendAttrAddTime, FriendAttrRemarkTime:
			attrs[k] = v
		}
	}

	return
}

// GetSNSCustomAttrs 获取SNS自定义关系数据（自定义字段需要单独申请，请在 IM 控制台 >应用配置>功能配置申请自定义好友字段，申请提交后，自定义好友字段将在5分钟内生效）
func (f *Friend) GetSNSCustomAttrs() (attrs map[string]interface{}) {
	attrs = make(map[string]interface{})

	if f.customAttrs == nil {
		return
	}

	for k, v := range f.customAttrs {
		switch k {
		case FriendAttrAddSource, FriendAttrRemark, FriendAttrGroup, FriendAttrAddWording, FriendAttrAddTime, FriendAttrRemarkTime:
		default:
			attrs[k] = v
		}
	}

	return
}

// checkError 检测参数错误
func (f *Friend) checkError() error {
	if f.GetUserId() == "" {
		return errNotSetAccount
	}

	if _, exist := f.GetSrcAddSource(); !exist {
		return errNotSetAddSource
	}

	return nil
}
