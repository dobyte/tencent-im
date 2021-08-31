/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/5/27 19:32
 * @Desc: Message Api Implementation.
 */

package api

const (
    serviceMessage         = "openim"
    commandSendMsg         = "sendmsg"
    commandBatchSendMsg    = "batchsendmsg"
    commandImportMsg       = "importmsg"
    commandGetRoamMsg      = "admin_getroammsg"
    commandWithdrawMsg     = "admin_msgwithdraw"
    commandSetMsgRead      = "admin_set_msg_read"
    commandGetUnreadMsgNum = "get_c2c_unread_msg_num"
)

//type Message struct {
//    rest        core.Rest
//    serviceName string
//}
//
//func NewMessage(rest core.Rest) *Message {
//    return &Message{
//        rest:        rest,
//        serviceName: serviceMessage,
//    }
//}
//
//// SendMsg Message to account.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/2282
//func (i *Message) SendMsg(req *SendMsgReq) (*SendMsgResp, error) {
//    resp := &SendMsgResp{}
//
//    if err := i.rest.Post(i.serviceName, commandSendMsg, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// BatchSendMsg Send batches and chat messages.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/1612
//func (i *Message) BatchSendMsg(req *BatchSendMsgReq) (*BatchSendMsgResp, error) {
//    resp := &BatchSendMsgResp{}
//
//    if err := i.rest.Post(i.serviceName, commandBatchSendMsg, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// ImportMsg Import historical chat messages to IM.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/2568
//func (i *Message) ImportMsg(req *ImportMsgReq) (*ImportMsgResp, error) {
//    resp := &ImportMsgResp{}
//
//    if err := i.rest.Post(i.serviceName, commandImportMsg, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// GetRoamMsg Query the message record of a single chat session based on conditions.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/42794
//func (i *Message) GetRoamMsg(req *GetRoamMsgReq) (*GetRoamMsgResp, error) {
//    resp := &GetRoamMsgResp{}
//
//    if err := i.rest.Post(i.serviceName, commandGetRoamMsg, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// WithdrawMsg Withdraw single chat message.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/38980
//func (i *Message) WithdrawMsg(req *WithdrawMsgReq) (*WithdrawMsgResp, error) {
//    resp := &WithdrawMsgResp{}
//
//    if err := i.rest.Post(i.serviceName, commandWithdrawMsg, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// SetMsgRead Set that all messages of a user’s single chat session have been read.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/50349
//func (i *Message) SetMsgRead(req *SetMsgReadReq) (*SetMsgReadResp, error) {
//    resp := &SetMsgReadResp{}
//
//    if err := i.rest.Post(i.serviceName, commandSetMsgRead, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// GetUnreadMsgNum Query the total unread value of the single chat of a specific account.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/56043
//func (i *Message) GetUnreadMsgNum(req *GetUnreadMsgNumReq) (*GetUnreadMsgNumResp, error) {
//    resp := &GetUnreadMsgNumResp{}
//
//    if err := i.rest.Post(i.serviceName, commandGetUnreadMsgNum, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
