/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/9/7 17:11
 * @Desc: TODO
 */

package group

import (
    "errors"
)

var (
    errNotSetGroupType          = errors.New("group type is not set")
    errNotSetGroupName          = errors.New("group name is not set")
    errGroupNameTooLong         = errors.New("group name is too long")
    errInvalidGroupType         = errors.New("invalid group type")
    errGroupIntroductionTooLong = errors.New("group introduction is too long")
    errGroupNotificationTooLong = errors.New("group notification is too long")
)

type Group struct {
    groupId           string                 // 群ID
    groupName         string                 // 群名称
    groupType         string                 // 群类型
    groupOwner        string                 // 群主ID
    groupIntroduction string                 // 群简介
    groupNotification string                 // 群公告
    groupAvatar       string                 // 群头像
    maxMemberCount    uint                   // 最大群成员数量
    applyJoinOption   string                 // 申请加群处理方式
    groupMembers      []*Member              // 群成员
    groupCustomData   map[string]interface{} // 群自定义数据
}

func NewGroup() *Group {
    return &Group{}
}

// SetId 设置群ID
func (g *Group) SetId(groupId string) {
    g.groupId = groupId
}

// GetId 获取群ID
func (g *Group) GetId() string {
    return g.groupId
}

// SetOwner 设置群主ID
func (g *Group) SetOwner(userId string) {
    g.groupOwner = userId
}

// GetOwner 获取群主ID
func (g *Group) GetOwner() string {
    return g.groupOwner
}

// SetName 设置群名称
func (g *Group) SetName(name string) {
    g.groupName = name
}

// GetName 获取群名称
func (g *Group) GetName() string {
    return g.groupName
}

// SetType 设置群类型
func (g *Group) SetType(groupType GroupType) {
    g.groupType = string(groupType)
}

// GetType 获取群类型
func (g *Group) GetType() string {
    return g.groupType
}

// SetIntroduction 设置群简介
func (g *Group) SetIntroduction(groupIntroduction string) {
    g.groupIntroduction = groupIntroduction
}

// GetIntroduction 获取群简介
func (g *Group) GetIntroduction() string {
    return g.groupIntroduction
}

// SetNotification 设置群公告
func (g *Group) SetNotification(groupNotification string) {
    g.groupNotification = groupNotification
}

// GetNotification 获取群公告
func (g *Group) GetNotification() string {
    return g.groupNotification
}

// SetAvatar 设置群头像
func (g *Group) SetAvatar(groupAvatar string) {
    g.groupAvatar = groupAvatar
}

// GetAvatar 获取群头像
func (g *Group) GetAvatar() string {
    return g.groupAvatar
}

// SetMaxMemberCount 设置最大群成员数量
func (g *Group) SetMaxMemberCount(maxMemberCount uint) {
    g.maxMemberCount = maxMemberCount
}

// GetMaxMemberCount 获取最大群成员数量
func (g *Group) GetMaxMemberCount() uint {
    return g.maxMemberCount
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
    if g.groupMembers == nil {
        g.groupMembers = make([]*Member, 0)
    }
    
    g.groupMembers = append(g.groupMembers, member...)
}

// SetMembers 设置群成员
func (g *Group) SetMembers(member ...*Member) {
    if g.groupMembers != nil {
        g.groupMembers = g.groupMembers[0:0]
    }
    
    g.AddMembers(member...)
}

// SetCustomData 设置自定义数据
func (g *Group) SetCustomData(name string, value interface{}) {
    if g.groupCustomData == nil {
        g.groupCustomData = make(map[string]interface{})
    }
    
    g.groupCustomData[name] = value
}

// GetCustomData 获取自定义数据
func (g *Group) GetCustomData(name string) (val interface{}, exist bool) {
    if g.groupCustomData == nil {
        return
    }
    
    val, exist = g.groupCustomData[name]
    
    return
}

// GetAllCustomData 获取所有自定义数据
func (g *Group) GetAllCustomData() map[string]interface{}  {
    return g.groupCustomData
}

// GetMembers 获取群成员
func (g *Group) GetMembers() []*Member {
    return g.groupMembers
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
    if g.groupName == "" {
        return errNotSetGroupName
    }
    
    if len(g.groupName) > 30 {
        return errGroupNameTooLong
    }
    
    return nil
}

// 检测群类型参数错误
func (g *Group) checkTypeArgError() error {
    if g.groupType == "" {
        return errNotSetGroupType
    }
    
    switch GroupType(g.groupType) {
    case GroupTypePublic, GroupTypePrivate, GroupTypeChatRoom, GroupTypeAVChatRoom:
    default:
        return errInvalidGroupType
    }
    
    return nil
}

// 检测群简介参数错误
func (g *Group) checkIntroductionArgError() error {
    if len(g.groupIntroduction) > 240 {
        return errGroupIntroductionTooLong
    }
    
    return nil
}

// 检测群公告参数错误
func (g *Group) checkNotificationArgError() error {
    if len(g.groupNotification) > 300 {
        return errGroupNotificationTooLong
    }
    
    return nil
}
