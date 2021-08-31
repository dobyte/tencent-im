/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/29 17:42
 * @Desc: Message Api Request And Response Type Definition.
 */

package private

import "github.com/dobyte/tencent-im/types"

type (
	SendMsgReq struct {
		SyncOtherMachine      int                   `json:"SyncOtherMachine"`
		FromAccount           string                `json:"From_Account"`
		ToAccount             string                `json:"To_Account"`
		MsgLifeTime           int                   `json:"MsgLifeTime"`
		MsgRandom             int                   `json:"MsgRandom"`
		MsgTimeStamp          int                   `json:"MsgTimeStamp"`
		MsgBody               []types.MsgBody       `json:"MsgBody"`
		CloudCustomData       string                `json:"CloudCustomData"`
		ForbidCallbackControl []string              `json:"ForbidCallbackControl"`
		OfflinePushInfo       types.OfflinePushInfo `json:"OfflinePushInfo"`
	}

	SendMsgResp struct {
		types.ActionBaseResp
		MsgTime int    `json:"MsgTime"`
		MsgKey  string `json:"MsgKey"`
	}

	BatchSendMsgReq struct {
		SyncOtherMachine      int                   `json:"SyncOtherMachine"`
		FromAccount           string                `json:"From_Account"`
		ToAccount             []string              `json:"To_Account"`
		MsgRandom             int                   `json:"MsgRandom"`
		MsgBody               []types.MsgBody       `json:"MsgBody"`
		CloudCustomData       string                `json:"CloudCustomData"`
		ForbidCallbackControl []string              `json:"ForbidCallbackControl"`
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
		SyncOtherMachine int             `json:"SyncOtherMachine"`
		FromAccount      string          `json:"From_Account"`
		ToAccount        string          `json:"To_Account"`
		MsgRandom        int             `json:"MsgRandom"`
		MsgTimeStamp     int             `json:"MsgTimeStamp"`
		MsgBody          []types.MsgBody `json:"MsgBody"`
		CloudCustomData  string          `json:"CloudCustomData"`
	}

	ImportMsgResp struct {
		types.ActionBaseResp
	}

	// FetchMessagesArg 拉取消息参数
	FetchMessagesArg struct {
		FromUserId string `json:"From_Account"` // （必填）会话其中一方的 UserID，若已指定发送消息方帐号，则为消息发送方
		ToUserId   string `json:"To_Account"`   // （必填）会话其中一方的 UserID
		MaxLimited int    `json:"MaxCnt"`       // （必填）请求的消息条数
		MinTime    int64  `json:"MinTime"`      // （必填）请求的消息时间范围的最小值
		MaxTime    int64  `json:"MaxTime"`      // （必填）请求的消息时间范围的最大值
		LastMsgKey string `json:"LastMsgKey"`   // （选填）上一次拉取到的最后一条消息的 MsgKey，续拉时需要填该字段，填写方法见上方
	}

	// 拉取消息参数（响应）
	fetchMessagesResp struct {
		types.ActionBaseResp
		Complete    int           `json:"Complete"`    // 是否全部拉取，0表示未全部拉取，需要续拉，1表示已全部拉取
		LastMsgTime int64         `json:"LastMsgTime"` // 本次拉取到的消息里的最后一条消息的时间
		LastMsgKey  string        `json:"LastMsgKey"`  // 本次拉取到的消息里的最后一条消息的标识
		MsgCount    int           `json:"MsgCnt"`      // 本次拉取到的消息条数
		MsgList     []MessageItem `json:"MsgList"`     // 消息列表
	}

	// FetchMessagesRet 消息结果
	FetchMessagesRet struct {
		IsOver      bool          // 是否拉取结束
		LastMsgTime int64         // 本次拉取到的消息里的最后一条消息的时间
		LastMsgKey  string        // 本次拉取到的消息里的最后一条消息的标识
		MsgCount    int           // 本次拉取到的消息条数
		MsgList     []MessageItem // 消息列表
	}

	// MessageItem 消息项
	MessageItem struct {
		FromUserId      string          `json:"From_Account"`
		ToUserId        string          `json:"To_Account"`
		MsgSeq          int             `json:"MsgSeq"`
		MsgRandom       int             `json:"MsgRandom"`
		MsgTimeStamp    int64           `json:"MsgTimeStamp"`
		MsgFlagBits     int             `json:"MsgFlagBits"`
		MsgKey          string          `json:"MsgKey"`
		MsgBody         []types.MsgBody `json:"MsgBody"`
		CloudCustomData string          `json:"CloudCustomData"`
	}

	// 撤销消息（请求）
	revokeMessageReq struct {
		FromUserId string `json:"From_Account"` // （必填）消息发送方UserID
		ToUserId   string `json:"To_Account"`   // （必填）消息接收方UserID
		MsgKey     string `json:"MsgKey"`       // （必填）待撤回消息的唯一标识。该字段由 REST API 接口 单发单聊消息 和 批量发单聊消息 返回
	}

	// 设置单聊消息已读（请求）
	setMessageReadReq struct {
		UserId     string `json:"Report_Account"` // （必填）进行消息已读的用户UserId
		PeerUserId string `json:"Peer_Account"`   // （必填）进行消息已读的单聊会话的另一方用户UserId
	}

	// 查询单聊未读消息计数（请求）
	getUnreadMessageNumReq struct {
		UserId      string   `json:"To_Account"`             // （必填）待查询的用户UserId
		PeerUserIds []string `json:"Peer_Account,omitempty"` // （选填）待查询的单聊会话对端的用户UserId
	}

	// 查询单聊未读消息计数（响应）
	getUnreadMessageNumResp struct {
		types.ActionBaseResp
		AllUnreadMsgNum   int                  `json:"AllC2CUnreadMsgNum"`  // 单聊消息总未读数
		PeerUnreadMsgNums []unreadMessageNum   `json:"C2CUnreadMsgNumList"` // 单聊消息未读对端列表
		PeerErrors        []unreadMessageError `json:"ErrorList"`           // 查询错误列表
	}

	// 未读消息数
	unreadMessageNum struct {
		UserId       string `json:"Peer_Account"`    // 单聊会话对端UserId
		UnreadMsgNum int    `json:"C2CUnreadMsgNum"` // 该单聊会话的未读数
	}

	// 查询错误项
	unreadMessageError struct {
		UserId    string `json:"Peer_Account"` // 查询错误的目标UserId
		ErrorCode int    `json:"ErrorCode"`    // 查询错误的错误码。若目标帐号的错误码为70107表示该帐号不存在
	}

	// 未读消息结果
	UnreadMessageRet struct {
		Total      int                  // 单聊消息总未读数
		UnreadList map[string]int       // 未读消息数列表
		ErrorList  []unreadMessageError // 错误消息列表
	}
)
