/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:46
 * @Desc: Group Api Implementation.
 */

package group

import (
    "github.com/dobyte/tencent-im/internal/core"
    "github.com/dobyte/tencent-im/internal/types"
)

const (
    serviceGroup                       = "group_open_http_svc"
    commandGetGroupList                = "get_appid_group_list"
    commandCreateGroup                 = "create_group"
    commandDestroyGroup                = "destroy_group"
    commandGetGroups                   = "get_group_info"
    commandGetGroupMemberInfo          = "get_group_member_info"
    commandModifyGroupBaseInfo         = "modify_group_base_info"
    commandAddGroupMembers             = "add_group_member"
    commandDeleteGroupMember           = "delete_group_member"
    commandModifyGroupMemberInfo       = "modify_group_member_info"
    commandGetJoinedGroupList          = "get_joined_group_list"
    commandGetRoleInGroup              = "get_role_in_group"
    commandForbidSendMsg               = "forbid_send_msg"
    commandGetGroupShuttedUin          = "get_group_shutted_uin"
    commandSendGroupMsg                = "send_group_msg"
    commandSendGroupSystemNotification = "send_group_system_notification"
    commandChangeGroupOwner            = "change_group_owner"
    commandRecallGroupMsg              = "group_msg_recall"
    commandImportGroup                 = "import_group"
    commandImportGroupMsg              = "import_group_msg"
    commandImportGroupMember           = "import_group_member"
    commandSetUnreadMsgNum             = "set_unread_msg_num"
    commandDeleteGroupMsgBySender      = "delete_group_msg_by_sender"
    commandGetGroupSimpleMsg           = "group_msg_get_simple"
    commandGetOnlineMemberNum          = "get_online_member_num"
)

type API interface {
    // CreateGroup 创建群组
    // App 管理员可以通过该接口创建群组。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1615
    CreateGroup(group *Group) (groupId string, err error)
    
    // DestroyGroup 解散群组
    // App管理员通过该接口解散群。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1624
    DestroyGroup(groupId string) (err error)
    
    // GetGroup 获取单个群详细资料
    // 本方法由“获取多个群详细资料（GetGroups）”拓展而来
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1616
    GetGroup(groupId string, filter ...*Filter) (group *Group, err error)
    
    // GetGroups 获取群详细资料
    // App 管理员可以根据群组 ID 获取群组的详细信息。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1616
    GetGroups(groupIds []string, filter ...*Filter) (groups []*Group, err error)
    
    // AddGroupMembers 增加群成员
    // App管理员可以通过该接口向指定的群中添加新成员。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1621
    AddGroupMembers(groupId string, userIds []string, silence ...bool) (results []AddMembersResult, err error)
    
    // DeleteGroupMembers 删除群成员
    // App管理员可以通过该接口删除群成员。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1622
    DeleteGroupMembers(groupId string, userIds []string, reasonAndSilence ...interface{}) (err error)
}

type api struct {
    client core.Client
}

func NewAPI(client core.Client) API {
    return &api{client: client}
}

// GetGroupList Get the IDs of all groups in the app.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1614
func (a *api) GetGroupList(req *GetGroupListReq) (*GetGroupListResp, error) {
    resp := &GetGroupListResp{}
    
    if err := a.client.Post(serviceGroup, commandGetGroupList, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// CreateGroup 创建群组
// App 管理员可以通过该接口创建群组。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1615
func (a *api) CreateGroup(group *Group) (groupId string, err error) {
    if err = group.checkError(); err != nil {
        return
    }
    
    req := createGroupReq{}
    req.GroupId = group.id
    req.OwnerUserId = group.owner
    req.Type = group.types
    req.Name = group.name
    req.FaceUrl = group.avatar
    req.Introduction = group.introduction
    req.Notification = group.notification
    req.MaxMemberCount = group.maxMemberCount
    req.ApplyJoinOption = group.applyJoinOption
    
    if data := group.GetAllCustomData(); data != nil {
        req.AppDefinedData = make([]customData, 0, len(data))
        for key, val := range data {
            req.AppDefinedData = append(req.AppDefinedData, customData{
                Key:   key,
                Value: val,
            })
        }
    }
    
    if len(group.members) > 0 {
        req.MemberList = make([]memberInfo, 0, len(group.members))
        for _, member := range group.members {
            if err = member.checkError(); err != nil {
                return
            } else {
                req.MemberList = append(req.MemberList, memberInfo{
                    UserId:   member.userId,
                    Role:     member.role,
                    JoinTime: member.joinTime,
                    NameCard: member.nameCard,
                })
            }
        }
    }
    
    resp := &createGroupResp{}
    
    if err = a.client.Post(serviceGroup, commandCreateGroup, req, resp); err != nil {
        return
    } else {
        groupId = resp.GroupId
    }
    
    return
}

// DestroyGroup 解散群组
// App管理员通过该接口解散群。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1624
func (a *api) DestroyGroup(groupId string) (err error) {
    req := destroyGroupReq{GroupId: groupId}
    
    if err = a.client.Post(serviceGroup, commandDestroyGroup, req, &types.ActionBaseResp{}); err != nil {
        return
    }
    
    return
}

// GetGroup 获取单个群详细资料
// 本方法由“获取多个群详细资料（GetGroups）”拓展而来
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1616
func (a *api) GetGroup(groupId string, filter ...*Filter) (group *Group, err error) {
    var groups []*Group
    
    if groups, err = a.GetGroups([]string{groupId}, filter...); err != nil {
        return
    }
    
    if groups != nil && len(groups) > 0 {
        if err = groups[0].err; err != nil {
            return
        }
        
        group = groups[0]
    }
    
    return
}

// GetGroups 获取多个群详细资料
// App 管理员可以根据群组 ID 获取群组的详细信息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1616
func (a *api) GetGroups(groupIds []string, filter ...*Filter) (groups []*Group, err error) {
    if len(groupIds) > 0 {
        req := getGroupsReq{GroupIds: groupIds}
        resp := &getGroupsResp{}
        
        if len(filter) > 0 {
            req.ResponseFilter = &responseFilter{
                GroupBaseInfoFilter:    filter[0].GetAllBaseInfoFilterFields(),
                MemberInfoFilter:       filter[0].GetAllMemberInfoFilterFields(),
                GroupCustomDataFilter:  filter[0].GetAllGroupCustomDataFilterFields(),
                MemberCustomDataFilter: filter[0].GetAllMemberCustomDataFilterFields(),
            }
        }
        
        if err = a.client.Post(serviceGroup, commandGetGroups, req, resp); err != nil {
            return
        }
        
        groups = make([]*Group, 0, len(resp.GroupInfos))
        for _, item := range resp.GroupInfos {
            group := NewGroup()
            group.setError(item.ErrorCode, item.ErrorInfo)
            if group.err == nil {
                group.id = item.GroupId
                group.name = item.Name
                group.types = item.Type
                group.owner = item.OwnerUserId
                group.avatar = item.FaceUrl
                group.memberNum = item.MemberNum
                group.maxMemberCount = item.MaxMemberNum
                group.applyJoinOption = item.ApplyJoinOption
                group.createTime = item.CreateTime
                group.lastInfoTime = item.LastInfoTime
                group.lastMsgTime = item.LastMsgTime
                group.shutUpStatus = item.ShutUpAllMember
                group.nextMsgSeq = item.NextMsgSeq
                
                if item.AppDefinedData != nil && len(item.AppDefinedData) > 0 {
                    for _, v := range item.AppDefinedData {
                        group.SetCustomData(v.Key, v.Value)
                    }
                }
                
                if item.MemberList != nil && len(item.MemberList) > 0 {
                    for _, m := range item.MemberList {
                        member := &Member{
                            userId:          m.UserId,
                            role:            m.Role,
                            joinTime:        m.JoinTime,
                            nameCard:        m.NameCard,
                            msgSeq:          m.MsgSeq,
                            msgFlag:         m.MsgFlag,
                            lastSendMsgTime: m.LastSendMsgTime,
                        }
                        
                        if m.AppMemberDefinedData != nil && len(m.AppMemberDefinedData) > 0 {
                            for _, v := range m.AppMemberDefinedData {
                                member.SetCustomData(v.Key, v.Value)
                            }
                        }
                        
                        group.AddMembers(member)
                    }
                }
                
                groups = append(groups, group)
            }
        }
    }
    
    return
}

// GetGroupMemberInfo Get group member's data based on group ID.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1617
func (a *api) GetGroupMemberInfo(req *GetGroupMemberInfoReq) (*GetGroupMemberInfoResp, error) {
    resp := &GetGroupMemberInfoResp{}
    
    if err := a.client.Post(serviceGroup, commandGetGroupMemberInfo, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// ModifyGroupBaseInfo Modify the basic information of the specified group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1620
func (a *api) ModifyGroupBaseInfo(req *ModifyGroupBaseInfoReq) (*ModifyGroupBaseInfoResp, error) {
    resp := &ModifyGroupBaseInfoResp{}
    
    if err := a.client.Post(serviceGroup, commandModifyGroupBaseInfo, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// AddGroupMembers 增加群成员
// App管理员可以通过该接口向指定的群中添加新成员。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1621
func (a *api) AddGroupMembers(groupId string, userIds []string, silence ...bool) (results []AddMembersResult, err error) {
    req := addGroupMembersReq{}
    req.GroupId = groupId
    req.MemberList = make([]addMemberItem, 0, len(userIds))
    for _, userId := range userIds {
        req.MemberList = append(req.MemberList, addMemberItem{
            UserId: userId,
        })
    }
    if len(silence) > 0 && silence[0] {
        req.Silence = 1
    }
    
    resp := &addGroupMembersResp{}
    
    if err = a.client.Post(serviceGroup, commandAddGroupMembers, req, resp); err != nil {
        return
    }
    
    results = resp.MemberList
    
    return
}

// DeleteGroupMembers 删除群成员
// App管理员可以通过该接口删除群成员。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1622
func (a *api) DeleteGroupMembers(groupId string, userIds []string, reasonAndSilence ...interface{}) (err error) {
    req := deleteGroupMembersReq{}
    req.GroupId = groupId
    req.UserIds = userIds
    
    if len(reasonAndSilence) > 0 {
        for i, val := range reasonAndSilence {
            if i > 1 {
                break
            }
            
            switch v := val.(type) {
            case string:
                req.Reason = v
            case bool:
                if v {
                    req.Silence = 1
                }
            }
        }
    }
    
    if err = a.client.Post(serviceGroup, commandDeleteGroupMember, req, &types.ActionBaseResp{}); err != nil {
        return
    }
    
    return
}

// ModifyGroupMemberInfo Modify group member information.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1623
func (a *api) ModifyGroupMemberInfo(req *ModifyGroupMemberInfoReq) (*ModifyGroupMemberInfoResp, error) {
    resp := &ModifyGroupMemberInfoResp{}
    
    if err := a.client.Post(serviceGroup, commandModifyGroupMemberInfo, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// GetJoinedGroupList Get group information that a user has joined.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1625
func (a *api) GetJoinedGroupList(req *GetJoinedGroupListReq) (*GetJoinedGroupListResp, error) {
    resp := &GetJoinedGroupListResp{}
    
    if err := a.client.Post(serviceGroup, commandGetJoinedGroupList, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// GetRoleInGroup Get the identities of a batch of users in the group, that is, "member roles".
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1626
func (a *api) GetRoleInGroup(req *GetRoleInGroupReq) (*GetRoleInGroupResp, error) {
    resp := &GetRoleInGroupResp{}
    
    if err := a.client.Post(serviceGroup, commandGetRoleInGroup, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// ForbidSendMsg Mute and unmute a group of users.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1627
func (a *api) ForbidSendMsg(req *ForbidSendMsgReq) (*ForbidSendMsgResp, error) {
    resp := &ForbidSendMsgResp{}
    
    if err := a.client.Post(serviceGroup, commandForbidSendMsg, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// GetGroupShuttedUin Get the list of banned users in the group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/2925
func (a *api) GetGroupShuttedUin(req *GetGroupShuttedUinReq) (*GetGroupShuttedUinResp, error) {
    resp := &GetGroupShuttedUinResp{}
    
    if err := a.client.Post(serviceGroup, commandGetGroupShuttedUin, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// SendGroupMsg Send a normal message in the group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1629
func (a *api) SendGroupMsg(req *SendGroupMsgReq) (*SendGroupMsgResp, error) {
    resp := &SendGroupMsgResp{}
    
    if err := a.client.Post(serviceGroup, commandSendGroupMsg, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// SendGroupSystemNotification Send system notifications in groups.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1630
func (a *api) SendGroupSystemNotification(req *SendGroupSystemNotificationReq) (*SendGroupSystemNotificationResp, error) {
    resp := &SendGroupSystemNotificationResp{}
    
    if err := a.client.Post(serviceGroup, commandSendGroupSystemNotification, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// ChangeGroupOwner Transfer the identity of the group owner to others.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1633
func (a *api) ChangeGroupOwner(req *ChangeGroupOwnerReq) (*ChangeGroupOwnerResp, error) {
    resp := &ChangeGroupOwnerResp{}
    
    if err := a.client.Post(serviceGroup, commandChangeGroupOwner, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// RecallGroupMsg Withdraw the message of the specified group, the message must be within the roaming validity period.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/12341
func (a *api) RecallGroupMsg(req *RecallGroupMsgReq) (*RecallGroupMsgResp, error) {
    resp := &RecallGroupMsgResp{}
    
    if err := a.client.Post(serviceGroup, commandRecallGroupMsg, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// ImportGroup Import group basic information.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1634
func (a *api) ImportGroup(req *ImportGroupReq) (*ImportGroupResp, error) {
    resp := &ImportGroupResp{}
    
    if err := a.client.Post(serviceGroup, commandImportGroup, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// ImportGroupMsg Import group messages.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1635
func (a *api) ImportGroupMsg(req *ImportGroupMsgReq) (*ImportGroupMsgResp, error) {
    resp := &ImportGroupMsgResp{}
    
    if err := a.client.Post(serviceGroup, commandImportGroupMsg, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// ImportGroupMember Import group members.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1636
func (a *api) ImportGroupMember(req *ImportGroupMemberReq) (*ImportGroupMemberResp, error) {
    resp := &ImportGroupMemberResp{}
    
    if err := a.client.Post(serviceGroup, commandImportGroupMember, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// SetUnreadMsgNum Set member unread message count.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1637
func (a *api) SetUnreadMsgNum(req *SetUnreadMsgNumReq) (*SetUnreadMsgNumResp, error) {
    resp := &SetUnreadMsgNumResp{}
    
    if err := a.client.Post(serviceGroup, commandSetUnreadMsgNum, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// DeleteGroupMsgBySender Withdraw the message sent by the specified user.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1637
func (a *api) DeleteGroupMsgBySender(req *DeleteGroupMsgBySenderReq) (*DeleteGroupMsgBySenderResp, error) {
    resp := &DeleteGroupMsgBySenderResp{}
    
    if err := a.client.Post(serviceGroup, commandDeleteGroupMsgBySender, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// GetGroupSimpleMsg Pull group history messages.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/2738
func (a *api) GetGroupSimpleMsg(req *GetGroupSimpleMsgReq) (*GetGroupSimpleMsgResp, error) {
    resp := &GetGroupSimpleMsgResp{}
    
    if err := a.client.Post(serviceGroup, commandGetGroupSimpleMsg, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}

// GetOnlineMemberNum Get the number of people online in the live broadcast group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/49180
func (a *api) GetOnlineMemberNum(req *GetOnlineMemberNumReq) (*GetOnlineMemberNumResp, error) {
    resp := &GetOnlineMemberNumResp{}
    
    if err := a.client.Post(serviceGroup, commandGetOnlineMemberNum, req, resp); err != nil {
        return nil, err
    }
    
    return resp, nil
}
