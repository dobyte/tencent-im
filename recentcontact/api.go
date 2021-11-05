/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/10/28 16:05
 * @Desc: TODO
 */

package recentcontact

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

	// PullSessions 续拉取会话列表
	// 本API是借助"拉取会话列表"API进行扩展实现
	// 支持分页拉取会话列表
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/62118
	PullSessions(arg *PullSessionsArg, fn func(ret *FetchSessionsRet)) (err error)

	// DeleteSession 删除单个会话
	// 删除指定会话，支持同步清理漫游消息。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/62119
	DeleteSession(fromUserId, toUserId string, SessionType SessionType, isClearRamble ...bool) (err error)
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
		UserId:        arg.UserId,
		TimeStamp:     arg.TimeStamp,
		StartIndex:    arg.StartIndex,
		TopTimeStamp:  arg.TopTimeStamp,
		TopStartIndex: arg.TopStartIndex,
	}

	if arg.IsAllowTopSession {
		req.AssistFlags += 1 << 0
	}

	if arg.IsReturnEmptySession {
		req.AssistFlags += 1 << 1
	}

	if arg.IsAllowTopSessionPaging {
		req.AssistFlags += 1 << 2
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
		List:          resp.Sessions,
		HasMore:       resp.CompleteFlag == 0,
	}

	return
}

// PullSessions 续拉取会话列表
// 本API是借助"拉取会话列表"API进行扩展实现
// 支持分页拉取会话列表
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/62118
func (a *api) PullSessions(arg *PullSessionsArg, fn func(ret *FetchSessionsRet)) (err error) {
	var (
		ret *FetchSessionsRet
		req = &FetchSessionsArg{
			UserId:                  arg.UserId,
			IsAllowTopSession:       arg.IsAllowTopSession,
			IsReturnEmptySession:    arg.IsReturnEmptySession,
			IsAllowTopSessionPaging: arg.IsAllowTopSessionPaging,
		}
	)

	for ret == nil || ret.HasMore {
		ret, err = a.FetchSessions(req)
		if err != nil {
			return
		}

		fn(ret)

		if ret.HasMore {
			req.TimeStamp = ret.TimeStamp
			req.StartIndex = ret.StartIndex
			req.TopTimeStamp = ret.TopTimeStamp
			req.TopStartIndex = ret.TopStartIndex
		}
	}

	return
}

// DeleteSession 删除单个会话
// 删除指定会话，支持同步清理漫游消息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/62119
func (a *api) DeleteSession(fromUserId, toUserId string, SessionType SessionType, isClearRamble ...bool) (err error) {
	req := &deleteSessionReq{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		Type:       SessionType,
	}

	if len(isClearRamble) > 0 && isClearRamble[0] {
		req.ClearRamble = 1
	}

	if err = a.client.Post(service, commandDeleteSession, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}
