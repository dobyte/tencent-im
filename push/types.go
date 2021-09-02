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
