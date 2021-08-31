/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/5/29 17:42
 * @Desc: Message Api Request And Response Type Definition.
 */

package api

import "github.com/dobyte/tencent-im/types"

type (
    SendMsgReq struct {
        SyncOtherMachine      int                   `json:"SyncOtherMachine"`
        FromAccount           string                `json:"From_Account"`
        ToAccount             string                `json:"To_Account"`
        MsgLifeTime           int                   `json:"MsgLifeTime"`
        MsgRandom             int                   `json:"MsgRandom"`
        MsgTimeStamp          int                    `json:"MsgTimeStamp"`
        MsgBody               []types.MsgBody       `json:"MsgBody"`
        CloudCustomData       string                 `json:"CloudCustomData"`
        ForbidCallbackControl []string               `json:"ForbidCallbackControl"`
        OfflinePushInfo       types.OfflinePushInfo `json:"OfflinePushInfo"`
    }
    
    SendMsgResp struct {
        types.ActionBaseResp
        MsgTime int    `json:"MsgTime"`
        MsgKey  string `json:"MsgKey"`
    }
    
    BatchSendMsgReq struct {
        SyncOtherMachine      int                  `json:"SyncOtherMachine"`
        FromAccount           string               `json:"From_Account"`
        ToAccount             []string             `json:"To_Account"`
        MsgRandom             int                    `json:"MsgRandom"`
        MsgBody               []types.MsgBody       `json:"MsgBody"`
        CloudCustomData       string                 `json:"CloudCustomData"`
        ForbidCallbackControl []string               `json:"ForbidCallbackControl"`
        OfflinePushInfo       types.OfflinePushInfo `json:"OfflinePushInfo"`
    }
    
    BatchSendMsgErrorItem struct {
        ToAccount string `json:"To_Account"`
        ErrorCode int    `json:"ErrorCode"`
    }
    
    BatchSendMsgResp struct {
        types.ActionBaseResp
        MsgKey    string                  `json:"MsgKey"`
        ErrorList []BatchSendMsgErrorItem `json:"ErrorList"`
    }
    
    ImportMsgReq struct {
        SyncOtherMachine int            `json:"SyncOtherMachine"`
        FromAccount      string         `json:"From_Account"`
        ToAccount        string         `json:"To_Account"`
        MsgRandom        int            `json:"MsgRandom"`
        MsgTimeStamp     int              `json:"MsgTimeStamp"`
        MsgBody          []types.MsgBody `json:"MsgBody"`
        CloudCustomData  string           `json:"CloudCustomData"`
    }
    
    ImportMsgResp struct {
        types.ActionBaseResp
    }
    
    GetRoamMsgReq struct {
        FromAccount string `json:"From_Account"`
        ToAccount   string `json:"To_Account"`
        MaxCnt      int    `json:"MaxCnt"`
        MinTime     int    `json:"MinTime"`
        MaxTime     int    `json:"MaxTime"`
    }
    
    RoamMsgItem struct {
        FromAccount     string         `json:"From_Account"`
        ToAccount       string         `json:"To_Account"`
        MsgSeq          int            `json:"MsgSeq"`
        MsgRandom       int            `json:"MsgRandom"`
        MsgTimeStamp    int            `json:"MsgTimeStamp"`
        MsgFlagBits     int            `json:"MsgFlagBits"`
        MsgKey          string           `json:"MsgKey"`
        MsgBody         []types.MsgBody `json:"MsgBody"`
        CloudCustomData string           `json:"CloudCustomData"`
    }
    
    GetRoamMsgResp struct {
        types.ActionBaseResp
        Complete    int           `json:"Complete"`
        MsgCnt      int           `json:"MsgCnt"`
        LastMsgTime int           `json:"LastMsgTime"`
        LastMsgKey  string        `json:"LastMsgKey"`
        MsgList     []RoamMsgItem `json:"MsgList"`
    }
    
    WithdrawMsgReq struct {
        FromAccount string `json:"From_Account"`
        ToAccount   string `json:"To_Account"`
        MsgKey      string `json:"MsgKey"`
    }
    
    WithdrawMsgResp struct {
        types.ActionBaseResp
    }
    
    SetMsgReadReq struct {
        ReportAccount string `json:"Report_Account"`
        PeerAccount   string `json:"Peer_Account"`
    }
    
    SetMsgReadResp struct {
        types.ActionBaseResp
    }
    
    GetUnreadMsgNumReq struct {
        ToAccount   string   `json:"To_Account"`
        PeerAccount []string `json:"Peer_Account"`
    }
    
    UnreadMsgNumItem struct {
        PeerAccount     string `json:"Peer_Account"`
        C2CUnreadMsgNum int    `json:"C2CUnreadMsgNum"`
    }
    
    GetUnreadMsgNumResp struct {
        types.ActionBaseResp
        AllC2CUnreadMsgNum  int                `json:"AllC2CUnreadMsgNum"`
        C2CUnreadMsgNumList []UnreadMsgNumItem `json:"C2CUnreadMsgNumList"`
    }
)
