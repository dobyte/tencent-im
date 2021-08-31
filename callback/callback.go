/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 14:24
 * @Desc: TODO
 */

package callback

import (
    "sync"
)

const (
    CommandStateStateChange           = "State.StateChange"
    CommandSnsFriendAdd               = "Sns.CallbackFriendAdd"
    CommandSnsFriendDelete            = "Sns.CallbackFriendDelete"
    CommandSnsBlackListAdd            = "Sns.CallbackBlackListAdd"
    CommandSnsBlackListDelete         = "Sns.CallbackBlackListDelete"
    CommandC2CBeforeSendMsg           = "C2C.CallbackBeforeSendMsg"
    CommandC2CAfterSendMsg            = "C2C.CallbackAfterSendMsg"
    CommandGroupBeforeCreateGroup     = "Group.CallbackBeforeCreateGroup"
    CommandGroupAfterCreateGroup      = "Group.CallbackAfterCreateGroup"
    CommandGroupBeforeApplyJoinGroup  = "Group.CallbackBeforeApplyJoinGroup"
    CommandGroupBeforeInviteJoinGroup = "Group.CallbackBeforeInviteJoinGroup"
    CommandGroupAfterNewMemberJoin    = "Group.CallbackAfterNewMemberJoin"
    CommandGroupAfterMemberExit       = "Group.CallbackAfterMemberExit"
    CommandGroupBeforeSendMsg         = "Group.CallbackBeforeSendMsg"
    CommandGroupAfterSendMsg          = "Group.CallbackAfterSendMsg"
    CommandGroupAfterGroupFull        = "Group.CallbackAfterGroupFull"
    CommandGroupAfterGroupDestroyed   = "Group.CallbackAfterGroupDestroyed"
    CommandGroupAfterGroupInfoChanged = "Group.CallbackAfterGroupInfoChanged"
)

const (
    EventStateStateChange EventType = iota + 1
    EventSnsFriendAdd
    EventSnsFriendDelete
    EventSnsBlackListAdd
    EventSnsBlackListDelete
    EventC2CBeforeSendMsg
    EventC2CAfterSendMsg
    EventGroupBeforeCreateGroup
    EventGroupAfterCreateGroup
    EventGroupBeforeApplyJoinGroup
    EventGroupBeforeInviteJoinGroup
    EventGroupAfterNewMemberJoin
    EventGroupAfterMemberExit
    EventGroupBeforeSendMsg
    EventGroupAfterSendMsg
    EventGroupAfterGroupFull
    EventGroupAfterGroupDestroyed
    EventGroupAfterGroupInfoChanged
)

const (
    ReplySuccess = "OK"
    ReplyFailure = "FAIL"
    
    queryAppId       = "SdkAppid"
    queryCommand     = "CallbackCommand"
    queryClientId    = "ClientIP"
    queryOptPlatform = "OptPlatform"
    queryContentType = "contenttype"
)

type (
    EventType        int
    EventHandlerFunc func(data interface{}) (*Reply, error)
    Options          struct {
        SdkAppId int
    }
    Reply struct {
        ActionStatus string `json:"ActionStatus"`
        ErrorInfo    string `json:"ErrorInfo"`
        ErrorCode    int    `json:"ErrorCode"`
    }
    Callback struct {
        opt      *Options
        lock     sync.Mutex
        handlers map[EventType]EventHandlerFunc
    }
)

func NewCallback(opt *Options) *Callback {
    return &Callback{
        opt:      opt,
        handlers: make(map[EventType]EventHandlerFunc),
    }
}

//// Register register an event handler.
//func (c *Callback) Register(event EventType, handler EventHandlerFunc) {
//    c.lock.Lock()
//    defer c.lock.Unlock()
//    c.handlers[event] = handler
//}
//
//// SyncListen wait a callback request.
//func (c *Callback) SyncListen(request *http.Request) {
//    queries := request.URL.Query()
//
//    appId, ok := c.GetQuery(queries, queryAppId)
//    if !ok || appId != c.opt.SdkAppId {
//        c.ReplyFailure(0, "invalid sdk appid")
//        return
//    }
//
//    command, ok := c.GetQuery(queries, queryCommand)
//    if !ok {
//        c.ReplyFailure(0, "invalid callback command")
//        return
//    }
//
//    if event, data, err := c.parseCommand(command); err != nil {
//        c.ReplyFailure(0, err.Error())
//        return
//    } else {
//        if ret, err := c.handleEvent(event, data); err != nil {
//            c.ReplyFailure(0, err.Error())
//            return
//        } else {
//            c.Reply(ret)
//            return
//        }
//    }
//}
//
//func (c *Callback) AsyncListen(request *http.Request) {
//
//}
//
//// handleEvent handle an event.
//func (c *Callback) handleEvent(event EventType, data interface{}) (*Reply, error) {
//    if fn, ok := c.handlers[event]; ok {
//        return fn(data)
//    }
//
//    return nil, nil
//}
//
//// parseCommand parse command and body package.
//func (c *Callback) parseCommand(command string) (EventType, interface{}, error) {
//    var (
//        err   error
//        event EventType
//        body  []byte
//        data  interface{}
//    )
//
//    switch command {
//    case CommandStateStateChange:
//        event = EventStateStateChange
//        data = StateStateChange{}
//
//    case CommandSnsFriendAdd:
//        event = EventSnsFriendAdd
//        data = SnsFriendAdd{}
//
//    case CommandSnsFriendDelete:
//        event = EventSnsFriendDelete
//        data = SnsFriendDelete{}
//
//    case CommandSnsBlackListAdd:
//        event = EventSnsBlackListAdd
//        data = SnsBlackListAdd{}
//
//    case CommandSnsBlackListDelete:
//        event = EventSnsBlackListDelete
//        data = SnsBlackListDelete{}
//
//    case CommandC2CBeforeSendMsg:
//        event = EventC2CBeforeSendMsg
//        data = C2CBeforeSendMsg{}
//
//    case CommandC2CAfterSendMsg:
//        event = EventC2CAfterSendMsg
//        data = C2CAfterSendMsg{}
//
//    case CommandGroupBeforeCreateGroup:
//        event = EventGroupBeforeCreateGroup
//        data = GroupBeforeCreateGroup{}
//
//    case CommandGroupAfterCreateGroup:
//        event = EventGroupAfterCreateGroup
//        data = GroupAfterCreateGroup{}
//
//    case CommandGroupBeforeApplyJoinGroup:
//        event = EventGroupBeforeApplyJoinGroup
//        data = GroupBeforeApplyJoinGroup{}
//
//    case CommandGroupBeforeInviteJoinGroup:
//        event = EventGroupBeforeInviteJoinGroup
//        data = GroupBeforeInviteJoinGroup{}
//
//    case CommandGroupAfterNewMemberJoin:
//        event = EventGroupAfterNewMemberJoin
//        data = GroupAfterNewMemberJoin{}
//
//    case CommandGroupAfterMemberExit:
//        event = EventGroupAfterMemberExit
//        data = GroupAfterMemberExit{}
//
//    case CommandGroupBeforeSendMsg:
//        event = EventGroupBeforeSendMsg
//        data = GroupBeforeSendMsg{}
//
//    case CommandGroupAfterSendMsg:
//        event = EventGroupAfterSendMsg
//        data = GroupAfterSendMsg{}
//
//    case CommandGroupAfterGroupFull:
//        event = EventGroupAfterGroupFull
//        data = GroupAfterGroupFull{}
//
//    case CommandGroupAfterGroupDestroyed:
//        event = EventGroupAfterGroupDestroyed
//        data = GroupAfterGroupDestroyed{}
//
//    case CommandGroupAfterGroupInfoChanged:
//        event = EventGroupAfterGroupInfoChanged
//        data = GroupAfterGroupInfoChanged{}
//    default:
//        return 0, nil, errors.New("invalid callback command")
//    }
//
//    if err = json.Unmarshal(body, &data); err != nil {
//        return 0, nil, err
//    }
//
//    return event, data, nil
//}
//
//// ReplySuccess reply success to client.
//func (c *Callback) ReplySuccess(code int, info string) {
//    c.Reply(&Reply{
//        ActionStatus: ReplySuccess,
//        ErrorCode:    code,
//        ErrorInfo:    info,
//    })
//}
//
//// ReplyFailure reply failure to client.
//func (c *Callback) ReplyFailure(code int, info string) {
//    c.Reply(&Reply{
//        ActionStatus: ReplyFailure,
//        ErrorCode:    code,
//        ErrorInfo:    info,
//    })
//}
//
//func (c *Callback) Reply(reply *Reply) {
//
//}
//
//func (c *Callback) Response(request *http.Request) {
//    // request.Response.Write()
//
//}
//
//// GetQuery Retrieves a query value from the request.
//func (c *Callback) GetQuery(queries url.Values, key string) (string, bool) {
//    if values, ok := queries[key]; ok {
//        return values[0], ok
//    }
//    return "", false
//}
