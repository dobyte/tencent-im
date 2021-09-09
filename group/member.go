/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/9/7 18:06
 * @Desc: TODO
 */

package group

import (
	"errors"
	"time"
)

var errNotSetUserId = errors.New("group member userid is not set")

type (
	// MsgFlag 消息接收选项
	MsgFlag string
)

const (
	MsgFlagAcceptAndNotify MsgFlag = "AcceptAndNotify" // 接收并提示
	MsgFlagAcceptNotNotify MsgFlag = "AcceptNotNotify" // 接收不提示（不会触发 APNs 远程推送）
	MsgFlagDiscard         MsgFlag = "Discard"         // 屏蔽群消息（不会向客户端推送消息）
)

type Member struct {
	userId          string                 // 成员ID
	role            string                 // 群内身份
	joinTime        int64                  // 加入时间
	nameCard        string                 // 群名片
	msgSeq          int                    // 成员当前已读消息Seq
	msgFlag         MsgFlag                // 消息接收选项
	lastSendMsgTime int64                  // 最后发送消息的时间
	shutUpUntil     *int64                 // 需禁言时间，单位为秒，0表示取消禁言
	customData      map[string]interface{} // 自定义数据
}

func NewMember(userId ...string) *Member {
	member := &Member{}
	if len(userId) > 0 {
		member.SetUserId(userId[0])
	}
	return member
}

// SetUserId 设置群成员ID
func (m *Member) SetUserId(userId string) {
	m.userId = userId
}

// GetUserId 获取群成员ID
func (m *Member) GetUserId() string {
	return m.userId
}

// SetRole 设置群内身份
func (m *Member) SetRole(role string) {
	m.role = role
}

// GetRole 获取群内身份
func (m *Member) GetRole() string {
	return m.role
}

// SetJoinTime 设置加入时间
func (m *Member) SetJoinTime(joinTime time.Time) {
	m.joinTime = joinTime.Unix()
}

// GetJoinTime 获取加入时间
func (m *Member) GetJoinTime() time.Time {
	return time.Unix(m.joinTime, 0)
}

// SetNameCard 设置群名片
func (m *Member) SetNameCard(nameCard string) {
	m.nameCard = nameCard
}

// GetNameCard 获取群名片
func (m *Member) GetNameCard() string {
	return m.nameCard
}

// GetMsgSeq 获取成员当前已读消息Seq
func (m *Member) GetMsgSeq() int {
	return m.msgSeq
}

// SetMsgFlag 设置消息接收选项
func (m *Member) SetMsgFlag(msgFlag MsgFlag) {
	m.msgFlag = msgFlag
}

// GetMsgFlag 获取消息接收选项
func (m *Member) GetMsgFlag() MsgFlag {
	return m.msgFlag
}

// SetShutUpUntil 设置需禁言时间，单位为秒，0表示取消禁言
func (m *Member) SetShutUpUntil(shutUpUntil int64) {
	m.shutUpUntil = &shutUpUntil
}

// GetShutUpUntil 获取需禁言时间，单位为秒，0表示取消禁言
func (m *Member) GetShutUpUntil() int64 {
	if m.shutUpUntil == nil {
		return 0
	} else {
		return *m.shutUpUntil
	}
}

// SetCustomData 设置自定义数据
func (m *Member) SetCustomData(name string, value interface{}) {
	if m.customData == nil {
		m.customData = make(map[string]interface{})
	}

	m.customData[name] = value
}

// GetCustomData 获取自定义数据
func (m *Member) GetCustomData(name string) (val interface{}, exist bool) {
	if m.customData == nil {
		return
	}

	val, exist = m.customData[name]

	return
}

// GetAllCustomData 获取所有自定义数据
func (m *Member) GetAllCustomData() map[string]interface{} {
	return m.customData
}

// 检测参数错误
func (m *Member) checkError() (err error) {
	if m.userId == "" {
		return errNotSetUserId
	}

	return nil
}
