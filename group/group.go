/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/9/7 17:11
 * @Desc: TODO
 */

package group

import (
	"errors"
	"time"

	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/internal/enum"
)

var (
	errNotSetGroupType          = errors.New("group type is not set")
	errNotSetGroupName          = errors.New("group name is not set")
	errGroupNameTooLong         = errors.New("group name is too long")
	errInvalidGroupType         = errors.New("invalid group type")
	errGroupIntroductionTooLong = errors.New("group introduction is too long")
	errGroupNotificationTooLong = errors.New("group notification is too long")
)

type (
	// Type 群类型
	Type string

	// ApplyJoinOption 申请加群处理方式
	ApplyJoinOption string

	// ShutUpStatus 全员禁言状态
	ShutUpStatus string
)

const (
	TypePublic   Type = "Public"     // Public（陌生人社交群）
	TypePrivate  Type = "Private"    // Private（即 Work，好友工作群）
	TypeChatRoom Type = "ChatRoom"   // ChatRoom（即 Meeting，会议群）
	TypeLiveRoom Type = "AVChatRoom" // AVChatRoom（直播群）

	ApplyJoinOptionFreeAccess     ApplyJoinOption = "FreeAccess"     // 自由加入
	ApplyJoinOptionNeedPermission ApplyJoinOption = "NeedPermission" // 需要验证
	ApplyJoinOptionDisableApply   ApplyJoinOption = "DisableApply"   // 禁止加群

	ShutUpStatusOn  ShutUpStatus = "On"  // 开启
	ShutUpStatusOff ShutUpStatus = "Off" // 关闭
)

type Group struct {
	err             error
	id              string                 // 群ID
	name            string                 // 群名称
	types           string                 // 群类型
	owner           string                 // 群主ID
	introduction    string                 // 群简介
	notification    string                 // 群公告
	avatar          string                 // 群头像
	memberNum       uint                   // 群成员数
	maxMemberNum    uint                   // 最大群成员数量
	applyJoinOption string                 // 申请加群处理方式
	members         []*Member              // 群成员
	customData      map[string]interface{} // 群自定义数据
	createTime      int64                  // 群创建时间
	lastInfoTime    int64                  // 最后群资料变更时间
	lastMsgTime     int64                  // 群内最后一条消息的时间
	nextMsgSeq      int                    // 群内下一条消息的Seq
	shutUpStatus    string                 // 群全员禁言状态
}

func NewGroup(id ...string) *Group {
	group := &Group{}
	if len(id) > 0 {
		group.SetId(id[0])
	}
	return group
}

// SetId 设置群ID
func (g *Group) SetId(id string) {
	g.id = id
}

// GetId 获取群ID
func (g *Group) GetId() string {
	return g.id
}

// SetOwner 设置群主ID
func (g *Group) SetOwner(owner string) {
	g.owner = owner
}

// GetOwner 获取群主ID
func (g *Group) GetOwner() string {
	return g.owner
}

// SetName 设置群名称
func (g *Group) SetName(name string) {
	g.name = name
}

// GetName 获取群名称
func (g *Group) GetName() string {
	return g.name
}

// SetType 设置群类型
func (g *Group) SetType(types Type) {
	g.types = string(types)
}

// GetType 获取群类型
func (g *Group) GetType() Type {
	return Type(g.types)
}

// SetIntroduction 设置群简介
func (g *Group) SetIntroduction(introduction string) {
	g.introduction = introduction
}

// GetIntroduction 获取群简介
func (g *Group) GetIntroduction() string {
	return g.introduction
}

// SetNotification 设置群公告
func (g *Group) SetNotification(notification string) {
	g.notification = notification
}

// GetNotification 获取群公告
func (g *Group) GetNotification() string {
	return g.notification
}

// SetAvatar 设置群头像
func (g *Group) SetAvatar(avatar string) {
	g.avatar = avatar
}

// GetAvatar 获取群头像
func (g *Group) GetAvatar() string {
	return g.avatar
}

// SetMaxMemberNum 设置最大群成员数量
func (g *Group) SetMaxMemberNum(maxMemberNum uint) {
	g.maxMemberNum = maxMemberNum
}

// GetMaxMemberNum 获取最大群成员数量
func (g *Group) GetMaxMemberNum() uint {
	return g.maxMemberNum
}

// GetMemberNum 获取群成员数
func (g *Group) GetMemberNum() uint {
	return g.memberNum
}

// SetApplyJoinOption 设置申请加群处理方式
func (g *Group) SetApplyJoinOption(applyJoinOption ApplyJoinOption) {
	g.applyJoinOption = string(applyJoinOption)
}

// GetApplyJoinOption 获取申请加群处理方式
func (g *Group) GetApplyJoinOption() string {
	return g.applyJoinOption
}

// AddMembers 添加群成员
func (g *Group) AddMembers(member ...*Member) {
	if g.members == nil {
		g.members = make([]*Member, 0)
	}

	g.members = append(g.members, member...)
}

// SetMembers 设置群成员
func (g *Group) SetMembers(member ...*Member) {
	if g.members != nil {
		g.members = g.members[0:0]
	}

	g.AddMembers(member...)
}

// SetCustomData 设置自定义数据
func (g *Group) SetCustomData(name string, value interface{}) {
	if g.customData == nil {
		g.customData = make(map[string]interface{})
	}

	g.customData[name] = value
}

// GetCustomData 获取自定义数据
func (g *Group) GetCustomData(name string) (val interface{}, exist bool) {
	if g.customData == nil {
		return
	}

	val, exist = g.customData[name]

	return
}

// GetAllCustomData 获取所有自定义数据
func (g *Group) GetAllCustomData() map[string]interface{} {
	return g.customData
}

// GetMembers 获取群成员
func (g *Group) GetMembers() []*Member {
	return g.members
}

// GetGroupCreateTime 获取群创建时间
func (g *Group) GetGroupCreateTime() time.Time {
	return time.Unix(g.createTime, 0)
}

// GetLastInfoTime 获取最后群资料变更时间
func (g *Group) GetLastInfoTime() time.Time {
	return time.Unix(g.lastInfoTime, 0)
}

// GetLastMsgTime 获取群内最后一条消息的时间
func (g *Group) GetLastMsgTime() time.Time {
	return time.Unix(g.lastMsgTime, 0)
}

// GetNextMsgSeq 获取群内下一条消息的Seq
func (g *Group) GetNextMsgSeq() int {
	return g.nextMsgSeq
}

// SetShutUpStatus 设置全员禁言状态
func (g *Group) SetShutUpStatus(shutUpStatus ShutUpStatus) {
	g.shutUpStatus = string(shutUpStatus)
}

// GetShutUpStatus 获取群全员禁言状态
func (g *Group) GetShutUpStatus() string {
	return g.shutUpStatus
}

// IsValid 检测用户是否有效
func (g *Group) IsValid() bool {
	return g.err == nil
}

// GetError 获取异常错误
func (g *Group) GetError() error {
	return g.err
}

// 设置异常错误
func (g *Group) setError(code int, message string) {
	if code != enum.SuccessCode {
		g.err = core.NewError(code, message)
	}
}

// 检测错误
func (g *Group) checkError() (err error) {
	if err = g.checkTypeArgError(); err != nil {
		return
	}

	if err = g.checkNameArgError(); err != nil {
		return
	}

	if err = g.checkIntroductionArgError(); err != nil {
		return
	}

	if err = g.checkNotificationArgError(); err != nil {
		return
	}

	return
}

// 检测群名称参数错误
func (g *Group) checkNameArgError() error {
	if g.name == "" {
		return errNotSetGroupName
	}

	if len(g.name) > 30 {
		return errGroupNameTooLong
	}

	return nil
}

// 检测群类型参数错误
func (g *Group) checkTypeArgError() error {
	if g.types == "" {
		return errNotSetGroupType
	}

	switch Type(g.types) {
	case TypePublic, TypePrivate, TypeChatRoom, TypeLiveRoom:
	default:
		return errInvalidGroupType
	}

	return nil
}

// 检测群简介参数错误
func (g *Group) checkIntroductionArgError() error {
	if len(g.introduction) > 240 {
		return errGroupIntroductionTooLong
	}

	return nil
}

// 检测群公告参数错误
func (g *Group) checkNotificationArgError() error {
	if len(g.notification) > 300 {
		return errGroupNotificationTooLong
	}

	return nil
}
