/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/27 1:40 下午
 * @Desc: TODO
 */

package im_test

import (
    "fmt"
    "strconv"
    "testing"
    "time"
    
    im "github.com/dobyte/tencent-im"
    "github.com/dobyte/tencent-im/account"
    "github.com/dobyte/tencent-im/operation"
    "github.com/dobyte/tencent-im/private"
    "github.com/dobyte/tencent-im/profile"
    "github.com/dobyte/tencent-im/push"
    "github.com/dobyte/tencent-im/sns"
    "github.com/dobyte/tencent-im/types"
    "github.com/dobyte/tencent-im/user"
)

// func NewIM() im.IM {
// 	return im.NewIM(im.Options{
// 		AppId:     1400564830,
// 		AppSecret: "0d2a321b087fdb8fd5ed5ea14fe0489139086eb1b03541774fc9feeab8f2bfd3",
// 		UserId:    "administrator",
// 	})
// }

func NewIM() im.IM {
    return im.NewIM(im.Options{
        AppId:     1400567367,
        AppSecret: "9c16200a10f74a12100b7987344687a27a09dd3d24c3df57e3d98569dff7acbe",
        UserId:    "administrator",
    })
}

// 导入单个账号
func TestIm_Account_ImportAccount(t *testing.T) {
    if err := NewIM().Account().ImportAccount(account.AccountInfo{
        UserId:    "assistant",
        Nickname:  "小助手",
        AvatarUrl: "http://www.qq.com",
    }); err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 导入多个帐号
func TestIm_Account_ImportAccounts(t *testing.T) {
    failedAccounts, err := NewIM().Account().ImportAccounts([]string{
        "test1",
        "test2",
        "test3",
        "test4",
        "test5",
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(failedAccounts)
}

// 删除多个帐号
func TestIm_Account_DeleteAccounts(t *testing.T) {
    deleteResults, err := NewIM().Account().DeleteAccounts([]string{
        "test1",
        "test2",
        "test3",
        "test4",
        "test5",
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(deleteResults)
}

// 查询多个帐号
func TestIm_Account_CheckAccounts(t *testing.T) {
    checkResults, err := NewIM().Account().CheckAccounts([]string{
        "test1",
        "test2",
        "test3",
        "test4",
        "test5",
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(checkResults)
}

// 使帐号登录状态失效
func TestIm_Account_KickAccount(t *testing.T) {
    if err := NewIM().Account().KickAccount("test1"); err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 使帐号登录状态失效
func TestIm_Account_QueryAccountsOnlineStatus(t *testing.T) {
    resp, err := NewIM().Account().GetAccountsOnlineState([]string{
        "assistant",
        "test1",
        "test2",
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(resp.Results)
    t.Log(resp.Errors)
}

// 全员推送
func TestIm_Push_Push(t *testing.T) {
    taskId, err := NewIM().Push().Push(push.PushArgument{
        MsgContents: []interface{}{
            types.MsgTextContent{
                Text: "Hello Tencent IM",
            },
        },
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(taskId)
}

// 设置应用属性名称
func TestIm_Push_SetAttrNames(t *testing.T) {
    if err := NewIM().Push().SetAttrNames(map[int]string{
        0: "age",
        1: "city",
    }); err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 获取应用属性名称
func TestIm_Push_GetAttrNames(t *testing.T) {
    ret, err := NewIM().Push().GetAttrNames()
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(ret)
}

// 获取用户属性
func TestIm_Push_GetUserAttrs(t *testing.T) {
    ret, err := NewIM().Push().GetUserAttrs("test1")
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(ret)
}

// 设置用户属性
func TestIm_Push_SetUserAttrs(t *testing.T) {
    err := NewIM().Push().SetUserAttrs(map[string]map[string]interface{}{
        "test1": {
            "age":  20,
            "city": "成都",
        },
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 删除用户属性
func TestIm_Push_DeleteUserAttrs(t *testing.T) {
    err := NewIM().Push().DeleteUserAttrs(map[string][]string{
        "test1": {"age", "city"},
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 获取用户标签
func TestIm_Push_GetUserTags(t *testing.T) {
    ret, err := NewIM().Push().GetUserTags("test1")
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(ret)
}

// 添加用户标签
func TestIm_Push_AddUserTags(t *testing.T) {
    err := NewIM().Push().AddUserTags(map[string][]string{
        "test1": {"chengdu"},
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 删除用户标签
func TestIm_Push_DeleteUserTags(t *testing.T) {
    err := NewIM().Push().DeleteUserTags(map[string][]string{
        "test1": {"chengdu"},
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 删除用户所有标签
func TestIm_Push_DeleteUserAllTags(t *testing.T) {
    err := NewIM().Push().DeleteUserAllTags("test1", "test2")
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 设置资料
func TestIm_Profile_SetProfile(t *testing.T) {
    p := profile.NewProfile("assistant")
    p.SetAvatar("http://www.qq.com")
    p.SetGender(profile.GenderTypeMale)
    // p.SetLocation(1, 23, 27465, 92)
    p.SetLocation(1, 23, 2, 92)
    p.SetLanguage(20)
    
    if err := NewIM().Profile().SetProfile(p); err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 获取资料
func TestIm_Profile_GetProfile(t *testing.T) {
    profiles, err := NewIM().Profile().GetProfiles([]string{
        "assistant",
    }, []string{
        profile.AttrNickname,
        profile.AttrGender,
        profile.AttrBirthday,
        profile.AttrLocation,
        profile.AttrLanguage,
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    for _, p := range profiles {
        t.Log(p.GetUserId())
        t.Log(p.GetNickname())
        t.Log(p.GetGender())
        t.Log(p.GetBirthday())
        t.Log(p.GetLocation())
        t.Log(p.GetLanguage())
    }
}

// 拉取运营数据
func TestIm_Operation_GetOperationData(t *testing.T) {
    data, err := NewIM().Operation().GetOperationData([]string{})
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(data)
}

// 拉取运营数据
func TestIm_Operation_GetHistoryData(t *testing.T) {
    files, err := NewIM().Operation().GetHistoryData(operation.ChatTypePrivate, time.Date(2021, time.August, 22, 14, 0, 0, 0, time.Local))
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(files)
}

// 获取服务器IP地址
func TestIm_Operation_GetIpList(t *testing.T) {
    ips, err := NewIM().Operation().GetIPList()
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(ips)
}

// 设置全局禁言
func TestIm_Mute_SetNoSpeaking(t *testing.T) {
    var privateMuteTime uint = 400
    var groupMuteTime uint = 200
    if err := NewIM().Mute().SetNoSpeaking("assistant", &privateMuteTime, &groupMuteTime); err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 查询全局禁言
func TestIm_Mute_GetNoSpeaking(t *testing.T) {
    privateMuteTime, groupMuteTime, err := NewIM().Mute().GetNoSpeaking("assistant")
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(privateMuteTime)
    t.Log(groupMuteTime)
}

// 添加好友
func TestIm_SNS_AddFriends(t *testing.T) {
    friends := make([]*sns.Friend, 0)
    
    var friend *sns.Friend
    var userIds []string
    var userId string
    for i := 0; i < 10; i++ {
        userId = "test" + strconv.Itoa(i)
        friend = sns.NewFriend()
        friend.SetUserId(userId)
        friend.SetAddSource("android")
        friends = append(friends, friend)
        userIds = append(userIds, userId)
    }
    
    failUserIds, err := NewIM().Account().ImportAccounts(userIds)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(failUserIds)
    
    results, err := NewIM().SNS().AddFriends("assistant", friends, true, false)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(results)
}

// 导入好友
func TestIm_SNS_ImportFriends(t *testing.T) {
    friends := make([]*sns.Friend, 0)
    
    var friend *sns.Friend
    var userIds []string
    var userId string
    var now = time.Now().Unix()
    for i := 20; i < 30; i++ {
        userId = "test" + strconv.Itoa(i)
        friend = sns.NewFriend()
        friend.SetUserId(userId)
        friend.SetAddSource("android")
        friend.SetGroup("测试组")
        friend.SetAddWording("测试一下")
        friend.SetAddTime(now)
        friend.SetRemark("测试好友")
        friend.SetRemarkTime(now)
        friends = append(friends, friend)
        userIds = append(userIds, userId)
    }
    
    failUserIds, err := NewIM().Account().ImportAccounts(userIds)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(failUserIds)
    
    results, err := NewIM().SNS().ImportFriends("assistant", friends)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(results)
}

// 更新好友
func TestIm_SNS_UpdateFriends(t *testing.T) {
    friends := make([]*sns.Friend, 0)
    
    var friend *sns.Friend
    var userIds []string
    var userId string
    var now = time.Now().Unix()
    for i := 0; i < 10; i++ {
        userId = "test" + strconv.Itoa(i)
        friend = sns.NewFriend()
        friend.SetUserId(userId)
        friend.SetAddSource("android")
        friend.SetGroup("测试组")
        friend.SetAddWording("测试一下")
        friend.SetAddTime(now)
        friend.SetRemark("测试好友")
        friend.SetRemarkTime(now)
        friends = append(friends, friend)
        userIds = append(userIds, userId)
    }
    
    failUserIds, err := NewIM().Account().ImportAccounts(userIds)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(failUserIds)
    
    results, err := NewIM().SNS().UpdateFriends("assistant", friends)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(results)
}

// 删除好友
func TestIm_SNS_DeleteFriends(t *testing.T) {
    var userIds []string
    var userId string
    for i := 0; i < 10; i++ {
        userId = "test" + strconv.Itoa(i)
        userIds = append(userIds, userId)
    }
    
    results, err := NewIM().SNS().DeleteFriends("assistant", userIds)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(results)
}

// 删除所有好友
func TestIm_SNS_DeleteAllFriends(t *testing.T) {
    err := NewIM().SNS().DeleteAllFriends("assistant")
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 校验好友
func TestIm_SNS_CheckFriends(t *testing.T) {
    var userIds []string
    var userId string
    for i := 0; i < 10; i++ {
        userId = "test" + strconv.Itoa(i)
        userIds = append(userIds, userId)
    }
    
    results, err := NewIM().SNS().CheckFriends("assistant", userIds, sns.CheckTypeSingle)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(results)
}

// 拉取指定好友
func TestIm_SNS_GetFriends(t *testing.T) {
    var userIds []string
    var userId string
    for i := 0; i < 10; i++ {
        userId = "test" + strconv.Itoa(i)
        userIds = append(userIds, userId)
    }
    
    friends, err := NewIM().SNS().GetFriends("assistant", userIds, []string{
        sns.SNSAttrAddSource,
        sns.SNSAttrRemark,
        sns.SNSAttrRemarkTime, // 此Tag无效，GetFriends内部忽略了
        sns.SNSAttrAddTime,
        sns.SNSAttrAddWording,
        sns.SNSAttrGroup,
        user.StandardAttrNickname,
        user.StandardAttrBirthday,
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    // 第一种获取方式
    for _, friend := range friends {
        if friend.IsValid() {
            t.Log(friend.GetUserId())
            t.Log(friend.GetAddSource())
            t.Log(friend.GetRemark())
            t.Log(friend.GetRemarkTime())
            t.Log(friend.GetGroup())
            t.Log(friend.GetNickname())
            t.Log(friend.GetBirthday())
            fmt.Println()
        }
    }
    
    // 第二种获取方式
    for _, friend := range friends {
        if err := friend.GetError(); err != nil {
            t.Log(fmt.Sprintf("获取账号%s失败", friend.GetUserId()))
        } else {
            t.Log(friend.GetUserId())
            t.Log(friend.GetAddSource())
            t.Log(friend.GetRemark())
            t.Log(friend.GetRemarkTime())
            t.Log(friend.GetGroup())
            t.Log(friend.GetNickname())
            t.Log(friend.GetBirthday())
            fmt.Println()
        }
    }
}

// 拉取好友
func TestIm_SNS_FetchFriends(t *testing.T) {
    var (
        err              error
        ret              *sns.FetchFriendsRet
        s                = NewIM().SNS()
        startIndex       = 0
        standardSequence = 0
        customSequence   = 0
    )
    
    for ret == nil || !ret.IsOver {
        ret, err = s.FetchFriends("assistant", startIndex, standardSequence, customSequence)
        if err != nil {
            t.Error(err)
            return
        }
        
        startIndex = ret.NextStartIndex
        standardSequence = ret.StandardSequence
        customSequence = ret.CustomSequence
        
        t.Log("下一个开始点：", ret.NextStartIndex)
        t.Log("是否拉取完毕：", ret.IsOver)
        t.Log("标准排序：", ret.StandardSequence)
        t.Log("自定义排序：", ret.CustomSequence)
        t.Log("好友总数：", ret.FriendNum)
        t.Log("好友列表：")
        fmt.Println()
        for _, friend := range ret.Friends {
            if friend.IsValid() {
                t.Log(friend.GetUserId())
                t.Log(friend.GetAddSource())
                t.Log(friend.GetRemark())
                t.Log(friend.GetRemarkTime())
                t.Log(friend.GetGroup())
                t.Log(friend.GetNickname())
                t.Log(friend.GetBirthday())
                fmt.Println()
            }
        }
    }
}

// 添加黑名单
func TestIm_SNS_AddBlacklist(t *testing.T) {
    var userIds []string
    var userId string
    for i := 0; i < 10; i++ {
        userId = "test" + strconv.Itoa(i)
        userIds = append(userIds, userId)
    }
    
    results, err := NewIM().SNS().AddBlacklist("assistant", userIds)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(results)
}

// 删除黑名单
func TestIm_SNS_DeleteBlacklist(t *testing.T) {
    var userIds []string
    var userId string
    for i := 0; i < 5; i++ {
        userId = "test" + strconv.Itoa(i)
        userIds = append(userIds, userId)
    }
    
    results, err := NewIM().SNS().DeleteBlacklist("assistant", userIds)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(results)
}

// 拉取黑名单
func TestIm_SNS_FetchBlacklist(t *testing.T) {
    var (
        err              error
        ret              *sns.FetchBlacklistRet
        s                = NewIM().SNS()
        startIndex       = 0
        maxLimited       = 2
        standardSequence = 0
    )
    
    for ret == nil || !ret.IsOver {
        ret, err = s.FetchBlacklist("assistant", startIndex, maxLimited, standardSequence)
        if err != nil {
            t.Error(err)
            return
        }
        
        startIndex = ret.NextStartIndex
        standardSequence = ret.StandardSequence
        
        t.Log("下一个开始点：", startIndex)
        t.Log("标准排序：", standardSequence)
        t.Log("黑名单列表：")
        fmt.Println()
        for _, blacklist := range ret.Blacklists {
            t.Log(blacklist.UserId)
            t.Log(blacklist.Time)
            fmt.Println()
        }
    }
}

// 校验黑名单
func TestIm_SNS_CheckBlacklist(t *testing.T) {
    var userIds []string
    var userId string
    for i := 0; i < 10; i++ {
        userId = "test" + strconv.Itoa(i)
        userIds = append(userIds, userId)
    }
    
    results, err := NewIM().SNS().CheckBlacklist("assistant", userIds, sns.BlacklistCheckTypeSingle)
    if err != nil {
        t.Error(err)
        return
    }
    
    for _, result := range results {
        if result.ResultCode == 0 {
            t.Log(result.UserId)
            t.Log(result.Relation)
            switch result.Relation {
            case sns.BlackCheckResultTypeNO:
                t.Log("From_Account 的黑名单中没有 To_Account，但无法确定 To_Account 的黑名单中是否有 From_Account")
            case sns.BlackCheckResultTypeAWithB:
                t.Log("From_Account 的黑名单中有 To_Account，但无法确定 To_Account 的黑名单中是否有 From_Account")
            }
        } else {
            t.Log(result.ResultCode)
            t.Log(result.ResultInfo)
        }
        fmt.Println()
    }
}

// 添加分组
func TestIm_SNS_AddGroups(t *testing.T) {
    var userIds []string
    var userId string
    for i := 0; i < 10; i++ {
        userId = "test" + strconv.Itoa(i)
        userIds = append(userIds, userId)
    }
    
    _, results, err := NewIM().SNS().AddGroups("assistant", []string{
        "测试3",
        "测试4",
    }, userIds)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(results)
}

// 删除分组
func TestIm_SNS_DeleteGroups(t *testing.T) {
    _, err := NewIM().SNS().DeleteGroups("assistant", []string{
        "测试3",
        "测试4",
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 拉取分组
func TestIm_SNS_GetGroups(t *testing.T) {
    var (
        err          error
        lastSequence int
        groupName    = []string{
            "测试1",
            "测试2",
        }
        results []sns.GroupResult
    )
    
    lastSequence, results, err = NewIM().SNS().GetGroups("assistant", lastSequence, true, groupName)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(results)
}

// 发送单聊消息
func TestIm_Private_SendMessage(t *testing.T) {
    message := private.NewMessage()
    message.SetSender("assistant")
    message.SetReceiver("test1")
    message.SetLifeTime(30000)
    message.SetTimeStamp(time.Now().Unix())
    message.SetContent(types.MsgTextContent{
        Text: "Hello world",
    })
    
    ret, err := NewIM().Private().SendMessage(message)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(ret.MsgKey)
    t.Log(ret.MsgTime)
}

// 批量发单聊消息
func TestIm_Private_SendMessages(t *testing.T) {
    message := private.NewMessage()
    message.SetSender("assistant")
    message.AddReceiver("test1", "test2")
    message.SetContent(types.MsgTextContent{
        Text: "Hello world",
    })
    
    ret, err := NewIM().Private().SendMessages(message)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(ret.MsgKey)
    t.Log(ret.Errors)
}

// 导入单聊消息
func TestIm_Private_ImportMessage(t *testing.T) {
    message := private.NewMessage()
    message.SetSender("assistant")
    message.SetReceiver("test1")
    message.SetTimeStamp(time.Now().Unix())
    message.SetSyncOtherMachine()
    message.SetContent(types.MsgTextContent{
        Text: "Hello world",
    })
    
    err := NewIM().Private().ImportMessage(message)
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 查询单聊消息
func TestIm_Private_FetchMessages(t *testing.T) {
    var (
        err error
        p   = NewIM().Private()
        ret *private.FetchMessagesRet
        arg = private.FetchMessagesArg{
            FromUserId: "test1",
            ToUserId:   "assistant",
            MaxLimited: 5,
            MinTime:    time.Now().Add(-20 * time.Hour).Unix(),
            MaxTime:    time.Now().Unix(),
        }
    )
    
    for ret == nil || !ret.IsOver {
        ret, err = p.FetchMessages(arg)
        if err != nil {
            t.Error(err)
            return
        }
        
        if !ret.IsOver {
            arg.LastMsgKey = ret.LastMsgKey
            arg.MaxTime = ret.LastMsgTime
        }
        
        t.Log(ret.IsOver)
        t.Log(ret.LastMsgKey)
        t.Log(ret.LastMsgTime)
        t.Log(ret.MsgCount)
        t.Log(ret.MsgList)
        fmt.Println()
    }
}

// 分页拉取所有消息
func TestIm_Private_PullMessages(t *testing.T) {
    err := NewIM().Private().PullMessages(private.PullMessagesArg{
        FromUserId: "test1",
        ToUserId:   "assistant",
        MaxLimited: 5,
        MinTime:    time.Now().Add(-30 * time.Hour).Unix(),
        MaxTime:    time.Now().Unix(),
    }, func(ret *private.FetchMessagesRet) {
        t.Log(ret.IsOver)
        t.Log(ret.LastMsgKey)
        t.Log(ret.LastMsgTime)
        t.Log(ret.MsgCount)
        t.Log(ret.MsgList)
        fmt.Println()
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 撤销消息
func TestIm_Private_RevokeMessage(t *testing.T) {
    err := NewIM().Private().RevokeMessage("assistant", "test1", "31906_833502_1572869830")
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 设置单聊消息已读
func TestIm_Private_SetMessageRead(t *testing.T) {
    err := NewIM().Private().SetMessageRead("assistant", "test1")
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log("Success")
}

// 获取未读消息数
func TestIm_Private_GetUnreadMessageNum(t *testing.T) {
    ret, err := NewIM().Private().GetUnreadMessageNum("assistant", []string{
        "test1",
        "test2",
    })
    if err != nil {
        t.Error(err)
        return
    }
    
    t.Log(ret.Total)
    t.Log(ret.UnreadList)
    t.Log(ret.ErrorList)
}
