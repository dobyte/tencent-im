/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/30 2:41 上午
 * @Desc: 全局禁言数据类型
 */

package mute

import "github.com/dobyte/tencent-im/internal/types"

type (
    // 设置全局禁言（请求）
    setNoSpeakingReq struct {
        UserId          string `json:"Set_Account"`                      // （必填）设置禁言配置的帐号
        PrivateMuteTime *uint  `json:"C2CmsgNospeakingTime,omitempty"`   // （选填）单聊消息禁言时间，单位为秒，非负整数，最大值为4294967295（十六进制 0xFFFFFFFF） 0表示取消该帐号的单聊消息禁言;4294967295表示该帐号被设置永久禁言;其它值表示该帐号具体的禁言时间
        GroupMuteTime   *uint  `json:"GroupmsgNospeakingTime,omitempty"` // （选填）单聊消息禁言时间，单位为秒，非负整数，最大值为4294967295（十六进制 0xFFFFFFFF） 0表示取消该帐号的单聊消息禁言;4294967295表示该帐号被设置永久禁言;其它值表示该帐号具体的禁言时间
    }
    
    // 设置全局禁言（请求）
    getNoSpeakingReq struct {
        UserId string `json:"Get_Account"` // （必填）查询禁言信息的帐号
    }
    
    // 设置全局禁言（响应）
    getNoSpeakingResp struct {
        types.BaseResp
        PrivateMuteTime uint `json:"C2CmsgNospeakingTime"`
        GroupMuteTime   uint `json:"GroupmsgNospeakingTime"`
    }
)
