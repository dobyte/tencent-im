/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 14:24
 * @Desc: TODO
 */

package callback

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	
	"github.com/dobyte/tencent-im/types"
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
	AckSuccess = "OK"
	AckFailure = "FAIL"
	
	queryAppId       = "SdkAppid"
	queryCommand     = "CallbackCommand"
	queryClientId    = "ClientIP"
	queryOptPlatform = "OptPlatform"
	queryContentType = "contenttype"
)

//type (
//	EventType        int
//	EventHandlerFunc func(data interface{}) (*Reply, error)
//	Options          struct {
//		SdkAppId int
//	}
//
//	API interface {
//	}
//
//	api struct {
//		appId    int
//		lock     sync.Mutex
//		handlers map[EventType]EventHandlerFunc
//	}
//)
//
//func NewAPI(appId int) *API {
//	return &api{
//		appId:    appId,
//		handlers: make(map[EventType]EventHandlerFunc),
//	}
//}

//// Register 注册事件
//func (a *api) Register(event EventType, handler EventHandlerFunc) {
//	a.lock.Lock()
//	defer a.lock.Unlock()
//	a.handlers[event] = handler
//}
//
//// SyncListen wait a callback request.
//func (a *api) SyncListen(r *http.Request) {
//	appId, ok := a.getQuery(r, queryAppId)
//	if !ok || appId != strconv.Itoa(a.appId) {
//		_ = a.AckFailure(r.Response. , 0 , "invalid sdk appid")
//		c.ReplyFailure(0, "invalid sdk appid")
//		return
//	}
//
//	command, ok := c.GetQuery(r, queryCommand)
//	if !ok {
//		c.ReplyFailure(0, "invalid callback command")
//		return
//	}
//
//	if event, data, err := c.parseCommand(command); err != nil {
//		c.ReplyFailure(0, err.Error())
//		return
//	} else {
//		if ret, err := c.handleEvent(event, data); err != nil {
//			c.ReplyFailure(0, err.Error())
//			return
//		} else {
//			c.Reply(ret)
//			return
//		}
//	}
//}
//
//func (c *api) AsyncListen(request *http.Request) {
//
//}
//
//// handleEvent 处理监听事件
//func (c *api) handleEvent(event EventType, data interface{}) (*Reply, error) {
//	if fn, ok := c.handlers[event]; ok {
//		return fn(data)
//	}
//
//	return nil, nil
//}
//
//// parseCommand parse command and body package.
//func (c *api) parseCommand(command string) (EventType, interface{}, error) {
//	var (
//		err   error
//		event EventType
//		body  []byte
//		data  interface{}
//	)
//
//	switch command {
//	case CommandStateStateChange:
//		event = EventStateStateChange
//		data = StateChange{}
//	case CommandSnsFriendAdd:
//		event = EventSnsFriendAdd
//		data = SnsFriendAdd{}
//	case CommandSnsFriendDelete:
//		event = EventSnsFriendDelete
//		data = SnsFriendDelete{}
//	case CommandSnsBlackListAdd:
//		event = EventSnsBlackListAdd
//		data = SnsBlackListAdd{}
//	case CommandSnsBlackListDelete:
//		event = EventSnsBlackListDelete
//		data = SnsBlackListDelete{}
//	case CommandC2CBeforeSendMsg:
//		event = EventC2CBeforeSendMsg
//		data = C2CBeforeSendMsg{}
//	case CommandC2CAfterSendMsg:
//		event = EventC2CAfterSendMsg
//		data = C2CAfterSendMsg{}
//	case CommandGroupBeforeCreateGroup:
//		event = EventGroupBeforeCreateGroup
//		data = GroupBeforeCreateGroup{}
//	case CommandGroupAfterCreateGroup:
//		event = EventGroupAfterCreateGroup
//		data = GroupAfterCreateGroup{}
//	case CommandGroupBeforeApplyJoinGroup:
//		event = EventGroupBeforeApplyJoinGroup
//		data = GroupBeforeApplyJoinGroup{}
//	case CommandGroupBeforeInviteJoinGroup:
//		event = EventGroupBeforeInviteJoinGroup
//		data = GroupBeforeInviteJoinGroup{}
//	case CommandGroupAfterNewMemberJoin:
//		event = EventGroupAfterNewMemberJoin
//		data = GroupAfterNewMemberJoin{}
//	case CommandGroupAfterMemberExit:
//		event = EventGroupAfterMemberExit
//		data = GroupAfterMemberExit{}
//	case CommandGroupBeforeSendMsg:
//		event = EventGroupBeforeSendMsg
//		data = GroupBeforeSendMsg{}
//	case CommandGroupAfterSendMsg:
//		event = EventGroupAfterSendMsg
//		data = GroupAfterSendMsg{}
//	case CommandGroupAfterGroupFull:
//		event = EventGroupAfterGroupFull
//		data = GroupAfterGroupFull{}
//	case CommandGroupAfterGroupDestroyed:
//		event = EventGroupAfterGroupDestroyed
//		data = GroupAfterGroupDestroyed{}
//	case CommandGroupAfterGroupInfoChanged:
//		event = EventGroupAfterGroupInfoChanged
//		data = GroupAfterGroupInfoChanged{}
//	default:
//		return 0, nil, errors.New("invalid callback command")
//	}
//
//	if err = json.Unmarshal(body, &data); err != nil {
//		return 0, nil, err
//	}
//
//	return event, data, nil
//}
//
//// AckSuccess 应答失败
//func (a *api) AckFailure(w http.ResponseWriter, code int, message ...string) error {
//	return a.Ack(w, AckFailure, code, message...)
//}
//
//// AckSuccess 应答成功
//func (a *api) AckSuccess(w http.ResponseWriter, code int, message ...string) error {
//	return a.Ack(w, AckSuccess, code, message...)
//}
//
//// Ack 应答
//func (a *api) Ack(w http.ResponseWriter, status string, code int, message ...string) error {
//	ack := types.ActionBaseResp{}
//	ack.ActionStatus = status
//	ack.ErrorCode = code
//	if len(message) > 0 {
//		ack.ErrorInfo = message[0]
//	}
//	b, _ := json.Marshal(ack)
//	w.WriteHeader(http.StatusOK)
//	_, err := w.Write(b)
//	return err
//}
//
//// GetQuery 获取查询参数
//func (a *api) getQuery(r *http.Request, key string) (string, bool) {
//	if values, ok := r.URL.Query()[key]; ok {
//		return values[0], ok
//	} else {
//		return "", false
//	}
//}
