/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/9/3 18:19
 * @Desc: 基础消息实体
 */

package entity

import (
    "errors"
    "math/rand"
    
    "github.com/dobyte/tencent-im/internal/enum"
    "github.com/dobyte/tencent-im/internal/types"
)

var (
    errInvalidMsgContent  = errors.New("invalid message content")
    errInvalidMsgLifeTime = errors.New("invalid message life time")
    errNotSetMsgContent   = errors.New("message content is not set")
)

type Message struct {
    sender      string          // 发送方UserId
    lifeTime    int             // 消息离线保存时长（单位：秒），最长为7天（604800秒）
    random      uint32          // 消息随机数，由随机函数产生
    body        []types.MsgBody // 消息体
    offlinePush *offlinePush    // 推送实体
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

// AddContent 添加消息内容（添加会累加之前的消息内容）
func (m *Message) AddContent(msgContent ...interface{}) {
    if m.body == nil {
        m.body = make([]types.MsgBody, 0)
    }
    
    if len(msgContent) > 0 {
        var msgType string
        for _, content := range msgContent {
            switch content.(type) {
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
                msgType = ""
            }
            
            m.body = append(m.body, types.MsgBody{
                MsgType:    msgType,
                MsgContent: content,
            })
        }
    }
}

// SetContent 设置消息内容（设置会冲掉之前的消息内容）
func (m *Message) SetContent(msgContent ...interface{}) {
    if m.body != nil {
        m.body = m.body[0:0]
    }
    m.AddContent(msgContent...)
}

// GetBody 获取消息体
func (m *Message) GetBody() []types.MsgBody {
    return m.body
}

// OfflinePush 新建离线推送对象
func (m *Message) OfflinePush() *offlinePush {
    if m.offlinePush == nil {
        m.offlinePush = newOfflinePush()
    }
    
    return m.offlinePush
}

// GetOfflinePushInfo 获取离线推送消息
func (m *Message) GetOfflinePushInfo() *types.OfflinePushInfo {
    if m.offlinePush == nil {
        return nil
    }
    
    return &types.OfflinePushInfo{
        PushFlag:    m.offlinePush.pushFlag,
        Title:       m.offlinePush.title,
        Desc:        m.offlinePush.desc,
        Ext:         m.offlinePush.ext,
        AndroidInfo: m.offlinePush.androidInfo,
        ApnsInfo:    m.offlinePush.apnsInfo,
    }
}

// CheckArgError 检测参数错误
func (m *Message) CheckArgError() error {
    if m.lifeTime < 0 || m.lifeTime > 604800 {
        return errInvalidMsgLifeTime
    }
    
    if m.body != nil && len(m.body) > 0 {
        for _, item := range m.body {
            if item.MsgType == "" {
                return errInvalidMsgContent
            }
        }
    } else {
        return errNotSetMsgContent
    }
    
    return nil
}
