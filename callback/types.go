/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 16:36
 * @Desc: Callback request struct defined.
 */

package callback

import "github.com/dobyte/tencent-im/internal/types"

type (
	BaseResp struct {
		ErrorCode    int    `json:"ErrorCode"`
		ErrorInfo    string `json:"ErrorInfo"`
		ActionStatus string `json:"ActionStatus"`
	}

	// StateChange 状态变更回调
	StateChange struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		EventTime       int64  `json:"EventTime"`       // 触发本次回调的时间戳，单位为毫秒
		Info            struct {
			UserId string `json:"To_Account"` // 用户 UserID
			Action string `json:"Action"`     // 用户上线或者下线的动作，Login 表示上线（TCP 建立），Logout 表示下线（TCP 断开），Disconnect 表示网络断开（TCP 断开）
			Reason string `json:"Reason"`     // 用户上下线触发的原因
		} `json:"Info"` // 用户上下线的信息
		KickedDevice []struct {
			Platform string `json:"Platform"` // 被踢下线的设备的平台类型，可能的取值有"iOS", "Android", "Web", "Windows", "iPad", "Mac", "Linux"。
		} `json:"KickedDevice"` // 此字段表示其他被踢下线的设备的信息
	}

	// BeforeFriendAdd 添加好友之前回调
	BeforeFriendAdd struct {
		CallbackCommand string `json:"CallbackCommand"`   // 回调命令
		EventTime       int64  `json:"EventTime"`         // 触发本次回调的时间戳，单位为毫秒
		RequesterUserId string `json:"Requester_Account"` // 请求发起方的 UserID
		FromUserId      string `json:"From_Account"`      // 请求添加好友的用户的 UserID（A添加B为好友中的A）
		AddType         string `json:"AddType"`           // 加好友方式（默认双向加好友方式，Add_Type_Single：表示单向加好友 Add_Type_Both：表示双向加好友）
		ForceAddFlags   int    `json:"ForceAddFlags"`     // 管理员强制加好友标记（1：表示强制加好友 0：表示常规加好友方式）
		Friends         []struct {
			ToAccount  string `json:"To_Account"` // 请求添加的用户的 UserID
			Remark     string `json:"Remark"`     // From_Account 对 To_Account 设置的好友备注
			GroupName  string `json:"GroupName"`  // From_Account 对 To_Account 设置的好友分组
			AddSource  string `json:"AddSource"`  // 加好友来源
			AddWording string `json:"AddWording"` // 加好友附言
		} `json:"FriendItem"` // 加好友请求的参数
	}

	// BeforeFriendAddResp 添加好友之前回调应答
	BeforeFriendAddResp struct {
		BaseResp
		Results []*BeforeFriendAddResult `json:"ResultItem"` // App 后台的处理结果
	}

	// BeforeFriendAddResult App后台的处理结果
	BeforeFriendAddResult struct {
		UserId     string `json:"To_Account"` // （必填）请求添加的用户的 UserID
		ResultCode int    `json:"ResultCode"` // （必填）错误码：0表示允许加好友; 非0值表示不允许加好友; 如果不允许加好友，请将错误码设置在[38000, 39000]
		ResultInfo string `json:"ResultInfo"` // （必填）错误信息
	}

	// BeforeFriendResponse 添加好友回应之前回调
	BeforeFriendResponse struct {
		CallbackCommand string `json:"CallbackCommand"`   // 回调命令
		EventTime       int64  `json:"EventTime"`         // 触发本次回调的时间戳，单位为毫秒
		RequesterUserId string `json:"Requester_Account"` // 请求发起方的 UserID
		FromUserId      string `json:"From_Account"`      // 请求加好友回应的用户的 UserID
		Friends         []struct {
			ToAccount      string `json:"To_Account"`     // 请求回应的用户的 UserID
			Remark         string `json:"Remark"`         // From_Account 对 To_Account 设置的好友备注
			TagName        string `json:"TagName"`        // From_Account 对 To_Account 设置的好友分组
			ResponseAction string `json:"ResponseAction"` // 加好友回应方式，Response_Action_AgreeAndAdd 表示同意且添加对方为好友；Response_Action_Agree 表示同意对方加自己为好友；Response_Action_Reject 表示拒绝对方的加好友请求
		} `json:"ResponseFriendItem"` // 加好友回应请求的参数
	}

	// BeforeFriendResponseResp 添加好友之前回调应答
	BeforeFriendResponseResp struct {
		BaseResp
		Results []*BeforeFriendAddResult `json:"ResultItem"` // App 后台的处理结果
	}

	// BeforeFriendResponseResult App后台的处理结果
	BeforeFriendResponseResult struct {
		UserId     string `json:"To_Account"` // （必填）请求添加的用户的 UserID
		ResultCode int    `json:"ResultCode"` // （必填）错误码：0表示允许加好友; 非0值表示不允许加好友; 如果不允许加好友，请将错误码设置在[38000, 39000]
		ResultInfo string `json:"ResultInfo"` // （必填）错误信息
	}

	// AfterFriendAdd 添加好友之后
	AfterFriendAdd struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		ClientCmd       string `json:"ClientCmd"`       // 触发回调的命令字：加好友请求，合理的取值如下：friend_add、FriendAdd; 加好友回应，合理的取值如下：friend_response、FriendResponse
		AdminUserId     string `json:"Admin_Account"`   // 如果当前请求是后台触发的加好友请求，则该字段被赋值为管理员帐号；否则为空
		ForceFlag       int    `json:"ForceFlag"`       // 管理员强制加好友标记：1 表示强制加好友；0 表示常规加好友方式
		PairList        []struct {
			FromUserId      string `json:"From_Account"`      // From_Account 的好友表中增加了 To_Account
			ToUserId        string `json:"To_Account"`        // To_Account 被增加到了 From_Account 的好友表中
			InitiatorUserId string `json:"Initiator_Account"` // 发起加好友请求的用户的 UserID
		} `json:"PairList"` // 成功添加的好友对
	}

	// AfterFriendDelete 删除好友之后回调
	AfterFriendDelete struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		PairList        []struct {
			FromUserId string `json:"From_Account"` // From_Account 的好友表中删除了 To_Account
			ToUserId   string `json:"To_Account"`   // To_Account 从 From_Account 的好友表中删除
		} `json:"PairList"` // 成功删除的好友
	}

	// AfterBlacklistAdd 添加黑名单之后回调
	AfterBlacklistAdd struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		PairList        []struct {
			FromUserId string `json:"From_Account"` // From_Account 的黑名单列表中添加了 To_Account
			ToUserId   string `json:"To_Account"`   // To_Account 被加入到 From_Account 的黑名单列表中
		} `json:"PairList"` // 成功添加的黑名单关系链对
	}

	// AfterBlacklistDelete 删除黑名单之后回调
	AfterBlacklistDelete struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		PairList        []struct {
			FromUserId string `json:"From_Account"` // From_Account 的黑名单列表中删除了 To_Account
			ToUserId   string `json:"To_Account"`   // To_Account 从 From_Account 的黑名单列表中删除
		} `json:"PairList"` // 成功删除的黑名单对
	}

	// BeforePrivateMessageSend 发单聊消息之前回调
	BeforePrivateMessageSend struct {
		CallbackCommand string           `json:"CallbackCommand"` // 回调命令
		FromUserId      string           `json:"From_Account"`    // 消息发送者 UserID
		ToUserId        string           `json:"To_Account"`      // 消息接收者 UserID
		MsgSeq          int              `json:"MsgSeq"`          // 消息序列号，用于标记该条消息（32位无符号整数）
		MsgRandom       int              `json:"MsgRandom"`       // 消息随机数，用于标记该条消息（32位无符号整数）
		MsgTime         int64            `json:"MsgTime"`         // 消息的发送时间戳，单位为秒，单聊消息优先使用 MsgTime 进行排序，同一秒发送的消息则按 MsgSeq 排序，MsgSeq 值越大消息越靠后
		MsgKey          string           `json:"MsgKey"`          // 该条消息的唯一标识，可根据该标识进行 REST API 撤回单聊消息
		OnlineOnlyFlag  int              `json:"OnlineOnlyFlag"`  // 在线消息，为1，否则为0
		MsgBody         []*types.MsgBody `json:"MsgBody"`         // 消息体
		CloudCustomData string           `json:"CloudCustomData"` // 消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到）
	}

	// BeforePrivateMessageSendResp 发单聊消息之前回调应答
	BeforePrivateMessageSendResp struct {
		BaseResp
		MsgBody         []*types.MsgBody `json:"MsgBody,omitempty"`         // （选填）App 修改之后的消息，如果没有，则默认使用用户发送的消息
		CloudCustomData string           `json:"CloudCustomData,omitempty"` // （选填）经过 App 修改之后的消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到），即时通信 IM 后台将把修改后的消息发送给接收方
	}

	// AfterPrivateMessageSend 发单聊消息之后回调
	AfterPrivateMessageSend struct {
		CallbackCommand string           `json:"CallbackCommand"` // 回调命令
		FromUserId      string           `json:"From_Account"`    // 消息发送者 UserID
		ToUserId        string           `json:"To_Account"`      // 消息接收者 UserID
		MsgSeq          int              `json:"MsgSeq"`          // 消息序列号，用于标记该条消息（32位无符号整数）
		MsgRandom       int              `json:"MsgRandom"`       // 消息随机数，用于标记该条消息（32位无符号整数）
		MsgTime         int64            `json:"MsgTime"`         // 消息的发送时间戳，单位为秒，单聊消息优先使用 MsgTime 进行排序，同一秒发送的消息则按 MsgSeq 排序，MsgSeq 值越大消息越靠后
		MsgKey          string           `json:"MsgKey"`          // 该条消息的唯一标识，可根据该标识进行 REST API 撤回单聊消息
		OnlineOnlyFlag  int              `json:"OnlineOnlyFlag"`  // 在线消息，为1，否则为0
		MsgBody         []*types.MsgBody `json:"MsgBody"`         // 消息体
		CloudCustomData string           `json:"CloudCustomData"` // 消息自定义数据（云端保存，会发送到对端，程序卸载重装后还能拉取到）
		SendMsgResult   int              `json:"SendMsgResult"`   // 该条消息的下发结果，0表示下发成功，非0表示下发失败
		ErrorInfo       string           `json:"ErrorInfo"`       // 该条消息下发失败的错误信息，若消息发送成功，则为"send msg succeed"
		UnreadMsgNum    int              `json:"UnreadMsgNum"`    // To_Account 未读的单聊消息总数量（包含所有的单聊会话）。若该条消息下发失败（例如被脏字过滤），该字段值为-1
	}

	// AfterPrivateMessageReport 单聊消息已读上报后回调
	AfterPrivateMessageReport struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		ReportUserId    string `json:"Report_Account"`  // 已读上报方 UserID
		PeerUserId      string `json:"Peer_Account"`    // 会话对端 UserID
		LastReadTime    int64  `json:"LastReadTime"`    // 已读时间
		UnreadMsgNum    int    `json:"UnreadMsgNum"`    // Report_Account 未读的单聊消息总数量（包含所有的单聊会话）
	}

	// AfterPrivateMessageRevoke 单聊消息撤回后回调
	AfterPrivateMessageRevoke struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		FromUserId      string `json:"From_Account"`    // 消息发送者 UserID
		ToUserId        string `json:"To_Account"`      // 消息接收者 UserID
		MsgKey          string `json:"MsgKey"`          // 消息的唯一标识
		UnreadMsgNum    int    `json:"UnreadMsgNum"`    // To_Account 未读的单聊消息总数量（包含所有的单聊会话）
	}

	// BeforeGroupCreate 创建群组之前回调
	BeforeGroupCreate struct {
		CallbackCommand string `json:"CallbackCommand"`  // 回调命令
		OperatorUserId  string `json:"Operator_Account"` // 操作者
		OwnerUserId     string `json:"Owner_Account"`    // 群主
		Type            string `json:"Type"`             // 群组类型
		Name            string `json:"Name"`             // 请求创建的群组的名称
		CreateGroupNum  int    `json:"CreateGroupNum"`   // 该用户已创建的同类的群组个数
		MemberList      []struct {
			UserId string `json:"Member_Account"` // 成员 UserID
		} `json:"MemberList"` // 请求创建的群组的初始化成员列表
	}

	// AfterGroupCreate 创建群组之后回调
	AfterGroupCreate struct {
		CallbackCommand string `json:"CallbackCommand"`  // 回调命令
		OperatorUserId  string `json:"Operator_Account"` // 操作者
		OwnerUserId     string `json:"Owner_Account"`    // 群主
		GroupId         string `json:"GroupId"`          // 群ID
		Type            string `json:"Type"`             // 群组类型
		Name            string `json:"Name"`             // 请求创建的群组的名称
		CreateGroupNum  int    `json:"CreateGroupNum"`   // 该用户已创建的同类的群组个数
		MemberList      []struct {
			UserId string `json:"Member_Account"` // 成员 UserID
		} `json:"MemberList"` // 请求创建的群组的初始化成员列表
		UserDefinedDataList []struct {
			Key   string `json:"Key"`
			Value string `json:"Value"`
		} `json:"UserDefinedDataList"` // 用户建群时的自定义字段
	}

	// BeforeApplyJoinGroup 申请入群之前回调
	BeforeApplyJoinGroup struct {
		CallbackCommand string `json:"CallbackCommand"`   // 回调命令
		GroupId         string `json:"GroupId"`           // 群ID
		Type            string `json:"Type"`              // 群组类型
		RequestorUserId string `json:"Requestor_Account"` // 申请者
	}

	// BeforeInviteJoinGroup 拉人入群之前回调
	BeforeInviteJoinGroup struct {
		CallbackCommand string `json:"CallbackCommand"`  // 回调命令
		GroupId         string `json:"GroupId"`          // 群ID
		Type            string `json:"Type"`             // 群组类型
		OperatorUserId  string `json:"Operator_Account"` // 操作者
		MemberList      []struct {
			UserId string `json:"Member_Account"` // 成员 UserID
		} `json:"DestinationMembers"` // 要拉入群组的 UserID 集合
	}

	// BeforeInviteJoinGroupResp 拉人入群之前回调应答
	BeforeInviteJoinGroupResp struct {
		BaseResp
		RefusedMemberUserIds []string `json:"RefusedMembers_Account,omitempty"` // 拒绝加入的用户列表
	}

	// AfterNewMemberJoinGroup 新成员入群之后回调
	AfterNewMemberJoinGroup struct {
		CallbackCommand string `json:"CallbackCommand"`  // 回调命令
		GroupId         string `json:"GroupId"`          // 群ID
		Type            string `json:"Type"`             // 群组类型
		JoinType        string `json:"JoinType"`         // 入群方式：Apply（申请入群）；Invited（邀请入群）
		OperatorUserId  string `json:"Operator_Account"` // 操作者
		MemberList      []struct {
			UserId string `json:"Member_Account"` // 成员 UserID
		} `json:"NewMemberList"` // 新入群成员列表
	}

	// AfterMemberExitGroup 群成员离开之后回调
	AfterMemberExitGroup struct {
		CallbackCommand string `json:"CallbackCommand"`  // 回调命令
		GroupId         string `json:"GroupId"`          // 群ID
		Type            string `json:"Type"`             // 群组类型
		ExitType        string `json:"ExitType"`         // 成员离开方式：Kicked-被踢；Quit-主动退群
		OperatorUserId  string `json:"Operator_Account"` // 操作者
		MemberList      []struct {
			UserId string `json:"Member_Account"` // 成员 UserID
		} `json:"ExitMemberList"` // 离开群的成员列表
	}

	// BeforeGroupMessageSend 群内发言之前回调
	BeforeGroupMessageSend struct {
		CallbackCommand string           `json:"CallbackCommand"`  // 回调命令
		GroupId         string           `json:"GroupId"`          // 群ID
		Type            string           `json:"Type"`             // 群组类型
		FromUserId      string           `json:"From_Account"`     // 发送者
		OperatorUserId  string           `json:"Operator_Account"` // 请求的发起者
		OnlineOnlyFlag  int              `json:"OnlineOnlyFlag"`   // 在线消息，为1，否则为0；直播群忽略此属性，为默认值0。
		MsgRandom       int              `json:"Random"`           // 随机数
		MsgBody         []*types.MsgBody `json:"MsgBody"`          // 消息体
	}

	// BeforeGroupMessageSendResp 群内发言之前回调应答
	BeforeGroupMessageSendResp struct {
		BaseResp
		MsgBody []*types.MsgBody `json:"MsgBody,omitempty"` // （选填）App 修改之后的消息，如果没有，则默认使用用户发送的消息
	}

	// AfterGroupMessageSend 群内发言之后回调
	AfterGroupMessageSend struct {
		CallbackCommand string           `json:"CallbackCommand"`  // 回调命令
		GroupId         string           `json:"GroupId"`          // 群ID
		Type            string           `json:"Type"`             // 群组类型
		FromUserId      string           `json:"From_Account"`     // 发送者
		OperatorUserId  string           `json:"Operator_Account"` // 请求的发起者
		OnlineOnlyFlag  int              `json:"OnlineOnlyFlag"`   // 在线消息，为1，否则为0；直播群忽略此属性，为默认值0。
		MsgSeq          int              `json:"MsgSeq"`           // 消息的序列号
		MsgRandom       int              `json:"Random"`           // 随机数
		MsgTime         int64            `json:"MsgTime"`          // 消息的时间
		MsgBody         []*types.MsgBody `json:"MsgBody"`          // 消息体
	}

	// AfterGroupFull 群组满员之后回调
	AfterGroupFull struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		GroupId         string `json:"GroupId"`         // 群ID
	}

	// AfterGroupDestroyed 群组解散之后回调
	AfterGroupDestroyed struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		GroupId         string `json:"GroupId"`         // 群ID
		Type            string `json:"Type"`            // 群组类型
		Name            string `json:"Name"`            // 群组名称
		OwnerUserId     string `json:"Owner_Account"`   // 操作者
		MemberList      []struct {
			UserId string `json:"Member_Account"` // 成员 UserID
		} `json:"MemberList"` // 被解散的群组中的成员
	}

	// AfterGroupInfoChanged 群组资料修改之后回调
	AfterGroupInfoChanged struct {
		CallbackCommand string `json:"CallbackCommand"`  // 回调命令
		GroupId         string `json:"GroupId"`          // 群ID
		Type            string `json:"Type"`             // 群组类型
		Notification    string `json:"Notification"`     // 修改后的群公告
		OperatorUserId  string `json:"Operator_Account"` // 请求的发起者
	}
)
