/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 16:36
 * @Desc: Callback request struct defined.
 */

package callback

type (
	CommandBaseReq struct {
		CallbackCommand string `json:"CallbackCommand"` // 回调命令
		EventTime       int    `json:"EventTime"`       // 触发本次回调的时间戳，单位为毫秒
	}

	// KickedDevice 其他被踢下线的设备的信息
	KickedDevice struct {
		Platform string `json:"Platform"` // 被踢下线的设备的平台类型，可能的取值有"iOS", "Android", "Web", "Windows", "iPad", "Mac", "Linux"。
	}

	// StateChangeInfo 用户上下线的信息
	StateChangeInfo struct {
		UserId string `json:"To_Account"` // 用户 UserID
		Action string `json:"Action"`     // 用户上线或者下线的动作，Login 表示上线（TCP 建立），Logout 表示下线（TCP 断开），Disconnect 表示网络断开（TCP 断开）
		Reason string `json:"Reason"`     // 用户上下线触发的原因
	}

	// StateChange 状态变更回调
	StateChange struct {
		CommandBaseReq
		Info         StateChangeInfo `json:"Info"`         // 用户上下线的信息
		KickedDevice []KickedDevice  `json:"KickedDevice"` // 此字段表示其他被踢下线的设备的信息
	}

	// BeforeFriendAdd 添加好友之前回调
	BeforeFriendAdd struct {
		CommandBaseReq
		RequesterAccount string                      `json:"Requester_Account"` // 请求发起方的 UserID
		FromAccount      string                      `json:"From_Account"`      // 请求添加好友的用户的 UserID（A添加B为好友中的A）
		FriendItem       []BeforeFriendAddFriendItem `json:"FriendItem"`        // 加好友请求的参数
		AddType          string                      `json:"AddType"`           // 加好友方式（默认双向加好友方式，Add_Type_Single：表示单向加好友 Add_Type_Both：表示双向加好友）
		ForceAddFlags    int                         `json:"ForceAddFlags"`     // 管理员强制加好友标记（1：表示强制加好友 0：表示常规加好友方式）
	}

	// BeforeFriendAddFriendItem 加好友请求的参数
	BeforeFriendAddFriendItem struct {
		ToAccount  string `json:"To_Account"` // 请求添加的用户的 UserID
		Remark     string `json:"Remark"`     // From_Account 对 To_Account 设置的好友备注
		GroupName  string `json:"GroupName"`  // From_Account 对 To_Account 设置的好友分组
		AddSource  string `json:"AddSource"`  // 加好友来源
		AddWording string `json:"AddWording"` // 加好友附言
	}

	// BeforeFriendResponse 添加好友回应之前回调
	BeforeFriendResponse struct {
		CommandBaseReq
		RequesterAccount string                           `json:"Requester_Account"`  // 请求发起方的 UserID
		FromAccount      string                           `json:"From_Account"`       // 请求加好友回应的用户的 UserID
		FriendItem       []BeforeFriendResponseFriendItem `json:"ResponseFriendItem"` // 加好友回应请求的参数
	}

	// BeforeFriendResponseFriendItem 加好友回应请求的参数
	BeforeFriendResponseFriendItem struct {
		ToAccount      string `json:"To_Account"`     // 请求回应的用户的 UserID
		Remark         string `json:"Remark"`         // From_Account 对 To_Account 设置的好友备注
		TagName        string `json:"TagName"`        // From_Account 对 To_Account 设置的好友分组
		ResponseAction string `json:"ResponseAction"` // 加好友回应方式，Response_Action_AgreeAndAdd 表示同意且添加对方为好友；Response_Action_Agree 表示同意对方加自己为好友；Response_Action_Reject 表示拒绝对方的加好友请求
	}

	FriendItem struct {
		FromUserId string `json:"From_Account"`
		UserId     string `json:"To_Account"` // 请求添加的用户的UserID
		Remark     string `json:"Remark"`     // From_Account 对 To_Account 设置的好友备注
	}

	// FriendAdd 添加好友之前回调
	FriendAdd struct {
		CommandBaseReq
		PairList     []FriendItem `json:"PairList"`
		ClientCmd    string       `json:"ClientCmd"`
		AdminAccount string       `json:"Admin_Account"`
		ForceFlag    int          `json:"ForceFlag"`
	}

	SnsFriendDeleteItem struct {
		FromAccount string `json:"From_Account"`
		ToAccount   string `json:"To_Account"`
	}

	// SnsFriendDelete Sns.CallbackFriendDelete callback request package.
	SnsFriendDelete struct {
		CallbackCommand string                `json:"CallbackCommand"`
		PairList        []SnsFriendDeleteItem `json:"PairList"`
	}

	SnsBlackListAddItem struct {
		FromAccount string `json:"From_Account"`
		ToAccount   string `json:"To_Account"`
	}

	// SnsBlackListAdd Sns.CallbackFriendDelete callback request package.
	SnsBlackListAdd struct {
		CallbackCommand string                `json:"CallbackCommand"`
		PairList        []SnsBlackListAddItem `json:"PairList"`
	}

	SnsBlackListDeleteItem struct {
		FromAccount string `json:"From_Account"`
		ToAccount   string `json:"To_Account"`
	}

	// SnsBlackListDelete Sns.CallbackBlackListDelete callback request package.
	SnsBlackListDelete struct {
		CallbackCommand string                   `json:"CallbackCommand"`
		PairList        []SnsBlackListDeleteItem `json:"PairList"`
	}

	MsgBodyContentImageInfo struct {
		Type   int    `json:"Type"`
		Size   int    `json:"Size"`
		Width  int    `json:"Width"`
		Height int    `json:"Height"`
		URL    string `json:"URL"`
	}

	MsgBodyContent struct {
		Text           string                    `json:"Text"`
		Desc           string                    `json:"Desc"`
		Latitude       float64                   `json:"Latitude"`
		Longitude      float64                   `json:"Longitude"`
		Index          int                       `json:"Index"`
		Data           string                    `json:"Data"`
		Ext            string                    `json:"Ext"`
		Sound          string                    `json:"Sound"`
		Url            string                    `json:"Url"`
		Size           int                       `json:"Size"`
		Second         int                       `json:"Second"`
		DownloadFlag   int                       `json:"Download_Flag"`
		UUID           string                    `json:"UUID"`
		ImageFormat    int                       `json:"ImageFormat"`
		ImageInfoArray []MsgBodyContentImageInfo `json:"ImageInfoArray"`
	}

	MsgBody struct {
		MsgType    string         `json:"MsgType"`
		MsgContent MsgBodyContent `json:"MsgContent"`
	}

	// C2CBeforeSendMsg Sns.CallbackBeforeSendMsg callback request package.
	C2CBeforeSendMsg struct {
		CallbackCommand string    `json:"CallbackCommand"`
		FromAccount     string    `json:"From_Account"`
		ToAccount       string    `json:"To_Account"`
		MsgSeq          int       `json:"MsgSeq"`
		MsgRandom       int       `json:"MsgRandom"`
		MsgTime         int       `json:"MsgTime"`
		MsgKey          string    `json:"MsgKey"`
		MsgBody         []MsgBody `json:"MsgBody"`
		CloudCustomData string    `json:"CloudCustomData"`
	}

	// C2CAfterSendMsg Sns.CallbackAfterSendMsg callback request package.
	C2CAfterSendMsg struct {
		CallbackCommand string    `json:"CallbackCommand"`
		FromAccount     string    `json:"From_Account"`
		ToAccount       string    `json:"To_Account"`
		MsgSeq          int       `json:"MsgSeq"`
		MsgRandom       int       `json:"MsgRandom"`
		MsgTime         int       `json:"MsgTime"`
		MsgKey          string    `json:"MsgKey"`
		SendMsgResult   int       `json:"SendMsgResult"`
		ErrorInfo       string    `json:"ErrorInfo"`
		UnreadMsgNum    int       `json:"UnreadMsgNum"`
		MsgBody         []MsgBody `json:"MsgBody"`
		CloudCustomData string    `json:"CloudCustomData"`
	}

	GroupMember struct {
		MemberAccount string `json:"Member_Account"`
	}

	// GroupBeforeCreateGroup Group.CallbackBeforeCreateGroup callback request package.
	GroupBeforeCreateGroup struct {
		CallbackCommand string        `json:"CallbackCommand"`
		OperatorAccount string        `json:"Operator_Account"`
		OwnerAccount    string        `json:"Owner_Account"`
		Type            string        `json:"Type"`
		Name            string        `json:"Name"`
		CreatedGroupNum int           `json:"CreatedGroupNum"`
		MemberList      []GroupMember `json:"MemberList"`
	}

	GroupAfterCreateGroupUserDefinedDataItem struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	}

	// GroupAfterCreateGroup Group.CallbackAfterCreateGroup callback request package.
	GroupAfterCreateGroup struct {
		CallbackCommand     string                                     `json:"CallbackCommand"`
		GroupId             string                                     `json:"GroupId"`
		OperatorAccount     string                                     `json:"Operator_Account"`
		OwnerAccount        string                                     `json:"Owner_Account"`
		Type                string                                     `json:"Type"`
		Name                string                                     `json:"Name"`
		MemberList          []GroupMember                              `json:"MemberList"`
		UserDefinedDataList []GroupAfterCreateGroupUserDefinedDataItem `json:"UserDefinedDataList"`
	}

	// GroupBeforeApplyJoinGroup Group.CallbackBeforeApplyJoinGroup callback request package.
	GroupBeforeApplyJoinGroup struct {
		CallbackCommand  string `json:"CallbackCommand"`
		GroupId          string `json:"GroupId"`
		Type             string `json:"Type"`
		RequestorAccount string `json:"Requestor_Account"`
	}

	// GroupBeforeInviteJoinGroup Group.CallbackBeforeInviteJoinGroup callback request package.
	GroupBeforeInviteJoinGroup struct {
		CallbackCommand    string        `json:"CallbackCommand"`
		GroupId            string        `json:"GroupId"`
		Type               string        `json:"Type"`
		OperatorAccount    string        `json:"Operator_Account"`
		DestinationMembers []GroupMember `json:"DestinationMembers"`
	}

	// GroupAfterNewMemberJoin Group.CallbackAfterNewMemberJoin callback request package.
	GroupAfterNewMemberJoin struct {
		CallbackCommand string        `json:"CallbackCommand"`
		GroupId         string        `json:"GroupId"`
		Type            string        `json:"Type"`
		JoinType        string        `json:"JoinType"`
		OperatorAccount string        `json:"Operator_Account"`
		NewMemberList   []GroupMember `json:"NewMemberList"`
	}

	// GroupAfterMemberExit Group.CallbackAfterMemberExit callback request package.
	GroupAfterMemberExit struct {
		CallbackCommand string        `json:"CallbackCommand"`
		GroupId         string        `json:"GroupId"`
		Type            string        `json:"Type"`
		ExitType        string        `json:"ExitType"`
		OperatorAccount string        `json:"Operator_Account"`
		ExitMemberList  []GroupMember `json:"NewMemberList"`
	}

	// GroupBeforeSendMsg Group.CallbackBeforeSendMsg callback request package.
	GroupBeforeSendMsg struct {
		CallbackCommand string    `json:"CallbackCommand"`
		GroupId         string    `json:"GroupId"`
		Type            string    `json:"Type"`
		FromAccount     string    `json:"From_Account"`
		OperatorAccount string    `json:"Operator_Account"`
		Random          int       `json:"Random"`
		MsgBody         []MsgBody `json:"MsgBody"`
	}

	// GroupAfterSendMsg Group.CallbackAfterSendMsg callback request package.
	GroupAfterSendMsg struct {
		CallbackCommand string    `json:"CallbackCommand"`
		GroupId         string    `json:"GroupId"`
		Type            string    `json:"Type"`
		FromAccount     string    `json:"From_Account"`
		OperatorAccount string    `json:"Operator_Account"`
		Random          int       `json:"Random"`
		MsgSeq          int       `json:"MsgSeq"`
		MsgTime         int       `json:"MsgTime"`
		MsgBody         []MsgBody `json:"MsgBody"`
	}

	// GroupAfterGroupFull Group.CallbackAfterGroupFull callback request package.
	GroupAfterGroupFull struct {
		CallbackCommand string `json:"CallbackCommand"`
		GroupId         string `json:"GroupId"`
	}

	// GroupAfterGroupDestroyed Group.CallbackAfterGroupDestroyed callback request package.
	GroupAfterGroupDestroyed struct {
		CallbackCommand string        `json:"CallbackCommand"`
		GroupId         string        `json:"GroupId"`
		Type            string        `json:"Type"`
		OwnerAccount    string        `json:"Owner_Account"`
		Name            string        `json:"Name"`
		MemberList      []GroupMember `json:"MemberList"`
	}

	// GroupAfterGroupInfoChanged Group.CallbackAfterGroupInfoChanged callback request package.
	GroupAfterGroupInfoChanged struct {
		CallbackCommand string `json:"CallbackCommand"`
		GroupId         string `json:"GroupId"`
		Type            string `json:"Type"`
		OperatorAccount string `json:"Operator_Account"`
		Notification    string `json:"Notification"`
	}
)
