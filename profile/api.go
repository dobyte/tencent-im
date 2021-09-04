/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 20:44
 * @Desc: 资料管理
 */

package profile

import (
	"github.com/dobyte/tencent-im/internal/core"
	"github.com/dobyte/tencent-im/types"
)

const (
	serviceName        = "profile"
	commandSetPortrait = "portrait_set"
	commandGetPortrait = "portrait_get"
)

type API interface {
	// SetProfile 设置资料
	// 支持 标配资料字段 和 自定义资料字段 的设置
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1640
	SetProfile(profile *profile) (err error)
	
	// GetProfiles 拉取资料
	// 支持拉取好友和非好友的资料字段。
	// 支持拉取 标配资料字段 和 自定义资料字段。
	// 建议每次拉取的用户数不超过100，避免因回包数据量太大导致回包失败。
	// 请确保请求中的所有帐号都已导入即时通信 IM，如果请求中含有未导入即时通信 IM 的帐号，即时通信 IM 后台将会提示错误。
	// 点击查看详细文档:
	// https://cloud.tencent.com/document/product/269/1639
	GetProfiles(userIds []string, attrs []string) (profiles []*profile, err error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// SetProfile 设置资料
// 支持 标配资料字段 和 自定义资料字段 的设置
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1640
func (a *api) SetProfile(profile *profile) (err error) {
	if len(profile.attrs) > 0 {
		req := setProfileReq{UserId: profile.userId}
		
		for tag, value := range profile.attrs {
			req.Profile = append(req.Profile, types.TagPair{
				Tag:   tag,
				Value: value,
			})
		}
		
		if err = a.client.Post(serviceName, commandSetPortrait, req, &types.ActionBaseResp{}); err != nil {
			return
		}
	}
	
	return
}

// GetProfiles 拉取资料
// 支持拉取好友和非好友的资料字段。
// 支持拉取 标配资料字段 和 自定义资料字段。
// 建议每次拉取的用户数不超过100，避免因回包数据量太大导致回包失败。
// 请确保请求中的所有帐号都已导入即时通信 IM，如果请求中含有未导入即时通信 IM 的帐号，即时通信 IM 后台将会提示错误。
// 点击查看详细文档:
// https://cloud.tencent.com/document/product/269/1639
func (a *api) GetProfiles(userIds []string, attrs []string) (profiles []*profile, err error) {
	req := &getProfileReq{UserIds: userIds, TagList: attrs}
	resp := &getProfileResp{}
	
	if err := a.client.Post(serviceName, commandGetPortrait, req, resp); err != nil {
		return nil, err
	}
	
	for _, account := range resp.UserProfiles {
		profile := NewProfile(account.UserId)
		profile.setAbnormal(account.ResultCode, account.ResultInfo)
		for _, item := range account.Profile {
			profile.setAttr(item.Tag, item.Value)
		}
		profiles = append(profiles, profile)
	}
	
	return
}
