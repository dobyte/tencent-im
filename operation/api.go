/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/29 18:38
 * @Desc: 运营管理
 */

package operation

import (
	"time"
	
	"github.com/dobyte/tencent-im/internal/core"
)

const (
	serviceOperation     = "openconfigsvr"
	serviceOpenMessage   = "open_msg_svc"
	serviceConfig        = "ConfigSvc"
	commandGetAppInfo    = "getappinfo"
	commandGetHistory    = "get_history"
	commandGetIPList     = "GetIPList"
)

type ChatType string

const (
	// 聊天类型
	ChatTypePrivate = "C2C"   // 单聊消息
	ChatTypeGroup   = "Group" // 群聊消息
)

type API interface {
	// GetOperationData 拉取运营数据
	// App 管理员可以通过该接口拉取最近30天的运营数据，可拉取的字段见下文可拉取的运营字段。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/4193
	GetOperationData(fields ...[]string) (data []OperationData, err error)
	
	// GetHistoryData 下载最近消息记录
	// App 管理员可以通过该接口获取 App 中最近7天中某天某小时的所有单发或群组消息记录的下载地址
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1650
	GetHistoryData(chatType ChatType, msgTime time.Time) (files []HistoryFile, err error)
	
	// GetIPList 获取服务器IP地址
	// 基于安全等考虑，您可能需要获知服务器的 IP 地址列表，以便进行相关限制。
	// App 管理员可以通过该接口获得 SDK、第三方回调所使用到的服务器 IP 地址列表或 IP 网段信息。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45438
	GetIPList() (ipList []string, err error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// GetOperationData 拉取运营数据
// App 管理员可以通过该接口拉取最近30天的运营数据，可拉取的字段见下文可拉取的运营字段。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/4193
func (a *api) GetOperationData(fields ...[]string) (data []OperationData, err error) {
	req := getOperationDataReq{}
	resp := &getOperationDataResp{}
	
	if len(fields) > 0 {
		req.Fields = fields[0]
	}
	
	if err = a.client.Post(serviceOperation, commandGetAppInfo, req, resp); err != nil {
		return
	} else {
		data = resp.Result
	}
	
	return
}

// GetHistoryData 下载最近消息记录
// App 管理员可以通过该接口获取 App 中最近7天中某天某小时的所有单发或群组消息记录的下载地址
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1650
func (a *api) GetHistoryData(chatType ChatType, msgTime time.Time) (files []HistoryFile, err error) {
	req := getHistoryDataReq{
		ChatType: string(chatType),
		MsgTime:  msgTime.Format("2006010215"),
	}
	resp := &getHistoryDataResp{}
	
	if err = a.client.Post(serviceOpenMessage, commandGetHistory, req, resp); err != nil {
		return
	} else {
		files = resp.File
	}
	
	return
}

// GetIPList 获取服务器IP地址
// 基于安全等考虑，您可能需要获知服务器的 IP 地址列表，以便进行相关限制。
// App 管理员可以通过该接口获得 SDK、第三方回调所使用到的服务器 IP 地址列表或 IP 网段信息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45438
func (a *api) GetIPList() (ipList []string, err error) {
	req := getIPListReq{}
	resp := &getIPListResp{}
	
	if err = a.client.Post(serviceConfig, commandGetIPList, req, resp); err != nil {
		return
	} else {
		ipList = resp.IPList
	}
	
	return
}
