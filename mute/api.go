/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/30 2:41 上午
 * @Desc: 全局禁言管理
 */

package mute

import (
    "github.com/dobyte/tencent-im/internal/core"
    "github.com/dobyte/tencent-im/internal/types"
)

const (
    serviceMute          = "openconfigsvr"
    commandSetNoSpeaking = "setnospeaking"
    commandGetNoSpeaking = "getnospeaking"
)

type API interface {
    // SetNoSpeaking 设置全局禁言
    // 设置帐号的单聊消息全局禁言。
    // 设置帐号的群组消息全局禁言。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/4230
    SetNoSpeaking(userId string, privateMuteTime, groupMuteTime *uint) (err error)
    
    // GetNoSpeaking 查询全局禁言
    // 查询帐号的单聊消息全局禁言。
    // 查询帐号的群组消息全局禁言。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/4229
    GetNoSpeaking(userId string) (privateMuteTime, groupMuteTime uint, err error)
}

type api struct {
    client core.Client
}

func NewAPI(client core.Client) API {
    return &api{client: client}
}

// SetNoSpeaking 设置全局禁言
// 设置帐号的单聊消息全局禁言。
// 设置帐号的群组消息全局禁言。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/4230
func (a *api) SetNoSpeaking(userId string, privateMuteTime, groupMuteTime *uint) (err error) {
    req := setNoSpeakingReq{
        UserId:          userId,
        PrivateMuteTime: privateMuteTime,
        GroupMuteTime:   groupMuteTime,
    }
    
    if err = a.client.Post(serviceMute, commandSetNoSpeaking, req, &types.BaseResp{}); err != nil {
        return
    }
    
    return
}

// GetNoSpeaking 查询全局禁言
// 查询帐号的单聊消息全局禁言。
// 查询帐号的群组消息全局禁言。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/4229
func (a *api) GetNoSpeaking(userId string) (privateMuteTime, groupMuteTime uint, err error) {
    req := getNoSpeakingReq{UserId: userId}
    resp := &getNoSpeakingResp{}
    
    if err = a.client.Post(serviceMute, commandGetNoSpeaking, req, resp); err != nil {
        return
    } else {
        privateMuteTime = resp.PrivateMuteTime
        groupMuteTime = resp.GroupMuteTime
    }
    
    return
}
