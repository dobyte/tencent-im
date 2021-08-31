/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/29 18:39
 * @Desc: Profile Api Request And Response Type Definition.
 */

package operation

import "github.com/dobyte/tencent-im/types"

type (
	// 拉取运营数据（请求）
	getOperationDataReq struct {
		Fields []string `json:"RequestField"`
	}
	
	// 拉取运营数据（响应）
	getOperationDataResp struct {
		types.BaseResp
		Result []OperationData `json:"Result"`
	}
	
	// 运营数据
	OperationData struct {
		AppId                string `json:"AppId"`                // 应用AppID
		AppName              string `json:"AppName"`              // 应用名称
		Company              string `json:"Company"`              // 所属客户名称
		ActiveUserNum        string `json:"ActiveUserNum"`        // 活跃用户数
		RegistUserNumOneDay  string `json:"RegistUserNumOneDay"`  // 新增注册人数
		RegistUserNumTotal   string `json:"RegistUserNumTotal"`   // 累计注册人数
		LoginTimes           string `json:"LoginTimes"`           // 登录次数
		LoginUserNum         string `json:"LoginUserNum"`         // 登录人数
		UpMsgNum             string `json:"UpMsgNum"`             // 上行消息数
		DownMsgNum           string `json:"DownMsgNum"`           // 下行消息数
		SendMsgUserNum       string `json:"SendMsgUserNum"`       // 发消息人数
		APNSMsgNum           string `json:"APNSMsgNum"`           // APNs推送数
		C2CUpMsgNum          string `json:"C2CUpMsgNum"`          // 上行消息数（C2C）
		C2CSendMsgUserNum    string `json:"C2CSendMsgUserNum"`    // 发消息人数（C2C）
		C2CAPNSMsgNum        string `json:"C2CAPNSMsgNum"`        // APNs推送数（C2C）
		C2CDownMsgNum        string `json:"C2CDownMsgNum"`        // 下行消息数（C2C）
		MaxOnlineNum         string `json:"MaxOnlineNum"`         // 最高在线人数
		ChainDecrease        string `json:"ChainDecrease"`        // 关系链对数删除量
		ChainIncrease        string `json:"ChainIncrease"`        // 关系链对数增加量
		GroupUpMsgNum        string `json:"GroupUpMsgNum"`        // 上行消息数（群）
		GroupDownMsgNum      string `json:"GroupDownMsgNum"`      // 下行消息数（群）
		GroupSendMsgUserNum  string `json:"GroupSendMsgUserNum"`  // 发消息人数（群）
		GroupAPNSMsgNum      string `json:"GroupAPNSMsgNum"`      // APNs推送数（群）
		GroupSendMsgGroupNum string `json:"GroupSendMsgGroupNum"` // 发消息群组数
		GroupJoinGroupTimes  string `json:"GroupJoinGroupTimes"`  // 入群总数
		GroupQuitGroupTimes  string `json:"GroupQuitGroupTimes"`  // 退群总数
		GroupNewGroupNum     string `json:"GroupNewGroupNum"`     // 新增群组数
		GroupAllGroupNum     string `json:"GroupAllGroupNum"`     // 累计群组数
		GroupDestroyGroupNum string `json:"GroupDestroyGroupNum"` // 解散群个数
		CallBackReq          string `json:"CallBackReq"`          // 回调请求数
		CallBackRsp          string `json:"CallBackRsp"`          // 回调应答数
		Date                 string `json:"Date"`                 // 日期
	}
	
	// 获取历史数据（请求）
	getHistoryDataReq struct {
		ChatType string `json:"ChatType"` // （必填）消息类型，C2C 表示单发消息 Group 表示群组消息
		MsgTime  string `json:"MsgTime"`  // （必填）需要下载的消息记录的时间段，2015120121表示获取2015年12月1日21:00 - 21:59的消息的下载地址。该字段需精确到小时。每次请求只能获取某天某小时的所有单发或群组消息记录
	}
	
	// 获取历史数据（响应）
	getHistoryDataResp struct {
		types.BaseResp
		File []HistoryFile `json:"File"` // 消息记录文件下载信息
	}
	
	// 历史数据文件
	HistoryFile struct {
		URL        string `json:"URL"`        // 消息记录文件下载地址
		ExpireTime string `json:"ExpireTime"` // 下载地址过期时间，请在过期前进行下载，若地址失效，请通过该接口重新获取
		FileSize   int    `json:"FileSize"`   // GZip 压缩前的文件大小（单位 Byte）
		FileMD5    string `json:"FileMD5"`    // GZip 压缩前的文件 MD5
		GzipSize   int    `json:"GzipSize"`   // GZip 压缩后的文件大小（单位 Byte）
		GzipMD5    string `json:"GzipMD5"`    // GZip 压缩后的文件 MD5
	}
	
	// 获取服务器IP地址（请求）
	getIPListReq struct {
	}
	
	// 获取服务器IP地址（响应）
	getIPListResp struct {
		types.BaseResp
		IPList []string `json:"IPList"` // 服务器IP列表
	}
)
