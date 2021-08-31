/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/5/27 16:36
 * @Desc: Callback request struct defined.
 */

package callback

type (
    StateStateChangeInfo struct {
        Action    string `json:"Action"`
        ToAccount string `json:"To_Account"`
        Reason    string `json:"Reason"`
    }
    
    // StateStateChange State.StateChange callback request package.
    StateStateChange struct {
        CallbackCommand string               `json:"CallbackCommand"`
        Info            StateStateChangeInfo `json:"Info"`
    }
    
    SnsFriendAddPairItem struct {
        FromAccount      string `json:"From_Account"`
        ToAccount        string `json:"To_Account"`
        InitiatorAccount string `json:"Initiator_Account"`
    }
    
    // SnsFriendAdd Sns.CallbackFriendAdd callback request package.
    SnsFriendAdd struct {
        CallbackCommand string                 `json:"CallbackCommand"`
        PairList        []SnsFriendAddPairItem `json:"PairList"`
        ClientCmd       string                 `json:"ClientCmd"`
        AdminAccount    string                 `json:"Admin_Account"`
        ForceFlag       int                    `json:"ForceFlag"`
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
