/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/10/28 16:05
 * @Desc: TODO
 */

package session

import (
	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/internal/types"
)

const (
	service              = "recentcontact"
	commandFetchSessions = "get_list"
	commandDeleteSession = "delete"
)

type API interface {
	// FetchSessions 拉取会话列表
	// 支持分页拉取会话列表
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/62118
	FetchSessions(arg *FetchSessionsArg) (ret *FetchSessionsRet, err error)

	// DeleteSession 删除单个会话
	// 删除指定会话，支持同步清理漫游消息。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/62119
	DeleteSession(fromAccount, toAccount string, SessionType SessionType, isClearRamble ...bool) (err error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// FetchSessions 拉取会话列表
// 支持分页拉取会话列表
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/62118
func (a *api) FetchSessions(arg *FetchSessionsArg) (ret *FetchSessionsRet, err error) {
	req := &fetchSessionsReq{
		Account:       arg.Account,
		TimeStamp:     arg.TimeStamp,
		StartIndex:    arg.StartIndex,
		TopTimeStamp:  arg.TopTimeStamp,
		TopStartIndex: arg.TopStartIndex,
	}

	if arg.IsAllowTopSession {
		req.AssistFlags += 1
	}

	if arg.IsReturnEmptySession {
		req.AssistFlags += 2
	}

	if arg.IsAllowTopSessionPaging {
		req.AssistFlags += 4
	}

	resp := &fetchSessionsResp{}

	if err = a.client.Post(service, commandFetchSessions, req, resp); err != nil {
		return
	}

	ret = &FetchSessionsRet{
		TimeStamp:     resp.TimeStamp,
		StartIndex:    resp.StartIndex,
		TopTimeStamp:  resp.TopTimeStamp,
		TopStartIndex: resp.TopStartIndex,
		Sessions:      resp.Sessions,
	}

	if resp.CompleteFlag == 1 {
		ret.IsOver = true
	}

	return
}

// DeleteSession 删除单个会话
// 删除指定会话，支持同步清理漫游消息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/62119
func (a *api) DeleteSession(fromAccount, toAccount string, SessionType SessionType, isClearRamble ...bool) (err error) {
	req := &deleteSessionReq{
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Type:        int(SessionType),
	}

	if len(isClearRamble) > 0 && isClearRamble[0] {
		req.ClearRamble = 1
	}

	if err = a.client.Post(service, commandDeleteSession, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}
