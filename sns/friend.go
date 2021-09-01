/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/30 2:55 下午
 * @Desc: 好友关系
 */

package sns

import (
	"fmt"
	"strings"
	
	"github.com/dobyte/tencent-im/user"
)

const (
	// 好友属性
	SNSAttrAddSource  = "Tag_SNS_IM_AddSource"  // 添加源
	SNSAttrRemark     = "Tag_SNS_IM_Remark"     // 备注
	SNSAttrGroup      = "Tag_SNS_IM_Group"      // 分组
	SNSAttrAddWording = "Tag_SNS_IM_AddWording" // 附言信息
	SNSAttrAddTime    = "Tag_SNS_IM_AddTime"    // 添加时间
	SNSAttrRemarkTime = "Tag_SNS_IM_RemarkTime" // 备注时间
	
	// 自定义好友字段前缀
	SNSCustomAttrPrefix = "Tag_SNS_Custom"
)

type Friend struct {
	user.User
	customAttrs map[string]interface{}
}

func NewFriend() *Friend {
	f := new(Friend)
	f.InitAttrs()
	f.customAttrs = make(map[string]interface{})
	return f
}

// SetAddSource 设置添加来源
func (f *Friend) SetAddSource(addSource string) {
	f.SetAttr(SNSAttrAddSource, "AddSource_Type_"+addSource)
}

// GetAddSource 获取添加来源
func (f *Friend) GetAddSource() (addSource string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(SNSAttrAddSource); exist {
		addSource = strings.TrimLeft(v.(string), "AddSource_Type_")
	}
	
	return
}

// GetAddSource 获取添加来源（原始的）
func (f *Friend) GetSrcAddSource() (addSource string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(SNSAttrAddSource); exist {
		addSource = v.(string)
	}
	
	return
}

// SetRemark 设置备注
func (f *Friend) SetRemark(remark string) {
	f.SetAttr(SNSAttrRemark, remark)
}

// GetRemark 获取备注
func (f *Friend) GetRemark() (remark string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(SNSAttrRemark); exist {
		remark = v.(string)
	}
	
	return
}

// SetGroup 设置分组
func (f *Friend) SetGroup(groupName ...string) {
	f.SetAttr(SNSAttrGroup, groupName)
}

// GetGroup 获取分组
func (f *Friend) GetGroup() (groups []string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(SNSAttrGroup); exist && v != nil {
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
	f.SetAttr(SNSAttrAddWording, addWording)
}

// GetAddWording 获取形成好友关系时的附言信息
func (f *Friend) GetAddWording() (addWording string, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(SNSAttrAddWording); exist {
		addWording = v.(string)
	}
	
	return
}

// SetAddTime 设置添加时间
func (f *Friend) SetAddTime(addTime int64) {
	f.SetAttr(SNSAttrAddTime, addTime)
}

// GetAddTime 获取添加时间
func (f *Friend) GetAddTime() (addTime int64, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(SNSAttrAddTime); exist {
		addTime = v.(int64)
	}
	
	return
}

// SetRemarkTime 设置备注时间
func (f *Friend) SetRemarkTime(remarkTime int64) {
	f.SetAttr(SNSAttrRemarkTime, remarkTime)
}

// GetRemarkTime 获取备注时间
func (f *Friend) GetRemarkTime() (remarkTime int64, exist bool) {
	var v interface{}
	if v, exist = f.GetAttr(SNSAttrRemarkTime); exist {
		remarkTime = v.(int64)
	}
	
	return
}

// SetSNSCustomAttr 设置SNS自定义关系数据（自定义字段需要单独申请，请在 IM 控制台 >应用配置>功能配置申请自定义好友字段，申请提交后，自定义好友字段将在5分钟内生效）
func (f *Friend) SetSNSCustomAttr(name string, value interface{}) {
	f.customAttrs[fmt.Sprintf("%s_%s", SNSCustomAttrPrefix, name)] = value
}

// SetSNSCustomAttr 设置SNS自定义关系数据 （自定义字段需要单独申请，请在 IM 控制台 >应用配置>功能配置申请自定义好友字段，申请提交后，自定义好友字段将在5分钟内生效）
func (f *Friend) GetSNSCustomAttr(name string) (value interface{}, exist bool) {
	value, exist = f.customAttrs[fmt.Sprintf("%s_%s", SNSCustomAttrPrefix, name)]
	return
}

// GetSNSAttrs 获取SNS标准关系数据
func (f *Friend) GetSNSAttrs() (attrs map[string]interface{}) {
	attrs = make(map[string]interface{})
	for k, v := range f.GetAttrs() {
		switch k {
		case SNSAttrAddSource, SNSAttrRemark, SNSAttrGroup, SNSAttrAddWording, SNSAttrAddTime, SNSAttrRemarkTime:
			attrs[k] = v
		}
	}
	
	return
}

// GetSNSCustomAttrs 获取SNS自定义关系数据（自定义字段需要单独申请，请在 IM 控制台 >应用配置>功能配置申请自定义好友字段，申请提交后，自定义好友字段将在5分钟内生效）
func (f *Friend) GetSNSCustomAttrs() (attrs map[string]interface{}) {
	attrs = make(map[string]interface{})
	for k, v := range f.customAttrs {
		switch k {
		case SNSAttrAddSource, SNSAttrRemark, SNSAttrGroup, SNSAttrAddWording, SNSAttrAddTime, SNSAttrRemarkTime:
		default:
			attrs[k] = v
		}
	}
	
	return
}
