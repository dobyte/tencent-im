/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/29 17:39
 * @Desc: Push Api Request And Response Type Definition.
 */

package push

import "github.com/dobyte/tencent-im/types"

type (
	// 推送条件
	condition struct {
		TagsAnd  []string               `json:"TagsAnd"`  // （选填）标签条件的交集。标签是一个不超过50字节的字符串。注意属性推送和标签推送不可同时作为推送条件。TagsAnd 条件中的标签个数不能超过10个
		TagsOr   []string               `json:"TagsOr"`   // （选填）标签条件的并集。标签是一个不超过50字节的字符串。注意属性推送和标签推送不可同时作为推送条件。TagsOr 条件中的标签个数不能超过10个
		AttrsAnd map[string]interface{} `json:"AttrsAnd"` // （选填）属性条件的交集。注意属性推送和标签推送不可同时作为推送条件
		AttrsOr  map[string]interface{} `json:"AttrsOr"`  // （选填）属性条件的并集。注意属性推送和标签推送不可同时作为推送条件
	}
	
	// 推送（请求）
	pushMessageReq struct {
		FromUserId      string                 `json:"From_Account,omitempty"`    // （选填）消息推送方帐号
		Condition       *condition             `json:"Condition,omitempty"`       // （选填）推送条件
		MsgRandom       uint32                 `json:"MsgRandom"`                 // （必填）消息随机数，由随机函数产生
		MsgBody         []types.MsgBody        `json:"MsgBody"`                   // （必填）消息内容
		MsgLifeTime     int                    `json:"MsgLifeTime,omitempty"`     // （选填）消息离线存储时间，单位秒，最多保存7天（604800秒）。默认为0，表示不离线存储
		OfflinePushInfo *types.OfflinePushInfo `json:"OfflinePushInfo,omitempty"` // （选填）离线推送信息配置
	}
	
	// 推送（响应）
	pushMessageResp struct {
		types.ActionBaseResp
		TaskId string `json:"TaskId"` // 推送任务ID
	}
	
	// 设置应用属性名称（请求）
	setAttrNamesReq struct {
		AttrNames map[string]string `json:"AttrNames"` // （必填）属性名
	}
	
	// 获取应用属性名称（请求）
	getAttrNamesReq struct {
	}
	
	// 获取应用属性名称（响应）
	getAttrNamesResp struct {
		types.ActionBaseResp
		AttrNames map[string]string `json:"AttrNames"` // 属性名
	}
	
	// 获取用户属性（请求）
	getUserAttrsReq struct {
		UserIds []string `json:"To_Account"` // （必填）目标用户帐号列表
	}
	
	// 获取用户属性（响应）
	getUserAttrsResp struct {
		types.ActionBaseResp
		UserAttrs []userAttr `json:"UserAttrs"`
	}
	
	// 用户属性
	userAttr struct {
		UserId string                 `json:"To_Account"` // 用户UserId
		Attrs  map[string]interface{} `json:"Attrs"`      // 用户属性
	}
	
	// 设置用户属性（请求）
	setUserAttrsReq struct {
		UserAttrs []userAttr `json:"UserAttrs"` // （必填）用户属性
	}
	
	// 删除用户属性（请求）
	deleteUserAttrsReq struct {
		UserAttrs []deleteUserAttr `json:"UserAttrs"` // （必填）用户属性
	}
	
	// 删除的用户属性
	deleteUserAttr struct {
		UserId string   `json:"To_Account"` // 用户UserId
		Attrs  []string `json:"Attrs"`      // 用户属性
	}
	
	// 获取用户标签（请求）
	getUserTagsReq struct {
		UserIds []string `json:"To_Account"` // （必填）目标用户帐号列表
	}
	
	// 获取用户标签（响应）
	getUserTagsResp struct {
		types.ActionBaseResp
		UserTags []userTag `json:"UserTags"` // 用户标签内容列表
	}
	
	// 用户标签
	userTag struct {
		UserId string   `json:"To_Account"`
		Tags   []string `json:"Tags"`
	}
	
	// 添加用户标签（请求）
	addUserTagsReq struct {
		UserTags []userTag `json:"UserTags"` // （必填）用户标签内容列表
	}
	
	// 删除用户标签（请求）
	deleteUserTagsReq struct {
		UserTags []userTag `json:"UserTags"` // （必填）用户标签内容列表
	}
	
	// 删除用户所有标签（请求）
	deleteUserAllTagsReq struct {
		UserIds []string `json:"To_Account"`
	}
)
