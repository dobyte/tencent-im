/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/8/31 18:04
 * @Desc: TODO
 */

package private

import (
	"errors"

	"github.com/dobyte/tencent-im/enum"
	"github.com/dobyte/tencent-im/types"
)

var invalidMsgContent = errors.New("invalid msg content")

type message struct {
	fromUserId         string          // 发送方UserId
	toUserId           string          // 接收方UserId
	isSyncOtherMachine bool            // 是否同步到其他机器
	lifeTime           int             // 消息离线保存时长（单位：秒），最长为7天（604800秒）
	timestamp          int             // 消息时间戳，UNIX 时间戳（单位：秒）
	body               []types.MsgBody // 消息体
	customData         string          // 自定义数据
	forbidCallbacks    map[string]bool // 禁用回调
	err                error
}

// SetContent 设置消息内容
func (m *message) SetContent(msgContent interface{}) *message {
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

// SetSender 设置发送方UserId
func (m *message) SetSender(userId string) *message {
	m.fromUserId = userId
	return m
}

// SetReceiver 设置接收方UserId
func (m *message) SetReceiver(userId string) *message {
	m.toUserId = userId
	return m
}

// SetSyncOtherMachine 设置同步到其他机器
func (m *message) SetSyncOtherMachine(isSync bool) *message {
	m.isSyncOtherMachine = isSync
	return m
}

// SetLifeTime 设置消息离线保存时长
func (m *message) SetLifeTime(lifeTime int) *message {
	m.lifeTime = lifeTime
	return m
}

// SetTimeStamp 设置消息的时间戳
func (m *message) SetTimeStamp(timestamp int) *message {
	m.timestamp = timestamp
	return m
}

// SetCustomData 设置自定义数据
func (m *message) SetCustomData(data string) *message {
	m.customData = data
	return m
}

// SetForbidBeforeSendMsgCallback 设置禁止发消息前回调
func (m *message) SetForbidBeforeSendMsgCallback() *message {
	m.forbidCallbacks["ForbidBeforeSendMsgCallback"] = true
	return m
}

// SetForbidAfterSendMsgCallback 设置禁止发消息后回调
func (m *message) SetForbidAfterSendMsgCallback() *message {
	m.forbidCallbacks["ForbidAfterSendMsgCallback"] = true
	return m
}
