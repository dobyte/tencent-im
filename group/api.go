/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:46
 * @Desc: Group Api Implementation.
 */

package group

import (
	"fmt"

	"github.com/dobyte/tencent-im/internal/conv"
	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/internal/enum"
	"github.com/dobyte/tencent-im/internal/types"
)

const (
	serviceGroup                       = "group_open_http_svc"
	commandFetchGroupIds               = "get_appid_group_list"
	commandCreateGroup                 = "create_group"
	commandDestroyGroup                = "destroy_group"
	commandGetGroups                   = "get_group_info"
	commandFetchGroupMembers           = "get_group_member_info"
	commandUpdateGroup                 = "modify_group_base_info"
	commandAddGroupMembers             = "add_group_member"
	commandDeleteGroupMember           = "delete_group_member"
	commandModifyGroupMemberInfo       = "modify_group_member_info"
	commandFetchMemberGroups           = "get_joined_group_list"
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

	batchGetGroupsLimit = 50 // 批量获取群组限制
)

type API interface {
	// FetchGroupIds 拉取App中的所有群组ID
	// App 管理员可以通过该接口获取App中所有群组的ID。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1614
	FetchGroupIds(limit int, next int, groupType ...GroupType) (ret *FetchGroupIdsRet, err error)

	// FetchGroups 拉取App中的所有群组
	// 本方法由“拉取App中的所有群组ID（FetchGroupIds）”拓展而来
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1614
	FetchGroups(limit int, next int, groupTypeAndFilter ...interface{}) (ret *FetchGroupsRet, err error)

	// PullGroups 续拉取App中的所有群组
	// 本方法由“拉取App中的所有群组（FetchGroups）”拓展而来
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1614
	PullGroups(arg *PullGroupsArg, fn func(ret *FetchGroupsRet)) (err error)

	// CreateGroup 创建群组
	// App 管理员可以通过该接口创建群组。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1615
	CreateGroup(group *Group) (groupId string, err error)

	// GetGroup 获取单个群详细资料
	// 本方法由“获取多个群详细资料（GetGroups）”拓展而来
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1616
	GetGroup(groupId string, filter ...*Filter) (group *Group, err error)

	// GetGroups 获取多个群详细资料
	// App 管理员可以根据群组 ID 获取群组的详细信息。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1616
	GetGroups(groupIds []string, filter ...*Filter) (groups []*Group, err error)

	// FetchMembers 拉取群成员详细资料
	// App管理员可以根据群组ID获取群组成员的资料。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1617
	FetchMembers(groupId string, limit, offset int, filter ...*Filter) (ret *FetchMembersRet, err error)

	// PullMembers 续拉取群成员详细资料
	// 本方法由“拉取群成员详细资料（FetchMembers）”拓展而来
	// App管理员可以根据群组ID获取群组成员的资料。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1617
	PullMembers(arg *PullMembersArg, fn func(ret *FetchMembersRet)) (err error)

	// UpdateGroup 修改群基础资料
	// App管理员可以通过该接口修改指定群组的基础信息。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1620
	UpdateGroup(group *Group) (err error)

	// AddMembers 增加群成员
	// App管理员可以通过该接口向指定的群中添加新成员。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1621
	AddMembers(groupId string, userIds []string, silence ...bool) (results []AddMembersResult, err error)

	// DeleteMembers 删除群成员
	// App管理员可以通过该接口删除群成员。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1622
	DeleteMembers(groupId string, userIds []string, reasonAndSilence ...interface{}) (err error)

	// UpdateMember 修改群成员资料
	// App管理员可以通过该接口修改群成员资料。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1623
	UpdateMember(groupId string, member *Member) (err error)

	// DestroyGroup 解散群组
	// App管理员通过该接口解散群。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1624
	DestroyGroup(groupId string) (err error)

	// FetchMemberGroups 拉取用户所加入的群组
	// App管理员可以通过本接口获取某一用户加入的群信息。默认不获取用户已加入但未激活好友工作群（Work）以及直播群（AVChatRoom）群信息。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1625
	FetchMemberGroups(arg *FetchMemberGroupsArg) (ret *FetchMemberGroupsRet, err error)

	// PullMemberGroups 续拉取用户所加入的群组
	// 本方法由“拉取用户所加入的群组（FetchMemberGroups）”拓展而来
	// App管理员可以通过本接口获取某一用户加入的群信息。默认不获取用户已加入但未激活好友工作群（Work）以及直播群（AVChatRoom）群信息。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1625
	PullMemberGroups(arg *PullMemberGroupsArg, fn func(ret *FetchMemberGroupsRet)) (err error)

	// GetRolesInGroup 查询用户在群组中的身份
	// App管理员可以通过该接口获取一批用户在群内的身份，即“成员角色”。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1626
	GetRolesInGroup(groupId string, userIds []string) (memberRoles map[string]string, err error)

	// ForbidSendMessage 批量禁言
	// App 管理员禁止指定群组中某些用户在一段时间内发言。
	// App 管理员取消对某些用户的禁言。
	// 被禁言用户退出群组之后再进入同一群组，禁言仍然有效。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1627
	ForbidSendMessage(groupId string, userIds []string, shutUpTime int64) (err error)

	// AllowSendMessage 取消禁言
	// 本方法由“批量禁言（ForbidSendMessage）”拓展而来
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1627
	AllowSendMessage(groupId string, userIds []string) (err error)

	// GetShuttedUpMembers 获取被禁言群成员列表
	// App管理员可以根据群组ID获取群组中被禁言的用户列表。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/2925
	GetShuttedUpMembers(groupId string) (shuttedUps map[string]int64, err error)

	// SendMessage 在群组中发送普通消息
	// App管理员可以通过该接口在群组中发送普通消息。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1629
	SendMessage(groupId string, message *Message) (ret *SendMessageRet, err error)

	// SendNotification 在群组中发送系统通知
	// App 管理员可以通过该接口在群组中发送系统通知。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1630
	SendNotification(groupId, content string, userId ...string) (err error)

	// ChangeGroupOwner 转让群主
	// App 管理员可以通过该接口将群主身份转移给他人。
	// 没有群主的群，App 管理员可以通过此接口指定他人作为群主。
	// 新群主必须为群内成员。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1633
	ChangeGroupOwner(groupId, userId string) (err error)

	// RevokeMessage 撤回单条群消息
	// 本方法由“撤回多条群消息（RevokeMessages）”拓展而来
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/12341
	RevokeMessage(groupId string, msgSeq int) (err error)

	// RevokeMessages 撤回多条群消息
	// App 管理员通过该接口撤回指定群组的消息，消息需要在漫游有效期以内。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/12341
	RevokeMessages(groupId string, msgSeq ...int) (results map[int]int, err error)

	// ImportGroup 导入群基础资料
	// App 管理员可以通过该接口导入群组，不会触发回调、不会下发通知；当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群组数据。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1634
	ImportGroup(group *Group) (groupId string, err error)

	// ImportMessages 导入群消息
	// 该 API 接口的作用是导入群组的消息，不会触发回调、不会下发通知。
	// 当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群消息数据。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1635
	ImportMessages(groupId string, messages ...*Message) (results []ImportMessagesResult, err error)

	// ImportMembers 导入多个群成员
	// 该 API 接口的作用是导入群组成员，不会触发回调、不会下发通知。
	// 当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群成员数据。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1636
	ImportMembers(groupId string, members ...*Member) (results []ImportMemberResult, err error)

	// SetMemberUnreadMsgNum 设置成员未读消息计数
	// App管理员使用该接口设置群组成员未读消息数，不会触发回调、不会下发通知。
	// 当App需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议设置群成员的未读消息计数。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1637
	SetMemberUnreadMsgNum(groupId, userId string, unreadMsgNum int) (err error)

	// RevokeMemberMessages 撤回指定用户发送的消息
	// 该API接口的作用是撤回最近1000条消息中指定用户发送的消息。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/2359
	RevokeMemberMessages(groupId, userId string) (err error)

	// FetchMessages 拉取群历史消息
	// 即时通信 IM 的群消息是按 Seq 排序的，按照 server 收到群消息的顺序分配 Seq，先发的群消息 Seq 小，后发的 Seq 大。
	// 如果用户想拉取一个群的全量消息，首次拉取时不用填拉取 Seq，Server 会自动返回最新的消息，以后拉取时拉取 Seq 填上次返回的最小 Seq 减1。
	// 如果返回消息的 IsPlaceMsg 为1，表示这个 Seq 的消息或者过期、或者存储失败、或者被删除了。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/2738
	FetchMessages(groupId string, limit int, msgSeq ...int) (ret *FetchMessagesRet, err error)

	// PullMessages 续拉取群历史消息
	// 本方法由“拉取群历史消息（FetchMessages）”拓展而来
	// 即时通信 IM 的群消息是按 Seq 排序的，按照 server 收到群消息的顺序分配 Seq，先发的群消息 Seq 小，后发的 Seq 大。
	// 如果用户想拉取一个群的全量消息，首次拉取时不用填拉取 Seq，Server 会自动返回最新的消息，以后拉取时拉取 Seq 填上次返回的最小 Seq 减1。
	// 如果返回消息的 IsPlaceMsg 为1，表示这个 Seq 的消息或者过期、或者存储失败、或者被删除了。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/2738
	PullMessages(groupId string, limit int, fn func(ret *FetchMessagesRet)) (err error)

	// GetOnlineMemberNum 获取直播群在线人数
	// App 管理员可以根据群组 ID 获取直播群在线人数。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/49180
	GetOnlineMemberNum(groupId string) (num int, err error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// FetchGroupIds 拉取App中的所有群组ID
// App 管理员可以通过该接口获取App中所有群组的ID。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1614
func (a *api) FetchGroupIds(limit int, next int, groupType ...GroupType) (ret *FetchGroupIdsRet, err error) {
	req := &fetchGroupIdsReq{Limit: limit, Next: next}

	if len(groupType) > 0 {
		req.GroupType = string(groupType[0])
	}

	resp := &fetchGroupIdsResp{}

	if err = a.client.Post(serviceGroup, commandFetchGroupIds, req, resp); err != nil {
		return
	}

	ret = &FetchGroupIdsRet{}
	ret.Next = resp.Next
	ret.Total = resp.TotalCount
	ret.HasMore = ret.Next != 0
	ret.List = make([]string, 0, len(resp.GroupIdList))

	for _, item := range resp.GroupIdList {
		ret.List = append(ret.List, item.GroupId)
	}

	return
}

// FetchGroups 拉取App中的所有群组
// 本方法由“拉取App中的所有群组ID（FetchGroupIds）”拓展而来
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1614
func (a *api) FetchGroups(limit int, next int, groupTypeAndFilter ...interface{}) (ret *FetchGroupsRet, err error) {
	if limit > batchGetGroupsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of groups id cannot exceed %d", batchGetGroupsLimit))
		return
	}

	var (
		resp      *FetchGroupIdsRet
		filter    *Filter
		groupType GroupType
	)

	if len(groupTypeAndFilter) > 0 {
		for i, val := range groupTypeAndFilter {
			if i > 1 {
				break
			}
			switch v := val.(type) {
			case GroupType:
				groupType = v
			case *Filter:
				filter = v
			}
		}
	}

	if resp, err = a.FetchGroupIds(limit, next, groupType); err != nil {
		return
	}

	ret = &FetchGroupsRet{Next: resp.Next, Total: resp.Total, HasMore: resp.HasMore}

	if len(resp.List) > 0 {
		if ret.List, err = a.GetGroups(resp.List, filter); err != nil {
			return
		}
	}

	return
}

// PullGroups 续拉取App中的所有群组
// 本方法由“拉取App中的所有群组（FetchGroups）”拓展而来
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1614
func (a *api) PullGroups(arg *PullGroupsArg, fn func(ret *FetchGroupsRet)) (err error) {
	var (
		limit     = arg.Limit
		groupType = arg.GroupType
		filter    = arg.Filter
		next      int
		ret       *FetchGroupsRet
	)

	for ret == nil || ret.HasMore {
		ret, err = a.FetchGroups(limit, next, groupType, filter)
		if err != nil {
			return
		}

		fn(ret)

		if ret.HasMore {
			next = ret.Next
			break
		}
	}

	return
}

// CreateGroup 创建群组
// App管理员可以通过该接口创建群组。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1615
func (a *api) CreateGroup(group *Group) (groupId string, err error) {
	if err = group.checkCreateError(); err != nil {
		return
	}

	req := &createGroupReq{}
	req.GroupId = group.id
	req.OwnerUserId = group.owner
	req.GroupType = group.groupType
	req.Name = group.name
	req.FaceUrl = group.avatar
	req.Introduction = group.introduction
	req.Notification = group.notification
	req.MaxMemberNum = group.maxMemberNum
	req.ApplyJoinOption = group.applyJoinOption

	if data := group.GetAllCustomData(); data != nil {
		req.AppDefinedData = make([]*customDataItem, 0, len(data))
		for key, val := range data {
			req.AppDefinedData = append(req.AppDefinedData, &customDataItem{
				Key:   key,
				Value: val,
			})
		}
	}

	if c := len(group.members); c > 0 {
		req.MemberList = make([]*memberItem, 0, c)

		var item *memberItem
		for _, member := range group.members {
			if err = member.checkError(); err != nil {
				return
			}

			item = &memberItem{
				UserId:               member.userId,
				Role:                 member.role,
				JoinTime:             member.joinTime,
				NameCard:             member.nameCard,
				AppMemberDefinedData: make([]*customDataItem, 0, len(member.GetAllCustomData())),
			}

			for k, v := range member.GetAllCustomData() {
				item.AppMemberDefinedData = append(item.AppMemberDefinedData, &customDataItem{
					Key:   k,
					Value: v,
				})
			}

			req.MemberList = append(req.MemberList, item)
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

// GetGroup 获取单个群详细资料
// 本方法由“获取多个群详细资料（GetGroups）”拓展而来
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1616
func (a *api) GetGroup(groupId string, filter ...*Filter) (group *Group, err error) {
	var groups []*Group

	if groups, err = a.GetGroups([]string{groupId}, filter...); err != nil {
		return
	}

	if len(groups) > 0 {
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
func (a *api) GetGroups(groupIds []string, filters ...*Filter) (groups []*Group, err error) {
	if c := len(groupIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the group's id is not set")
		return
	} else if c > batchGetGroupsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of group's id cannot exceed %d", batchGetGroupsLimit))
		return
	}

	req := &getGroupsReq{GroupIds: groupIds}
	resp := &getGroupsResp{}

	if len(filters) > 0 {
		if filter := filters[0]; filter != nil {
			req.ResponseFilter = &responseFilter{
				GroupBaseInfoFilter:    filter.GetAllBaseInfoFilterFields(),
				MemberInfoFilter:       filter.GetAllMemberInfoFilterFields(),
				GroupCustomDataFilter:  filter.GetAllGroupCustomDataFilterFields(),
				MemberCustomDataFilter: filter.GetAllMemberCustomDataFilterFields(),
			}
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
			group.groupType = item.GroupType
			group.owner = item.OwnerUserId
			group.avatar = item.FaceUrl
			group.memberNum = item.MemberNum
			group.maxMemberNum = item.MaxMemberNum
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
						msgFlag:         MsgFlag(m.MsgFlag),
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

	return
}

// FetchMembers 拉取群成员详细资料
// App管理员可以根据群组ID获取群组成员的资料。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1617
func (a *api) FetchMembers(groupId string, limit, offset int, filters ...*Filter) (ret *FetchMembersRet, err error) {
	req := &fetchMembersReq{GroupId: groupId, Limit: limit, Offset: offset}

	if len(filters) > 0 {
		if filter := filters[0]; filter != nil {
			req.MemberInfoFilter = filter.GetAllMemberInfoFilterFields()
			req.MemberRoleFilter = filter.GetAllMemberRoleFilterValues()
			req.MemberCustomDataFilter = filter.GetAllMemberCustomDataFilterFields()
		}
	}

	resp := &fetchMembersResp{}

	if err = a.client.Post(serviceGroup, commandFetchGroupMembers, req, resp); err != nil {
		return
	}

	ret = &FetchMembersRet{}
	ret.Total = resp.MemberNum
	ret.List = make([]*Member, 0, len(resp.MemberList))
	ret.HasMore = resp.MemberNum > limit+offset

	for _, m := range resp.MemberList {
		member := &Member{
			userId:          m.UserId,
			role:            m.Role,
			joinTime:        m.JoinTime,
			nameCard:        m.NameCard,
			msgSeq:          m.MsgSeq,
			msgFlag:         MsgFlag(m.MsgFlag),
			lastSendMsgTime: m.LastSendMsgTime,
		}

		if m.AppMemberDefinedData != nil && len(m.AppMemberDefinedData) > 0 {
			for _, v := range m.AppMemberDefinedData {
				member.SetCustomData(v.Key, v.Value)
			}
		}

		ret.List = append(ret.List, member)
	}

	return
}

// PullMembers 续拉取群成员详细资料
// 本方法由“拉取群成员详细资料（FetchMembers）”拓展而来
// App管理员可以根据群组ID获取群组成员的资料。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1617
func (a *api) PullMembers(arg *PullMembersArg, fn func(ret *FetchMembersRet)) (err error) {
	var (
		offset int
		ret    *FetchMembersRet
	)

	for ret == nil || ret.HasMore {
		ret, err = a.FetchMembers(arg.GroupId, arg.Limit, offset, arg.Filter)
		if err != nil {
			return
		}

		fn(ret)

		if ret.HasMore {
			offset += arg.Limit
		}
	}

	return
}

// UpdateGroup 修改群基础资料
// App管理员可以通过该接口修改指定群组的基础信息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1620
func (a *api) UpdateGroup(group *Group) (err error) {
	if err = group.checkUpdateError(); err != nil {
		return
	}

	req := &updateGroupReq{}
	req.GroupId = group.id
	req.Name = group.name
	req.FaceUrl = group.avatar
	req.Introduction = group.introduction
	req.Notification = group.notification
	req.MaxMemberNum = group.maxMemberNum
	req.ApplyJoinOption = group.applyJoinOption
	req.ShutUpAllMember = group.shutUpStatus

	if data := group.GetAllCustomData(); data != nil {
		req.AppDefinedData = make([]customDataItem, 0, len(data))
		for key, val := range data {
			req.AppDefinedData = append(req.AppDefinedData, customDataItem{
				Key:   key,
				Value: val,
			})
		}
	}

	if err = a.client.Post(serviceGroup, commandUpdateGroup, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// AddMembers 增加群成员
// App管理员可以通过该接口向指定的群中添加新成员。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1621
func (a *api) AddMembers(groupId string, userIds []string, silence ...bool) (results []AddMembersResult, err error) {
	req := &addMembersReq{}
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

	resp := &addMembersResp{}

	if err = a.client.Post(serviceGroup, commandAddGroupMembers, req, resp); err != nil {
		return
	}

	results = resp.MemberList

	return
}

// DeleteMembers 删除群成员
// App管理员可以通过该接口删除群成员。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1622
func (a *api) DeleteMembers(groupId string, userIds []string, reasonAndSilence ...interface{}) (err error) {
	req := &deleteMembersReq{}
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

// UpdateMember 修改群成员资料
// App管理员可以通过该接口修改群成员资料。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1623
func (a *api) UpdateMember(groupId string, member *Member) (err error) {
	if err = member.checkError(); err != nil {
		return
	}

	req := &updateMemberReq{}
	req.GroupId = groupId
	req.UserId = member.userId
	req.Role = member.role
	req.MsgFlag = string(member.msgFlag)
	req.NameCard = member.nameCard
	req.ShutUpUntil = member.shutUpUntil

	if data := member.GetAllCustomData(); data != nil {
		req.AppMemberDefinedData = make([]customDataItem, 0, len(data))
		for key, val := range data {
			req.AppMemberDefinedData = append(req.AppMemberDefinedData, customDataItem{
				Key:   key,
				Value: val,
			})
		}
	}

	if err = a.client.Post(serviceGroup, commandModifyGroupMemberInfo, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// DestroyGroup 解散群组
// App管理员通过该接口解散群。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1624
func (a *api) DestroyGroup(groupId string) (err error) {
	req := &destroyGroupReq{GroupId: groupId}

	if err = a.client.Post(serviceGroup, commandDestroyGroup, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// FetchMemberGroups 拉取用户所加入的群组
// App管理员可以通过本接口获取某一用户加入的群信息。默认不获取用户已加入但未激活好友工作群（Work）以及直播群（AVChatRoom）群信息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1625
func (a *api) FetchMemberGroups(arg *FetchMemberGroupsArg) (ret *FetchMemberGroupsRet, err error) {
	req := &fetchMemberGroupsReq{UserId: arg.UserId, Limit: arg.Limit, Offset: arg.Offset, GroupType: arg.GroupType}

	if arg.Filter != nil {
		req.ResponseFilter = &responseFilter{
			GroupBaseInfoFilter: arg.Filter.GetAllBaseInfoFilterFields(),
			SelfInfoFilter:      arg.Filter.GetAllMemberInfoFilterFields(),
		}
	}

	if arg.IsWithNoActiveGroups {
		req.WithNoActiveGroups = 1
	}

	if arg.IsWithLiveRoomGroups {
		req.WithHugeGroups = 1
	}

	resp := &fetchMemberGroupsResp{}

	if err = a.client.Post(serviceGroup, commandFetchMemberGroups, req, resp); err != nil {
		return
	}

	ret = &FetchMemberGroupsRet{}
	ret.Total = resp.TotalCount
	ret.List = make([]*Group, 0, len(resp.GroupList))

	if arg.Limit == 0 {
		ret.HasMore = false
	} else {
		ret.HasMore = arg.Limit+arg.Offset < resp.TotalCount
	}

	for _, item := range resp.GroupList {
		group := NewGroup()
		group.id = item.GroupId
		group.name = item.Name
		group.groupType = item.GroupType
		group.owner = item.OwnerUserId
		group.avatar = item.FaceUrl
		group.memberNum = item.MemberNum
		group.maxMemberNum = item.MaxMemberNum
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

		if item.MemberInfo != nil {
			member := &Member{
				userId:          arg.UserId,
				role:            item.MemberInfo.Role,
				joinTime:        item.MemberInfo.JoinTime,
				nameCard:        item.MemberInfo.NameCard,
				msgSeq:          item.MemberInfo.MsgSeq,
				msgFlag:         MsgFlag(item.MemberInfo.MsgFlag),
				lastSendMsgTime: item.MemberInfo.LastSendMsgTime,
			}

			if item.MemberInfo.AppMemberDefinedData != nil && len(item.MemberInfo.AppMemberDefinedData) > 0 {
				for _, v := range item.MemberInfo.AppMemberDefinedData {
					member.SetCustomData(v.Key, v.Value)
				}
			}

			group.AddMembers(member)
		}

		ret.List = append(ret.List, group)
	}

	return
}

// PullMemberGroups 续拉取用户所加入的群组
// 本方法由“拉取用户所加入的群组（FetchMemberGroups）”拓展而来
// App管理员可以通过本接口获取某一用户加入的群信息。默认不获取用户已加入但未激活好友工作群（Work）以及直播群（AVChatRoom）群信息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1625
func (a *api) PullMemberGroups(arg *PullMemberGroupsArg, fn func(ret *FetchMemberGroupsRet)) (err error) {
	var (
		ret *FetchMemberGroupsRet
		req = &FetchMemberGroupsArg{
			UserId:               arg.UserId,
			Limit:                arg.Limit,
			GroupType:            arg.GroupType,
			Filter:               arg.Filter,
			IsWithNoActiveGroups: arg.IsWithNoActiveGroups,
			IsWithLiveRoomGroups: arg.IsWithLiveRoomGroups,
		}
	)

	for ret == nil || ret.HasMore {
		ret, err = a.FetchMemberGroups(req)
		if err != nil {
			return
		}

		fn(ret)

		if ret.HasMore {
			req.Offset += arg.Limit
		}
	}

	return
}

// GetRolesInGroup 查询用户在群组中的身份
// App管理员可以通过该接口获取一批用户在群内的身份，即“成员角色”。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1626
func (a *api) GetRolesInGroup(groupId string, userIds []string) (roles map[string]string, err error) {
	req := &getRolesInGroupReq{GroupId: groupId, UserIds: userIds}
	resp := &getRolesInGroupResp{}

	if err = a.client.Post(serviceGroup, commandGetRoleInGroup, req, resp); err != nil {
		return
	}

	roles = make(map[string]string, len(resp.MemberRoleList))
	for _, item := range resp.MemberRoleList {
		roles[item.UserId] = item.Role
	}

	return
}

// ForbidSendMessage 批量禁言
// App 管理员禁止指定群组中某些用户在一段时间内发言。
// App 管理员取消对某些用户的禁言。
// 被禁言用户退出群组之后再进入同一群组，禁言仍然有效。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1627
func (a *api) ForbidSendMessage(groupId string, userIds []string, shutUpTime int64) (err error) {
	req := &forbidSendMessageReq{
		GroupId:    groupId,
		UserIds:    userIds,
		ShutUpTime: shutUpTime,
	}

	if err = a.client.Post(serviceGroup, commandForbidSendMsg, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// AllowSendMessage 取消禁言
// 本方法由“批量禁言（ForbidSendMessage）”拓展而来
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1627
func (a *api) AllowSendMessage(groupId string, userIds []string) (err error) {
	return a.ForbidSendMessage(groupId, userIds, 0)
}

// GetShuttedUpMembers 获取被禁言群成员列表
// App管理员可以根据群组ID获取群组中被禁言的用户列表。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2925
func (a *api) GetShuttedUpMembers(groupId string) (shuttedUps map[string]int64, err error) {
	req := &getShuttedUpMembersReq{GroupId: groupId}
	resp := &getShuttedUpMembersResp{}

	if err = a.client.Post(serviceGroup, commandGetGroupShuttedUin, req, resp); err != nil {
		return
	}

	shuttedUps = make(map[string]int64)
	for _, item := range resp.ShuttedUpList {
		shuttedUps[item.UserId] = item.ShuttedUntil
	}

	return
}

// SendMessage 在群组中发送普通消息
// App管理员可以通过该接口在群组中发送普通消息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1629
func (a *api) SendMessage(groupId string, message *Message) (ret *SendMessageRet, err error) {
	if err = message.checkSendError(); err != nil {
		return
	}

	req := &sendMessageReq{}
	req.GroupId = groupId
	req.FromUserId = message.GetSender()
	req.OfflinePushInfo = message.GetOfflinePushInfo()
	req.MsgPriority = string(message.GetPriority())
	req.MsgBody = message.GetBody()
	req.Random = message.GetRandom()
	req.CustomData = conv.String(message.GetCustomData())
	req.SendMsgControl = message.GetSendMsgControl()
	req.ForbidCallbackControl = message.GetForbidCallbackControl()
	req.OnlineOnlyFlag = int(message.GetOnlineOnlyFlag())

	if message.atMembers != nil && len(message.atMembers) > 0 {
		req.GroupAtInfo = make([]atInfo, 0, len(message.atMembers))

		for userId, _ := range message.atMembers {
			if userId == AtAllMembersFlag {
				req.GroupAtInfo = append(req.GroupAtInfo, atInfo{
					GroupAtAllFlag: 1,
				})
			} else {
				req.GroupAtInfo = append(req.GroupAtInfo, atInfo{
					GroupAtAllFlag: 0,
					GroupAtUserId:  userId,
				})
			}
		}
	}

	resp := &sendMessageResp{}

	if err = a.client.Post(serviceGroup, commandSendGroupMsg, req, resp); err != nil {
		return
	} else {
		ret = &SendMessageRet{
			MsgSeq:  resp.MsgSeq,
			MsgTime: resp.MsgTime,
		}
	}

	return
}

// SendNotification 在群组中发送系统通知
// App 管理员可以通过该接口在群组中发送系统通知。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1630
func (a *api) SendNotification(groupId, content string, userIds ...string) (err error) {
	req := &sendNotificationReq{GroupId: groupId, Content: content, UserIds: userIds}

	if err = a.client.Post(serviceGroup, commandSendGroupSystemNotification, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// ChangeGroupOwner 转让群主
// App 管理员可以通过该接口将群主身份转移给他人。
// 没有群主的群，App 管理员可以通过此接口指定他人作为群主。
// 新群主必须为群内成员。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1633
func (a *api) ChangeGroupOwner(groupId, userId string) (err error) {
	req := &changeGroupOwnerReq{GroupId: groupId, OwnerUserId: userId}

	if err = a.client.Post(serviceGroup, commandChangeGroupOwner, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// RevokeMessage 撤回单条群消息
// 本方法由“撤回多条群消息（RevokeMessages）”拓展而来
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/12341
func (a *api) RevokeMessage(groupId string, msgSeq int) (err error) {
	var results map[int]int

	if results, err = a.RevokeMessages(groupId, msgSeq); err != nil {
		return
	}

	if results[msgSeq] != enum.SuccessCode {
		err = core.NewError(results[msgSeq], "message revoke failed")
		return
	}

	return
}

// RevokeMessages 撤回多条群消息
// App 管理员通过该接口撤回指定群组的消息，消息需要在漫游有效期以内。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/12341
func (a *api) RevokeMessages(groupId string, msgSeq ...int) (results map[int]int, err error) {
	req := revokeMessagesReq{}
	req.GroupId = groupId
	req.MsgSeqList = make([]msgSeqItem, 0, len(msgSeq))
	for _, seq := range msgSeq {
		req.MsgSeqList = append(req.MsgSeqList, msgSeqItem{
			MsgSeq: seq,
		})
	}

	resp := &revokeMessagesResp{}

	if err = a.client.Post(serviceGroup, commandRecallGroupMsg, req, resp); err != nil {
		return
	}

	results = make(map[int]int)
	for _, item := range resp.Results {
		results[item.MsgSeq] = item.RetCode
	}

	return
}

// ImportGroup 导入群基础资料
// App 管理员可以通过该接口导入群组，不会触发回调、不会下发通知；当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群组数据。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1634
func (a *api) ImportGroup(group *Group) (groupId string, err error) {
	if err = group.checkImportError(); err != nil {
		return
	}

	req := &importGroupReq{}
	req.GroupId = group.id
	req.OwnerUserId = group.owner
	req.GroupType = group.groupType
	req.Name = group.name
	req.FaceUrl = group.avatar
	req.Introduction = group.introduction
	req.Notification = group.notification
	req.MaxMemberNum = group.maxMemberNum
	req.ApplyJoinOption = group.applyJoinOption
	req.CreateTime = group.createTime

	if data := group.GetAllCustomData(); data != nil {
		req.AppDefinedData = make([]*customDataItem, 0, len(data))
		for key, val := range data {
			req.AppDefinedData = append(req.AppDefinedData, &customDataItem{
				Key:   key,
				Value: val,
			})
		}
	}

	resp := &importGroupResp{}

	if err = a.client.Post(serviceGroup, commandImportGroup, req, resp); err != nil {
		return
	} else {
		groupId = resp.GroupId
	}

	return
}

// ImportMessages 导入群消息
// 该 API 接口的作用是导入群组的消息，不会触发回调、不会下发通知。
// 当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群消息数据。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1635
func (a *api) ImportMessages(groupId string, messages ...*Message) (results []ImportMessagesResult, err error) {
	req := &importMessagesReq{GroupId: groupId, Messages: make([]messageItem, 0, len(messages))}

	for _, message := range messages {
		if err = message.checkImportError(); err != nil {
			return
		}

		req.Messages = append(req.Messages, messageItem{
			FromUserId: message.GetSender(),
			MsgBody:    message.GetBody(),
			SendTime:   message.GetSendTime(),
			Random:     message.GetRandom(),
		})
	}

	resp := &importMessagesResp{}

	if err = a.client.Post(serviceGroup, commandImportGroupMsg, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// ImportMembers 导入多个群成员
// 该 API 接口的作用是导入群组成员，不会触发回调、不会下发通知。
// 当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群成员数据。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1636
func (a *api) ImportMembers(groupId string, members ...*Member) (results []ImportMemberResult, err error) {
	req := &importMembersReq{GroupId: groupId, Members: make([]*memberItem, 0, len(members))}
	resp := &importMembersResp{}

	for _, member := range members {
		req.Members = append(req.Members, &memberItem{
			UserId:       member.userId,
			Role:         member.role,
			JoinTime:     member.joinTime,
			UnreadMsgNum: member.unreadMsgNum,
		})
	}

	if err = a.client.Post(serviceGroup, commandImportGroupMember, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// SetMemberUnreadMsgNum 设置成员未读消息计数
// App管理员使用该接口设置群组成员未读消息数，不会触发回调、不会下发通知。
// 当App需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议设置群成员的未读消息计数。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1637
func (a *api) SetMemberUnreadMsgNum(groupId, userId string, unreadMsgNum int) (err error) {
	req := &setMemberUnreadMsgNumReq{GroupId: groupId, UserId: userId, UnreadMsgNum: unreadMsgNum}

	if err = a.client.Post(serviceGroup, commandSetUnreadMsgNum, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// RevokeMemberMessages 撤回指定用户发送的消息
// 该API接口的作用是撤回最近1000条消息中指定用户发送的消息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2359
func (a *api) RevokeMemberMessages(groupId, userId string) (err error) {
	req := &revokeMemberMessagesReq{GroupId: groupId, UserId: userId}

	if err = a.client.Post(serviceGroup, commandDeleteGroupMsgBySender, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// FetchMessages 拉取群历史消息
// 即时通信 IM 的群消息是按 Seq 排序的，按照 server 收到群消息的顺序分配 Seq，先发的群消息 Seq 小，后发的 Seq 大。
// 如果用户想拉取一个群的全量消息，首次拉取时不用填拉取 Seq，Server 会自动返回最新的消息，以后拉取时拉取 Seq 填上次返回的最小 Seq 减1。
// 如果返回消息的 IsPlaceMsg 为1，表示这个 Seq 的消息或者过期、或者存储失败、或者被删除了。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2738
func (a *api) FetchMessages(groupId string, limit int, msgSeq ...int) (ret *FetchMessagesRet, err error) {
	req := &fetchMessagesReq{GroupId: groupId, ReqMsgNumber: limit}

	if len(msgSeq) > 0 {
		req.ReqMsgSeq = msgSeq[0]
	}

	resp := &fetchMessagesResp{}

	if err = a.client.Post(serviceGroup, commandGetGroupSimpleMsg, req, resp); err != nil {
		return
	}

	ret = &FetchMessagesRet{}
	ret.IsFinished = resp.IsFinished

	if ret.IsFinished == 0 {
		ret.HasMore = true
	}

	if count := len(resp.RspMsgList); count > 0 {
		ret.NextSeq = resp.RspMsgList[count-1].MsgSeq - 1

		if ret.IsFinished == 1 && count == limit {
			ret.HasMore = true
		}
	}

	ret.List = make([]*Message, 0, len(resp.RspMsgList))
	for _, item := range resp.RspMsgList {
		message := NewMessage()
		message.SetSender(item.FromUserId)
		message.SetRandom(item.MsgRandom)
		message.seq = item.MsgSeq
		message.timestamp = item.MsgTimeStamp
		message.status = MsgStatus(item.IsPlaceMsg)
		switch item.MsgPriority {
		case 1:
			message.priority = MsgPriorityHigh
		case 2:
			message.priority = MsgPriorityNormal
		case 3:
			message.priority = MsgPriorityLow
		case 4:
			message.priority = MsgPriorityLowest
		}
	}

	return
}

// PullMessages 续拉取群历史消息
// 本方法由“拉取群历史消息（FetchMessages）”拓展而来
// 即时通信 IM 的群消息是按 Seq 排序的，按照 server 收到群消息的顺序分配 Seq，先发的群消息 Seq 小，后发的 Seq 大。
// 如果用户想拉取一个群的全量消息，首次拉取时不用填拉取 Seq，Server 会自动返回最新的消息，以后拉取时拉取 Seq 填上次返回的最小 Seq 减1。
// 如果返回消息的 IsPlaceMsg 为1，表示这个 Seq 的消息或者过期、或者存储失败、或者被删除了。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2738
func (a *api) PullMessages(groupId string, limit int, fn func(ret *FetchMessagesRet)) (err error) {
	var (
		ret    *FetchMessagesRet
		msgSeq int
	)

	for ret == nil || ret.HasMore {
		ret, err = a.FetchMessages(groupId, limit, msgSeq)
		if err != nil {
			return
		}

		fn(ret)

		if ret.HasMore {
			msgSeq = ret.NextSeq
		}
	}

	return
}

// GetOnlineMemberNum 获取直播群在线人数
// App 管理员可以根据群组 ID 获取直播群在线人数。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/49180
func (a *api) GetOnlineMemberNum(groupId string) (num int, err error) {
	req := &getOnlineMemberNumReq{GroupId: groupId}
	resp := &getOnlineMemberNumResp{}

	if err = a.client.Post(serviceGroup, commandGetOnlineMemberNum, req, resp); err != nil {
		return
	}

	num = resp.OnlineMemberNum

	return
}
