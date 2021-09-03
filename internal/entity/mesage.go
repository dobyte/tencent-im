/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/9/3 18:19
 * @Desc: TODO
 */

package entity

import (
    "errors"
    "math/rand"
    
    "github.com/dobyte/tencent-im/types"
)

var (
    errInvalidMsgContent  = errors.New("invalid msg content")
    errInvalidMsgLifeTime = errors.New("invalid msg life time")
    errNotSetMsgContent   = errors.New("do not set msg content")
    errNotSetMsgReceiver  = errors.New("do not set msg receiver")
)

const (
    invalidMsgContent = iota
    invalidMsgLifeTime
    notSetMsgContent
    notSetMsgReceiver
)

type Message struct {
    errs               map[int]error   //
    sender             string          // 发送方UserId
    toUserIds          []string        // 接收方UserId（可以为多个）
    isSyncOtherMachine bool            // 是否同步到其他机器
    lifeTime           int             // 消息离线保存时长（单位：秒），最长为7天（604800秒）
    random             uint32          // 消息随机数，由随机函数产生
    timestamp          int64           // 消息时间戳，UNIX 时间戳（单位：秒）
    seq                int             // 消息序列号
    body               []types.MsgBody // 消息体
    customData         interface{}     // 自定义数据
    forbidCallbacks    map[string]bool // 禁用回调
    sendControls       map[string]bool // 发送消息控制
    // offlinePushInfo    *types.OfflinePushInfo // 离线推送信息配置
    offlinePush *OfflinePush
}

// SetSender 设置发送方UserId
func (m *Message) SetSender(userId string) {
    m.sender = userId
}

// GetSender 获取发送者
func (m *Message) GetSender() string {
    return m.sender
}

// SetLifeTime 设置消息离线保存时长
func (m *Message) SetLifeTime(lifeTime int) {
    m.lifeTime = lifeTime
}

// GetLifeTime 获取消息离线保存时长
func (m *Message) GetLifeTime() int {
    if m.lifeTime < 0 || m.lifeTime > 604800 {
        m.errs[invalidMsgLifeTime] = errInvalidMsgLifeTime
    }
    return m.lifeTime
}

// SetRandom 设置消息随机数
func (m *Message) SetRandom(random uint32) {
    m.random = random
}

// GetRandom 获取消息随机数
func (m *Message) GetRandom() uint32 {
    if m.random == 0 {
        m.random = rand.Uint32()
    }
    
    return m.random
}

// OfflinePush 新建离线推送对象
func (m *Message) OfflinePush() *OfflinePush {
    m.offlinePush = NewOfflinePush()
    return m.offlinePush
}

func (m *Message) GetOfflinePushInfo() {

}
