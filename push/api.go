/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:42
 * @Desc: 全员推送
 */

package push

import (
	"fmt"
	"strconv"

	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/internal/enum"
	"github.com/dobyte/tencent-im/internal/types"
)

const (
	service                  = "all_member_push"
	commandPushMessage       = "im_push"
	commandSetAttrNames      = "im_set_attr_name"
	commandGetAttrNames      = "im_get_attr_name"
	commandGetUserAttrs      = "im_get_attr"
	commandSetUserAttrs      = "im_set_attr"
	commandDeleteUserAttrs   = "im_remove_attr"
	commandGetUserTags       = "im_get_tag"
	commandAddUserTags       = "im_add_tag"
	commandDeleteUserTags    = "im_remove_tag"
	commandDeleteUserAllTags = "im_remove_all_tags"

	batchSetAttrNamesLimit          = 10  // 批量设置应用属性名限制
	batchGetUserAttrsLimit          = 100 // 批量获取用户属性限制
	batchSetUserAttrsLimit          = 100 // 批量设置用户属性限制
	batchDeleteUserAttrsLimit       = 100 // 批量删除用户属性限制
	batchAddUserTagsLimit           = 100 // 批量添加用户标签限制
	batchGetUserTagsLimit           = 100 // 批量获取用户标签限制
	batchDeleteUserTagsLimit        = 100 // 批量删除用户标签限制
	batchDeleteUserAllTagsUserLimit = 100 // 批量删除用户所有标签的用户限制
)

type API interface {
	// PushMessage 全员推送
	// 支持全员推送。
	// 支持按用户属性推送。
	// 支持按用户标签推送。
	// 管理员推送消息，接收方看到消息发送者是管理员。
	// 管理员指定某一帐号向其他帐号推送消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
	// 支持消息离线存储，不支持漫游。
	// 由于全员推送需要下发的帐号数量巨大，下发完全部帐号需要一定时间（根据帐号总数而定，一般在一分钟内）。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45934
	PushMessage(message *Message) (taskId string, err error)

	// SetAttrNames 设置应用属性名称
	// 每个应用可以设置自定义的用户属性，最多可以有10个。通过本接口可以设置每个属性的名称，设置完成后，即可用于按用户属性推送等。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45935
	SetAttrNames(attrNames map[int]string) (err error)

	// GetAttrNames 获取应用属性名称
	// 管理员获取应用属性名称。使用前请先 设置应用属性名称 。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45936
	GetAttrNames() (attrNames map[int]string, err error)

	// GetUserAttrs 获取用户属性
	// 获取用户属性（必须以管理员帐号调用）；每次最多只能获取100个用户的属性。使用前请先 设置应用属性名称 。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45937
	GetUserAttrs(userIds ...string) (attrs map[string]map[string]interface{}, err error)

	// SetUserAttrs 设置用户属性
	// 管理员给用户设置属性。每次最多只能给100个用户设置属性。使用前请先 设置应用属性名称 。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45938
	SetUserAttrs(attrs map[string]map[string]interface{}) (err error)

	// DeleteUserAttrs 删除用户属性
	// 管理员给用户删除属性。注意每次最多只能给100个用户删除属性。使用前请先 设置应用属性名称。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45939
	DeleteUserAttrs(attrs map[string][]string) (err error)

	// GetUserTags 获取用户标签
	// 获取用户标签（必须以管理员帐号调用）。每次最多只能获取100个用户的标签。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45940
	GetUserTags(userIds ...string) (tags map[string][]string, err error)

	// AddUserTags 添加用户标签
	// 管理员给用户添加标签。
	// 每次请求最多只能给100个用户添加标签，请求体中单个用户添加标签数最多为10个。
	// 单个用户可设置最大标签数为100个，若用户当前标签超过100，则添加新标签之前请先删除旧标签。
	// 单个标签最大长度为50字节。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45941
	AddUserTags(tags map[string][]string) (err error)

	// DeleteUserTags 删除用户标签
	// 管理员给用户删除标签。注意每次最多只能给100个用户删除标签。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45942
	DeleteUserTags(tags map[string][]string) (err error)

	// DeleteUserAllTags 删除用户所有标签
	// 管理员给用户删除所有标签。注意每次最多只能给100个用户删除所有标签。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/45943
	DeleteUserAllTags(userIds ...string) (err error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// PushMessage 全员推送
// 支持全员推送。
// 支持按用户属性推送。
// 支持按用户标签推送。
// 管理员推送消息，接收方看到消息发送者是管理员。
// 管理员指定某一帐号向其他帐号推送消息，接收方看到发送者不是管理员，而是管理员指定的帐号。
// 支持消息离线存储，不支持漫游。
// 由于全员推送需要下发的帐号数量巨大，下发完全部帐号需要一定时间（根据帐号总数而定，一般在一分钟内）。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45934
func (a *api) PushMessage(message *Message) (taskId string, err error) {
	if err = message.checkError(); err != nil {
		return
	}

	req := &pushMessageReq{}
	req.FromUserId = message.GetSender()
	req.MsgLifeTime = message.GetLifeTime()
	req.OfflinePushInfo = message.GetOfflinePushInfo()
	req.MsgBody = message.GetBody()
	req.MsgRandom = message.GetRandom()
	req.Condition = message.GetCondition()

	resp := &pushMessageResp{}

	if err = a.client.Post(service, commandPushMessage, req, resp); err != nil {
		return
	}

	taskId = resp.TaskId

	return
}

// SetAttrNames 设置应用属性名称
// 每个应用可以设置自定义的用户属性，最多可以有10个。通过本接口可以设置每个属性的名称，设置完成后，即可用于按用户属性推送等。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45935
func (a *api) SetAttrNames(attrNames map[int]string) (err error) {
	if c := len(attrNames); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the attribute names is not set")
		return
	} else if c > batchSetAttrNamesLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of attribute names to be set cannot exceed %d", batchSetAttrNamesLimit))
		return
	}

	req := &setAttrNamesReq{AttrNames: make(map[string]string, len(attrNames))}

	for i, attrName := range attrNames {
		req.AttrNames[strconv.Itoa(i)] = attrName
	}

	if err = a.client.Post(service, commandSetAttrNames, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// GetAttrNames 获取应用属性名称
// 管理员获取应用属性名称。使用前请先 设置应用属性名称 。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45936
func (a *api) GetAttrNames() (attrNames map[int]string, err error) {
	req := &getAttrNamesReq{}
	resp := &getAttrNamesResp{}

	if err = a.client.Post(service, commandGetAttrNames, req, resp); err != nil {
		return
	}

	if len(resp.AttrNames) > 0 {
		var i int
		attrNames = make(map[int]string, len(resp.AttrNames))
		for key, attrName := range resp.AttrNames {
			i, _ = strconv.Atoi(key)
			attrNames[i] = attrName
		}
	}

	return
}

// GetUserAttrs 获取用户属性
// 获取用户属性（必须以管理员帐号调用）；每次最多只能获取100个用户的属性。使用前请先 设置应用属性名称 。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45937
func (a *api) GetUserAttrs(userIds ...string) (attrs map[string]map[string]interface{}, err error) {
	if c := len(userIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the accounts is not set")
		return
	} else if c > batchGetUserAttrsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of accounts being queried cannot exceed %d", batchGetUserAttrsLimit))
		return
	}

	req := &getUserAttrsReq{UserIds: userIds}
	resp := &getUserAttrsResp{}

	if err = a.client.Post(service, commandGetUserAttrs, req, resp); err != nil {
		return
	}

	attrs = make(map[string]map[string]interface{}, len(resp.Attrs))
	for _, attr := range resp.Attrs {
		attrs[attr.UserId] = attr.Attrs
	}

	return
}

// SetUserAttrs 设置用户属性
// 管理员给用户设置属性。每次最多只能给100个用户设置属性。使用前请先 设置应用属性名称 。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45938
func (a *api) SetUserAttrs(attrs map[string]map[string]interface{}) (err error) {
	if c := len(attrs); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the attributes is not set")
		return
	} else if c > batchSetUserAttrsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of attributes to be set cannot exceed %d", batchSetUserAttrsLimit))
		return
	}

	req := &setUserAttrsReq{Attrs: make([]*userAttrItem, 0, len(attrs))}
	for userId, attrs := range attrs {
		req.Attrs = append(req.Attrs, &userAttrItem{
			UserId: userId,
			Attrs:  attrs,
		})
	}

	if err = a.client.Post(service, commandSetUserAttrs, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// DeleteUserAttrs 删除用户属性
// 管理员给用户删除属性。注意每次最多只能给100个用户删除属性。使用前请先 设置应用属性名称。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45939
func (a *api) DeleteUserAttrs(attrs map[string][]string) (err error) {
	if c := len(attrs); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the attributes is not set")
		return
	} else if c > batchDeleteUserAttrsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of attributes to be delete cannot exceed %d", batchDeleteUserAttrsLimit))
		return
	}

	req := &deleteUserAttrsReq{Attrs: make([]deleteUserAttr, 0, len(attrs))}
	for userId, attrs := range attrs {
		req.Attrs = append(req.Attrs, deleteUserAttr{
			UserId: userId,
			Attrs:  attrs,
		})
	}

	if err = a.client.Post(service, commandDeleteUserAttrs, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// GetUserTags 获取用户标签
// 获取用户标签（必须以管理员帐号调用）。每次最多只能获取100个用户的标签。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45940
func (a *api) GetUserTags(userIds ...string) (tags map[string][]string, err error) {
	if c := len(userIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the accounts is not set")
		return
	} else if c > batchGetUserTagsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of tags being queried cannot exceed %d", batchGetUserTagsLimit))
		return
	}

	req := &getUserTagsReq{UserIds: userIds}
	resp := &getUserTagsResp{}

	if err = a.client.Post(service, commandGetUserTags, req, resp); err != nil {
		return
	}

	if len(resp.Tags) > 0 {
		tags = make(map[string][]string, len(resp.Tags))
		for _, item := range resp.Tags {
			tags[item.UserId] = item.Tags
		}
	}

	return
}

// AddUserTags 添加用户标签
// 管理员给用户添加标签。
// 每次请求最多只能给100个用户添加标签，请求体中单个用户添加标签数最多为10个。
// 单个用户可设置最大标签数为100个，若用户当前标签超过100，则添加新标签之前请先删除旧标签。
// 单个标签最大长度为50字节。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45941
func (a *api) AddUserTags(tags map[string][]string) (err error) {
	if c := len(tags); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the tags of user is not set")
		return
	} else if c > batchAddUserTagsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of tags to be add cannot exceed %d", batchAddUserTagsLimit))
		return
	}

	req := &addUserTagsReq{Tags: make([]*userTag, 0, len(tags))}
	for userId, tags := range tags {
		req.Tags = append(req.Tags, &userTag{
			UserId: userId,
			Tags:   tags,
		})
	}

	if err = a.client.Post(service, commandAddUserTags, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// DeleteUserTags 删除用户标签
// 管理员给用户删除标签。注意每次最多只能给100个用户删除标签。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45942
func (a *api) DeleteUserTags(tags map[string][]string) (err error) {
	if c := len(tags); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the tags of user is not set")
		return
	} else if c > batchDeleteUserTagsLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of tags to be delete cannot exceed %d", batchDeleteUserTagsLimit))
		return
	}

	req := &deleteUserTagsReq{Tags: make([]*userTag, 0, len(tags))}
	for userId, tags := range tags {
		req.Tags = append(req.Tags, &userTag{
			UserId: userId,
			Tags:   tags,
		})
	}

	if err = a.client.Post(service, commandDeleteUserTags, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}

// DeleteUserAllTags 删除用户所有标签
// 管理员给用户删除所有标签。注意每次最多只能给100个用户删除所有标签。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/45943
func (a *api) DeleteUserAllTags(userIds ...string) (err error) {
	if c := len(userIds); c == 0 {
		err = core.NewError(enum.InvalidParamsCode, "the accounts is not set")
		return
	} else if c > batchDeleteUserAllTagsUserLimit {
		err = core.NewError(enum.InvalidParamsCode, fmt.Sprintf("the number of accounts to be delete cannot exceed %d", batchDeleteUserAllTagsUserLimit))
		return
	}

	req := &deleteUserAllTagsReq{UserIds: userIds}

	if err = a.client.Post(service, commandDeleteUserAllTags, req, &types.ActionBaseResp{}); err != nil {
		return
	}

	return
}
