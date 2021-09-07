/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:07
 * @Desc: 账号管理
 */

package account

import (
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
)

type API interface {
    // ImportAccount 导入单个帐号
    // 本接口用于将 App 自有帐号导入即时通信 IM 帐号系统，
    // 为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1608
    ImportAccount(account Info) (err error)
    
    // ImportAccounts 导入多个帐号
    // 本接口用于批量将 App 自有帐号导入即时通信 IM 帐号系统，
    // 为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/4919
    ImportAccounts(userId ...string) (failUserIds []string, err error)
    
    // DeleteAccount 删除账号
    // 本方法拓展于“删除多个帐号（DeleteAccounts）”方法。
    // 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/36443
    DeleteAccount(userId string) error
    
    // DeleteAccounts 删除多个帐号
    // 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/36443
    DeleteAccounts(userId ...string) (results []DeleteResult, err error)
    
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
    CheckAccounts(userId ...string) (results []CheckResult, err error)
    
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
func (a *api) ImportAccount(account Info) (err error) {
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
func (a *api) ImportAccounts(userId ...string) (failUserIds []string, err error) {
    if len(userId) > 0 {
        req := importAccountsReq{UserIds: userId}
        resp := &importAccountsResp{}
        
        if err = a.client.Post(serviceAccount, commandImportAccounts, req, resp); err != nil {
            return
        } else {
            failUserIds = resp.FailAccounts
        }
    }
    
    return
}

// DeleteAccount 删除账号
// 本方法拓展于“删除多个帐号（DeleteAccounts）”方法。
// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/36443
func (a *api) DeleteAccount(userId string) error {
    results, err := a.DeleteAccounts(userId)
    if err != nil {
        return err
    }
    
    for _, result := range results {
        if result.UserId == userId && result.ResultCode != enum.SuccessCode {
            return core.NewError(result.ResultCode, result.ResultInfo)
        }
    }
    
    return nil
}

// DeleteAccounts 删除多个帐号
// 仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/36443
func (a *api) DeleteAccounts(userId ...string) (results []DeleteResult, err error) {
    if len(userId) > 0 {
        req := deleteAccountsReq{}
        resp := &deleteAccountsResp{}
        
        for _, uid := range userId {
            req.DeleteItem = append(req.DeleteItem, accountItem{
                UserId: uid,
            })
        }
        
        if err = a.client.Post(serviceAccount, commandDeleteAccounts, req, resp); err != nil {
            return
        } else {
            results = resp.ResultItem
        }
    }
    
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
                return result.AccountStatus == ImportedStatusYes, nil
            }
        }
    }
    
    return false, nil
}

// CheckAccounts 查询多个帐号导入状态.
// 用于查询自有帐号是否已导入即时通信 IM，支持批量查询。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/38417
func (a *api) CheckAccounts(userId ...string) (results []CheckResult, err error) {
    if len(userId) > 0 {
        req := checkAccountsReq{}
        resp := &checkAccountsResp{}
        
        for _, uid := range userId {
            req.CheckItem = append(req.CheckItem, accountItem{
                UserId: uid,
            })
        }
        
        if err = a.client.Post(serviceAccount, commandCheckAccounts, req, resp); err != nil {
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
    
    if ret.Errors != nil && len(ret.Errors) > 0 {
        for _, item := range ret.Errors {
            if item.UserId == userId && item.ErrorCode != enum.SuccessCode {
                return nil, core.NewError(item.ErrorCode, "account exception")
            }
        }
    }
    
    if ret.Results != nil && len(ret.Results) > 0 {
        for _, item := range ret.Results {
            if item.UserId == userId {
                return &item, nil
            }
        }
    }
    
    return nil, nil
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
