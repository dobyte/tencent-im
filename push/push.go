/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/5/27 20:42
 * @Desc: Push Api Implementation.
 */

package push

import (
	"errors"
	"time"
	
	"github.com/dobyte/tencent-im/enum"
	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/types"
)

const (
	servicePush          = "all_member_push"
	commandPush          = "im_push"
	commandSetAttrName   = "im_set_attr_name"
	commandGetAttrName   = "im_get_attr_name"
	commandGetAttr       = "im_get_attr"
	commandSetAttr       = "im_set_attr"
	commandRemoveAttr    = "im_remove_attr"
	commandGetTag        = "im_get_tag"
	commandAddTag        = "im_add_tag"
	commandRemoveTag     = "im_remove_tag"
	commandRemoveAllTags = "im_remove_all_tags"
)

var invalidMsgContent = errors.New("invalid msg content")

type Push interface {
	// Push 全员推送
	// 支持全员推送。
	// 支持按用户属性推送。
	// 支持按用户标签推送。
	// 管理员推送消息，接收方看到消息发送者是管理员。
	// 管理员指定某一帐号向其他帐号推送消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
	// 支持消息离线存储，不支持漫游。
	// 由于全员推送需要下发的帐号数量巨大，下发完全部帐号需要一定时间（根据帐号总数而定，一般在一分钟内）。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45934
	Push(arg PushArgument) (taskId string, err error)
}

type push struct {
	client core.Client
}

func NewPush(client core.Client) Push {
	return &push{client: client}
}

// Push 全员推送
// 支持全员推送。
// 支持按用户属性推送。
// 支持按用户标签推送。
// 管理员推送消息，接收方看到消息发送者是管理员。
// 管理员指定某一帐号向其他帐号推送消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
// 支持消息离线存储，不支持漫游。
// 由于全员推送需要下发的帐号数量巨大，下发完全部帐号需要一定时间（根据帐号总数而定，一般在一分钟内）。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45934
func (i *push) Push(arg PushArgument) (taskId string, err error) {
	var msgType string
	var msgBody = make([]types.MsgBody, 0)
	
	for _, msgContent := range arg.MsgContents {
		switch msgContent.(type) {
		case types.MsgTextContent, *types.MsgTextContent:
			msgType = enum.MsgText
		case types.MsgLocationContent, *types.MsgLocationContent:
			msgType = enum.MsgLocation
		case types.MsgFaceContent, *types.MsgFaceContent:
			msgType = enum.MsgFace
		case types.MsgCustomContent, *types.MsgCustomContent:
			msgType = enum.MsgCustom
		case types.MsgSoundContent, *types.MsgSoundContent:
			msgType = enum.MsgSound
		case types.MsgImageContent, *types.MsgImageContent:
			msgType = enum.MsgImage
		case types.MsgFileContent, *types.MsgFileContent:
			msgType = enum.MsgFile
		case types.MsgVideoContent, *types.MsgVideoContent:
			msgType = enum.MsgVideo
		default:
			return "", invalidMsgContent
		}
		
		msgBody = append(msgBody, types.MsgBody{
			MsgType:    msgType,
			MsgContent: msgContent,
		})
	}
	
	resp := &pushResp{}
	
	if err = i.client.Post(servicePush, commandPush, pushReq{
		FromUserId:      arg.FromUserId,
		MsgRandom:       time.Now().UnixNano(),
		MsgBody:         msgBody,
		MsgLifeTime:     arg.MsgLifeTime,
		Condition:       arg.Condition,
		OfflinePushInfo: arg.OfflinePushInfo,
	}, resp); err != nil {
		return
	} else {
		taskId = resp.TaskId
	}
	
	return
}

//
//// SetAttrName Set application property name.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/45935
//func (i *Push) SetAttrName(req *SetAttrNameReq) (*SetAttrNameResp, error) {
//    resp := &SetAttrNameResp{}
//
//    if err := i.rest.Post(i.serviceName, commandSetAttrName, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// GetAttrName Get application property name.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/45936
//func (i *Push) GetAttrName(req *GetAttrNameReq) (*GetAttrNameResp, error) {
//    resp := &GetAttrNameResp{}
//
//    if err := i.rest.Post(i.serviceName, commandGetAttrName, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// GetAttr Get user attributes.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/45937
//func (i *Push) GetAttr(req *GetAttrReq) (*GetAttrResp, error) {
//    resp := &GetAttrResp{}
//
//    if err := i.rest.Post(i.serviceName, commandGetAttr, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// SetAttr Set user attributes.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/45938
//func (i *Push) SetAttr(req *SetAttrReq) (*SetAttrResp, error) {
//    resp := &SetAttrResp{}
//
//    if err := i.rest.Post(i.serviceName, commandSetAttr, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// RemoveAttr Remove user attributes.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/45939
//func (i *Push) RemoveAttr(req *RemoveAttrReq) (*RemoveAttrResp, error) {
//    resp := &RemoveAttrResp{}
//
//    if err := i.rest.Post(i.serviceName, commandRemoveAttr, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// GetTag Get user tags.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/45940
//func (i *Push) GetTag(req *GetTagReq) (*GetTagResp, error) {
//    resp := &GetTagResp{}
//
//    if err := i.rest.Post(i.serviceName, commandGetTag, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// AddTag Add user tags.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/45941
//func (i *Push) AddTag(req *AddTagReq) (*AddTagResp, error) {
//    resp := &AddTagResp{}
//
//    if err := i.rest.Post(i.serviceName, commandAddTag, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// RemoveTag Remove user tags.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/45942
//func (i *Push) RemoveTag(req *RemoveTagReq) (*RemoveTagResp, error) {
//    resp := &RemoveTagResp{}
//
//    if err := i.rest.Post(i.serviceName, commandRemoveTag, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
//
//// RemoveAllTags Remove user all tags.
//// click here to view the document:
//// https://cloud.tencent.com/document/product/269/45943
//func (i *Push) RemoveAllTags(req *RemoveAllTagsReq) (*RemoveAllTagsResp, error) {
//    resp := &RemoveAllTagsResp{}
//
//    if err := i.rest.Post(i.serviceName, commandRemoveAllTags, req, resp); err != nil {
//        return nil, err
//    }
//
//    return resp, nil
//}
