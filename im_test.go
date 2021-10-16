/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/27 1:40 下午
 * @Desc: TODO
 */

package im_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
	
	"github.com/dobyte/tencent-im"
	"github.com/dobyte/tencent-im/account"
	"github.com/dobyte/tencent-im/group"
	"github.com/dobyte/tencent-im/operation"
	"github.com/dobyte/tencent-im/private"
	"github.com/dobyte/tencent-im/profile"
	"github.com/dobyte/tencent-im/push"
	"github.com/dobyte/tencent-im/sns"
)

const (
	assistant = "assistant"
	test1     = "test1"
	test2     = "test2"
	test3     = "test3"
	test4     = "test4"
	test5     = "test5"
	test6     = "test6"
	test7     = "test7"
	test8     = "test8"
	test9     = "test9"
)

func NewIM() im.IM {
	return im.NewIM(im.Options{
		AppId:     1400564830,
		AppSecret: "0d2a321b087fdb8fd5ed5ea14fe0489139086eb1b03541774fc9feeab8f2bfd3",
		UserId:    "administrator",
	})
}

// func NewIM() im.IM {
//     return im.NewIM(im.Options{
//         AppId:     0,
//         AppSecret: "",
//         UserId:    "administrator",
//     })
// }

// GetUserSig 获取UserSig签名
func TestIm_GetUserSig(t *testing.T) {
	tim := NewIM()
	
	for i := 0; i < 1000; i++ {
		tim.GetUserSig()
	}
	
	t.Log("Success")
}

// 导入单个账号
func TestIm_Account_ImportAccount(t *testing.T) {
	if err := NewIM().Account().ImportAccount(account.Info{
		UserId:    assistant,
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
	failedAccounts, err := NewIM().Account().ImportAccounts(
		test1,
		test2,
		test3,
		test4,
		test5,
		test6,
		test7,
		test8,
		test9,
	)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(failedAccounts)
}

// 删除单个账号
func TestIm_Account_DeleteAccount(t *testing.T) {
	err := NewIM().Account().DeleteAccount(test1)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 删除多个帐号
func TestIm_Account_DeleteAccounts(t *testing.T) {
	deleteResults, err := NewIM().Account().DeleteAccounts(
		test1,
		test2,
		test3,
		test4,
		test5,
		test6,
		test7,
		test8,
		test9,
	)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(deleteResults)
}

// 查询多个帐号
func TestIm_Account_CheckAccounts(t *testing.T) {
	checkResults, err := NewIM().Account().CheckAccounts(
		test1,
		test2,
		test3,
		test4,
		test5,
		test6,
		test7,
		test8,
		test9,
	)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(checkResults)
}

// 使帐号登录状态失效
func TestIm_Account_KickAccount(t *testing.T) {
	if err := NewIM().Account().KickAccount(test1); err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 查询帐号在线状态
func TestIm_Account_QueryAccountOnlineStatus(t *testing.T) {
	ret, err := NewIM().Account().GetAccountOnlineState(test2, true)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret)
}

// 查询多个帐号在线状态
func TestIm_Account_QueryAccountsOnlineStatus(t *testing.T) {
	resp, err := NewIM().Account().GetAccountsOnlineState([]string{
		test1,
		test2,
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
	message := push.NewMessage()
	message.SetSender(assistant)
	message.SetLifeTime(5000)
	message.SetContent(push.MsgTextContent{
		Text: "Hello Tencent IM",
	})
	message.OfflinePush().SetTitle("你好腾讯IM")
	message.OfflinePush().SetDesc("你好腾讯IM，我来了~~~")
	message.OfflinePush().SetPushFlag(push.PushFlagYes)
	message.OfflinePush().SetExt(map[string]interface{}{
		"url": "http://www.tencent.com",
	})
	message.OfflinePush().SetAndroidExtAsHuaweiIntentParam(push.HuaweiIntentParamIntent)
	message.OfflinePush().SetApnsBadgeMode(push.BadgeModeNormal)
	
	taskId, err := NewIM().Push().PushMessage(message)
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
	ret, err := NewIM().Push().GetUserAttrs(test1)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret)
}

// 设置用户属性
func TestIm_Push_SetUserAttrs(t *testing.T) {
	err := NewIM().Push().SetUserAttrs(map[string]map[string]interface{}{
		test1: {
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
		test1: {"age", "city"},
	})
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 获取用户标签
func TestIm_Push_GetUserTags(t *testing.T) {
	ret, err := NewIM().Push().GetUserTags(test1)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret)
}

// 添加用户标签
func TestIm_Push_AddUserTags(t *testing.T) {
	err := NewIM().Push().AddUserTags(map[string][]string{
		test1: {"chengdu"},
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
		test1: {"chengdu"},
	})
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 删除用户所有标签
func TestIm_Push_DeleteUserAllTags(t *testing.T) {
	err := NewIM().Push().DeleteUserAllTags(test1, test2)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 设置资料
func TestIm_Profile_SetProfile(t *testing.T) {
	p := profile.NewProfile()
	p.SetUserId(assistant)
	p.SetAvatar("http://www.qq.com")
	p.SetGender(profile.GenderTypeMale)
	p.SetLocation(1, 23, 27465, 92)
	// p.SetLocation(1, 23, 2, 92)
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
		assistant,
	}, []string{
		profile.StandardAttrNickname,
		profile.StandardAttrGender,
		profile.StandardAttrBirthday,
		profile.StandardAttrLocation,
		profile.StandardAttrLanguage,
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
	if err := NewIM().Mute().SetNoSpeaking(assistant, &privateMuteTime, &groupMuteTime); err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 查询全局禁言
func TestIm_Mute_GetNoSpeaking(t *testing.T) {
	privateMuteTime, groupMuteTime, err := NewIM().Mute().GetNoSpeaking(assistant)
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
	
	failUserIds, err := NewIM().Account().ImportAccounts(userIds...)
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
	
	failUserIds, err := NewIM().Account().ImportAccounts(userIds...)
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
	
	failUserIds, err := NewIM().Account().ImportAccounts(userIds...)
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
	
	results, err := NewIM().SNS().DeleteFriends("assistant", userIds, false)
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
		sns.FriendAttrAddSource,
		sns.FriendAttrRemark,
		sns.FriendAttrRemarkTime, // 此Tag无效，GetFriends内部忽略了
		sns.FriendAttrAddTime,
		sns.FriendAttrAddWording,
		sns.FriendAttrGroup,
		sns.StandardAttrNickname,
		sns.StandardAttrBirthday,
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
	message.SetSender(assistant)
	message.SetReceivers(test1)
	message.SetLifeTime(30000)
	message.SetTimestamp(time.Now().Unix())
	message.SetContent(private.MsgTextContent{
		Text: "Hello world",
	})
	message.OfflinePush().SetTitle("你好腾讯IM")
	message.OfflinePush().SetDesc("你好腾讯IM，我来了~~~")
	message.OfflinePush().SetPushFlag(private.PushFlagYes)
	message.OfflinePush().SetExt(map[string]interface{}{
		"url": "http://www.tencent.com",
	})
	message.OfflinePush().SetAndroidExtAsHuaweiIntentParam(private.HuaweiIntentParamIntent)
	message.OfflinePush().SetApnsBadgeMode(private.BadgeModeNormal)
	
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
	message.AddReceivers("test1", "test2")
	message.SetContent(private.MsgTextContent{
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
	message.SetReceivers("test1")
	message.SetTimestamp(time.Now().Unix())
	message.SetSyncOtherMachine(private.SyncOtherMachineYes)
	message.SetContent(private.MsgTextContent{
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

// 创建群组
func TestIm_Group_CreateGroup(t *testing.T) {
	g := group.NewGroup()
	g.SetName("测试群1")
	g.SetType(group.TypeLiveRoom)
	g.SetMaxMemberNum(30)
	g.SetAvatar("http://www.baidu.com")
	// g.SetId("test_group1")
	g.SetIntroduction("这是一个测试群")
	g.SetNotification("这是一个测试群公告")
	
	for i := 1; i < 10; i++ {
		member := group.NewMember()
		member.SetUserId("test" + strconv.Itoa(i))
		member.SetJoinTime(time.Now())
		g.AddMembers(member)
	}
	
	groupId, err := NewIM().Group().CreateGroup(g)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(groupId)
}

// 解散单个群
func TestIm_Group_DestroyGroup(t *testing.T) {
	err := NewIM().Group().DestroyGroup("test_group1")
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 获取单个群详细资料
func TestIm_Group_GetGroup(t *testing.T) {
	g, err := NewIM().Group().GetGroup("test_group1")
	if err != nil {
		t.Error(err)
		return
	}
	
	if g != nil {
		t.Log(g.GetId())
		t.Log(g.GetName())
		t.Log(g.GetType())
		t.Log(g.GetOwner())
		t.Log(g.GetAvatar())
	}
}

// 获取多个群详细资料
func TestIm_Group_GetGroups(t *testing.T) {
	groups, err := NewIM().Group().GetGroups([]string{
		"test_group1",
	})
	if err != nil {
		t.Error(err)
		return
	}
	
	for _, g := range groups {
		if err = g.GetError(); err != nil {
			t.Error(err)
		} else {
			t.Log(g.GetId())
			t.Log(g.GetName())
			t.Log(g.GetType())
			t.Log(g.GetOwner())
			t.Log(g.GetAvatar())
		}
	}
}

// 添加群成员
func TestIm_Group_AddGroupMembers(t *testing.T) {
	results, err := NewIM().Group().AddMembers("test_group1", []string{
		test1,
		test2,
	}, true)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(results)
}

// 删除群成员
func TestIm_Group_DeleteGroupMembers(t *testing.T) {
	err := NewIM().Group().DeleteMembers("test_group1", []string{
		test1,
		test2,
		test3,
	}, "测试删除", true)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 转让群组
func TestIm_Group_ChangeGroupOwner(t *testing.T) {
	err := NewIM().Group().ChangeGroupOwner("test_group1", test1)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 修改群基础资料
func TestIm_Group_UpdateGroup(t *testing.T) {
	g := group.NewGroup()
	g.SetName("测试群1")
	g.SetType(group.TypePublic)
	g.SetMaxMemberNum(30)
	g.SetAvatar("http://www.baidu.com")
	g.SetId("test_group1")
	g.SetIntroduction("这是一个测试群")
	g.SetNotification("这是一个测试群公告")
	
	err := NewIM().Group().UpdateGroup(g)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 查询用户在群组中的身份
func TestIm_Group_GetRolesInGroup(t *testing.T) {
	ret, err := NewIM().Group().GetRolesInGroup("test_group1", []string{
		test1,
		test2,
		test3,
	})
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret)
}

// 拉取群成员详细资料
func TestIm_Group_FetchGroupMembers(t *testing.T) {
	ret, err := NewIM().Group().FetchMembers("test_group1", 3, 2)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret.HasMore)
	t.Log(ret.Total)
	
	for _, member := range ret.List {
		t.Log(member.GetUserId())
	}
}

// 拉取App中的所有群组
func TestIm_Group_FetchGroupIds(t *testing.T) {
	ret, err := NewIM().Group().FetchGroupIds(3, 7964653962)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret.Total)
	t.Log(ret.Next)
	t.Log(ret.HasMore)
	t.Log(ret.List)
}

// 拉取App中的所有群组
func TestIm_Group_FetchGroups(t *testing.T) {
	ret, err := NewIM().Group().FetchGroups(3, 7964653962)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret.Total)
	t.Log(ret.Next)
	t.Log(ret.HasMore)
	
	for _, g := range ret.List {
		t.Log(g.GetId())
		t.Log(g.GetOwner())
		t.Log(g.GetName())
	}
}

// 修改群成员资料
func TestIm_Group_UpdateGroupMember(t *testing.T) {
	member := group.NewMember(test1)
	member.SetRole("Admin")
	member.SetNameCard("这是一个测试名片信息")
	member.SetMsgFlag(group.MsgFlagAcceptAndNotify)
	
	err := NewIM().Group().UpdateMember("test_group1", member)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 拉取用户所加入的群组
func TestIm_Group_FetchMemberGroups(t *testing.T) {
	ret, err := NewIM().Group().FetchMemberGroups(group.FetchMemberGroupsArg{
		UserId:               test1,
		Limit:                3,
		Offset:               2,
		IsWithLiveRoomGroups: true,
		IsWithNoActiveGroups: true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret.Total)
	t.Log(ret.HasMore)
	t.Log(ret.List)
}

// 批量禁言
func TestIm_Group_ForbidSendMessage(t *testing.T) {
	err := NewIM().Group().ForbidSendMessage("test_group1", []string{
		test1,
	}, 1000)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 取消禁言
func TestIm_Group_AllowSendMessage(t *testing.T) {
	err := NewIM().Group().AllowSendMessage("test_group1", []string{
		test1,
		test2,
	})
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 获取被禁言群成员列表
func TestIm_Group_GetShuttedUpMembers(t *testing.T) {
	shuttedUps, err := NewIM().Group().GetShuttedUpMembers("test_group1")
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(shuttedUps)
}

// 撤回单条群消息
func TestIm_Group_RevokeMessage(t *testing.T) {
	err := NewIM().Group().RevokeMessage("test_group1", 100)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 撤回多条群消息
func TestIm_Group_RevokeMessages(t *testing.T) {
	results, err := NewIM().Group().RevokeMessages("test_group1", 100)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(results)
}

// 设置成员未读消息计数
func TestIm_Group_SetMemberUnreadMsgNum(t *testing.T) {
	err := NewIM().Group().SetMemberUnreadMsgNum("test_group1", test1, 100)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 撤回指定用户发送的消息
func TestIm_Group_RevokeMemberMessages(t *testing.T) {
	err := NewIM().Group().RevokeMemberMessages("test_group1", test1)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 在群组中发送普通消息
func TestIm_Group_SendMessage(t *testing.T) {
	message := group.NewMessage()
	message.SetSender(test1)
	message.SetNoLastMsg()
	message.SetPriority("first")
	// message.SetOnlineOnlyFlag(group.OnlineOnlyFlagYes)
	message.SetContent(private.MsgTextContent{
		Text: "Hello world",
	})
	message.OfflinePush().SetTitle("你好腾讯IM")
	message.OfflinePush().SetDesc("你好腾讯IM，我来了~~~")
	message.OfflinePush().SetPushFlag(private.PushFlagYes)
	message.OfflinePush().SetExt(map[string]interface{}{
		"url": "http://www.tencent.com",
	})
	message.OfflinePush().SetAndroidExtAsHuaweiIntentParam(private.HuaweiIntentParamIntent)
	message.OfflinePush().SetApnsBadgeMode(private.BadgeModeNormal)
	message.AtAllMembers()
	message.AtMembers(test2)
	message.ClearAtMembers()
	
	ret, err := NewIM().Group().SendMessage("test_group1", message)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret.MsgSeq)
	t.Log(ret.MsgTime)
}

// 在群组中发送普通消息
func TestIm_Group_SendNotification(t *testing.T) {
	err := NewIM().Group().SendNotification("test_group1", "Hello welcome to the test group", test1)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log("Success")
}

// 导入群基础资料
func TestIm_Group_ImportGroup(t *testing.T) {
	g := group.NewGroup()
	g.SetName("测试群1")
	g.SetType(group.TypePublic)
	g.SetMaxMemberNum(30)
	g.SetAvatar("http://www.baidu.com")
	g.SetIntroduction("这是一个测试群")
	g.SetNotification("这是一个测试群公告")
	
	groupId, err := NewIM().Group().ImportGroup(g)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(groupId)
}

// 导入群消息
func TestIm_Group_ImportMessages(t *testing.T) {
	message := group.NewMessage()
	message.SetSender(test1)
	message.SetSendTime(time.Now().Unix())
	message.SetRandom(rand.Uint32())
	message.SetContent(private.MsgTextContent{
		Text: "Hello world",
	})
	
	results, err := NewIM().Group().ImportMessages("test_group1", message)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(results)
}

// 导入多个群成员
func TestIm_Group_ImportMembers(t *testing.T) {
	members := make([]*group.Member, 0)
	for i := 1; i < 10; i++ {
		member := group.NewMember()
		member.SetUserId("test" + strconv.Itoa(i))
		member.SetRole("Admin")
		member.SetUnreadMsgNum(10)
		member.SetJoinTime(time.Now())
		members = append(members, member)
	}
	
	results, err := NewIM().Group().ImportMembers("@TGS#25J4AWNHA", members...)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(results)
}

// 拉取群历史消息
func TestIm_Group_FetchMessages(t *testing.T) {
	ret, err := NewIM().Group().FetchMessages("test_group1", 5)
	if err != nil {
		t.Error(err)
		return
	}
	
	t.Log(ret)
}
