/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/29 17:43
 * @Desc: Group Api Request And Response Type Definition.
 */

package group

import "github.com/dobyte/tencent-im/internal/types"

type (
    // 拉取App中的所有群组（请求）
    fetchGroupIdsReq struct {
        Limit     int    `json:"Limit,omitempty"`
        Next      int    `json:"Next"`
        GroupType string `json:"GroupType,omitempty"`
    }
    
    // 拉取App中的所有群组（响应）
    fetchGroupIdsResp struct {
        types.ActionBaseResp
        Next        int       `json:"Next"`
        TotalCount  int       `json:"TotalCount"`
        GroupIdList []groupId `json:"GroupIdList"`
    }
    
    // FetchGroupIdsRet 拉取App中的所有群组ID返回
    FetchGroupIdsRet struct {
        Total   int
        Next    int
        HasMore bool
        List    []string
    }
    
    FetchGroupsRet struct {
        Total   int
        Next    int
        HasMore bool
        List    []*Group
    }
    
    groupId struct {
        GroupId string `json:"GroupId"`
    }
    
    // 自定义数据
    customData struct {
        Key   string      `json:"Key"`
        Value interface{} `json:"Value"`
    }
    
    // 创建群（请求）
    createGroupReq struct {
        OwnerUserId     string       `json:"Owner_Account,omitempty"`   // （选填）群主 ID（需是 已导入 的账号）。填写后自动添加到群成员中；如果不填，群没有群主
        GroupId         string       `json:"GroupId,omitempty"`         // （选填）为了使得群组 ID 更加简单，便于记忆传播，腾讯云支持 App 在通过 REST API 创建群组时 自定义群组 ID
        Type            string       `json:"Type"`                      // （必填）群组形态，包括 Public（陌生人社交群），Private（即 Work，好友工作群），ChatRoom（即 Meeting，会议群），AVChatRoom（直播群）
        Name            string       `json:"Name"`                      // （必填）群名称，最长30字节，使用 UTF-8 编码，1个汉字占3个字节
        Introduction    string       `json:"Introduction,omitempty"`    // （选填）群简介，最长240字节，使用 UTF-8 编码，1个汉字占3个字节
        Notification    string       `json:"Notification,omitempty"`    // （选填）群公告，最长300字节，使用 UTF-8 编码，1个汉字占3个字节
        FaceUrl         string       `json:"FaceUrl,omitempty"`         // （选填）群头像 URL，最长100字节
        MaxMemberNum    uint         `json:"MaxMemberCount,omitempty"`  // （选填）最大群成员数量，缺省时的默认值：付费套餐包上限，例如体验版是20，如果升级套餐包，需按照修改群基础资料修改这个字段
        ApplyJoinOption string       `json:"ApplyJoinOption,omitempty"` // （选填）申请加群处理方式。包含 FreeAccess（自由加入），NeedPermission（需要验证），DisableApply（禁止加群），不填默认为 NeedPermission（需要验证） 仅当创建支持申请加群的 群组 时，该字段有效
        AppDefinedData  []customData `json:"AppDefinedData,omitempty"`  // （选填）群组维度的自定义字段，默认情况是没有的，可以通过 即时通信 IM 控制台 进行配置，详情请参阅 自定义字段
        MemberList      []memberInfo `json:"MemberList,omitempty"`      // （选填）初始群成员列表，最多100个；成员信息字段详情请参阅 群成员资料
    }
    
    // 创建群（响应）
    createGroupResp struct {
        types.ActionBaseResp
        GroupId string `json:"GroupId"` // 群ID
    }
    
    // 群成员信息
    memberInfo struct {
        UserId               string       `json:"Member_Account"`            // 群成员ID
        Role                 string       `json:"Role,omitempty"`            // 群内身份
        JoinTime             int64        `json:"JoinTime,omitempty"`        // 入群时间
        MsgSeq               int          `json:"MsgSeq,omitempty"`          // 该成员当前已读消息Seq
        MsgFlag              string       `json:"MsgFlag,omitempty"`         // 消息接收选项
        LastSendMsgTime      int64        `json:"LastSendMsgTime,omitempty"` // 最后发送消息的时间
        NameCard             string       `json:"NameCard,omitempty"`        // 群名片
        ShutUpUntil          int64        `json:"ShutUpUntil"`               // 禁言截至时间
        AppMemberDefinedData []customData `json:"AppMemberDefinedData"`      // 群成员自定义数据
    }
    
    // 解散群（请求）
    destroyGroupReq struct {
        GroupId string `json:"GroupId"` // （必填）操作的群 ID
    }
    
    responseFilter struct {
        GroupBaseInfoFilter    []string `json:"GroupBaseInfoFilter,omitempty"`
        MemberInfoFilter       []string `json:"MemberInfoFilter,omitempty"`
        GroupCustomDataFilter  []string `json:"AppDefinedDataFilter_Group,omitempty"`
        MemberCustomDataFilter []string `json:"AppDefinedDataFilter_GroupMember,omitempty"`
    }
    
    // 获取群详细资料（请求）
    getGroupsReq struct {
        GroupIds       []string        `json:"GroupIdList"`
        ResponseFilter *responseFilter `json:"ResponseFilter,omitempty"`
    }
    
    // 获取群详细资料（响应）
    getGroupsResp struct {
        types.ActionBaseResp
        GroupInfos []groupInfo `json:"GroupInfo"`
    }
    
    groupInfo struct {
        GroupId         string       `json:"GroupId"`
        ErrorCode       int          `json:"ErrorCode"`
        ErrorInfo       string       `json:"ErrorInfo"`
        Type            string       `json:"Type"`
        Name            string       `json:"Name"`
        AppId           int          `json:"Appid"`
        Introduction    string       `json:"Introduction"`
        Notification    string       `json:"Notification"`
        FaceUrl         string       `json:"FaceUrl"`
        OwnerUserId     string       `json:"Owner_Account"`
        CreateTime      int64        `json:"CreateTime"`
        LastInfoTime    int64        `json:"LastInfoTime"`
        LastMsgTime     int64        `json:"LastMsgTime"`
        NextMsgSeq      int          `json:"NextMsgSeq"`
        MemberNum       uint         `json:"MemberNum"`
        MaxMemberNum    uint         `json:"MaxMemberNum"`
        ApplyJoinOption string       `json:"ApplyJoinOption"`
        ShutUpAllMember string       `json:"ShutUpAllMember"`
        AppDefinedData  []customData `json:"AppDefinedData"`
        MemberList      []memberInfo `json:"MemberList"`
    }
    
    // 获取群成员详细资料（请求）
    fetchGroupMembersReq struct {
        GroupId                string   `json:"GroupId"`
        Limit                  int      `json:"Limit"`
        Offset                 int      `json:"Offset"`
        MemberInfoFilter       []string `json:"MemberInfoFilter"`
        MemberRoleFilter       []string `json:"MemberRoleFilter"`
        MemberCustomDataFilter []string `json:"AppDefinedDataFilter_GroupMember"`
    }
    
    // 获取群成员详细资料（响应）
    fetchGroupMembersResp struct {
        types.ActionBaseResp
        MemberNum  int          `json:"MemberNum"`
        MemberList []memberInfo `json:"MemberList"`
    }
    
    // FetchGroupMembersRet 拉取群成员结果
    FetchGroupMembersRet struct {
        Total   int       // 成员数量
        HasMore bool      // 是否还有更多数据
        List    []*Member // 成员列表
    }
    
    // 修改群基础资料（请求）
    updateGroupReq struct {
        GroupId         string       `json:"GroupId"`
        Name            string       `json:"Name,omitempty"`
        Introduction    string       `json:"Introduction,omitempty"`
        Notification    string       `json:"Notification,omitempty"`
        FaceUrl         string       `json:"FaceUrl,omitempty"`
        MaxMemberNum    uint         `json:"MaxMemberNum,omitempty"`
        ApplyJoinOption string       `json:"ApplyJoinOption,omitempty"`
        ShutUpAllMember string       `json:"ShutUpAllMember,omitempty"`
        AppDefinedData  []customData `json:"AppDefinedData,omitempty"`
    }
    
    // 添加群成员（请求）
    addGroupMembersReq struct {
        GroupId    string          `json:"GroupId"`
        Silence    int             `json:"Silence,omitempty"`
        MemberList []addMemberItem `json:"MemberList"`
    }
    
    // 添加群成员（响应）
    addGroupMembersResp struct {
        types.ActionBaseResp
        MemberList []AddMembersResult `json:"MemberList"`
    }
    
    addMemberItem struct {
        UserId string `json:"Member_Account"`
    }
    
    // AddMembersResult 添加群成员结果
    AddMembersResult struct {
        UserId string `json:"Member_Account"`
        Result int    `json:"Result"`
    }
    
    // 删除群成员（请求）
    deleteGroupMembersReq struct {
        GroupId string   `json:"GroupId"`             // （必填）操作的群ID
        Silence int      `json:"Silence"`             // （选填）是否静默删人
        Reason  string   `json:"Reason"`              // （选填）踢出用户原因
        UserIds []string `json:"MemberToDel_Account"` // （必填）待删除的群成员
    }
    
    ModifyGroupMemberInfoReq struct {
        GroupId             string       `json:"GroupId"`
        MemberAccount       string       `json:"Member_Account"`
        Role                string       `json:"Role"`
        NameCard            string       `json:"NameCard"`
        MsgFlag             string       `json:"MsgFlag"`
        ShutUpUntil         int          `json:"ShutUpUntil"`
        AppMembertypesdData []customData `json:"AppMembertypesdData"`
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
    
    getRolesInGroupReq struct {
        GroupId string   `json:"GroupId"`
        UserIds []string `json:"User_Account"`
    }
    
    getRolesInGroupResp struct {
        types.ActionBaseResp
        MemberRoleList []memberRole `json:"UserIdList"`
    }
    
    memberRole struct {
        UserId string `json:"Member_Account"`
        Role   string `json:"Role"`
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
        GroupId               string                `json:"GroupId"`
        FromAccount           string                `json:"From_Account"`
        SendMsgControl        []string              `json:"SendMsgControl"`
        MsgPriority           string                `json:"MsgPriority"`
        OnlineOnlyFlag        int                   `json:"OnlineOnlyFlag"`
        ForbidCallbackControl []string              `json:"ForbidCallbackControl"`
        Random                int                   `json:"Random"`
        MsgBody               []types.MsgBody       `json:"MsgBody"`
        OfflinePushInfo       types.OfflinePushInfo `json:"OfflinePushInfo"`
        GroupAtInfo           []GroupAtInfoItem     `json:"GroupAtInfo"`
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
    
    // 转让群主（请求）
    changeGroupOwnerReq struct {
        GroupId     string `json:"GroupId"`
        OwnerUserId string `json:"NewOwner_Account"`
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
        OwnerAccount    string       `json:"Owner_Account"`
        GroupId         string       `json:"GroupId"`
        Type            string       `json:"Type"`
        Name            string       `json:"Name"`
        Introduction    string       `json:"Introduction"`
        Notification    string       `json:"Notification"`
        FaceUrl         string       `json:"FaceUrl"`
        MaxMemberCount  int          `json:"MaxMemberCount"`
        CreateTime      int          `json:"CreateTime"`
        ApplyJoinOption string       `json:"ApplyJoinOption"`
        ApptypesdData   []customData `json:"ApptypesdData"`
    }
    
    ImportGroupResp struct {
        types.ActionBaseResp
        GroupId string `json:"GroupId"`
    }
    
    GroupMsgItem struct {
        FromAccount string        `json:"From_Account"`
        SendTime    int           `json:"SendTime"`
        Random      int           `json:"Random"`
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
        FromAccount  string          `json:"From_Account"`
        IsPlaceMsg   int             `json:"IsPlaceMsg"`
        MsgBody      []types.MsgBody `json:"MsgBody"`
        MsgPriority  int             `json:"MsgPriority"`
        MsgRandom    int             `json:"MsgRandom"`
        MsgSeq       int             `json:"MsgSeq"`
        MsgTimeStamp int             `json:"MsgTimeStamp"`
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
