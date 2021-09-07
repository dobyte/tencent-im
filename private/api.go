/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 19:32
 * @Desc: 单聊消息
 */

package private

import (
    "github.com/dobyte/tencent-im/internal/conv"
    "github.com/dobyte/tencent-im/internal/core"
    "github.com/dobyte/tencent-im/internal/types"
)

const (
    serviceMessage             = "openim"
    commandSendMessage         = "sendmsg"
    commandSendMessages        = "batchsendmsg"
    commandImportMessage       = "importmsg"
    commandFetchMessages       = "admin_getroammsg"
    commandRevokeMessage       = "admin_msgwithdraw"
    commandSetMessageRead      = "admin_set_msg_read"
    commandGetUnreadMessageNum = "get_c2c_unread_msg_num"
)

type API interface {
    // SendMessage 单发单聊消息
    // 管理员向帐号发消息，接收方看到消息发送者是管理员。
    // 管理员指定某一帐号向其他帐号发消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
    // 该接口不会检查发送者和接收者的好友关系（包括黑名单），同时不会检查接收者是否被禁言。
    // 单聊消息 MsgSeq 字段的作用及说明：该字段在发送消息时由用户自行指定，该值可以重复，非后台生成，非全局唯一。与群聊消息的 MsgSeq 字段不同，群聊消息的 MsgSeq 由后台生成，每个群都维护一个 MsgSeq，从1开始严格递增。单聊消息历史记录对同一个会话的消息先以时间戳排序，同秒内的消息再以 MsgSeq 排序。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/2282
    SendMessage(message *Message) (ret *SendMessageRet, err error)
    
    // SendMessages 批量发单聊消息
    // 支持一次对最多500个用户进行单发消息。
    // 与单发消息相比，该接口更适用于营销类消息、系统通知 tips 等时效性较强的消息。
    // 管理员指定某一帐号向目标帐号批量发消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
    // 该接口不触发回调请求。
    // 该接口不会检查发送者和接收者的好友关系（包括黑名单），同时不会检查接收者是否被禁言。
    // 单聊消息 MsgSeq 字段的作用及说明：该字段在发送消息时由用户自行指定，该值可以重复，非后台生成，非全局唯一。与群聊消息的 MsgSeq 字段不同，群聊消息的 MsgSeq 由后台生成，每个群都维护一个 MsgSeq，从1开始严格递增。单聊消息历史记录对同一个会话的消息先以时间戳排序，同秒内的消息再以 MsgSeq 排序。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/1612
    SendMessages(message *Message) (ret *SendMessagesRet, err error)
    
    // ImportMessage 导入单聊消息
    // 导入历史单聊消息到即时通信 IM。
    // 平滑过渡期间，将原有即时通信实时单聊消息导入到即时通信 IM。
    // 该接口不会触发回调。
    // 该接口会根据 From_Account ， To_Account ，MsgSeq ， MsgRandom ， MsgTimeStamp 字段的值对导入的消息进行去重。仅当这五个字段的值都对应相同时，才判定消息是重复的，消息是否重复与消息内容本身无关。
    // 重复导入的消息不会覆盖之前已导入的消息（即消息内容以首次导入的为准）。
    // 单聊消息 MsgSeq 字段的作用及说明：该字段在发送消息时由用户自行指定，该值可以重复，非后台生成，非全局唯一。与群聊消息的 MsgSeq 字段不同，群聊消息的 MsgSeq 由后台生成，每个群都维护一个 MsgSeq，从1开始严格递增。单聊消息历史记录对同一个会话的消息先以时间戳排序，同秒内的消息再以 MsgSeq 排序。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/2568
    ImportMessage(message *Message) (err error)
    
    // FetchMessages 查询单聊消息
    // 管理员按照时间范围查询某单聊会话的消息记录。
    // 查询的单聊会话由请求中的 From_Account 和 To_Account 指定。查询结果包含会话双方互相发送的消息，具体每条消息的发送方和接收方由每条消息里的 From_Account 和 To_Account 指定。
    // 一般情况下，请求中的 From_Account 和 To_Account 字段值互换，查询结果不变。但通过 单发单聊消息 或 批量发单聊消息 接口发送的消息，如果指定 SyncOtherMachine 值为2，则需要指定正确的 From_Account 和 To_Account 字段值才能查询到该消息。
    // 例如，通过 单发单聊消息 接口指定帐号 A 给帐号 B 发一条消息，同时指定 SyncOtherMachine 值为2。则调用本接口时，From_Account 必须设置为帐号 B，To_Account 必须设置为帐号 A 才能查询到该消息。
    // 查询结果包含被撤回的消息，由消息里的 MsgFlagBits 字段标识。
    // 若想通过 REST API 撤回单聊消息 接口撤回某条消息，可先用本接口查询出该消息的 MsgKey，然后再调用撤回接口进行撤回。
    // 可查询的消息记录的时间范围取决于漫游消息存储时长，默认是7天。支持在控制台修改消息漫游时长，延长消息漫游时长是增值服务。具体请参考 漫游消息存储。
    // 若请求时间段内的消息总大小超过应答包体大小限制（目前为13K）时，则需要续拉。您可以通过应答中的 Complete 字段查看是否已拉取请求的全部消息。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/42794
    FetchMessages(arg FetchMessagesArg) (ret *FetchMessagesRet, err error)
    
    // PullMessages 续拉取单聊消息
    // 本API是借助"查询单聊消息"API进行扩展实现
    // 管理员按照时间范围查询某单聊会话的全部消息记录
    // 查询的单聊会话由请求中的 From_Account 和 To_Account 指定。查询结果包含会话双方互相发送的消息，具体每条消息的发送方和接收方由每条消息里的 From_Account 和 To_Account 指定。
    // 一般情况下，请求中的 From_Account 和 To_Account 字段值互换，查询结果不变。但通过 单发单聊消息 或 批量发单聊消息 接口发送的消息，如果指定 SyncOtherMachine 值为2，则需要指定正确的 From_Account 和 To_Account 字段值才能查询到该消息。
    // 例如，通过 单发单聊消息 接口指定帐号 A 给帐号 B 发一条消息，同时指定 SyncOtherMachine 值为2。则调用本接口时，From_Account 必须设置为帐号 B，To_Account 必须设置为帐号 A 才能查询到该消息。
    // 查询结果包含被撤回的消息，由消息里的 MsgFlagBits 字段标识。
    // 若想通过 REST API 撤回单聊消息 接口撤回某条消息，可先用本接口查询出该消息的 MsgKey，然后再调用撤回接口进行撤回。
    // 可查询的消息记录的时间范围取决于漫游消息存储时长，默认是7天。支持在控制台修改消息漫游时长，延长消息漫游时长是增值服务。具体请参考 漫游消息存储。
    // 若请求时间段内的消息总大小超过应答包体大小限制（目前为13K）时，则需要续拉。您可以通过应答中的 Complete 字段查看是否已拉取请求的全部消息。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/42794
    PullMessages(arg PullMessagesArg, fn func(ret *FetchMessagesRet)) error
    
    // RevokeMessage 撤回单聊消息
    // 管理员撤回单聊消息。
    // 该接口可以撤回所有单聊消息，包括客户端发出的单聊消息，由 REST API 单发 和 批量发 接口发出的单聊消息。
    // 若需要撤回由客户端发出的单聊消息，您可以开通 发单聊消息之前回调 或 发单聊消息之后回调 ，通过该回调接口记录每条单聊消息的 MsgKey ，然后填在本接口的 MsgKey 字段进行撤回。您也可以通过 查询单聊消息 查询出待撤回的单聊消息的 MsgKey 后，填在本接口的 MsgKey 字段进行撤回。
    // 若需要撤回由 REST API 单发 和 批量发 接口发出的单聊消息，需要记录这些接口回包里的 MsgKey 字段以进行撤回。
    // 调用该接口撤回消息后，该条消息的离线、漫游存储，以及消息发送方和接收方的客户端的本地缓存都会被撤回。
    // 该接口可撤回的单聊消息没有时间限制，即可以撤回任何时间的单聊消息。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/38980
    RevokeMessage(fromUserId, toUserId, msgKey string) (err error)
    
    // SetMessageRead 设置单聊消息已读
    // 设置用户的某个单聊会话的消息全部已读。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/50349
    SetMessageRead(userId, peerUserId string) (err error)
    
    // GetUnreadMessageNum 查询单聊未读消息计数
    // App 后台可以通过该接口查询特定账号的单聊总未读数（包含所有的单聊会话）或者单个单聊会话的未读数。
    // 点击查看详细文档:
    // https://cloud.tencent.com/document/product/269/56043
    GetUnreadMessageNum(userId string, peerUserIds ...[]string) (ret *UnreadMessageRet, err error)
}

type api struct {
    client core.Client
}

func NewAPI(client core.Client) API {
    return &api{client: client}
}

// SendMessage 单发单聊消息
// 管理员向帐号发消息，接收方看到消息发送者是管理员。
// 管理员指定某一帐号向其他帐号发消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
// 该接口不会检查发送者和接收者的好友关系（包括黑名单），同时不会检查接收者是否被禁言。
// 单聊消息 MsgSeq 字段的作用及说明：该字段在发送消息时由用户自行指定，该值可以重复，非后台生成，非全局唯一。与群聊消息的 MsgSeq 字段不同，群聊消息的 MsgSeq 由后台生成，每个群都维护一个 MsgSeq，从1开始严格递增。单聊消息历史记录对同一个会话的消息先以时间戳排序，同秒内的消息再以 MsgSeq 排序。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2282
func (a *api) SendMessage(message *Message) (ret *SendMessageRet, err error) {
    if err = message.CheckError(); err != nil {
        return
    }
    
    req := sendMessageReq{}
    req.FromUserId = message.GetSender()
    req.ToUserId = message.GetLastReceiver()
    req.MsgLifeTime = message.GetLifeTime()
    req.MsgTimeStamp = message.GetTimestamp()
    req.OfflinePushInfo = message.GetOfflinePushInfo()
    req.CustomData = conv.String(message.customData)
    req.MsgSeq = message.GetSerialNo()
    req.MsgBody = message.GetBody()
    req.MsgRandom = message.GetRandom()
    req.SendMsgControl = message.GetSendMsgControl()
    req.ForbidCallbackControl = message.GetForbidCallbackControl()
    req.SyncOtherMachine = message.GetSyncOtherMachine()
    
    resp := &sendMessageResp{}
    
    if err = a.client.Post(serviceMessage, commandSendMessage, req, resp); err != nil {
        return
    } else {
        ret = &SendMessageRet{
            MsgKey:  resp.MsgKey,
            MsgTime: resp.MsgTime,
        }
    }
    
    return
}

// SendMessages 批量发单聊消息
// 支持一次对最多500个用户进行单发消息。
// 与单发消息相比，该接口更适用于营销类消息、系统通知 tips 等时效性较强的消息。
// 管理员指定某一帐号向目标帐号批量发消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
// 该接口不触发回调请求。
// 该接口不会检查发送者和接收者的好友关系（包括黑名单），同时不会检查接收者是否被禁言。
// 单聊消息 MsgSeq 字段的作用及说明：该字段在发送消息时由用户自行指定，该值可以重复，非后台生成，非全局唯一。与群聊消息的 MsgSeq 字段不同，群聊消息的 MsgSeq 由后台生成，每个群都维护一个 MsgSeq，从1开始严格递增。单聊消息历史记录对同一个会话的消息先以时间戳排序，同秒内的消息再以 MsgSeq 排序。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1612
func (a *api) SendMessages(message *Message) (ret *SendMessagesRet, err error) {
    if err = message.CheckError(); err != nil {
        return
    }
    
    req := sendMessagesReq{}
    req.FromUserId = message.GetSender()
    req.ToUserIds = message.GetReceivers()
    req.OfflinePushInfo = message.GetOfflinePushInfo()
    req.CustomData = message.GetCustomData()
    req.MsgSeq = message.GetSerialNo()
    req.MsgBody = message.GetBody()
    req.MsgRandom = message.GetRandom()
    req.SendMsgControl = message.GetSendMsgControl()
    req.SyncOtherMachine = message.GetSyncOtherMachine()
    
    resp := &sendMessagesResp{}
    
    if err = a.client.Post(serviceMessage, commandSendMessages, req, resp); err != nil {
        return
    } else {
        ret = &SendMessagesRet{
            MsgKey: resp.MsgKey,
            Errors: resp.Errors,
        }
    }
    
    return
}

// ImportMessage 导入单聊消息
// 导入历史单聊消息到即时通信 IM。
// 平滑过渡期间，将原有即时通信实时单聊消息导入到即时通信 IM。
// 该接口不会触发回调。
// 该接口会根据 From_Account ， To_Account ，MsgSeq ， MsgRandom ， MsgTimeStamp 字段的值对导入的消息进行去重。仅当这五个字段的值都对应相同时，才判定消息是重复的，消息是否重复与消息内容本身无关。
// 重复导入的消息不会覆盖之前已导入的消息（即消息内容以首次导入的为准）。
// 单聊消息 MsgSeq 字段的作用及说明：该字段在发送消息时由用户自行指定，该值可以重复，非后台生成，非全局唯一。与群聊消息的 MsgSeq 字段不同，群聊消息的 MsgSeq 由后台生成，每个群都维护一个 MsgSeq，从1开始严格递增。单聊消息历史记录对同一个会话的消息先以时间戳排序，同秒内的消息再以 MsgSeq 排序。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/2568
func (a *api) ImportMessage(message *Message) (err error) {
    if err = message.CheckError(); err != nil {
        return
    }
    
    req := importMessageReq{}
    req.FromUserId = message.GetSender()
    req.ToUserId = message.GetLastReceiver()
    req.MsgTimeStamp = message.GetTimestamp()
    req.CustomData = message.GetCustomData()
    req.MsgSeq = message.GetSerialNo()
    req.MsgBody = message.GetBody()
    req.MsgRandom = message.GetRandom()
    req.SyncFromOldSystem = message.GetSyncOtherMachine()
    
    if err = a.client.Post(serviceMessage, commandImportMessage, req, &types.ActionBaseResp{}); err != nil {
        return
    }
    
    return
}

// FetchMessages 查询单聊消息
// 管理员按照时间范围查询某单聊会话的消息记录。
// 查询的单聊会话由请求中的 From_Account 和 To_Account 指定。查询结果包含会话双方互相发送的消息，具体每条消息的发送方和接收方由每条消息里的 From_Account 和 To_Account 指定。
// 一般情况下，请求中的 From_Account 和 To_Account 字段值互换，查询结果不变。但通过 单发单聊消息 或 批量发单聊消息 接口发送的消息，如果指定 SyncOtherMachine 值为2，则需要指定正确的 From_Account 和 To_Account 字段值才能查询到该消息。
// 例如，通过 单发单聊消息 接口指定帐号 A 给帐号 B 发一条消息，同时指定 SyncOtherMachine 值为2。则调用本接口时，From_Account 必须设置为帐号 B，To_Account 必须设置为帐号 A 才能查询到该消息。
// 查询结果包含被撤回的消息，由消息里的 MsgFlagBits 字段标识。
// 若想通过 REST API 撤回单聊消息 接口撤回某条消息，可先用本接口查询出该消息的 MsgKey，然后再调用撤回接口进行撤回。
// 可查询的消息记录的时间范围取决于漫游消息存储时长，默认是7天。支持在控制台修改消息漫游时长，延长消息漫游时长是增值服务。具体请参考 漫游消息存储。
// 若请求时间段内的消息总大小超过应答包体大小限制（目前为13K）时，则需要续拉。您可以通过应答中的 Complete 字段查看是否已拉取请求的全部消息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/42794
func (a *api) FetchMessages(arg FetchMessagesArg) (ret *FetchMessagesRet, err error) {
    resp := &fetchMessagesResp{}
    
    if err = a.client.Post(serviceMessage, commandFetchMessages, arg, resp); err != nil {
        return
    } else {
        ret = &FetchMessagesRet{
            LastMsgKey:  resp.LastMsgKey,
            LastMsgTime: resp.LastMsgTime,
            MsgCount:    resp.MsgCount,
            MsgList:     resp.MsgList,
        }
        
        if resp.Complete == 1 {
            ret.IsOver = true
        }
    }
    
    return
}

// PullMessages 续拉取单聊消息
// 本API是借助"查询单聊消息"API进行扩展实现
// 管理员按照时间范围查询某单聊会话的全部消息记录
// 查询的单聊会话由请求中的 From_Account 和 To_Account 指定。查询结果包含会话双方互相发送的消息，具体每条消息的发送方和接收方由每条消息里的 From_Account 和 To_Account 指定。
// 一般情况下，请求中的 From_Account 和 To_Account 字段值互换，查询结果不变。但通过 单发单聊消息 或 批量发单聊消息 接口发送的消息，如果指定 SyncOtherMachine 值为2，则需要指定正确的 From_Account 和 To_Account 字段值才能查询到该消息。
// 例如，通过 单发单聊消息 接口指定帐号 A 给帐号 B 发一条消息，同时指定 SyncOtherMachine 值为2。则调用本接口时，From_Account 必须设置为帐号 B，To_Account 必须设置为帐号 A 才能查询到该消息。
// 查询结果包含被撤回的消息，由消息里的 MsgFlagBits 字段标识。
// 若想通过 REST API 撤回单聊消息 接口撤回某条消息，可先用本接口查询出该消息的 MsgKey，然后再调用撤回接口进行撤回。
// 可查询的消息记录的时间范围取决于漫游消息存储时长，默认是7天。支持在控制台修改消息漫游时长，延长消息漫游时长是增值服务。具体请参考 漫游消息存储。
// 若请求时间段内的消息总大小超过应答包体大小限制（目前为13K）时，则需要续拉。您可以通过应答中的 Complete 字段查看是否已拉取请求的全部消息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/42794
func (a *api) PullMessages(arg PullMessagesArg, fn func(ret *FetchMessagesRet)) error {
    var (
        err error
        ret *FetchMessagesRet
        req = FetchMessagesArg{
            FromUserId: arg.FromUserId,
            ToUserId:   arg.ToUserId,
            MaxLimited: arg.MaxLimited,
            MinTime:    arg.MinTime,
            MaxTime:    arg.MaxTime,
        }
    )
    
    for ret == nil || !ret.IsOver {
        ret, err = a.FetchMessages(req)
        if err != nil {
            return err
        }
        
        go fn(ret)
        
        if !ret.IsOver {
            req.LastMsgKey = ret.LastMsgKey
            req.MaxTime = ret.LastMsgTime
        }
    }
    
    return nil
}

// RevokeMessage 撤回单聊消息
// 管理员撤回单聊消息。
// 该接口可以撤回所有单聊消息，包括客户端发出的单聊消息，由 REST API 单发 和 批量发 接口发出的单聊消息。
// 若需要撤回由客户端发出的单聊消息，您可以开通 发单聊消息之前回调 或 发单聊消息之后回调 ，通过该回调接口记录每条单聊消息的 MsgKey ，然后填在本接口的 MsgKey 字段进行撤回。您也可以通过 查询单聊消息 查询出待撤回的单聊消息的 MsgKey 后，填在本接口的 MsgKey 字段进行撤回。
// 若需要撤回由 REST API 单发 和 批量发 接口发出的单聊消息，需要记录这些接口回包里的 MsgKey 字段以进行撤回。
// 调用该接口撤回消息后，该条消息的离线、漫游存储，以及消息发送方和接收方的客户端的本地缓存都会被撤回。
// 该接口可撤回的单聊消息没有时间限制，即可以撤回任何时间的单聊消息。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/38980
func (a *api) RevokeMessage(fromUserId, toUserId, msgKey string) (err error) {
    req := revokeMessageReq{FromUserId: fromUserId, ToUserId: toUserId, MsgKey: msgKey}
    
    if err = a.client.Post(serviceMessage, commandRevokeMessage, req, &types.ActionBaseResp{}); err != nil {
        return
    }
    
    return
}

// SetMessageRead 设置单聊消息已读
// 设置用户的某个单聊会话的消息全部已读。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/50349
func (a *api) SetMessageRead(userId, peerUserId string) (err error) {
    req := setMessageReadReq{UserId: userId, PeerUserId: peerUserId}
    
    if err = a.client.Post(serviceMessage, commandSetMessageRead, req, &types.ActionBaseResp{}); err != nil {
        return
    }
    
    return
}

// GetUnreadMessageNum 查询单聊未读消息计数
// App 后台可以通过该接口查询特定账号的单聊总未读数（包含所有的单聊会话）或者单个单聊会话的未读数。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/56043
func (a *api) GetUnreadMessageNum(userId string, peerUserIds ...[]string) (ret *UnreadMessageRet, err error) {
    req := getUnreadMessageNumReq{UserId: userId}
    resp := &getUnreadMessageNumResp{}
    
    if len(peerUserIds) > 0 {
        req.PeerUserIds = peerUserIds[0]
    }
    
    if err = a.client.Post(serviceMessage, commandGetUnreadMessageNum, req, resp); err != nil {
        return
    } else {
        ret = &UnreadMessageRet{
            Total:      resp.AllUnreadMsgNum,
            UnreadList: make(map[string]int),
            ErrorList:  resp.PeerErrors,
        }
        
        if len(resp.PeerUnreadMsgNums) > 0 {
            for _, item := range resp.PeerUnreadMsgNums {
                ret.UnreadList[item.UserId] = item.UnreadMsgNum
            }
        }
    }
    
    return
}
