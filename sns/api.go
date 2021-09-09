/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:45
 * @Desc: 关系链管理
 */

package sns

import (
    "github.com/dobyte/tencent-im/internal/core"
    "github.com/dobyte/tencent-im/internal/enum"
    "github.com/dobyte/tencent-im/internal/types"
)

const (
    serviceSNS             = "sns"
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
)

type API interface {
    // AddFriend 添加单个好友
    // 本方法拓展于“添加多个好友（AddFriends）”方法。
    // 添加好友，仅支持添加单个好友
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1643
    AddFriend(userId string, friend *Friend, isSingleAdd bool, isForceAdd bool) (err error)
    
    // AddFriends 添加多个好友
    // 添加好友，支持批量添加好友
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1643
    AddFriends(userId string, friends []*Friend, isSingleAdd bool, isForceAdd bool) (results []Result, err error)
    
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
    ImportFriends(userId string, friends []*Friend) (results []Result, err error)
    
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
    UpdateFriends(userId string, friends []*Friend) (results []Result, err error)
    
    // DeleteFriend 删除单个好友
    // 本方法拓展于“删除多个好友（DeleteFriends）”方法。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1644
    DeleteFriend(userId string, deleteUserId string, deleteType ...DeleteType) (err error)
    
    // DeleteFriends 删除多个好友
    // 删除好友，支持单向删除好友和双向删除好友。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1644
    DeleteFriends(userId string, deleteUserIds []string, deleteType ...DeleteType) (results []Result, err error)
    
    // DeleteAllFriends 删除所有好友
    // 清除指定用户的标配好友数据和自定义好友数据。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1645
    DeleteAllFriends(userId string, deleteType ...DeleteType) (err error)
    
    // CheckFriend 校验单个好友
    // 本方法拓展于“校验多个好友（CheckFriends）”方法。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1646
    CheckFriend(userId string, checkUserId string, checkType CheckType) (relation string, err error)
    
    // CheckFriends 校验多个好友
    // 支持批量校验好友关系。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1646
    CheckFriends(userId string, checkUserIds []string, checkType CheckType) (results []CheckResult, err error)
    
    // GetFriend 拉取单个指定好友
    // 本方法拓展于“拉取多个指定好友（GetFriends）”方法。
    // 支持拉取指定好友的好友数据和资料数据。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/8609
    GetFriend(userId string, friendUserId string, tagList []string) (friend *Friend, err error)
    
    // GetFriends 拉取多个指定好友
    // 支持拉取指定好友的好友数据和资料数据。
    // 建议每次拉取的好友数不超过100，避免因数据量太大导致回包失败。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/8609
    GetFriends(userId string, friendUserIds []string, tagList []string) (friends []*Friend, err error)
    
    // FetchFriends 拉取好友
    // 分页拉取全量好友数据。
    // 不支持资料数据的拉取。
    // 不需要指定请求拉取的字段，默认返回全量的标配好友数据和自定义好友数据。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1647
    FetchFriends(userId string, startIndex int, sequence ...int) (ret *FetchFriendsRet, err error)
    
    // AddBlacklist 添加黑名单
    // 添加黑名单，支持批量添加黑名单。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/3718
    AddBlacklist(userId string, blackUserIds []string) (results []Result, err error)
    
    // DeleteBlacklist 删除黑名单
    // 删除指定黑名单。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/3719
    DeleteBlacklist(userId string, deleteUserIds []string) (results []Result, err error)
    
    // FetchBlacklist 拉取黑名单
    // 支持分页拉取所有黑名单。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/3722
    FetchBlacklist(userId string, startIndex, maxLimited, lastSequence int) (ret *FetchBlacklistRet, err error)
    
    // CheckBlacklist 校验黑名单
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/3725
    CheckBlacklist(userId string, checkUserIds []string, checkType BlacklistCheckType) (results []CheckResult, err error)
    
    // AddGroups 添加分组
    // 添加分组，支持批量添加分组，并将指定好友加入到新增分组中。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/10107
    AddGroups(userId string, groupNames []string, groupUserIds ...[]string) (currentSequence int, results []Result, err error)
    
    // DeleteGroups 删除分组
    // 删除指定分组。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/10108
    DeleteGroups(userId string, groupNames []string) (currentSequence int, err error)
    
    // GetGroups 拉取分组
    // 拉取分组，支持指定分组以及拉取分组下的好友列表。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/54763
    GetGroups(userId string, lastSequence int, isGetFriends bool, groupName []string) (currentSequence int, results []GroupResult, err error)
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
func (a *api) AddFriend(userId string, friend *Friend, isSingleAdd bool, isForceAdd bool) (err error) {
    var results []Result
    
    if results, err = a.AddFriends(userId, []*Friend{friend}, isSingleAdd, isForceAdd); err != nil {
        return
    }
    
    if results != nil && len(results) > 0 {
        for _, result := range results {
            if result.UserId == friend.GetUserId() && result.ResultCode != enum.SuccessCode {
                return core.NewError(result.ResultCode, result.ResultInfo)
            }
        }
    }
    
    return
}

// AddFriends 添加多个好友
// 添加好友，支持批量添加好友
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1643
func (a *api) AddFriends(userId string, friends []*Friend, isSingleAdd bool, isForceAdd bool) (results []Result, err error) {
    req := addFriendsReq{UserId: userId}
    resp := &addFriendsResp{}
    
    for _, friend := range friends {
        if err = friend.checkError(); err != nil {
            return
        }
        
        info := friendInfo{}
        info.UserId = friend.GetUserId()
        info.Remark, _ = friend.GetRemark()
        info.AddWording, _ = friend.GetAddWording()
        info.AddSource, _ = friend.GetSrcAddSource()
        if groups, exist := friend.GetGroup(); exist {
            info.GroupName = groups[0]
        }
        
        req.Friends = append(req.Friends, info)
    }
    
    if isSingleAdd {
        req.AddType = string(AddTypeSingle)
    } else {
        req.AddType = string(AddTypeBoth)
    }
    
    if isForceAdd {
        req.ForceAddFlags = 1
    } else {
        req.ForceAddFlags = 0
    }
    
    if err = a.client.Post(serviceSNS, commandAddFriend, req, resp); err != nil {
        return
    } else {
        results = resp.Results
    }
    
    return
}

// ImportFriend 导入单个好友
// 本方法拓展于“添加多个好友（ImportFriends）”方法。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/8301
func (a *api) ImportFriend(userId string, friend *Friend) (err error) {
    var results []Result
    
    if results, err = a.ImportFriends(userId, []*Friend{friend}); err != nil {
        return
    }
    
    if results != nil && len(results) > 0 {
        for _, result := range results {
            if result.UserId == friend.GetUserId() && result.ResultCode != enum.SuccessCode {
                return core.NewError(result.ResultCode, result.ResultInfo)
            }
        }
    }
    
    return
}

// ImportFriends 导入多个好友
// 支持批量导入单向好友。
// 往同一个用户导入好友时建议采用批量导入的方式，避免并发写导致的写冲突。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/8301
func (a *api) ImportFriends(userId string, friends []*Friend) (results []Result, err error) {
    req := importFriendsReq{UserId: userId}
    resp := &importFriendsResp{}
    
    for _, friend := range friends {
        if err = friend.checkError(); err != nil {
            return
        }
        
        info := importFriend{}
        info.UserId = friend.GetUserId()
        info.Remark, _ = friend.GetRemark()
        info.AddWording, _ = friend.GetAddWording()
        info.AddTime, _ = friend.GetAddTime()
        info.RemarkTime, _ = friend.GetRemarkTime()
        info.AddSource, _ = friend.GetSrcAddSource()
        
        if groups, exist := friend.GetGroup(); exist {
            info.GroupName = groups
        }
        
        if customAttrs := friend.GetSNSCustomAttrs(); len(customAttrs) > 0 {
            info.CustomData = make([]types.TagPair, 0)
            for k, v := range customAttrs {
                info.CustomData = append(info.CustomData, types.TagPair{
                    Tag:   k,
                    Value: v,
                })
            }
        }
        
        req.Friends = append(req.Friends, info)
    }
    
    if err = a.client.Post(serviceSNS, commandImportFriend, req, resp); err != nil {
        return
    } else {
        results = resp.Results
    }
    
    return
}

// UpdateFriend 更新单个好友
// 本方法拓展于“更新多个好友（UpdateFriends）”方法。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/12525
func (a *api) UpdateFriend(userId string, friend *Friend) (err error) {
    var results []Result
    
    if results, err = a.UpdateFriends(userId, []*Friend{friend}); err != nil {
        return
    }
    
    if results != nil && len(results) > 0 {
        for _, result := range results {
            if result.UserId == friend.GetUserId() && result.ResultCode != enum.SuccessCode {
                return core.NewError(result.ResultCode, result.ResultInfo)
            }
        }
    }
    
    return
}

// UpdateFriends 更新多个好友
// 支持批量更新同一用户的多个好友的关系链数据。
// 更新一个用户多个好友时，建议采用批量方式，避免并发写导致的写冲突。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/12525
func (a *api) UpdateFriends(userId string, friends []*Friend) (results []Result, err error) {
    req := updateFriendsReq{UserId: userId}
    resp := &updateFriendsResp{}
    
    for _, friend := range friends {
        info := updateFriend{}
        info.UserId = friend.GetUserId()
        for k, v := range friend.GetSNSAttrs() {
            switch k {
            case FriendAttrAddSource, FriendAttrAddTime, FriendAttrRemarkTime, FriendAttrAddWording:
            default:
                info.Attrs = append(info.Attrs, types.TagPair{
                    Tag:   k,
                    Value: v,
                })
            }
        }
        
        for k, v := range friend.GetSNSCustomAttrs() {
            info.Attrs = append(info.Attrs, types.TagPair{
                Tag:   k,
                Value: v,
            })
        }
        
        req.Friends = append(req.Friends, info)
    }
    
    if err = a.client.Post(serviceSNS, commandUpdateFriend, req, resp); err != nil {
        return
    } else {
        results = resp.Results
    }
    
    return
}

// DeleteFriend 删除单个好友
// 本方法拓展于“删除多个好友（DeleteFriends）”方法。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1644
func (a *api) DeleteFriend(userId string, deleteUserId string, deleteType ...DeleteType) (err error) {
    var results []Result
    
    if results, err = a.DeleteFriends(userId, []string{deleteUserId}, deleteType...); err != nil {
        return
    }
    
    if results != nil && len(results) > 0 {
        for _, result := range results {
            if result.UserId == deleteUserId && result.ResultCode != enum.SuccessCode {
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
func (a *api) DeleteFriends(userId string, deleteUserIds []string, deleteType ...DeleteType) (results []Result, err error) {
    req := deleteFriendsReq{UserId: userId, DeleteUserIds: deleteUserIds}
    resp := &deleteFriendsResp{}
    
    if len(deleteType) > 0 {
        req.DeleteType = string(deleteType[0])
    } else {
        req.DeleteType = string(DeleteTypeSingle)
    }
    
    if err = a.client.Post(serviceSNS, commandDeleteFriend, req, resp); err != nil {
        return
    } else {
        results = resp.Results
    }
    
    return
}

// DeleteAllFriends 删除所有好友
// 清除指定用户的标配好友数据和自定义好友数据。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1645
func (a *api) DeleteAllFriends(userId string, deleteType ...DeleteType) (err error) {
    req := deleteAllFriendsReq{UserId: userId}
    
    if len(deleteType) > 0 {
        req.DeleteType = string(deleteType[0])
    } else {
        req.DeleteType = string(DeleteTypeSingle)
    }
    
    if err = a.client.Post(serviceSNS, commandDeleteAllFriend, req, &types.ActionBaseResp{}); err != nil {
        return
    }
    
    return
}

// CheckFriend 校验单个好友
// 本方法拓展于“校验多个好友（CheckFriends）”方法。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1646
func (a *api) CheckFriend(userId string, checkUserId string, checkType CheckType) (relation string, err error) {
    var results []CheckResult
    
    if results, err = a.CheckFriends(userId, []string{checkUserId}, checkType); err != nil {
        return
    }
    
    if results != nil && len(results) > 0 {
        for _, result := range results {
            if result.UserId == checkUserId {
                if result.ResultCode != enum.SuccessCode {
                    err = core.NewError(result.ResultCode, result.ResultInfo)
                    return
                } else {
                    return result.Relation, nil
                }
            }
        }
    }
    
    return
}

// CheckFriends 校验多个好友
// 支持批量校验好友关系。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1646
func (a *api) CheckFriends(userId string, checkUserIds []string, checkType CheckType) (results []CheckResult, err error) {
    req := checkFriendsReq{UserId: userId, CheckUserIds: checkUserIds, CheckType: string(checkType)}
    resp := &checkFriendsResp{}
    
    if err = a.client.Post(serviceSNS, commandCheckFriend, req, resp); err != nil {
        return
    } else {
        results = resp.Results
    }
    
    return
}

// GetFriend 拉取单个指定好友
// 本方法拓展于“拉取多个指定好友（GetFriends）”方法。
// 支持拉取指定好友的好友数据和资料数据。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/8609
func (a *api) GetFriend(userId string, friendUserId string, tagList []string) (friend *Friend, err error) {
    var friends []*Friend
    
    if friends, err = a.GetFriends(userId, []string{friendUserId}, tagList); err != nil {
        return
    }
    
    if friends != nil && len(friends) > 0 {
        friend = friends[0]
    }
    
    return
}

// GetFriends 拉取多个指定好友
// 支持拉取指定好友的好友数据和资料数据。
// 建议每次拉取的好友数不超过100，避免因数据量太大导致回包失败。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/8609
func (a *api) GetFriends(userId string, friendUserIds []string, tagList []string) (friends []*Friend, err error) {
    req := getFriendsReq{UserId: userId, FriendUserIds: friendUserIds}
    resp := &getFriendsResp{}
    
    for _, tag := range tagList {
        switch tag {
        case FriendAttrRemarkTime:
        default:
            req.TagList = append(req.TagList, tag)
        }
    }
    
    if err = a.client.Post(serviceSNS, commandGetFriend, req, resp); err != nil {
        return
    }
    
    friends = make([]*Friend, 0, len(resp.Friends))
    
    var friend *Friend
    for _, item := range resp.Friends {
        friend = NewFriend()
        friend.SetUserId(item.UserId)
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
    req := fetchFriendsReq{UserId: userId, StartIndex: startIndex}
    resp := &fetchFriendsResp{}
    
    if len(sequence) > 0 {
        req.StandardSequence = sequence[0]
    }
    
    if len(sequence) > 1 {
        req.CustomSequence = sequence[1]
    }
    
    if err = a.client.Post(serviceSNS, commandFetchFriend, req, resp); err != nil {
        return
    } else {
        ret = &FetchFriendsRet{
            StandardSequence: resp.StandardSequence,
            CustomSequence:   resp.CustomSequence,
            FriendNum:        resp.FriendNum,
            NextStartIndex:   resp.NextStartIndex,
        }
        
        if resp.CompleteFlag != 0 {
            ret.IsOver = true
        }
        
        var friend *Friend
        for _, item := range resp.Friends {
            friend = NewFriend()
            friend.SetUserId(item.UserId)
            for _, v := range item.Values {
                friend.SetAttr(v.Tag, v.Value)
            }
            ret.Friends = append(ret.Friends, friend)
        }
    }
    
    return
}

// AddBlacklist 添加黑名单
// 添加黑名单，支持批量添加黑名单。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3718
func (a *api) AddBlacklist(userId string, blackUserIds []string) (results []Result, err error) {
    req := addBlacklistReq{UserId: userId, BlackUserIds: blackUserIds}
    resp := &addBlacklistResp{}
    
    if err = a.client.Post(serviceSNS, commandAddBlackList, req, resp); err != nil {
        return
    } else {
        results = resp.Results
    }
    
    return
}

// DeleteBlacklist 删除黑名单
// 删除指定黑名单。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3719
func (a *api) DeleteBlacklist(userId string, deleteUserIds []string) (results []Result, err error) {
    req := deleteBlacklistReq{UserId: userId, DeleteUserIds: deleteUserIds}
    resp := &deleteBlacklistResp{}
    
    if err = a.client.Post(serviceSNS, commandDeleteBlackList, req, resp); err != nil {
        return
    } else {
        results = resp.Results
    }
    
    return
}

// FetchBlacklist 拉取黑名单
// 支持分页拉取所有黑名单。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3722
func (a *api) FetchBlacklist(userId string, startIndex, maxLimited, lastSequence int) (ret *FetchBlacklistRet, err error) {
    req := fetchBlacklistReq{UserId: userId, StartIndex: startIndex, MaxLimited: maxLimited, LastSequence: lastSequence}
    resp := &fetchBlacklistResp{}
    
    if err = a.client.Post(serviceSNS, commandGetBlackList, req, resp); err != nil {
        return
    } else {
        ret = &FetchBlacklistRet{
            NextStartIndex:   resp.StartIndex,
            StandardSequence: resp.CurrentSequence,
            Blacklists:       resp.Blacklists,
        }
        
        if resp.StartIndex == 0 {
            ret.IsOver = true
        }
    }
    
    return
}

// CheckBlacklist 校验黑名单
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3725
func (a *api) CheckBlacklist(userId string, checkUserIds []string, checkType BlacklistCheckType) (results []CheckResult, err error) {
    req := checkBlacklistReq{UserId: userId, CheckUserIds: checkUserIds, CheckType: string(checkType)}
    resp := &checkBlacklistResp{}
    
    if err = a.client.Post(serviceSNS, commandCheckBlackList, req, resp); err != nil {
        return
    } else {
        results = resp.CheckResults
    }
    
    return
}

// AddGroups 添加分组
// 添加分组，支持批量添加分组，并将指定好友加入到新增分组中。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/10107
func (a *api) AddGroups(userId string, groupNames []string, groupUserIds ...[]string) (currentSequence int, results []Result, err error) {
    req := addGroupsReq{UserId: userId, GroupNames: groupNames}
    resp := &addGroupsResp{}
    
    if len(groupUserIds) > 0 {
        req.GroupUserIds = groupUserIds[0]
    }
    
    if err = a.client.Post(serviceSNS, commandAddGroup, req, resp); err != nil {
        return
    } else {
        currentSequence = resp.CurrentSequence
        results = resp.Results
    }
    
    return
}

// DeleteGroups 删除分组
// 删除指定分组。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/10108
func (a *api) DeleteGroups(userId string, groupNames []string) (currentSequence int, err error) {
    req := deleteGroupsReq{UserId: userId, GroupNames: groupNames}
    resp := &deleteGroupsResp{}
    
    if err = a.client.Post(serviceSNS, commandDeleteGroup, req, resp); err != nil {
        return
    } else {
        currentSequence = resp.CurrentSequence
    }
    
    return
}

// GetGroups 拉取分组
// 拉取分组，支持指定分组以及拉取分组下的好友列表。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/54763
func (a *api) GetGroups(userId string, lastSequence int, isGetFriends bool, groupName []string) (currentSequence int, results []GroupResult, err error) {
    req := getGroupsReq{UserId: userId, LastSequence: lastSequence, GroupName: groupName}
    resp := &getGroupsResp{}
    
    if isGetFriends {
        req.NeedFriend = "Need_Friend_Type_Yes"
    } else {
        req.NeedFriend = "Need_Friend_Type_No"
    }
    
    if err = a.client.Post(serviceSNS, commandGetGroup, req, resp); err != nil {
        return
    } else {
        currentSequence = resp.CurrentSequence
        results = resp.Results
    }
    
    return
}
