/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/29 17:43
 * @Desc: Group Api Request And Response Type Definition.
 */

package api

import "github.com/dobyte/tencent-im/types"

type (
	GetGroupListReq struct {
		Limit     int    `json:"Limit"`
		Next      int    `json:"Next"`
		GroupType string `json:"GroupType"`
	}
	
	GroupIdItem struct {
		GroupId string `json:"GroupId"`
	}
	
	GetGroupListResp struct {
		types.ActionBaseResp
		Next        int           `json:"Next"`
		TotalCount  int           `json:"TotalCount"`
		GroupIdList []GroupIdItem `json:"GroupIdList"`
	}
	
	ApptypesdItem struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	}
	
	MemberItem struct {
		MemberAccount        string           `json:"Member_Account"`
		Role                 string           `json:"Role"`
		AppMembertypesdData []ApptypesdItem `json:"AppMembertypesdData"`
	}
	
	CreateGroupReq struct {
		OwnerAccount    string           `json:"Owner_Account"`
		GroupId         string           `json:"GroupId"`
		Type            string           `json:"Type"`
		Name            string           `json:"Name"`
		Introduction    string           `json:"Introduction"`
		Notification    string           `json:"Notification"`
		FaceUrl         string           `json:"FaceUrl"`
		MaxMemberCount  int              `json:"MaxMemberCount"`
		ApplyJoinOption string           `json:"ApplyJoinOption"`
		ApptypesdData  []ApptypesdItem `json:"ApptypesdData"`
		MemberList      []MemberItem     `json:"MemberList"`
	}
	
	CreateGroupResp struct {
		types.ActionBaseResp
		GroupId string `json:"GroupId"`
	}
	
	DestroyGroupReq struct {
		GroupId string `json:"GroupId"`
	}
	
	DestroyGroupResp struct {
		types.ActionBaseResp
	}
	
	GetGroupInfoResponseFilter struct {
		GroupBaseInfoFilter             []string `json:"GroupBaseInfoFilter"`
		MemberInfoFilter                []string `json:"MemberInfoFilter"`
		ApptypesdDataFilterGroup       []string `json:"ApptypesdDataFilter_Group"`
		ApptypesdDataFilterGroupMember []string `json:"ApptypesdDataFilter_GroupMember"`
	}
	
	GetGroupInfoReq struct {
		GroupIdList    []string                   `json:"GroupIdList"`
		ResponseFilter GetGroupInfoResponseFilter `json:"ResponseFilter"`
	}
	
	GroupInfoItem struct {
		GroupId         string           `json:"GroupId"`
		ErrorCode       int              `json:"ErrorCode"`
		ErrorInfo       string           `json:"ErrorInfo"`
		Type            string           `json:"Type"`
		Name            string           `json:"Name"`
		AppId           int              `json:"Appid"`
		Introduction    string           `json:"Introduction"`
		Notification    string           `json:"Notification"`
		FaceUrl         string           `json:"FaceUrl"`
		OwnerAccount    string           `json:"Owner_Account"`
		CreateTime      int              `json:"CreateTime"`
		LastInfoTime    int              `json:"LastInfoTime"`
		LastMsgTime     int              `json:"LastMsgTime"`
		NextMsgSeq      int              `json:"NextMsgSeq"`
		MemberNum       int              `json:"MemberNum"`
		MaxMemberNum    int              `json:"MaxMemberNum"`
		ApplyJoinOption string           `json:"ApplyJoinOption"`
		ShutUpAllMember string           `json:"ShutUpAllMember"`
		ApptypesdData  []ApptypesdItem `json:"ApptypesdData"`
		MemberList      []MemberItem     `json:"MemberList"`
	}
	
	GetGroupInfoResp struct {
		types.ActionBaseResp
		GroupInfo []GroupInfoItem `json:"GroupInfo"`
	}
	
	GetGroupMemberInfoReq struct {
		GroupId                         string   `json:"GroupId"`
		Limit                           int      `json:"Limit"`
		Offset                          int      `json:"Offset"`
		MemberInfoFilter                []string `json:"MemberInfoFilter"`
		MemberRoleFilter                []string `json:"MemberRoleFilter"`
		ApptypesdDataFilterGroupMember []string `json:"ApptypesdDataFilter_GroupMember"`
	}
	
	MemberInfoItem struct {
		MemberAccount        string           `json:"Member_Account"`
		Role                 string           `json:"Role"`
		JoinTime             int              `json:"JoinTime"`
		MsgSeq               int              `json:"MsgSeq"`
		MsgFlag              string           `json:"MsgFlag"`
		LastSendMsgTime      int              `json:"LastSendMsgTime"`
		ShutUpUntil          int              `json:"ShutUpUntil"`
		AppMembertypesdData []ApptypesdItem `json:"AppMembertypesdData"`
	}
	
	GetGroupMemberInfoResp struct {
		types.ActionBaseResp
		MemberNum  int              `json:"MemberNum"`
		MemberList []MemberInfoItem `json:"MemberList"`
	}
	
	ModifyGroupBaseInfoReq struct {
		GroupId         string           `json:"GroupId"`
		Name            string           `json:"Name"`
		Introduction    string           `json:"Introduction"`
		Notification    string           `json:"Notification"`
		FaceUrl         string           `json:"FaceUrl"`
		MaxMemberNum    int              `json:"MaxMemberNum"`
		ApplyJoinOption string           `json:"ApplyJoinOption"`
		ShutUpAllMember string           `json:"ShutUpAllMember"`
		ApptypesdData  []ApptypesdItem `json:"ApptypesdData"`
	}
	
	ModifyGroupBaseInfoResp struct {
		types.ActionBaseResp
	}
	
	AddMemberItem struct {
		MemberAccount string `json:"Member_Account"`
	}
	
	AddGroupMemberReq struct {
		GroupId    string          `json:"GroupId"`
		Silence    int             `json:"Silence"`
		MemberList []AddMemberItem `json:"MemberList"`
	}
	
	AddMemberRetItem struct {
		MemberAccount string `json:"Member_Account"`
		Result        int    `json:"Result"`
	}
	
	AddGroupMemberResp struct {
		types.ActionBaseResp
		MemberList []AddMemberRetItem `json:"MemberList"`
	}
	
	DeleteGroupMemberReq struct {
		GroupId            string   `json:"GroupId"`
		Silence            int      `json:"Silence"`
		Reason             string   `json:"Reason"`
		MemberToDelAccount []string `json:"MemberToDel_Account"`
	}
	
	DeleteGroupMemberResp struct {
		types.ActionBaseResp
	}
	
	ModifyGroupMemberInfoReq struct {
		GroupId              string           `json:"GroupId"`
		MemberAccount        string           `json:"Member_Account"`
		Role                 string           `json:"Role"`
		NameCard             string           `json:"NameCard"`
		MsgFlag              string           `json:"MsgFlag"`
		ShutUpUntil          int              `json:"ShutUpUntil"`
		AppMembertypesdData []ApptypesdItem `json:"AppMembertypesdData"`
	}
	
	ModifyGroupMemberInfoResp struct {
		types.ActionBaseResp
	}
	
	GetJoinedGroupListResponseFilter struct {
		GroupBaseInfoFilter []string `json:"GroupBaseInfoFilter"`
		SelfInfoFilter      []string `json:"SelfInfoFilter"`
	}
	
	GetJoinedGroupListReq struct {
		MemberAccount      string                           `json:"Member_Account"`
		Limit              int                              `json:"Limit"`
		Offset             int                              `json:"Offset"`
		WithHugeGroups     int                              `json:"WithHugeGroups"`
		WithNoActiveGroups int                              `json:"WithNoActiveGroups"`
		ResponseFilter     GetJoinedGroupListResponseFilter `json:"ResponseFilter"`
	}
	
	SelfInfo struct {
		JoinTime int    `json:"JoinTime"`
		MsgFlag  string `json:"MsgFlag"`
		Role     string `json:"Role"`
		MsgSeq   int    `json:"MsgSeq"`
	}
	
	GroupIdList struct {
		GroupId         string   `json:"GroupId"`
		ApplyJoinOption string   `json:"ApplyJoinOption"`
		CreateTime      int      `json:"CreateTime"`
		FaceUrl         string   `json:"FaceUrl"`
		Introduction    string   `json:"Introduction"`
		Notification    string   `json:"Notification"`
		LastInfoTime    int      `json:"LastInfoTime"`
		LastMsgTime     int      `json:"LastMsgTime"`
		MaxMemberNum    int      `json:"MaxMemberNum"`
		MemberNum       int      `json:"MemberNum"`
		Name            string   `json:"Name"`
		NextMsgSeq      int      `json:"NextMsgSeq"`
		OwnerAccount    string   `json:"Owner_Account"`
		ShutUpAllMember string   `json:"ShutUpAllMember"`
		Type            string   `json:"Type"`
		SelfInfo        SelfInfo `json:"SelfInfo"`
	}
	
	GetJoinedGroupListResp struct {
		types.ActionBaseResp
		TotalCount  int           `json:"TotalCount"`
		GroupIdList []GroupIdList `json:"GroupIdList"`
	}
	
	GetRoleInGroupReq struct {
		GroupId     string   `json:"GroupId"`
		UserAccount []string `json:"User_Account"`
	}
	
	UserIdList struct {
		MemberAccount string `json:"Member_Account"`
		Role          string `json:"Role"`
	}
	
	GetRoleInGroupResp struct {
		types.ActionBaseResp
		UserIdList []UserIdList `json:"UserIdList"`
	}
	
	ForbidSendMsgReq struct {
		GroupId        string   `json:"GroupId"`
		MembersAccount []string `json:"Members_Account"`
		ShutUpTime     int      `json:"ShutUpTime"`
	}
	
	ForbidSendMsgResp struct {
		types.ActionBaseResp
	}
	
	GetGroupShuttedUinReq struct {
		GroupId string `json:"GroupId"`
	}
	
	ShuttedUinItem struct {
		MembersAccount []string `json:"Members_Account"`
		ShuttedUntil   int      `json:"ShutUpTime"`
	}
	
	GetGroupShuttedUinResp struct {
		types.ActionBaseResp
		ShuttedUinList []ShuttedUinItem `json:"ShuttedUinItem"`
	}
	
	GroupAtInfoItem struct {
		GroupAtAllFlag int    `json:"GroupAtAllFlag"`
		GroupAtAccount string `json:"GroupAt_Account"`
	}
	
	SendGroupMsgReq struct {
		GroupId               string                 `json:"GroupId"`
		FromAccount           string                 `json:"From_Account"`
		SendMsgControl        []string               `json:"SendMsgControl"`
		MsgPriority           string                 `json:"MsgPriority"`
		OnlineOnlyFlag        int                    `json:"OnlineOnlyFlag"`
		ForbidCallbackControl []string               `json:"ForbidCallbackControl"`
		Random                int                    `json:"Random"`
		MsgBody               []types.MsgBody       `json:"MsgBody"`
		OfflinePushInfo       types.OfflinePushInfo `json:"OfflinePushInfo"`
		GroupAtInfo           []GroupAtInfoItem      `json:"GroupAtInfo"`
	}
	
	SendGroupMsgResp struct {
		types.ActionBaseResp
		MsgTime int `json:"MsgTime"`
		MsgSeq  int `json:"MsgSeq"`
	}
	
	SendGroupSystemNotificationReq struct {
		GroupId          string   `json:"GroupId"`
		Content          string   `json:"Content"`
		ToMembersAccount []string `json:"ToMembers_Account"`
	}
	
	SendGroupSystemNotificationResp struct {
		types.ActionBaseResp
	}
	
	ChangeGroupOwnerReq struct {
		GroupId         string `json:"GroupId"`
		NewOwnerAccount string `json:"NewOwner_Account"`
	}
	
	ChangeGroupOwnerResp struct {
		types.ActionBaseResp
	}
	
	MsgSeqItem struct {
		MsgSeq int `json:"MsgSeq"`
	}
	
	RecallGroupMsgReq struct {
		GroupId    string       `json:"GroupId"`
		MsgSeqList []MsgSeqItem `json:"MsgSeqList"`
	}
	
	RecallRetItem struct {
		MsgSeq  int `json:"MsgSeq"`
		RetCode int `json:"RetCode"`
	}
	
	RecallGroupMsgResp struct {
		types.ActionBaseResp
		RecallRetList []RecallRetItem `json:"RecallRetList"`
	}
	
	ImportGroupReq struct {
		OwnerAccount    string           `json:"Owner_Account"`
		GroupId         string           `json:"GroupId"`
		Type            string           `json:"Type"`
		Name            string           `json:"Name"`
		Introduction    string           `json:"Introduction"`
		Notification    string           `json:"Notification"`
		FaceUrl         string           `json:"FaceUrl"`
		MaxMemberCount  int              `json:"MaxMemberCount"`
		CreateTime      int              `json:"CreateTime"`
		ApplyJoinOption string           `json:"ApplyJoinOption"`
		ApptypesdData  []ApptypesdItem `json:"ApptypesdData"`
	}
	
	ImportGroupResp struct {
		types.ActionBaseResp
		GroupId string `json:"GroupId"`
	}
	
	GroupMsgItem struct {
		FromAccount string         `json:"From_Account"`
		SendTime    int            `json:"SendTime"`
		Random      int            `json:"Random"`
		MsgBody     types.MsgBody `json:"MsgBody"`
	}
	
	ImportGroupMsgReq struct {
		GroupId string         `json:"GroupId"`
		MsgList []GroupMsgItem `json:"MsgList"`
	}
	
	ImportGroupMsgResultItem struct {
		MsgSeq  int `json:"MsgSeq"`
		MsgTime int `json:"MsgTime"`
		Result  int `json:"Result"`
	}
	
	ImportGroupMsgResp struct {
		types.ActionBaseResp
		ImportMsgResult []ImportGroupMsgResultItem `json:"ImportMsgResult"`
	}
	
	ImportMemberItem struct {
		MemberAccount string `json:"Member_Account"`
		Role          string `json:"Role"`
		JoinTime      int    `json:"JoinTime"`
		UnreadMsgNum  int    `json:"UnreadMsgNum"`
	}
	
	ImportGroupMemberReq struct {
		GroupId    string             `json:"GroupId"`
		MemberList []ImportMemberItem `json:"MemberList"`
	}
	
	ImportMemberResultItem struct {
		MemberAccount string `json:"Member_Account"`
		Result        int    `json:"Result"`
	}
	
	ImportGroupMemberResp struct {
		types.ActionBaseResp
		MemberList []ImportMemberResultItem `json:"MemberList"`
	}
	
	SetUnreadMsgNumReq struct {
		GroupId       string `json:"GroupId"`
		MemberAccount string `json:"Member_Account"`
		UnreadMsgNum  int    `json:"UnreadMsgNum"`
	}
	
	SetUnreadMsgNumResp struct {
		types.ActionBaseResp
	}
	
	DeleteGroupMsgBySenderReq struct {
		GroupId       string `json:"GroupId"`
		SenderAccount string `json:"Sender_Account"`
	}
	
	DeleteGroupMsgBySenderResp struct {
		types.ActionBaseResp
	}
	
	GetGroupSimpleMsgReq struct {
		GroupId      string `json:"GroupId"`
		ReqMsgSeq    int    `json:"ReqMsgSeq"`
		ReqMsgNumber int    `json:"ReqMsgNumber"`
	}
	
	RspMsgItem struct {
		FromAccount  string           `json:"From_Account"`
		IsPlaceMsg   int              `json:"IsPlaceMsg"`
		MsgBody      []types.MsgBody `json:"MsgBody"`
		MsgPriority  int              `json:"MsgPriority"`
		MsgRandom    int              `json:"MsgRandom"`
		MsgSeq       int              `json:"MsgSeq"`
		MsgTimeStamp int              `json:"MsgTimeStamp"`
	}
	
	GetGroupSimpleMsgResp struct {
		types.ActionBaseResp
		GroupId    string       `json:"GroupId"`
		IsFinished int          `json:"IsFinished"`
		RspMsgList []RspMsgItem `json:"RspMsgList"`
	}
	
	GetOnlineMemberNumReq struct {
		GroupId string `json:"GroupId"`
	}
	
	GetOnlineMemberNumResp struct {
		types.ActionBaseResp
		OnlineMemberNum int `json:"OnlineMemberNum"`
	}
)
