/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/8/31 18:04
 * @Desc: 消息
 */

package private

import (
	"errors"
	
	"github.com/dobyte/tencent-im/enum"
	"github.com/dobyte/tencent-im/types"
)

var (
	invalidMsgContent  = errors.New("invalid msg content")
	invalidMsgLifeTime = errors.New("invalid msg life time")
)

type Message struct {
	err                error
	fromUserId         string                 // 发送方UserId
	toUserIds          []string               // 接收方UserId（可以为多个）
	isSyncOtherMachine bool                   // 是否同步到其他机器
	lifeTime           int                    // 消息离线保存时长（单位：秒），最长为7天（604800秒）
	timestamp          int64                  // 消息时间戳，UNIX 时间戳（单位：秒）
	seq                int                    // 消息序列号
	body               []types.MsgBody        // 消息体
	customData         interface{}            // 自定义数据
	forbidCallbacks    map[string]bool        // 禁用回调
	sendControls       map[string]bool        // 发送消息控制
	offlinePushInfo    *types.OfflinePushInfo // 离线推送信息配置
}

func NewMessage() *Message {
	return &Message{
		body:            make([]types.MsgBody, 0),
		toUserIds:       make([]string, 0),
		forbidCallbacks: make(map[string]bool),
		sendControls:    make(map[string]bool),
	}
}

// EmptyContent 清空消息内容
func (m *Message) AddContent(msgContent interface{}) *Message {
	var msgType string
	switch msgContent.(type) {
	case types.MsgTextContent, *types.MsgTextContent:
		msgType = enum.MsgText
	case types.MsgLocationContent, *types.MsgLocationContent:
		msgType = enum.MsgLocation
	case types.MsgFaceContent, *types.MsgFaceContent:
		msgType = enum.MsgFace
	case types.MsgCustomContent, *types.MsgCustomContent:
		msgType = enum.MsgCustom
	case types.MsgSoundContent, *types.MsgSoundContent:
		msgType = enum.MsgSound
	case types.MsgImageContent, *types.MsgImageContent:
		msgType = enum.MsgImage
	case types.MsgFileContent, *types.MsgFileContent:
		msgType = enum.MsgFile
	case types.MsgVideoContent, *types.MsgVideoContent:
		msgType = enum.MsgVideo
	default:
		m.err = invalidMsgContent
		return m
	}
	
	m.body = append(m.body, types.MsgBody{
		MsgType:    msgType,
		MsgContent: msgContent,
	})
	
	return m
}

// SetContent 设置消息内容
func (m *Message) SetContent(msgContent interface{}) *Message {
	m.body = m.body[0:0]
	if m.err == invalidMsgContent {
		m.err = nil
	}
	return m.AddContent(msgContent)
}

// SetSender 设置发送方UserId
func (m *Message) SetSender(userId string) *Message {
	m.fromUserId = userId
	return m
}

// AddReceiver 添加接收方UserId
func (m *Message) AddReceiver(userId string) *Message {
	m.toUserIds = append(m.toUserIds, userId)
	return m
}

// SetReceiver 设置接收方UserId
func (m *Message) SetReceiver(userId string) *Message {
	m.toUserIds = m.toUserIds[0:0]
	return m.AddReceiver(userId)
}

// SetSyncOtherMachine 设置同步到其他机器
func (m *Message) SetSyncOtherMachine() *Message {
	m.isSyncOtherMachine = true
	return m
}

// SetLifeTime 设置消息离线保存时长
func (m *Message) SetLifeTime(lifeTime int) *Message {
	if lifeTime < 0 || lifeTime > 604800 {
		m.err = invalidMsgLifeTime
	} else {
		m.lifeTime = lifeTime
	}
	return m
}

// SetSerialNo 设置消息序列号
func (m *Message) SetSerialNo(seq int) *Message {
	m.seq = seq
	return m
}

// SetTimeStamp 设置消息的时间戳
func (m *Message) SetTimeStamp(timestamp int64) *Message {
	m.timestamp = timestamp
	return m
}

// SetCustomData 设置自定义数据
func (m *Message) SetCustomData(data interface{}) *Message {
	m.customData = data
	return m
}

// SetOfflinePushInfo 设置离线推送信息配置
func (m *Message) SetOfflinePushInfo(offlinePushInfo types.OfflinePushInfo) *Message {
	m.offlinePushInfo = &offlinePushInfo
	return m
}

// SetForbidBeforeSendMsgCallback 设置禁止发消息前回调
func (m *Message) SetForbidBeforeSendMsgCallback() *Message {
	m.forbidCallbacks["ForbidBeforeSendMsgCallback"] = true
	return m
}

// SetForbidAfterSendMsgCallback 设置禁止发消息后回调
func (m *Message) SetForbidAfterSendMsgCallback() *Message {
	m.forbidCallbacks["ForbidAfterSendMsgCallback"] = true
	return m
}

// SetNoUnread 设置该条消息不计入未读数
func (m *Message) SetNoUnread() *Message {
	m.sendControls["NoUnread"] = true
	return m
}

// SetNoLastMsg 设置该条消息不更新会话列表
func (m *Message) SetNoLastMsg() *Message {
	m.sendControls["NoLastMsg"] = true
	return m
}

// IsValid 检测消息是否有效
func (m *Message) IsValid() bool {
	return m.err == nil
}

// GetError 获取异常错误
func (m *Message) GetError() error {
	return m.err
}
