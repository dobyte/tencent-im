/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/29 17:36
 * @Desc: Account Api Request And Response Type Definition.
 */

package account

import (
	"github.com/dobyte/tencent-im/types"
)

type (
	// AccountInfo 导入单个账号
	AccountInfo struct {
		UserId    string `json:"Identifier"` // （必填）用户名，长度不超过32字节
		Nickname  string `json:"Nick"`       // （选填）用户昵称
		AvatarUrl string `json:"FaceUrl"`    // （选填）用户头像 URL
	}

	// 批量导入账号（参数）
	importAccountsReq struct {
		UserIds []string `json:"Accounts"` // （必填）用户名，单个用户名长度不超过32字节，单次最多导入100个用户名
	}

	// 批量导入账号（响应）
	importAccountsResp struct {
		types.ActionBaseResp
		FailAccounts []string `json:"FailAccounts"` // 导入失败的帐号列表
	}

	// 账号项
	accountItem struct {
		UserId string `json:"UserID"` // 帐号的UserID
	}

	// 删除多个帐号（请求）
	deleteAccountsReq struct {
		DeleteItem []accountItem `json:"DeleteItem"` // 请求删除的帐号对象数组，单次请求最多支持100个帐号
	}

	// 删除多个账号（响应）
	deleteAccountsResp struct {
		types.ActionBaseResp
		ResultItem []DeleteResult `json:"ResultItem"`
	}

	// DeleteResult 删除多个账号结果项
	DeleteResult struct {
		ResultCode int    `json:"ResultCode"` // 单个帐号的错误码，0表示成功，非0表示失败
		ResultInfo string `json:"ResultInfo"` // 单个帐号删除失败时的错误描述信息
		UserId     string `json:"UserID"`     // 请求删除的帐号的 UserID
	}

	// 查询多个帐号（请求）
	checkAccountsReq struct {
		CheckItem []accountItem `json:"CheckItem"` // （必填）请求检查的帐号对象数组，单次请求最多支持100个帐号
	}

	// 查询多个帐号（响应）
	checkAccountsResp struct {
		types.ActionBaseResp
		Results []CheckResult `json:"ResultItem"` // 检测结果
	}

	// CheckResult 检测结果
	CheckResult struct {
		UserId        string `json:"UserID"`        // 请求检查的帐号的 UserID
		ResultCode    int    `json:"ResultCode"`    // 单个帐号的检查结果：0表示成功，非0表示失败
		ResultInfo    string `json:"ResultInfo"`    // 单个帐号检查失败时的错误描述信息
		AccountStatus string `json:"AccountStatus"` // 单个帐号的导入状态：Imported 表示已导入，NotImported 表示未导入
	}

	// 失效帐号登录状态（请求）
	kickAccountReq struct {
		UserId string `json:"Identifier"` // （必填）用户名
	}

	// 查询帐号在线状态（请求）
	queryAccountsOnlineStatusReq struct {
		UserIds      []string `json:"To_Account"`   // （必填）需要查询这些 UserID 的登录状态，一次最多查询500个 UserID 的状态
		IsNeedDetail int      `json:"IsNeedDetail"` // （选填）是否需要返回详细的登录平台信息。0表示不需要，1表示需要
	}

	// 查询帐号在线状态（响应）
	queryAccountsOnlineStatusResp struct {
		types.ActionBaseResp
		QueryResult []queryAccountsOnlineStatusResultItem `json:"QueryResult"` // 用户在线状态结构化信息
		ErrorList   []queryAccountsOnlineStatusErrorItem  `json:"ErrorList"`   // 状态查询失败的帐号列表，在此列表中的目标帐号，状态查询失败或目标帐号不存在。若状态全部查询成功，则 ErrorList 为空
	}

	// OnlineStatusRet 在线状态结果
	OnlineStatusRet struct {
		Results []queryAccountsOnlineStatusResultItem // 用户在线状态结构化信息
		Errors  []queryAccountsOnlineStatusErrorItem  // 状态查询失败的帐号列表，在此列表中的目标帐号，状态查询失败或目标帐号不存在。若状态全部查询成功，则 ErrorList 为空
	}

	// 详细的登录平台信息
	queryAccountsOnlineStatusResultDetailItem struct {
		Platform string `json:"Platform"` // 登录的平台类型。可能的返回值有："iPhone", "Android", "Web", "PC", "iPad", "Mac"。
		Status   string `json:"Status"`   // 该登录平台的状态
	}

	// 用户在线状态结构化信息项
	queryAccountsOnlineStatusResultItem struct {
		UserId string                                      `json:"To_Account"` // 用户的 UserID
		Status string                                      `json:"Status"`     // 用户状态，前台运行状态（Online）、后台运行状态（PushOnline）、未登录状态（Offline）
		Detail []queryAccountsOnlineStatusResultDetailItem `json:"Detail"`     // 详细的登录平台信息
	}

	// 状态查询失败的帐号项
	queryAccountsOnlineStatusErrorItem struct {
		UserId    string `json:"To_Account"` // 状态查询失败的目标帐号
		ErrorCode int    `json:"ErrorCode"`  // 状态查询失败的错误码，若目标帐号的错误码为70107，表示该帐号不存在
	}
)
