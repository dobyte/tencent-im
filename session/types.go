/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/10/28 16:11
 * @Desc: TODO
 */

package session

import "github.com/dobyte/tencent-im/internal/types"

// FetchSessionsArg 拉取会话列表（参数）
type FetchSessionsArg struct {
	Account                 string // （必填）请求拉取该用户的会话列表
	TimeStamp               int    // （必填）普通会话的起始时间，第一页填 0
	StartIndex              int    // （必填）普通会话的起始位置，第一页填 0
	TopTimeStamp            int    // （必填）置顶会话的起始时间，第一页填 0
	TopStartIndex           int    // （必填）置顶会话的起始位置，第一页填 0
	IsAllowTopSession       bool   // （选填）是否支持置顶会话
	IsReturnEmptySession    bool   // （选填）是否返回空会话
	IsAllowTopSessionPaging bool   // （选填）是否支持置顶会话分页
}

// FetchSessionsRet 拉取会话列表（返回）
type FetchSessionsRet struct {
	IsOver        bool           // 是否拉完了数据
	TimeStamp     int            // 普通会话下一页拉取的起始时间，分页拉取时通过请求包的 TimeStamp 字段带给移动通信后台
	StartIndex    int            // 普通会话下一页拉取的起始位置，分页拉取时通过请求包的 StartIndex 字段带给移动通信后台
	TopTimeStamp  int            // 置顶会话下一页拉取的起始时间，分页拉取时通过请求包的 TopTimeStamp 字段带给移动通信后台
	TopStartIndex int            // 置顶会话下一页拉取的起始位置，分页拉取时通过请求包的 TopStartIndex 字段带给移动通信后台
	Sessions      []*SessionItem // 会话对象数组
}

// fetchSessionsReq 拉取会话列表（请求）
type fetchSessionsReq struct {
	Account       string `json:"From_Account"`  // （必填）请求拉取该用户的会话列表
	TimeStamp     int    `json:"TimeStamp"`     // （必填）普通会话的起始时间，第一页填 0
	StartIndex    int    `json:"StartIndex"`    // （必填）普通会话的起始位置，第一页填 0
	TopTimeStamp  int    `json:"TopTimeStamp"`  // （必填）置顶会话的起始时间，第一页填 0
	TopStartIndex int    `json:"TopStartIndex"` // （必填）置顶会话的起始位置，第一页填 0
	AssistFlags   int    `json:"AssistFlags"`   // （必填）会话辅助标志位（bit 0 - 是否支持置顶会话；bit 1 - 是否返回空会话；bit 2 - 是否支持置顶会话分页）
}

// fetchSessionsResp 拉取会话列表（响应）
type fetchSessionsResp struct {
	types.ActionBaseResp
	CompleteFlag  int            `json:"CompleteFlag"`  // 结束标识：1 表示已返回全量会话，0 表示还有会话没拉完
	TimeStamp     int            `json:"TimeStamp"`     // 普通会话下一页拉取的起始时间，分页拉取时通过请求包的 TimeStamp 字段带给移动通信后台
	StartIndex    int            `json:"StartIndex"`    // 普通会话下一页拉取的起始位置，分页拉取时通过请求包的 StartIndex 字段带给移动通信后台
	TopTimeStamp  int            `json:"TopTimeStamp"`  // 置顶会话下一页拉取的起始时间，分页拉取时通过请求包的 TopTimeStamp 字段带给移动通信后台
	TopStartIndex int            `json:"TopStartIndex"` // 置顶会话下一页拉取的起始位置，分页拉取时通过请求包的 TopStartIndex 字段带给移动通信后台
	Sessions      []*SessionItem `json:"SessionItem"`   // 会话对象数组
}

// SessionItem 会话对象
type SessionItem struct {
	Type    SessionType `json:"Type"`                 // 会话类型：1 表示 C2C 会话；2 表示 G2C 会话
	Account string      `json:"To_Account,omitempty"` // C2C 会话才会返回，返回会话方的 UserID
	GroupId string      `json:"GroupId,omitempty"`    // G2C 会话才会返回，返回群 ID
	MsgTime int         `json:"MsgTime"`              // 会话时间
	TopFlag int         `json:"TopFlag"`              // 置顶标记：0 标识普通会话；1 标识置顶会话
}

// deleteSessionReq 删除单个会话（请求）
type deleteSessionReq struct {
	FromAccount string `json:"From_Account"`          // （必填）请求删除该 UserID 的会话
	Type        int    `json:"type"`                  // （必填）会话类型：1 表示 C2C 会话；2 表示 G2C 会话
	ToAccount   string `json:"To_Account"`            // （必填）待删除的会话的 UserID
	ClearRamble int    `json:"ClearRamble,omitempty"` // （选填）是否清理漫游消息：1 表示清理漫游消息；0 表示不清理漫游消息
}
