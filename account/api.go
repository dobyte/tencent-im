/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:07
 * @Desc: 账号管理
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
	// DeleteAccounts 删除多个帐号
	// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/36443
	DeleteAccounts(userIds []string) (results []DeleteResult, err error)
	// CheckAccounts 查询多个帐号.
	// 用于查询自有帐号是否已导入即时通信 IM，支持批量查询。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/38417
	CheckAccounts(userIds []string) (results []CheckResult, err error)
	// KickAccount 使帐号登录状态失效
	// 本接口适用于将 App 用户帐号的登录状态（例如 UserSig）失效。
	// 例如，开发者判断一个用户为恶意帐号后，可以调用本接口将该用户当前的登录状态失效，这样用户使用历史 UserSig 登录即时通信 IM 会失败。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/3853
	KickAccount(userId string) (err error)
	// GetAccountsOnlineState 查询多个帐号在线状态
	// 获取用户当前的登录状态。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/2566
	GetAccountsOnlineState(userIds []string, isNeedDetail ...bool) (ret *OnlineStatusRet, err error)
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
		req := importAccountsReq{UserIds: userIds}
		resp := &importAccountsResp{}

		if err = a.client.Post(serviceAccount, commandMultiImportAccount, req, resp); err != nil {
			return
		} else {
			failUserIds = resp.FailAccounts
		}
	}

	return
}

// DeleteAccounts 删除多个帐号
// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/36443
func (a *api) DeleteAccounts(userIds []string) (results []DeleteResult, err error) {
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
		} else {
			results = resp.ResultItem
		}
	}

	return
}

// CheckAccounts 查询多个帐号.
// 用于查询自有帐号是否已导入即时通信 IM，支持批量查询。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/38417
func (a *api) CheckAccounts(userIds []string) (results []CheckResult, err error) {
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
		} else {
			results = resp.Results
		}
	}

	return
}

// KickAccount 失效帐号登录状态
// 本接口适用于将 App 用户帐号的登录状态（例如 UserSig）失效。
// 例如，开发者判断一个用户为恶意帐号后，可以调用本接口将该用户当前的登录状态失效，这样用户使用历史 UserSig 登录即时通信 IM 会失败。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3853
func (a *api) KickAccount(userId string) (err error) {
	req := kickAccountReq{UserId: userId}

	if err = a.client.Post(serviceAccount, commandKickAccount, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// GetAccountsOnlineState 查询多个帐号在线状态
// 获取用户当前的登录状态。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2566
func (a *api) GetAccountsOnlineState(userIds []string, isNeedDetail ...bool) (ret *OnlineStatusRet, err error) {
	req := queryAccountsOnlineStatusReq{UserIds: userIds}
	resp := &queryAccountsOnlineStatusResp{}

	if len(isNeedDetail) > 0 && isNeedDetail[0] {
		req.IsNeedDetail = 1
	}

	if err = a.client.Post(serviceOpenIM, commandQueryAccountsOnlineStatus, req, resp); err != nil {
		return
	} else {
		ret = &OnlineStatusRet{
			Results: resp.QueryResult,
			Errors:  resp.ErrorList,
		}
	}

	return
}
