/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:07
 * @Desc: Account Api Implementation.
 */

package account

import (
	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/types"
)

const (
	serviceAccount                   = "im_open_login_svc"
	serviceOpenIM                    = "openim"
	commandImportAccount             = "account_import"
	commandMultiImportAccount        = "multiaccount_import"
	commandDeleteAccount             = "account_delete"
	commandCheckAccount              = "account_check"
	commandKickAccount               = "kick"
	commandQueryAccountsOnlineStatus = "query_online_status"
)

type API interface {
	// ImportAccount 导入单个帐号
	// 本接口用于将 App 自有帐号导入即时通信 IM 帐号系统，
	// 为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1608
	ImportAccount(account AccountInfo) (err error)
	// ImportAccounts 导入多个帐号
	// 本接口用于批量将 App 自有帐号导入即时通信 IM 帐号系统，
	// 为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/4919
	ImportAccounts(userIds []string) (failUserIds []string, err error)
	// DeleteAccount 删除多个帐号
	// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/36443
	DeleteAccounts(userIds []string) (results []DeleteResultItem, err error)
	// CheckAccount 查询多个帐号.
	// 用于查询自有帐号是否已导入即时通信 IM，支持批量查询。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/38417
	CheckAccounts(userIds []string) (results []CheckResultItem, err error)
	// KickAccount 使帐号登录状态失效
	// 本接口适用于将 App 用户帐号的登录状态（例如 UserSig）失效。
	// 例如，开发者判断一个用户为恶意帐号后，可以调用本接口将该用户当前的登录状态失效，这样用户使用历史 UserSig 登录即时通信 IM 会失败。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/3853
	KickAccount(userId string) (err error)
	// QueryAccountsOnlineStatus 查询多个帐号在线状态
	// 获取用户当前的登录状态。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/2566
	QueryAccountsOnlineStatus(userIds []string, isNeedDetail ...bool) (*OnlineStatusResult, error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{
		client: client,
	}
}

// ImportAccount 导入单个帐号
// 本接口用于将 App 自有帐号导入即时通信 IM 帐号系统，
// 为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1608
func (a *api) ImportAccount(account AccountInfo) (err error) {
	if err = a.client.Post(serviceAccount, commandImportAccount, account, &types.ActionBaseResp{}); err != nil {
		return
	}
	
	return
}

// ImportAccounts 导入多个帐号
// 本接口用于批量将 App 自有帐号导入即时通信 IM 帐号系统，
// 为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/4919
func (a *api) ImportAccounts(userIds []string) (failUserIds []string, err error) {
	if len(userIds) > 0 {
		resp := &importAccountsResp{}
		
		if err = a.client.Post(serviceAccount, commandMultiImportAccount, importAccountsReq{
			UserIds: userIds,
		}, resp); err != nil {
			return
		}
		
		failUserIds = resp.FailAccounts
	}
	
	return
}

// DeleteAccount 删除多个帐号
// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/36443
func (a *api) DeleteAccounts(userIds []string) (results []DeleteResultItem, err error) {
	if len(userIds) > 0 {
		req := deleteAccountsReq{}
		resp := &deleteAccountsResp{}
		
		for _, userId := range userIds {
			req.DeleteItem = append(req.DeleteItem, accountItem{
				UserId: userId,
			})
		}
		
		if err = a.client.Post(serviceAccount, commandDeleteAccount, req, resp); err != nil {
			return
		}
		
		results = resp.ResultItem
	}
	
	return
}

// CheckAccount 查询多个帐号.
// 用于查询自有帐号是否已导入即时通信 IM，支持批量查询。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/38417
func (a *api) CheckAccounts(userIds []string) (results []CheckResultItem, err error) {
	if len(userIds) > 0 {
		req := checkAccountsReq{}
		resp := &checkAccountsResp{}
		
		for _, userId := range userIds {
			req.CheckItem = append(req.CheckItem, accountItem{
				UserId: userId,
			})
		}
		
		if err = a.client.Post(serviceAccount, commandCheckAccount, req, resp); err != nil {
			return
		}
		
		results = resp.ResultItem
	}
	
	return
}

// KickAccount 失效帐号登录状态
// 本接口适用于将 App 用户帐号的登录状态（例如 UserSig）失效。
// 例如，开发者判断一个用户为恶意帐号后，可以调用本接口将该用户当前的登录状态失效，这样用户使用历史 UserSig 登录即时通信 IM 会失败。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3853
func (a *api) KickAccount(userId string) (err error) {
	if err = a.client.Post(serviceAccount, commandKickAccount, kickAccountReq{
		UserId: userId,
	}, &types.ActionBaseResp{}); err != nil {
		return
	}
	
	return
}

// QueryAccountsOnlineStatus 查询多个帐号在线状态
// 获取用户当前的登录状态。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2566
func (a *api) QueryAccountsOnlineStatus(userIds []string, isNeedDetail ...bool) (*OnlineStatusResult, error) {
	var isNeed int
	if len(isNeedDetail) > 0 && isNeedDetail[0] {
		isNeed = 1
	}
	
	resp := &queryAccountsOnlineStatusResp{}
	
	if err := a.client.Post(serviceOpenIM, commandQueryAccountsOnlineStatus, queryAccountsOnlineStatusReq{
		IsNeedDetail: isNeed,
		UserIds:      userIds,
	}, resp); err != nil {
		return nil, err
	}
	
	return &OnlineStatusResult{
		Results: resp.QueryResult,
		Errors:  resp.ErrorList,
	}, nil
}
