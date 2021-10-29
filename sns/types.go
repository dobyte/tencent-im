/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/29 17:34
 * @Desc: 关系链管理数据类型
 */

package sns

import "github.com/dobyte/tencent-im/internal/types"

type (
	// 添加的好友信息
	addFriendItem struct {
		UserId     string `json:"To_Account"`           // （必填）好友的UserID
		AddSource  string `json:"AddSource"`            // （必填）加好友来源
		Remark     string `json:"Remark,omitempty"`     // （选填）加好友备注
		GroupName  string `json:"GroupName,omitempty"`  // （选填）分组信息，添加好友时只允许设置一个分组，因此使用 String 类型即可
		AddWording string `json:"AddWording,omitempty"` // （选填）好友关系时的附言信息
	}

	// 添加好友（请求）
	addFriendsReq struct {
		UserId        string           `json:"From_Account"`  // （必填）需要为该UserID添加好友
		Friends       []*addFriendItem `json:"AddFriendItem"` // （必填）好友结构体对象
		AddType       AddType          `json:"AddType"`       // （选填）加好友方式（默认双向加好友方式）Add_Type_Single：表示单向加好友；Add_Type_Both：表示双向加好友
		ForceAddFlags ForceAddType     `json:"ForceAddFlags"` // （选填）管理员强制加好友标记：1表示强制加好友，0表示常规加好友方式
	}

	// 添加好友（响应）
	addFriendsResp struct {
		types.ActionBaseResp
		Results []*Result `json:"ResultItem"` // 批量加好友的结果对象数组
	}

	// Result 添加结果
	Result struct {
		UserId     string `json:"To_Account"` // 请求添加的好友的UserID
		ResultCode int    `json:"ResultCode"` // 处理结果，0表示成功，非0表示失败
		ResultInfo string `json:"ResultInfo"` // 错误描述信息
	}

	// 导入好友（请求）
	importFriendsReq struct {
		UserId  string              `json:"From_Account"`  // 需要为该UserID添加好友
		Friends []*importFriendItem `json:"AddFriendItem"` // 好友结构体对象
	}

	// 导入好友（响应）
	importFriendsResp struct {
		types.ActionBaseResp
		Results     []*Result `json:"ResultItem"`   // 结果对象数组
		FailUserIds []string  `json:"Fail_Account"` // 返回处理失败的用户列表，仅当存在失败用户时才返回该字段
	}

	// 导入好友
	importFriendItem struct {
		UserId     string           `json:"To_Account"`           // （必填）好友的UserID
		AddSource  string           `json:"AddSource"`            // （必填）加好友来源
		Remark     string           `json:"Remark,omitempty"`     // （选填）加好友备注
		GroupName  []string         `json:"GroupName,omitempty"`  // （选填）分组信息
		AddWording string           `json:"AddWording,omitempty"` // （选填）好友关系时的附言信息
		AddTime    int64            `json:"AddTime,omitempty"`    // （选填）形成好友关系的时间
		RemarkTime int64            `json:"RemarkTime,omitempty"` // （选填）好友备注时间
		CustomData []*types.TagPair `json:"CustomItem,omitempty"` // （选填）自定义好友数据
	}

	// 更新好友（请求）
	updateFriendsReq struct {
		UserId  string              `json:"From_Account"` // （必填）需要更新该UserID的关系链数据
		Friends []*updateFriendItem `json:"UpdateItem"`   // （必填）需要更新的好友对象数组
	}

	// 更新好友（响应）
	updateFriendsResp struct {
		types.ActionBaseResp
		Results     []*Result `json:"ResultItem"`   // 结果对象数组
		FailUserIds []string  `json:"Fail_Account"` // 返回处理失败的用户列表，仅当存在失败用户时才返回该字段
	}

	// 更新好友
	updateFriendItem struct {
		UserId string           `json:"To_Account"` // （必填）好友的UserID
		Attrs  []*types.TagPair `json:"SnsItem"`    // （必填）需要更新的关系链数据对象数组
	}

	// 删除好友（请求）
	deleteFriendsReq struct {
		UserId         string     `json:"From_Account"` // （必填）需要删除该UserID的好友
		DeletedUserIds []string   `json:"To_Account"`   // （必填）待删除的好友的 UserID 列表，单次请求的 To_Account 数不得超过1000
		DeleteType     DeleteType `json:"DeleteType"`   // （选填）删除模式
	}

	// 删除好友（响应）
	deleteFriendsResp struct {
		types.ActionBaseResp
		ErrorDisplay string    `json:"ErrorDisplay"` // 详细的客户端展示信息
		Results      []*Result `json:"ResultItem"`   // 返回处理失败的用户列表，仅当存在失败用户时才返回该字段
	}

	// 删除所有好友（请求）
	deleteAllFriendsReq struct {
		UserId     string     `json:"From_Account"` // （必填）需要删除该UserID的好友
		DeleteType DeleteType `json:"DeleteType"`   // （选填）删除类型
	}

	// 校验好友（请求）
	checkFriendsReq struct {
		UserId         string    `json:"From_Account"` // （必填）需要校验该 UserID 的好友
		CheckedUserIds []string  `json:"To_Account"`   // （必填）请求校验的好友的 UserID 列表，单次请求的 To_Account 数不得超过1000
		CheckType      CheckType `json:"CheckType"`    // （必填）校验模式
	}

	// 校验好友（响应）
	checkFriendsResp struct {
		types.ActionBaseResp
		Results     []*CheckResult `json:"InfoItem"`     // 结果对象数组
		FailUserIds []string       `json:"Fail_Account"` // 返回处理失败的用户列表，仅当存在失败用户时才返回该字段
	}

	// CheckResult 校验结果
	CheckResult struct {
		UserId     string `json:"To_Account"` // 请求校验的用户的 UserID
		Relation   string `json:"Relation"`   // 校验成功时 To_Account 与 From_Account 之间的好友关系
		ResultCode int    `json:"ResultCode"` // 处理结果，0表示成功，非0表示失败
		ResultInfo string `json:"ResultInfo"` // 描述信息，成功时该字段为空
	}

	// 拉取指定好友（请求）
	getFriendsReq struct {
		UserId        string   `json:"From_Account"` // （必填）指定要拉取好友数据的用户的 UserID
		FriendUserIds []string `json:"To_Account"`   // （必填）好友的 UserID 列表，建议每次请求的好友数不超过100，避免因数据量太大导致回包失败
		TagList       []string `json:"TagList"`      // （必填）指定要拉取的资料字段及好友字段
	}

	// 拉取指定好友（响应）
	getFriendsResp struct {
		types.ActionBaseResp
		FailUserIds []string     `json:"Fail_Account"` // 返回处理失败的用户列表，仅当存在失败用户时才返回该字段
		Friends     []friendData `json:"InfoItem"`     // 好友对象数组
	}

	// 拉取好友（请求）
	fetchFriendsReq struct {
		UserId           string `json:"From_Account"`     // （必填）需要拉取该 UserID 的黑名单
		StartIndex       int    `json:"StartIndex"`       // （必填）拉取的起始位置
		StandardSequence int    `json:"StandardSequence"` // （选填）上次拉好友数据时返回的 StandardSequence，如果 StandardSequence 字段的值与后台一致，后台不会返回标配好友数据
		CustomSequence   int    `json:"CustomSequence"`   // （选填）上次拉好友数据时返回的 CustomSequence，如果 CustomSequence 字段的值与后台一致，后台不会返回自定义好友数据
	}

	// 拉取好友（响应）
	fetchFriendsResp struct {
		types.ActionBaseResp
		Friends          []friendData `json:"UserDataItem"`     // 好友对象数组
		StandardSequence int          `json:"StandardSequence"` // 标配好友数据的 Sequence，客户端可以保存该 Sequence，下次请求时通过请求的 StandardSequence 字段返回给后台
		CustomSequence   int          `json:"CustomSequence"`   // 自定义好友数据的 Sequence，客户端可以保存该 Sequence，下次请求时通过请求的 CustomSequence 字段返回给后台
		FriendNum        int          `json:"FriendNum"`        // 好友总数
		CompleteFlag     int          `json:"CompleteFlag"`     // 分页的结束标识，非0值表示已完成全量拉取
		NextStartIndex   int          `json:"NextStartIndex"`   // 分页接口下一页的起始位置
	}

	// 好友信息
	friendData struct {
		UserId     string          `json:"To_Account"`     // 好友的 UserID
		Values     []types.TagPair `json:"ValueItem"`      // 好友数据的数组
		Profiles   []types.TagPair `json:"SnsProfileItem"` // 好友数据的数组
		ResultCode int             `json:"ResultCode"`     // 处理结果，0表示成功，非0表示失败
		ResultInfo string          `json:"ResultInfo"`     // 错误描述信息
	}

	// FetchFriendsRet 好友列表
	FetchFriendsRet struct {
		StandardSequence int       // 标准排序
		CustomSequence   int       // 自定义排序
		FriendNum        int       // 好友总数
		IsOver           bool      // 是否没有数据了
		NextStartIndex   int       // 分页接口下一页的起始位置
		Friends          []*Friend // 好友列表
	}

	// 添加黑名单（请求）
	addBlacklistReq struct {
		UserId         string   `json:"From_Account"` // （必填）请求为该 UserID 添加黑名单
		BlackedUserIds []string `json:"To_Account"`   // （必填）待添加为黑名单的用户 UserID 列表，单次请求的 To_Account 数不得超过1000
	}

	// 添加黑名单（响应）
	addBlacklistResp struct {
		types.ActionBaseResp
		FailUserIds []string  `json:"Fail_Account"` // 返回处理失败的用户列表，仅当存在失败用户时才返回该字段
		Results     []*Result `json:"ResultItem"`   // 批量添加黑名单的结果对象数组
	}

	// 删除黑名单（请求）
	deleteBlacklistReq struct {
		UserId         string   `json:"From_Account"` // （必填）需要删除该 UserID 的黑名单
		DeletedUserIds []string `json:"To_Account"`   // （必填）待删除的黑名单的 UserID 列表，单次请求的 To_Account 数不得超过1000
	}

	// 删除黑名单（响应）
	deleteBlacklistResp struct {
		types.ActionBaseResp
		FailUserIds []string  `json:"Fail_Account"` // 返回处理失败的用户列表，仅当存在失败用户时才返回该字段
		Results     []*Result `json:"ResultItem"`   // 批量添加黑名单的结果对象数组
	}

	// 拉取黑名单（请求）
	fetchBlacklistReq struct {
		UserId       string `json:"From_Account"` // （必填）需要拉取该 UserID 的黑名单
		StartIndex   int    `json:"StartIndex"`   // （必填）拉取的起始位置
		MaxLimited   int    `json:"MaxLimited"`   // （必填）每页最多拉取的黑名单数
		LastSequence int    `json:"LastSequence"` // （必填）上一次拉黑名单时后台返回给客户端的 Seq，初次拉取时为0
	}

	// 拉取黑名单（响应）
	fetchBlacklistResp struct {
		types.ActionBaseResp
		ErrorDisplay    string      `json:"ErrorDisplay"`    // 详细的客户端展示信息
		StartIndex      int         `json:"StartIndex"`      // 下页拉取的起始位置，0表示已拉完
		CurrentSequence int         `json:"CurrentSequence"` // 黑名单最新的 Seq
		Blacklists      []Blacklist `json:"BlackListItem"`   // 黑名单对象数组
	}

	// FetchBlacklistRet 拉取黑名单结果
	FetchBlacklistRet struct {
		StandardSequence int         // 标准排序
		IsOver           bool        // 是否没有数据了
		NextStartIndex   int         // 分页接口下一页的起始位置
		Blacklists       []Blacklist // 黑名单列表
	}

	// Blacklist 黑名单
	Blacklist struct {
		UserId string `json:"To_Account"`        // 黑名单的 UserID
		Time   int    `json:"AddBlackTimeStamp"` // 添加黑名单的时间
	}

	// 校验黑名单（请求）
	checkBlacklistReq struct {
		UserId         string             `json:"From_Account"` // （必填）需要校验该 UserID 的黑名单
		CheckedUserIds []string           `json:"To_Account"`   // （必填）待校验的黑名单的 UserID 列表，单次请求的 To_Account 数不得超过1000
		CheckType      BlacklistCheckType `json:"CheckType"`    // （必填）校验模式
	}

	// 校验黑名单（响应）
	checkBlacklistResp struct {
		types.ActionBaseResp
		FailUserIds []string       `json:"Fail_Account"`       // 返回处理失败的用户列表，仅当存在失败用户时才返回该字段
		Results     []*CheckResult `json:"BlackListCheckItem"` // 校验结果对象数组
	}

	// 添加分组（请求）
	addGroupsReq struct {
		UserId        string   `json:"From_Account"` // （必填）需要为该 UserID 添加新分组
		GroupNames    []string `json:"GroupName"`    // （必填）新增分组列表
		JoinedUserIds []string `json:"To_Account"`   // （必填）需要加入新增分组的好友的 UserID 列表
	}

	// 添加分组（响应）
	addGroupsResp struct {
		types.ActionBaseResp
		ErrorDisplay    string    `json:"ErrorDisplay"`    // 详细的客户端展示信息
		FailUserIds     []string  `json:"Fail_Account"`    // 返回处理失败的用户列表，仅当存在失败用户时才返回该字段
		CurrentSequence int       `json:"CurrentSequence"` // 返回最新的分组 Sequence
		Results         []*Result `json:"ResultItem"`      // 好友加入新增分组的结果对象数组
	}

	// 删除分组（请求）
	deleteGroupsReq struct {
		UserId     string   `json:"From_Account"` // （必填）需要删除该 UserID 的分组
		GroupNames []string `json:"GroupName"`    // （必填）要删除的分组列表
	}

	// 删除分组（响应）
	deleteGroupsResp struct {
		types.ActionBaseResp
		CurrentSequence int `json:"CurrentSequence"` // 返回最新的分组 Sequence
	}

	// 拉取分组（请求）
	getGroupsReq struct {
		UserId       string         `json:"From_Account"` // （必填）指定要拉取分组的用户的 UserID
		LastSequence int            `json:"LastSequence"` // （必填）上一次拉取分组时后台返回给客户端的 Seq，初次拉取时为0，只有 GroupName 为空时有效
		NeedFriend   NeedFriendType `json:"NeedFriend"`   // （选填）是否需要拉取分组下的 User 列表
		GroupNames   []string       `json:"GroupName"`    // （选填）要拉取的分组名称
	}

	// 拉取分组（响应）
	getGroupsResp struct {
		types.ActionBaseResp
		CurrentSequence int            `json:"CurrentSequence"` // 返回最新的分组 Sequence
		Results         []*GroupResult `json:"ResultItem"`      // 拉取分组的结果对象数组
	}

	// GroupResult 分组结果
	GroupResult struct {
		GroupName    string   `json:"GroupName"`    // 分组名
		FriendNumber int      `json:"FriendNumber"` // 该分组下的好友数量
		UserIds      []string `json:"To_Account"`   // 该分组下的好友的 UserID
	}
)
