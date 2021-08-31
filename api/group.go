/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/5/27 20:46
 * @Desc: Group Api Implementation.
 */

package api

import (
	"github.com/dobyte/tencent-im/internal/core"
)

const (
	serviceGroup                       = "group_open_http_svc"
	commandGetGroupList                = "get_appid_group_list"
	commandCreateGroup                 = "create_group"
	commandDestroyGroup                = "destroy_group"
	commandGetGroupInfo                = "get_group_info"
	commandGetGroupMemberInfo          = "get_group_member_info"
	commandModifyGroupBaseInfo         = "modify_group_base_info"
	commandAddGroupMember              = "add_group_member"
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

type Group interface {
}

type group struct {
	client core.Client
}

func NewGroup(client core.Client) Group {
	return &group{client: client}
}

// GetGroupList Get the IDs of all groups in the app.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1614
func (a *group) GetGroupList(req *GetGroupListReq) (*GetGroupListResp, error) {
	resp := &GetGroupListResp{}
	
	if err := a.client.Post(serviceGroup, commandGetGroupList, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// CreateGroup Create a group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1615
func (a *group) CreateGroup(req *CreateGroupReq) (*CreateGroupResp, error) {
	resp := &CreateGroupResp{}
	
	if err := a.client.Post(serviceGroup, commandCreateGroup, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// DestroyGroup Destroy the specified group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1624
func (a *group) DestroyGroup(req *DestroyGroupReq) (*DestroyGroupResp, error) {
	resp := &DestroyGroupResp{}
	
	if err := a.client.Post(serviceGroup, commandDestroyGroup, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// GetGroupInfo Get group details based on group ID.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1616
func (a *group) GetGroupInfo(req *GetGroupInfoReq) (*GetGroupInfoResp, error) {
	resp := &GetGroupInfoResp{}
	
	if err := a.client.Post(serviceGroup, commandGetGroupInfo, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// GetGroupMemberInfo Get group member's data based on group ID.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1617
func (a *group) GetGroupMemberInfo(req *GetGroupMemberInfoReq) (*GetGroupMemberInfoResp, error) {
	resp := &GetGroupMemberInfoResp{}
	
	if err := a.client.Post(serviceGroup, commandGetGroupMemberInfo, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// ModifyGroupBaseInfo Modify the basic information of the specified group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1620
func (a *group) ModifyGroupBaseInfo(req *ModifyGroupBaseInfoReq) (*ModifyGroupBaseInfoResp, error) {
	resp := &ModifyGroupBaseInfoResp{}
	
	if err := a.client.Post(serviceGroup, commandModifyGroupBaseInfo, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// AddGroupMember Add new members to the specified group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1621
func (a *group) AddGroupMember(req *AddGroupMemberReq) (*AddGroupMemberResp, error) {
	resp := &AddGroupMemberResp{}
	
	if err := a.client.Post(serviceGroup, commandAddGroupMember, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// DeleteGroupMember Delete members from the specified group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1622
func (a *group) DeleteGroupMember(req *DeleteGroupMemberReq) (*DeleteGroupMemberResp, error) {
	resp := &DeleteGroupMemberResp{}
	
	if err := a.client.Post(serviceGroup, commandDeleteGroupMember, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// ModifyGroupMemberInfo Modify group member information.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1623
func (a *group) ModifyGroupMemberInfo(req *ModifyGroupMemberInfoReq) (*ModifyGroupMemberInfoResp, error) {
	resp := &ModifyGroupMemberInfoResp{}
	
	if err := a.client.Post(serviceGroup, commandModifyGroupMemberInfo, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// GetJoinedGroupList Get group information that a user has joined.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1625
func (a *group) GetJoinedGroupList(req *GetJoinedGroupListReq) (*GetJoinedGroupListResp, error) {
	resp := &GetJoinedGroupListResp{}
	
	if err := a.client.Post(serviceGroup, commandGetJoinedGroupList, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// GetRoleInGroup Get the identities of a batch of users in the group, that is, "member roles".
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1626
func (a *group) GetRoleInGroup(req *GetRoleInGroupReq) (*GetRoleInGroupResp, error) {
	resp := &GetRoleInGroupResp{}
	
	if err := a.client.Post(serviceGroup, commandGetRoleInGroup, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// ForbidSendMsg Mute and unmute a group of users.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1627
func (a *group) ForbidSendMsg(req *ForbidSendMsgReq) (*ForbidSendMsgResp, error) {
	resp := &ForbidSendMsgResp{}
	
	if err := a.client.Post(serviceGroup, commandForbidSendMsg, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// GetGroupShuttedUin Get the list of banned users in the group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/2925
func (a *group) GetGroupShuttedUin(req *GetGroupShuttedUinReq) (*GetGroupShuttedUinResp, error) {
	resp := &GetGroupShuttedUinResp{}
	
	if err := a.client.Post(serviceGroup, commandGetGroupShuttedUin, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// SendGroupMsg Send a normal message in the group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1629
func (a *group) SendGroupMsg(req *SendGroupMsgReq) (*SendGroupMsgResp, error) {
	resp := &SendGroupMsgResp{}
	
	if err := a.client.Post(serviceGroup, commandSendGroupMsg, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// SendGroupSystemNotification Send system notifications in groups.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1630
func (a *group) SendGroupSystemNotification(req *SendGroupSystemNotificationReq) (*SendGroupSystemNotificationResp, error) {
	resp := &SendGroupSystemNotificationResp{}
	
	if err := a.client.Post(serviceGroup, commandSendGroupSystemNotification, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// ChangeGroupOwner Transfer the identity of the group owner to others.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1633
func (a *group) ChangeGroupOwner(req *ChangeGroupOwnerReq) (*ChangeGroupOwnerResp, error) {
	resp := &ChangeGroupOwnerResp{}
	
	if err := a.client.Post(serviceGroup, commandChangeGroupOwner, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// RecallGroupMsg Withdraw the message of the specified group, the message must be within the roaming validity period.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/12341
func (a *group) RecallGroupMsg(req *RecallGroupMsgReq) (*RecallGroupMsgResp, error) {
	resp := &RecallGroupMsgResp{}
	
	if err := a.client.Post(serviceGroup, commandRecallGroupMsg, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// ImportGroup Import group basic information.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1634
func (a *group) ImportGroup(req *ImportGroupReq) (*ImportGroupResp, error) {
	resp := &ImportGroupResp{}
	
	if err := a.client.Post(serviceGroup, commandImportGroup, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// ImportGroupMsg Import group messages.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1635
func (a *group) ImportGroupMsg(req *ImportGroupMsgReq) (*ImportGroupMsgResp, error) {
	resp := &ImportGroupMsgResp{}
	
	if err := a.client.Post(serviceGroup, commandImportGroupMsg, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// ImportGroupMember Import group members.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1636
func (a *group) ImportGroupMember(req *ImportGroupMemberReq) (*ImportGroupMemberResp, error) {
	resp := &ImportGroupMemberResp{}
	
	if err := a.client.Post(serviceGroup, commandImportGroupMember, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// SetUnreadMsgNum Set member unread message count.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1637
func (a *group) SetUnreadMsgNum(req *SetUnreadMsgNumReq) (*SetUnreadMsgNumResp, error) {
	resp := &SetUnreadMsgNumResp{}
	
	if err := a.client.Post(serviceGroup, commandSetUnreadMsgNum, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// DeleteGroupMsgBySender Withdraw the message sent by the specified user.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/1637
func (a *group) DeleteGroupMsgBySender(req *DeleteGroupMsgBySenderReq) (*DeleteGroupMsgBySenderResp, error) {
	resp := &DeleteGroupMsgBySenderResp{}
	
	if err := a.client.Post(serviceGroup, commandDeleteGroupMsgBySender, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// GetGroupSimpleMsg Pull group history messages.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/2738
func (a *group) GetGroupSimpleMsg(req *GetGroupSimpleMsgReq) (*GetGroupSimpleMsgResp, error) {
	resp := &GetGroupSimpleMsgResp{}
	
	if err := a.client.Post(serviceGroup, commandGetGroupSimpleMsg, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}

// GetOnlineMemberNum Get the number of people online in the live broadcast group.
// click here to view the document:
// https://cloud.tencent.com/document/product/269/49180
func (a *group) GetOnlineMemberNum(req *GetOnlineMemberNumReq) (*GetOnlineMemberNumResp, error) {
	resp := &GetOnlineMemberNumResp{}
	
	if err := a.client.Post(serviceGroup, commandGetOnlineMemberNum, req, resp); err != nil {
		return nil, err
	}
	
	return resp, nil
}
