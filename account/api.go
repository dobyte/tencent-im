/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:07
 * @Desc: 账号管理
 */

package account

import (
	"fmt"

	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/internal/enum"
	"github.com/dobyte/tencent-im/internal/types"
)

const (
	serviceAccount                   = "im_open_login_svc"
	serviceOpenIM                    = "openim"
	commandImportAccount             = "account_import"
	commandImportAccounts            = "multiaccount_import"
	commandDeleteAccounts            = "account_delete"
	commandCheckAccounts             = "account_check"
	commandKickAccount               = "kick"
	commandQueryAccountsOnlineStatus = "query_online_status"

	batchImportAccountsLimit = 100 // 导入账号限制
	batchDeleteAccountsLimit = 100 // 删除账号限制
	batchCheckAccountsLimit  = 100 // 查询账号限制
)

type API interface {
	// ImportAccount 导入单个帐号
	// 本接口用于将 App 自有帐号导入即时通信 IM 帐号系统，
	// 为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1608
	ImportAccount(account *Account) (err error)

	// ImportAccounts 导入多个帐号
	// 本接口用于批量将 App 自有帐号导入即时通信 IM 帐号系统，
	// 为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/4919
	ImportAccounts(userIds ...string) (failUserIds []string, err error)

	// DeleteAccount 删除账号
	// 本方法拓展于“删除多个帐号（DeleteAccounts）”方法。
	// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/36443
	DeleteAccount(userId string) (err error)

	// DeleteAccounts 删除多个帐号
	// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/36443
	DeleteAccounts(userIds ...string) (results []*DeleteResult, err error)

	// CheckAccount 查询帐号导入状态
	// 本方法拓展于“查询多个帐号导入状态（CheckAccounts）”方法。
	// 用于查询自有帐号是否已导入即时通信 IM。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/38417
	CheckAccount(userId string) (bool, error)

	// CheckAccounts 查询多个帐号导入状态
	// 用于查询自有帐号是否已导入即时通信 IM，支持批量查询。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/38417
	CheckAccounts(userIds ...string) (results []*CheckResult, err error)

	// KickAccount 使帐号登录状态失效
	// 本接口适用于将 App 用户帐号的登录状态（例如 UserSig）失效。
	// 例如，开发者判断一个用户为恶意帐号后，可以调用本接口将该用户当前的登录状态失效，这样用户使用历史 UserSig 登录即时通信 IM 会失败。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/3853
	KickAccount(userId string) (err error)

	// GetAccountOnlineState 查询帐号在线状态
	// 本方法拓展于“查询多个帐号在线状态（GetAccountsOnlineState）”方法。
	// 获取用户当前的登录状态。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/2566
	GetAccountOnlineState(userId string, isNeedDetail ...bool) (*OnlineStatusResult, error)

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
	return &api{client: client}
}

// ImportAccount 导入单个帐号
// 本接口用于将 App 自有帐号导入即时通信 IM 帐号系统，
// 为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1608
func (a *api) ImportAccount(account *Account) (err error) {
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
func (a *api) ImportAccounts(userIds ...string) (failUserIds []string, err error) {
	if c := len(userIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the userid is not set")
		return
	} else if c > batchImportAccountsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of imported accounts cannot exceed %d", batchImportAccountsLimit))
		return
	}

	req := &importAccountsReq{UserIds: userIds}
	resp := &importAccountsResp{}

	if err = a.client.Post(serviceAccount, commandImportAccounts, req, resp); err != nil {
		return
	}

	failUserIds = resp.FailUserIds

	return
}

// DeleteAccount 删除账号
// 本方法拓展于“删除多个帐号（DeleteAccounts）”方法。
// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/36443
func (a *api) DeleteAccount(userId string) (err error) {
	results, err := a.DeleteAccounts(userId)
	if err != nil {
		return
	}

	for _, result := range results {
		if result.UserId == userId && result.ResultCode != enum.SuccessCode {
			return core.NewError(result.ResultCode, result.ResultInfo)
		}
	}

	return
}

// DeleteAccounts 删除多个帐号
// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/36443
func (a *api) DeleteAccounts(userIds ...string) (results []*DeleteResult, err error) {
	if c := len(userIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the userid is not set")
		return
	} else if c > batchDeleteAccountsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of deleted accounts cannot exceed %d", batchDeleteAccountsLimit))
		return
	}

	req := &deleteAccountsReq{}
	resp := &deleteAccountsResp{}

	for _, userId := range userIds {
		req.Deletes = append(req.Deletes, &accountItem{userId})
	}

	if err = a.client.Post(serviceAccount, commandDeleteAccounts, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// CheckAccount 查询帐号导入状态.
// 本方法拓展于“查询多个帐号导入状态（CheckAccounts）”方法。
// 用于查询自有帐号是否已导入即时通信 IM。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/38417
func (a *api) CheckAccount(userId string) (bool, error) {
	results, err := a.CheckAccounts(userId)
	if err != nil {
		return false, err
	}

	for _, result := range results {
		if result.UserId == userId {
			if result.ResultCode != enum.SuccessCode {
				return false, core.NewError(result.ResultCode, result.ResultInfo)
			} else {
				return result.Status == ImportedStatusYes, nil
			}
		}
	}

	return false, nil
}

// CheckAccounts 查询多个帐号导入状态.
// 用于查询自有帐号是否已导入即时通信 IM，支持批量查询。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/38417
func (a *api) CheckAccounts(userIds ...string) (results []*CheckResult, err error) {
	if c := len(userIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the account is not set")
		return
	} else if c > batchCheckAccountsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of checked accounts cannot exceed %d", batchCheckAccountsLimit))
		return
	}

	req := &checkAccountsReq{}
	resp := &checkAccountsResp{}

	for _, userId := range userIds {
		req.Checks = append(req.Checks, &accountItem{userId})
	}

	if err = a.client.Post(serviceAccount, commandCheckAccounts, req, resp); err != nil {
		return
	}

	results = resp.Results

	return
}

// KickAccount 失效帐号登录状态
// 本接口适用于将 App 用户帐号的登录状态（例如 UserSig）失效。
// 例如，开发者判断一个用户为恶意帐号后，可以调用本接口将该用户当前的登录状态失效，这样用户使用历史 UserSig 登录即时通信 IM 会失败。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/3853
func (a *api) KickAccount(userId string) (err error) {
	if err = a.client.Post(serviceAccount, commandKickAccount, &kickAccountReq{userId}, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// GetAccountOnlineState 查询帐号在线状态
// 本方法拓展于“查询多个帐号在线状态（GetAccountsOnlineState）”方法。
// 获取用户当前的登录状态。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2566
func (a *api) GetAccountOnlineState(userId string, isNeedDetail ...bool) (*OnlineStatusResult, error) {
	ret, err := a.GetAccountsOnlineState([]string{userId}, isNeedDetail...)
	if err != nil {
		return nil, err
	}

	for _, item := range ret.Errors {
		if item.UserId == userId && item.ErrorCode != enum.SuccessCode {
			return nil, core.NewError(item.ErrorCode, "account exception")
		}
	}

	for _, item := range ret.Results {
		if item.UserId == userId {
			return &item, nil
		}
	}

	return nil, nil
}

// GetAccountsOnlineState 查询多个帐号在线状态
// 获取用户当前的登录状态。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2566
func (a *api) GetAccountsOnlineState(userIds []string, isNeedDetail ...bool) (ret *OnlineStatusRet, err error) {
	req := &queryAccountsOnlineStatusReq{UserIds: userIds}
	resp := &queryAccountsOnlineStatusResp{}

	if len(isNeedDetail) > 0 && isNeedDetail[0] {
		req.IsNeedDetail = 1
	}

	if err = a.client.Post(serviceOpenIM, commandQueryAccountsOnlineStatus, req, resp); err != nil {
		return
	}

	ret = &OnlineStatusRet{
		Results: resp.Results,
		Errors:  resp.Errors,
	}

	return
}
