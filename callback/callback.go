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

	"github.com/dobyte/tencent-im/internal/types"
)

const (
	CommandStateChange          = "State.StateChange"
	CommandBeforeFriendAdd      = "Sns.CallbackPrevFriendAdd"
	CommandBeforeFriendResponse = "Sns.CallbackPrevFriendResponse"
	CommandAfterFriendAdd       = "Sns.CallbackFriendAdd"

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
	EventStateChange Event = iota + 1
	EventBeforeFriendAdd
	EventBeforeFriendResponse
	EventAfterFriendAdd

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
	EventHandlerFunc func(data interface{}) error
	Options          struct {
		SdkAppId int
	}

	Callback interface {
	}

	callback struct {
		appId    int
		lock     sync.Mutex
		handlers map[Event]EventHandlerFunc
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
	c.lock.Lock()
	c.handlers[event] = handler
	c.lock.Unlock()
}

// Listen 监听事件
func (c *callback) Listen(w http.ResponseWriter, r *http.Request) {
	appId, ok := c.GetQuery(r, queryAppId)
	if !ok || appId != strconv.Itoa(c.appId) {
		_ = c.AckFailure(w, "invalid sdk appId")
		return
	}

	command, ok := c.GetQuery(r, queryCommand)
	if !ok {
		_ = c.AckFailure(w, "invalid callback command")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	_ = r.Body.Close()
	if err != nil {
		_ = c.AckFailure(w, err.Error())
		return
	}

	if event, data, err := c.parseCommand(command, body); err != nil {
		_ = c.AckFailure(w, err.Error())
	} else {
		if err = c.handleEvent(event, data); err != nil {
			_ = c.AckFailure(w, err.Error())
		} else {
			_ = c.AckSuccess(w)
		}
	}
}

// handleEvent 处理监听事件
func (c *callback) handleEvent(event Event, data interface{}) error {
	if fn, ok := c.handlers[event]; ok {
		return fn(data)
	}

	return nil
}

// parseCommand parse command and body package.
func (c *callback) parseCommand(command string, body []byte) (event Event, data interface{}, err error) {
	switch command {
	case CommandStateChange:
		event = EventStateChange
		data = &StateChange{}
	case CommandBeforeFriendAdd:
		event = EventBeforeFriendAdd
		data = &BeforeFriendAdd{}
	case CommandBeforeFriendResponse:
		event = EventBeforeFriendResponse
		data = &BeforeFriendResponse{}
	// case CommandAfterFriendAdd:
	//
	//
	//
	// case CommandSnsFriendAdd:
	//     event = EventSnsFriendAdd
	//     data = SnsFriendAdd{}
	case CommandSnsFriendDelete:
		event = EventSnsFriendDelete
		data = SnsFriendDelete{}
	case CommandSnsBlackListAdd:
		event = EventSnsBlackListAdd
		data = SnsBlackListAdd{}
	case CommandSnsBlackListDelete:
		event = EventSnsBlackListDelete
		data = SnsBlackListDelete{}
	case CommandC2CBeforeSendMsg:
		event = EventC2CBeforeSendMsg
		data = C2CBeforeSendMsg{}
	case CommandC2CAfterSendMsg:
		event = EventC2CAfterSendMsg
		data = C2CAfterSendMsg{}
	case CommandGroupBeforeCreateGroup:
		event = EventGroupBeforeCreateGroup
		data = GroupBeforeCreateGroup{}
	case CommandGroupAfterCreateGroup:
		event = EventGroupAfterCreateGroup
		data = GroupAfterCreateGroup{}
	case CommandGroupBeforeApplyJoinGroup:
		event = EventGroupBeforeApplyJoinGroup
		data = GroupBeforeApplyJoinGroup{}
	case CommandGroupBeforeInviteJoinGroup:
		event = EventGroupBeforeInviteJoinGroup
		data = GroupBeforeInviteJoinGroup{}
	case CommandGroupAfterNewMemberJoin:
		event = EventGroupAfterNewMemberJoin
		data = GroupAfterNewMemberJoin{}
	case CommandGroupAfterMemberExit:
		event = EventGroupAfterMemberExit
		data = GroupAfterMemberExit{}
	case CommandGroupBeforeSendMsg:
		event = EventGroupBeforeSendMsg
		data = GroupBeforeSendMsg{}
	case CommandGroupAfterSendMsg:
		event = EventGroupAfterSendMsg
		data = GroupAfterSendMsg{}
	case CommandGroupAfterGroupFull:
		event = EventGroupAfterGroupFull
		data = GroupAfterGroupFull{}
	case CommandGroupAfterGroupDestroyed:
		event = EventGroupAfterGroupDestroyed
		data = GroupAfterGroupDestroyed{}
	case CommandGroupAfterGroupInfoChanged:
		event = EventGroupAfterGroupInfoChanged
		data = GroupAfterGroupInfoChanged{}
	default:
		return 0, nil, errors.New("invalid callback command")
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return 0, nil, err
	}

	return event, data, nil
}

// AckFailure 应答失败
func (c *callback) AckFailure(w http.ResponseWriter, message ...string) error {
	return c.Ack(w, ackFailureStatus, ackFailureCode, message...)
}

// AckSuccess 应答成功
func (c *callback) AckSuccess(w http.ResponseWriter, message ...string) error {
	return c.Ack(w, ackSuccessStatus, ackSuccessCode, message...)
}

// Ack 应答
func (c *callback) Ack(w http.ResponseWriter, status string, code int, message ...string) error {
	ack := types.ActionBaseResp{}
	ack.ActionStatus = status
	ack.ErrorCode = code
	if len(message) > 0 {
		ack.ErrorInfo = message[0]
	}
	b, _ := json.Marshal(ack)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(b)
	return err
}

// GetQuery 获取查询参数
func (c *callback) GetQuery(r *http.Request, key string) (string, bool) {
	if values, ok := r.URL.Query()[key]; ok {
		return values[0], ok
	} else {
		return "", false
	}
}
