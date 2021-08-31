/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/5/29 17:39
 * @Desc: Push Api Request And Response Type Definition.
 */

package push

import "github.com/dobyte/tencent-im/types"

type (
	// 推送条件
	PushCondition struct {
		TagsAnd  []string               `json:"TagsAnd"`
		TagsOr   []string               `json:"TagsOr"`
		AttrsAnd map[string]interface{} `json:"AttrsAnd"`
		AttrsOr  map[string]interface{} `json:"AttrsOr"`
	}
	
	// 推送参数
	PushArgument struct {
		FromUserId      string                // （选填）消息推送方帐号
		Condition       *PushCondition        // （选填）推送条件
		MsgContents     []interface{}         // （必填）消息内容
		MsgLifeTime     int                   // （必填）消息离线存储时间，单位秒，最多保存7天（604800秒）。默认为0，表示不离线存储
		OfflinePushInfo types.OfflinePushInfo // （选填）离线推送信息配置
	}
	
	// 推送（请求）
	pushReq struct {
		FromUserId      string                `json:"From_Account"`    // （选填）消息推送方帐号
		Condition       *PushCondition        `json:"Condition"`       // （选填）推送条件
		MsgRandom       int64                 `json:"MsgRandom"`       // （必填）消息随机数，由随机函数产生
		MsgBody         []types.MsgBody       `json:"MsgBody"`         // （必填）消息内容
		MsgLifeTime     int                   `json:"MsgLifeTime"`     // （选填）消息离线存储时间，单位秒，最多保存7天（604800秒）。默认为0，表示不离线存储
		OfflinePushInfo types.OfflinePushInfo `json:"OfflinePushInfo"` // （选填）离线推送信息配置
	}
	
	// 推送（响应）
	pushResp struct {
		types.ActionBaseResp
		TaskId string `json:"TaskId"` // 推送任务ID
	}
	
	SetAttrNameReq struct {
		AttrNames map[string]string `json:"AttrNames"`
	}
	
	SetAttrNameResp struct {
		types.ActionBaseResp
	}
	
	GetAttrNameReq struct {
	}
	
	GetAttrNameResp struct {
		types.ActionBaseResp
		AttrNames map[string]string `json:"AttrNames"`
	}
	
	GetAttrReq struct {
		ToAccount []string `json:"To_Account"`
	}
	
	GetAttrResp struct {
		types.ActionBaseResp
		UserAttrs []UserAttrItem `json:"UserAttrs"`
	}
	
	UserAttrItem struct {
		ToAccount string                 `json:"To_Account"`
		Attrs     map[string]interface{} `json:"Attrs"`
	}
	
	SetAttrReq struct {
		UserAttrs []UserAttrItem `json:"UserAttrs"`
	}
	
	SetAttrResp struct {
		types.ActionBaseResp
	}
	
	RemoveAttrReq struct {
		UserAttrs []UserAttrItem `json:"UserAttrs"`
	}
	
	RemoveAttrResp struct {
		types.ActionBaseResp
	}
	
	UserTagItem struct {
		ToAccount string   `json:"To_Account"`
		Tags      []string `json:"Tags"`
	}
	
	GetTagReq struct {
		ToAccount []string `json:"To_Account"`
	}
	
	GetTagResp struct {
		types.ActionBaseResp
		UserTags []UserTagItem `json:"UserTags"`
	}
	
	AddTagReq struct {
		UserTags []UserTagItem `json:"UserTags"`
	}
	
	AddTagResp struct {
		types.ActionBaseResp
	}
	
	RemoveTagReq struct {
		UserTags []UserTagItem `json:"UserTags"`
	}
	
	RemoveTagResp struct {
		types.ActionBaseResp
	}
	
	RemoveAllTagsReq struct {
		ToAccount []string `json:"To_Account"`
	}
	
	RemoveAllTagsResp struct {
		types.ActionBaseResp
	}
)
