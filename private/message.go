/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/8/31 18:04
 * @Desc: 私聊消息实体
 */

package private

import (
    "errors"
    
    "github.com/dobyte/tencent-im/internal/conv"
    "github.com/dobyte/tencent-im/internal/entity"
    "github.com/dobyte/tencent-im/internal/types"
)

var errNotSetMsgReceiver = errors.New("message receiver is not set")

type Message struct {
    entity.Message
    receivers        []string        // 接收方UserId（可以为多个）
    syncOtherMachine int             // 同步到其他器
    timestamp        int64           // 消息时间戳，UNIX 时间戳（单位：秒）
    seq              int             // 消息序列号
    customData       interface{}     // 自定义数据
    sendControls     map[string]bool // 发送消息控制
    callbackControls map[string]bool // 禁用回调
}

func NewMessage() *Message {
    return &Message{}
}

// AddReceivers 添加接收方
func (m *Message) AddReceivers(userId ...string) {
    if m.receivers == nil {
        m.receivers = make([]string, 0)
    }
    m.receivers = append(m.receivers, userId...)
}

// SetReceivers 设置接收方
func (m *Message) SetReceivers(userId ...string) {
    if m.receivers != nil {
        m.receivers = m.receivers[0:0]
    }
    m.AddReceivers(userId...)
}

// GetReceivers 获取接收方
func (m *Message) GetReceivers() []string {
    return m.receivers
}

func (m *Message) GetLastReceiver() string {
    return m.receivers[0]
}

// SetSyncOtherMachine 设置同步到其他机器
func (m *Message) SetSyncOtherMachine(syncOtherMachine types.SyncOtherMachine) {
    m.syncOtherMachine = int(syncOtherMachine)
}

// GetSyncOtherMachine 获取同步至其他设备
func (m *Message) GetSyncOtherMachine() int {
    return m.syncOtherMachine
}

// SetSerialNo 设置消息序列号
func (m *Message) SetSerialNo(seq int) {
    m.seq = seq
}

// GetSerialNo 获取消息序列号
func (m *Message) GetSerialNo() int {
    return m.seq
}

// SetTimestamp 设置消息的时间戳
func (m *Message) SetTimestamp(timestamp int64) {
    m.timestamp = timestamp
}

// GetTimestamp 获取消息的时间戳
func (m *Message) GetTimestamp() int64 {
    return m.timestamp
}

// SetCustomData 设置自定义数据
func (m *Message) SetCustomData(data interface{}) {
    m.customData = data
}

// GetCustomData 获取自定义数据
func (m *Message) GetCustomData() string {
    return conv.String(m.customData)
}

// SetForbidBeforeSendMsgCallback 设置禁止发消息前回调
func (m *Message) SetForbidBeforeSendMsgCallback() {
    if m.callbackControls == nil {
        m.callbackControls = make(map[string]bool, 0)
    }
    m.callbackControls["ForbidBeforeSendMsgCallback"] = true
}

// SetForbidAfterSendMsgCallback 设置禁止发消息后回调
func (m *Message) SetForbidAfterSendMsgCallback() {
    if m.callbackControls == nil {
        m.callbackControls = make(map[string]bool, 0)
    }
    m.callbackControls["ForbidAfterSendMsgCallback"] = true
}

// GetForbidCallbackControl 获取消息回调禁止开关
func (m *Message) GetForbidCallbackControl() (controls []string) {
    if m.callbackControls != nil {
        if n := len(m.callbackControls); n > 0 {
            controls = make([]string, 0, n)
            for k := range m.callbackControls {
                controls = append(controls, k)
            }
        }
    }
    
    return
}

// SetNoUnread 设置该条消息不计入未读数
func (m *Message) SetNoUnread() {
    if m.sendControls == nil {
        m.sendControls = make(map[string]bool, 0)
    }
    m.sendControls["NoUnread"] = true
}

// SetNoLastMsg 设置该条消息不更新会话列表
func (m *Message) SetNoLastMsg() {
    if m.sendControls == nil {
        m.sendControls = make(map[string]bool, 0)
    }
    m.sendControls["NoLastMsg"] = true
}

// GetSendMsgControl 获取消息发送控制选项
func (m *Message) GetSendMsgControl() (controls []string) {
    if m.sendControls != nil {
        if n := len(m.sendControls); n > 0 {
            controls = make([]string, 0, n)
            for k := range m.sendControls {
                controls = append(controls, k)
            }
        }
    }
    
    return
}

// CheckError 检测错误
func (m *Message) CheckError() (err error) {
    if err = m.CheckLifeTimeArgError(); err != nil {
        return
    }

    if err = m.CheckBodyArgError(); err != nil {
        return
    }
    
    if err = m.checkReceiverArgError(); err != nil {
        return
    }
    
    return
}

// checkReceiverArgError 检测接收方参数
func (m *Message) checkReceiverArgError() error {
    if m.receivers == nil || len(m.receivers) == 0 {
        return errNotSetMsgReceiver
    }
    
    return nil
}
