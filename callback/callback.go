/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 14:24
 * @Desc: TODO
 */

package callback

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

const (
	commandStateChange               = "State.StateChange"
	commandBeforeFriendAdd           = "Sns.CallbackPrevFriendAdd"
	commandBeforeFriendResponse      = "Sns.CallbackPrevFriendResponse"
	commandAfterFriendAdd            = "Sns.CallbackFriendAdd"
	commandAfterFriendDelete         = "Sns.CallbackFriendDelete"
	commandAfterBlacklistAdd         = "Sns.CallbackBlackListAdd"
	commandAfterBlacklistDelete      = "Sns.CallbackBlackListDelete"
	commandBeforePrivateMessageSend  = "C2C.CallbackBeforeSendMsg"
	commandAfterPrivateMessageSend   = "C2C.CallbackAfterSendMsg"
	commandAfterPrivateMessageReport = "C2C.CallbackAfterMsgReport"
	commandAfterPrivateMessageRevoke = "C2C.CallbackAfterMsgWithDraw"
	commandBeforeGroupCreate         = "Group.CallbackBeforeCreateGroup"
	commandAfterGroupCreate          = "Group.CallbackAfterCreateGroup"
	commandBeforeApplyJoinGroup      = "Group.CallbackBeforeApplyJoinGroup"
	commandBeforeInviteJoinGroup     = "Group.CallbackBeforeInviteJoinGroup"
	commandAfterNewMemberJoinGroup   = "Group.CallbackAfterNewMemberJoin"
	commandAfterMemberExitGroup      = "Group.CallbackAfterMemberExit"
	commandBeforeGroupMessageSend    = "Group.CallbackBeforeSendMsg"
	commandAfterGroupMessageSend     = "Group.CallbackAfterSendMsg"
	commandAfterGroupFull            = "Group.CallbackAfterGroupFull"
	commandAfterGroupDestroyed       = "Group.CallbackAfterGroupDestroyed"
	commandAfterGroupInfoChanged     = "Group.CallbackAfterGroupInfoChanged"
)

const (
	EventStateChange Event = iota + 1
	EventBeforeFriendAdd
	EventBeforeFriendResponse
	EventAfterFriendAdd
	EventAfterFriendDelete
	EventAfterBlacklistAdd
	EventAfterBlacklistDelete
	EventBeforePrivateMessageSend
	EventAfterPrivateMessageSend
	EventAfterPrivateMessageReport
	EventAfterPrivateMessageRevoke
	EventBeforeGroupCreate
	EventAfterGroupCreate
	EventBeforeApplyJoinGroup
	EventBeforeInviteJoinGroup
	EventAfterNewMemberJoinGroup
	EventAfterMemberExitGroup
	EventBeforeGroupMessageSend
	EventAfterGroupMessageSend
	EventAfterGroupFull
	EventAfterGroupDestroyed
	EventAfterGroupInfoChanged
)

const (
	ackSuccessStatus = "OK"
	ackFailureStatus = "FAIL"

	ackSuccessCode = 0
	ackFailureCode = 1

	queryAppId       = "SdkAppid"
	queryCommand     = "CallbackCommand"
	queryClientId    = "ClientIP"
	queryOptPlatform = "OptPlatform"
	queryContentType = "contenttype"
)

type (
	Event            int
	EventHandlerFunc func(ack Ack, data interface{})
	Options          struct {
		SdkAppId int
	}

	Callback interface {
		// Register 注册事件
		Register(event Event, handler EventHandlerFunc)
		// Listen 监听事件
		Listen(w http.ResponseWriter, r *http.Request)
	}

	callback struct {
		appId    int
		mu       sync.Mutex
		handlers map[Event]EventHandlerFunc
	}

	Ack interface {
		// Ack 应答
		Ack(resp interface{}) error
		// AckFailure 失败应答
		AckFailure(message ...string) error
		// AckSuccess 成功应答
		AckSuccess(code int, message ...string) error
	}

	ack struct {
		w http.ResponseWriter
	}
)

func NewCallback(appId int) Callback {
	return &callback{
		appId:    appId,
		handlers: make(map[Event]EventHandlerFunc),
	}
}

// Register 注册事件
func (c *callback) Register(event Event, handler EventHandlerFunc) {
	c.mu.Lock()
	c.handlers[event] = handler
	c.mu.Unlock()
}

// Listen 监听事件
func (c *callback) Listen(w http.ResponseWriter, r *http.Request) {
	a := newAck(w)

	appId, ok := c.GetQuery(r, queryAppId)
	if !ok || appId != strconv.Itoa(c.appId) {
		_ = a.AckFailure("invalid sdk appId")
		return
	}

	command, ok := c.GetQuery(r, queryCommand)
	if !ok {
		_ = a.AckFailure("invalid callback command")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	_ = r.Body.Close()
	if err != nil {
		_ = a.AckFailure(err.Error())
		return
	}

	if event, data, err := c.parseCommand(command, body); err != nil {
		_ = a.AckFailure(err.Error())
	} else {
		if fn, ok := c.handlers[event]; ok {
			fn(a, data)
			return
		} else {
			_ = a.AckSuccess(ackSuccessCode)
		}
	}
}

// parseCommand parse command and body package.
func (c *callback) parseCommand(command string, body []byte) (event Event, data interface{}, err error) {
	switch command {
	case commandStateChange:
		event = EventStateChange
		data = &StateChange{}
	case commandBeforeFriendAdd:
		event = EventBeforeFriendAdd
		data = &BeforeFriendAdd{}
	case commandBeforeFriendResponse:
		event = EventBeforeFriendResponse
		data = &BeforeFriendResponse{}
	case commandAfterFriendAdd:
		event = EventAfterFriendAdd
		data = &AfterFriendAdd{}
	case commandAfterFriendDelete:
		event = EventAfterFriendDelete
		data = &AfterFriendDelete{}
	case commandAfterBlacklistAdd:
		event = EventAfterBlacklistAdd
		data = &AfterBlacklistAdd{}
	case commandAfterBlacklistDelete:
		event = EventAfterBlacklistDelete
		data = &AfterBlacklistDelete{}
	case commandBeforePrivateMessageSend:
		event = EventBeforePrivateMessageSend
		data = &BeforePrivateMessageSend{}
	case commandAfterPrivateMessageSend:
		event = EventAfterPrivateMessageSend
		data = &AfterPrivateMessageSend{}
	case commandAfterPrivateMessageReport:
		event = EventAfterPrivateMessageReport
		data = &AfterPrivateMessageReport{}
	case commandAfterPrivateMessageRevoke:
		event = EventAfterPrivateMessageRevoke
		data = &AfterPrivateMessageRevoke{}
	case commandBeforeGroupCreate:
		event = EventBeforeGroupCreate
		data = &BeforeGroupCreate{}
	case commandAfterGroupCreate:
		event = EventAfterGroupCreate
		data = &AfterGroupCreate{}
	case commandBeforeApplyJoinGroup:
		event = EventBeforeApplyJoinGroup
		data = &BeforeApplyJoinGroup{}
	case commandBeforeInviteJoinGroup:
		event = EventBeforeInviteJoinGroup
		data = &BeforeInviteJoinGroup{}
	case commandAfterNewMemberJoinGroup:
		event = EventAfterNewMemberJoinGroup
		data = &AfterNewMemberJoinGroup{}
	case commandAfterMemberExitGroup:
		event = EventAfterMemberExitGroup
		data = &AfterMemberExitGroup{}
	case commandBeforeGroupMessageSend:
		event = EventBeforeGroupMessageSend
		data = &BeforeGroupMessageSend{}
	case commandAfterGroupMessageSend:
		event = EventAfterGroupMessageSend
		data = &AfterGroupMessageSend{}
	case commandAfterGroupFull:
		event = EventAfterGroupFull
		data = &AfterGroupFull{}
	case commandAfterGroupDestroyed:
		event = EventAfterGroupDestroyed
		data = &AfterGroupDestroyed{}
	case commandAfterGroupInfoChanged:
		event = EventAfterGroupInfoChanged
		data = &AfterGroupInfoChanged{}
	default:
		return 0, nil, errors.New("invalid callback command")
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return 0, nil, err
	}

	return event, data, nil
}

// GetQuery 获取查询参数
func (c *callback) GetQuery(r *http.Request, key string) (string, bool) {
	if values, ok := r.URL.Query()[key]; ok {
		return values[0], ok
	} else {
		return "", false
	}
}

func newAck(w http.ResponseWriter) Ack {
	return &ack{w}
}

// Ack 应答
func (a *ack) Ack(resp interface{}) error {
	b, _ := json.Marshal(resp)
	a.w.WriteHeader(http.StatusOK)
	_, err := a.w.Write(b)
	return err
}

// AckFailure 应答失败
func (a *ack) AckFailure(message ...string) error {
	resp := BaseResp{}
	resp.ActionStatus = ackFailureStatus
	resp.ErrorCode = ackFailureCode
	if len(message) > 0 {
		resp.ErrorInfo = message[0]
	}

	return a.Ack(resp)
}

// AckSuccess 应答成功
func (a *ack) AckSuccess(code int, message ...string) error {
	resp := BaseResp{}
	resp.ActionStatus = ackSuccessStatus
	resp.ErrorCode = code
	if len(message) > 0 {
		resp.ErrorInfo = message[0]
	}

	return a.Ack(resp)
}
