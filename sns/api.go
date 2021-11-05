/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:45
 * @Desc: 关系链管理
 */

package sns

import (
	"fmt"

	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/internal/enum"
	"github.com/dobyte/tencent-im/internal/types"
)

const (
	service                = "sns"
	commandAddFriend       = "friend_add"
	commandImportFriend    = "friend_import"
	commandUpdateFriend    = "friend_update"
	commandDeleteFriend    = "friend_delete"
	commandDeleteAllFriend = "friend_delete_all"
	commandCheckFriend     = "friend_check"
	commandGetFriend       = "friend_get_list"
	commandFetchFriend     = "friend_get"
	commandAddBlackList    = "black_list_add"
	commandDeleteBlackList = "black_list_delete"
	commandGetBlackList    = "black_list_get"
	commandCheckBlackList  = "black_list_check"
	commandAddGroup        = "group_add"
	commandDeleteGroup     = "group_delete"
	commandGetGroup        = "group_get"

	batchCheckFriendsLimit    = 100  // 批量校验好友限制
	batchGetFriendsLimit      = 100  // 批量获取好友限制
	batchAddBlacklistLimit    = 1000 // 批量添加黑名单限制
	batchDeleteBlacklistLimit = 1000 // 批量删除黑名单限制
	batchCheckBlacklistLimit  = 1000 // 批量校验黑名单限制
	batchAddGroupsLimit       = 100  // 批量添加分组限制
	batchJoinGroupsLimit      = 1000 // 批量加入群组账号限制
	batchDeleteGroupsLimit    = 100  // 批量删除分组限制
	batchGetGroupsLimit       = 100  // 批量获取分组限制
)

type API interface {
	// AddFriend 添加单个好友
	// 本方法拓展于“添加多个好友（AddFriends）”方法。
	// 添加好友，仅支持添加单个好友
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1643
	AddFriend(userId string, isBothAdd, isForceAdd bool, friend *Friend) (err error)

	// AddFriends 添加多个好友
	// 添加好友，支持批量添加好友
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1643
	AddFriends(userId string, isBothAdd, isForceAdd bool, friends ...*Friend) (results []*Result, err error)

	// ImportFriend 导入单个好友
	// 本方法拓展于“添加多个好友（ImportFriends）”方法。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/8301
	ImportFriend(userId string, friend *Friend) (err error)

	// ImportFriends 导入多个好友
	// 支持批量导入单向好友。
	// 往同一个用户导入好友时建议采用批量导入的方式，避免并发写导致的写冲突。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/8301
	ImportFriends(userId string, friends ...*Friend) (results []*Result, err error)

	// UpdateFriend 更新单个好友
	// 本方法拓展于“更新多个好友（UpdateFriends）”方法。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/12525
	UpdateFriend(userId string, friend *Friend) (err error)

	// UpdateFriends 更新多个好友
	// 支持批量更新同一用户的多个好友的关系链数据。
	// 更新一个用户多个好友时，建议采用批量方式，避免并发写导致的写冲突。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/12525
	UpdateFriends(userId string, friends ...*Friend) (results []*Result, err error)

	// DeleteFriend 删除单个好友
	// 本方法拓展于“删除多个好友（DeleteFriends）”方法。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1644
	DeleteFriend(userId string, isBothDelete bool, deletedUserId string) (err error)

	// DeleteFriends 删除多个好友
	// 删除好友，支持单向删除好友和双向删除好友。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1644
	DeleteFriends(userId string, isBothDelete bool, deletedUserIds ...string) (results []*Result, err error)

	// DeleteAllFriends 删除所有好友
	// 清除指定用户的标配好友数据和自定义好友数据。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1645
	DeleteAllFriends(userId string, deleteType ...DeleteType) (err error)

	// CheckFriend 校验单个好友
	// 本方法拓展于“校验多个好友（CheckFriends）”方法。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1646
	CheckFriend(userId string, checkType CheckType, checkedUserId string) (relation string, err error)

	// CheckFriends 校验多个好友
	// 支持批量校验好友关系。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1646
	CheckFriends(userId string, checkType CheckType, checkedUserIds ...string) (results []*CheckResult, err error)

	// GetFriend 拉取单个指定好友
	// 本方法拓展于“拉取多个指定好友（GetFriends）”方法。
	// 支持拉取指定好友的好友数据和资料数据。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/8609
	GetFriend(userId string, tagList []string, friendUserId string) (friend *Friend, err error)

	// GetFriends 拉取多个指定好友
	// 支持拉取指定好友的好友数据和资料数据。
	// 建议每次拉取的好友数不超过100，避免因数据量太大导致回包失败。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/8609
	GetFriends(userId string, tagList []string, friendUserIds ...string) (friends []*Friend, err error)

	// FetchFriends 拉取好友
	// 分页拉取全量好友数据。
	// 不支持资料数据的拉取。
	// 不需要指定请求拉取的字段，默认返回全量的标配好友数据和自定义好友数据。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1647
	FetchFriends(userId string, startIndex int, sequence ...int) (ret *FetchFriendsRet, err error)

	// PullFriends 续拉取好友
	// 本API是借助"拉取好友"API进行扩展实现
	// 分页拉取全量好友数据。
	// 不支持资料数据的拉取。
	// 不需要指定请求拉取的字段，默认返回全量的标配好友数据和自定义好友数据。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1647
	PullFriends(userId string, fn func(ret *FetchFriendsRet)) (err error)

	// AddBlacklist 添加黑名单
	// 添加黑名单，支持批量添加黑名单。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/3718
	AddBlacklist(userId string, blackedUserIds ...string) (results []*Result, err error)

	// DeleteBlacklist 删除黑名单
	// 删除指定黑名单。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/3719
	DeleteBlacklist(userId string, deletedUserIds ...string) (results []*Result, err error)

	// FetchBlacklist 拉取黑名单
	// 支持分页拉取所有黑名单。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/3722
	FetchBlacklist(userId string, maxLimited int, startIndexAndSequence ...int) (ret *FetchBlacklistRet, err error)

	// PullBlacklist 拉取黑名单
	// 本API是借助"拉取黑名单"API进行扩展实现
	// 支持分页拉取所有黑名单。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/3722
	PullBlacklist(userId string, maxLimited int, fn func(ret *FetchBlacklistRet)) (err error)

	// CheckBlacklist 校验黑名单
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/3725
	CheckBlacklist(userId string, checkType BlacklistCheckType, checkedUserIds ...string) (results []*CheckResult, err error)

	// AddGroups 添加分组
	// 添加分组，支持批量添加分组，并将指定好友加入到新增分组中。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/10107
	AddGroups(userId string, groupNames []string, joinedUserIds ...[]string) (currentSequence int, results []*Result, err error)

	// DeleteGroups 删除分组
	// 删除指定分组。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/10108
	DeleteGroups(userId string, groupNames ...string) (currentSequence int, err error)

	// GetGroups 拉取分组
	// 拉取分组，支持指定分组以及拉取分组下的好友列表。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/54763
	GetGroups(userId string, lastSequence int, isGetFriends bool, groupNames ...string) (currentSequence int, results []*GroupResult, err error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// AddFriend 添加单个好友
// 本方法拓展于“添加多个好友（AddFriends）”方法。
// 添加好友，仅支持添加单个好友
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1643
func (a *api) AddFriend(userId string, isBothAdd, isForceAdd bool, friend *Friend) (err error) {
	var results []*Result

	if results, err = a.AddFriends(userId, isBothAdd, isForceAdd, friend); err != nil {
		return
	}

	for _, result := range results {
		if result.UserId == friend.GetUserId() && result.ResultCode != enum.SuccessCode {
			return core.NewError(result.ResultCode, result.ResultInfo)
		}
	}

	return
}

// AddFriends 添加多个好友
// 添加好友，支持批量添加好友
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1643
func (a *api) AddFriends(userId string, isBothAdd, isForceAdd bool, friends ...*Friend) (results []*Result, err error) {
	if len(friends) == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the friends is not set")
		return
	}

	req := &addFriendsReq{UserId: userId, Friends: make([]*addFriendItem, 0, len(friends))}

	for _, friend := range friends {
		if err = friend.checkError(); err != nil {
			return
		}

		item := new(addFriendItem)
		item.UserId = friend.GetUserId()
		item.Remark, _ = friend.GetRemark()
		item.AddWording, _ = friend.GetAddWording()
		item.AddSource, _ = friend.GetSrcAddSource()
		if groups, exist := friend.GetGroup(); exist {
			item.GroupName = groups[0]
		}

		req.Friends = append(req.Friends, item)
	}

	if isBothAdd {
		req.AddType = AddTypeBoth
	} else {
		req.AddType = AddTypeSingle
	}

	if isForceAdd {
		req.ForceAddFlags = ForceAddYes
	} else {
		req.ForceAddFlags = ForceAddNo
	}

	resp := &addFriendsResp{}

	if err = a.client.Post(service, commandAddFriend, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// ImportFriend 导入单个好友
// 本方法拓展于“添加多个好友（ImportFriends）”方法。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/8301
func (a *api) ImportFriend(userId string, friend *Friend) (err error) {
	var results []*Result

	if results, err = a.ImportFriends(userId, friend); err != nil {
		return
	}

	for _, result := range results {
		if result.UserId == friend.GetUserId() && result.ResultCode != enum.SuccessCode {
			return core.NewError(result.ResultCode, result.ResultInfo)
		}
	}

	return
}

// ImportFriends 导入多个好友
// 支持批量导入单向好友。
// 往同一个用户导入好友时建议采用批量导入的方式，避免并发写导致的写冲突。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/8301
func (a *api) ImportFriends(userId string, friends ...*Friend) (results []*Result, err error) {
	if len(friends) == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the friends is not set")
		return
	}

	req := &importFriendsReq{UserId: userId, Friends: make([]*importFriendItem, 0, len(friends))}

	for _, friend := range friends {
		if err = friend.checkError(); err != nil {
			return
		}

		item := new(importFriendItem)
		item.UserId = friend.GetUserId()
		item.Remark, _ = friend.GetRemark()
		item.AddWording, _ = friend.GetAddWording()
		item.AddTime, _ = friend.GetAddTime()
		item.RemarkTime, _ = friend.GetRemarkTime()
		item.AddSource, _ = friend.GetSrcAddSource()
		item.GroupName, _ = friend.GetGroup()

		if customAttrs := friend.GetSNSCustomAttrs(); len(customAttrs) > 0 {
			item.CustomData = make([]*types.TagPair, 0, len(customAttrs))
			for k, v := range customAttrs {
				item.CustomData = append(item.CustomData, &types.TagPair{
					Tag:   k,
					Value: v,
				})
			}
		}

		req.Friends = append(req.Friends, item)
	}

	resp := &importFriendsResp{}

	if err = a.client.Post(service, commandImportFriend, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// UpdateFriend 更新单个好友
// 本方法拓展于“更新多个好友（UpdateFriends）”方法。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/12525
func (a *api) UpdateFriend(userId string, friend *Friend) (err error) {
	var results []*Result

	if results, err = a.UpdateFriends(userId, friend); err != nil {
		return
	}

	for _, result := range results {
		if result.UserId == friend.GetUserId() && result.ResultCode != enum.SuccessCode {
			return core.NewError(result.ResultCode, result.ResultInfo)
		}
	}

	return
}

// UpdateFriends 更新多个好友
// 支持批量更新同一用户的多个好友的关系链数据。
// 更新一个用户多个好友时，建议采用批量方式，避免并发写导致的写冲突。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/12525
func (a *api) UpdateFriends(userId string, friends ...*Friend) (results []*Result, err error) {
	if len(friends) == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the friends is not set")
		return
	}

	req := &updateFriendsReq{UserId: userId, Friends: make([]*updateFriendItem, 0, len(friends))}

	for _, friend := range friends {
		item := new(updateFriendItem)
		item.UserId = friend.GetUserId()

		for k, v := range friend.GetSNSAttrs() {
			switch k {
			case FriendAttrAddSource, FriendAttrAddTime, FriendAttrRemarkTime, FriendAttrAddWording:
			default:
				item.Attrs = append(item.Attrs, &types.TagPair{
					Tag:   k,
					Value: v,
				})
			}
		}

		for k, v := range friend.GetSNSCustomAttrs() {
			item.Attrs = append(item.Attrs, &types.TagPair{
				Tag:   k,
				Value: v,
			})
		}

		req.Friends = append(req.Friends, item)
	}

	resp := &updateFriendsResp{}

	if err = a.client.Post(service, commandUpdateFriend, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// DeleteFriend 删除单个好友
// 本方法拓展于“删除多个好友（DeleteFriends）”方法。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1644
func (a *api) DeleteFriend(userId string, isBothDelete bool, deletedUserId string) (err error) {
	var results []*Result

	if results, err = a.DeleteFriends(userId, isBothDelete, deletedUserId); err != nil {
		return
	}

	if results != nil && len(results) > 0 {
		for _, result := range results {
			if result.UserId == deletedUserId && result.ResultCode != enum.SuccessCode {
				return core.NewError(result.ResultCode, result.ResultInfo)
			}
		}
	}

	return
}

// DeleteFriends 删除多个好友
// 删除好友，支持单向删除好友和双向删除好友。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1644
func (a *api) DeleteFriends(userId string, isBothDelete bool, deletedUserIds ...string) (results []*Result, err error) {
	req := &deleteFriendsReq{UserId: userId, DeletedUserIds: deletedUserIds}
	resp := &deleteFriendsResp{}

	if isBothDelete {
		req.DeleteType = DeleteTypeBoth
	} else {
		req.DeleteType = DeleteTypeSingle
	}

	if err = a.client.Post(service, commandDeleteFriend, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// DeleteAllFriends 删除所有好友
// 清除指定用户的标配好友数据和自定义好友数据。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1645
func (a *api) DeleteAllFriends(userId string, deleteType ...DeleteType) (err error) {
	req := &deleteAllFriendsReq{UserId: userId}

	if len(deleteType) > 0 {
		req.DeleteType = deleteType[0]
	} else {
		req.DeleteType = DeleteTypeSingle
	}

	if err = a.client.Post(service, commandDeleteAllFriend, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// CheckFriend 校验单个好友
// 本方法拓展于“校验多个好友（CheckFriends）”方法。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1646
func (a *api) CheckFriend(userId string, checkType CheckType, checkedUserId string) (relation string, err error) {
	var results []*CheckResult

	if results, err = a.CheckFriends(userId, checkType, checkedUserId); err != nil {
		return
	}

	if results != nil && len(results) > 0 {
		for _, result := range results {
			if result.UserId == checkedUserId {
				if result.ResultCode != enum.SuccessCode {
					err = core.NewError(result.ResultCode, result.ResultInfo)
					return
				}

				relation = result.Relation
				return
			}
		}
	}

	return
}

// CheckFriends 校验多个好友
// 支持批量校验好友关系。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1646
func (a *api) CheckFriends(userId string, checkType CheckType, checkedUserIds ...string) (results []*CheckResult, err error) {
	if c := len(checkedUserIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the accounts is not set")
		return
	} else if c > batchCheckFriendsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of checked accounts cannot exceed %d", batchCheckFriendsLimit))
		return
	}

	req := &checkFriendsReq{UserId: userId, CheckedUserIds: checkedUserIds, CheckType: checkType}
	resp := &checkFriendsResp{}

	if err = a.client.Post(service, commandCheckFriend, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// GetFriend 拉取单个指定好友
// 本方法拓展于“拉取多个指定好友（GetFriends）”方法。
// 支持拉取指定好友的好友数据和资料数据。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/8609
func (a *api) GetFriend(userId string, tagList []string, friendUserId string) (friend *Friend, err error) {
	var friends []*Friend

	if friends, err = a.GetFriends(userId, tagList, friendUserId); err != nil {
		return
	}

	if len(friends) > 0 {
		friend = friends[0]
	}

	return
}

// GetFriends 拉取多个指定好友
// 支持拉取指定好友的好友数据和资料数据。
// 建议每次拉取的好友数不超过100，避免因数据量太大导致回包失败。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/8609
func (a *api) GetFriends(userId string, tagList []string, friendUserIds ...string) (friends []*Friend, err error) {
	if c := len(friendUserIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the account of friends is not set")
		return
	} else if c > batchGetFriendsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of friend's account cannot exceed %d", batchGetFriendsLimit))
		return
	}

	req := &getFriendsReq{UserId: userId, FriendUserIds: friendUserIds}
	resp := &getFriendsResp{}

	for _, tag := range tagList {
		switch tag {
		case FriendAttrRemarkTime:
		default:
			req.TagList = append(req.TagList, tag)
		}
	}

	if err = a.client.Post(service, commandGetFriend, req, resp); err != nil {
		return
	}

	friends = make([]*Friend, 0, len(resp.Friends))

	for _, item := range resp.Friends {
		friend := NewFriend(item.UserId)
		friend.SetError(item.ResultCode, item.ResultInfo)
		for _, v := range item.Profiles {
			friend.SetAttr(v.Tag, v.Value)
		}
		friends = append(friends, friend)
	}

	return
}

// FetchFriends 拉取好友
// 分页拉取全量好友数据。
// 不支持资料数据的拉取。
// 不需要指定请求拉取的字段，默认返回全量的标配好友数据和自定义好友数据。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1647
func (a *api) FetchFriends(userId string, startIndex int, sequence ...int) (ret *FetchFriendsRet, err error) {
	req := &fetchFriendsReq{UserId: userId, StartIndex: startIndex}
	resp := &fetchFriendsResp{}

	if len(sequence) > 0 {
		req.StandardSequence = sequence[0]
	}

	if len(sequence) > 1 {
		req.CustomSequence = sequence[1]
	}

	if err = a.client.Post(service, commandFetchFriend, req, resp); err != nil {
		return
	}

	ret = &FetchFriendsRet{
		StandardSequence: resp.StandardSequence,
		CustomSequence:   resp.CustomSequence,
		StartIndex:       resp.NextStartIndex,
		Total:            resp.FriendNum,
		HasMore:          resp.CompleteFlag == 0,
		List:             make([]*Friend, 0, len(resp.Friends)),
	}

	for _, item := range resp.Friends {
		friend := NewFriend(item.UserId)
		for _, v := range item.Values {
			friend.SetAttr(v.Tag, v.Value)
		}
		ret.List = append(ret.List, friend)
	}

	return
}

// PullFriends 续拉取好友
// 本API是借助"拉取好友"API进行扩展实现
// 分页拉取全量好友数据。
// 不支持资料数据的拉取。
// 不需要指定请求拉取的字段，默认返回全量的标配好友数据和自定义好友数据。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1647
func (a *api) PullFriends(userId string, fn func(ret *FetchFriendsRet)) (err error) {
	var (
		ret              *FetchFriendsRet
		startIndex       int
		standardSequence int
		customSequence   int
	)

	for ret == nil || ret.HasMore {
		ret, err = a.FetchFriends(userId, startIndex, standardSequence, customSequence)
		if err != nil {
			return
		}

		fn(ret)

		if ret.HasMore {
			startIndex = ret.StartIndex
			standardSequence = ret.StandardSequence
			customSequence = ret.CustomSequence
		}
	}

	return
}

// AddBlacklist 添加黑名单
// 添加黑名单，支持批量添加黑名单。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3718
func (a *api) AddBlacklist(userId string, blackedUserIds ...string) (results []*Result, err error) {
	if c := len(blackedUserIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the blacked accounts is not set")
		return
	} else if c > batchAddBlacklistLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of blacked accounts cannot exceed %d", batchAddBlacklistLimit))
		return
	}

	req := &addBlacklistReq{UserId: userId, BlackedUserIds: blackedUserIds}
	resp := &addBlacklistResp{}

	if err = a.client.Post(service, commandAddBlackList, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// DeleteBlacklist 删除黑名单
// 删除指定黑名单。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3719
func (a *api) DeleteBlacklist(userId string, deletedUserIds ...string) (results []*Result, err error) {
	if c := len(deletedUserIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the deleted accounts is not set")
		return
	} else if c > batchDeleteBlacklistLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of deleted accounts cannot exceed %d", batchDeleteBlacklistLimit))
		return
	}

	req := &deleteBlacklistReq{UserId: userId, DeletedUserIds: deletedUserIds}
	resp := &deleteBlacklistResp{}

	if err = a.client.Post(service, commandDeleteBlackList, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// FetchBlacklist 拉取黑名单
// 支持分页拉取所有黑名单。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3722
func (a *api) FetchBlacklist(userId string, maxLimited int, startIndexAndSequence ...int) (ret *FetchBlacklistRet, err error) {
	req := &fetchBlacklistReq{UserId: userId, MaxLimited: maxLimited}

	if len(startIndexAndSequence) > 0 {
		req.StartIndex = startIndexAndSequence[0]
	}

	if len(startIndexAndSequence) > 1 {
		req.LastSequence = startIndexAndSequence[1]
	}

	resp := &fetchBlacklistResp{}

	if err = a.client.Post(service, commandGetBlackList, req, resp); err != nil {
		return
	}

	ret = &FetchBlacklistRet{
		StartIndex:       resp.StartIndex,
		StandardSequence: resp.CurrentSequence,
		List:             resp.Blacklists,
		HasMore:          resp.StartIndex != 0,
	}

	return
}

// PullBlacklist 拉取黑名单
// 本API是借助"拉取黑名单"API进行扩展实现
// 支持分页拉取所有黑名单。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3722
func (a *api) PullBlacklist(userId string, maxLimited int, fn func(ret *FetchBlacklistRet)) (err error) {
	var (
		ret              *FetchBlacklistRet
		startIndex       = 0
		standardSequence = 0
	)

	for ret == nil || ret.HasMore {
		ret, err = a.FetchBlacklist(userId, maxLimited, startIndex, standardSequence)
		if err != nil {
			return
		}

		fn(ret)

		if ret.HasMore {
			startIndex = ret.StartIndex
			standardSequence = ret.StandardSequence
		}
	}

	return
}

// CheckBlacklist 校验黑名单
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3725
func (a *api) CheckBlacklist(userId string, checkType BlacklistCheckType, checkedUserIds ...string) (results []*CheckResult, err error) {
	if c := len(checkedUserIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the checked accounts is not set")
		return
	} else if c > batchCheckBlacklistLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of checked accounts cannot exceed %d", batchCheckBlacklistLimit))
		return
	}

	req := &checkBlacklistReq{UserId: userId, CheckedUserIds: checkedUserIds, CheckType: checkType}
	resp := &checkBlacklistResp{}

	if err = a.client.Post(service, commandCheckBlackList, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// AddGroups 添加分组
// 添加分组，支持批量添加分组，并将指定好友加入到新增分组中。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/10107
func (a *api) AddGroups(userId string, groupNames []string, joinedUserIds ...[]string) (currentSequence int, results []*Result, err error) {
	if c := len(groupNames); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the added groups is not set")
		return
	} else if c > batchAddGroupsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of added groups cannot exceed %d", batchAddGroupsLimit))
		return
	}

	req := &addGroupsReq{UserId: userId, GroupNames: groupNames}

	if len(joinedUserIds) > 0 {
		if c := len(joinedUserIds[0]); c == 0 {
			err = core.NewError(enum.InvalidParamsCode, "the added groups is not set")
			return
		} else if c > batchJoinGroupsLimit {
			err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of accounts joining the group cannot exceed %d", batchJoinGroupsLimit))
			return
		}

		req.JoinedUserIds = joinedUserIds[0]
	}

	resp := &addGroupsResp{}

	if err = a.client.Post(service, commandAddGroup, req, resp); err != nil {
		return
	}

	currentSequence = resp.CurrentSequence
	results = resp.Results

	return
}

// DeleteGroups 删除分组
// 删除指定分组。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/10108
func (a *api) DeleteGroups(userId string, groupNames ...string) (currentSequence int, err error) {
	if c := len(groupNames); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the deleted groups is not set")
		return
	} else if c > batchDeleteGroupsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of deleted groups cannot exceed %d", batchDeleteGroupsLimit))
		return
	}

	req := &deleteGroupsReq{UserId: userId, GroupNames: groupNames}
	resp := &deleteGroupsResp{}

	if err = a.client.Post(service, commandDeleteGroup, req, resp); err != nil {
		return
	}

	currentSequence = resp.CurrentSequence

	return
}

// GetGroups 拉取分组
// 拉取分组，支持指定分组以及拉取分组下的好友列表。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/54763
func (a *api) GetGroups(userId string, lastSequence int, isGetFriends bool, groupNames ...string) (currentSequence int, results []*GroupResult, err error) {
	if c := len(groupNames); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the gotten groups is not set")
		return
	} else if c > batchGetGroupsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of gotten groups cannot exceed %d", batchGetGroupsLimit))
		return
	}

	req := &getGroupsReq{UserId: userId, LastSequence: lastSequence, GroupNames: groupNames}
	resp := &getGroupsResp{}

	if isGetFriends {
		req.NeedFriend = NeedFriendYes
	} else {
		req.NeedFriend = NeedFriendNo
	}

	if err = a.client.Post(service, commandGetGroup, req, resp); err != nil {
		return
	}

	currentSequence = resp.CurrentSequence
	results = resp.Results

	return
}
