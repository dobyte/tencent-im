/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/9/7 14:03
 * @Desc: TODO
 */

package profile

import (
	"github.com/dobyte/tencent-im/internal/enum"
	"github.com/dobyte/tencent-im/internal/types"
)

type (
	// GenderType 性别类型
	GenderType = types.GenderType

	// AllowType 加好友验证方式
	AllowType = types.AllowType

	// AdminForbidType 管理员禁止加好友标识类型
	AdminForbidType = types.AdminForbidType
)

const (
	// 性别类型
	GenderTypeUnknown = enum.GenderTypeUnknown // 没设置性别
	GenderTypeFemale  = enum.GenderTypeFemale  // 女性
	GenderTypeMale    = enum.GenderTypeMale    // 男性

	// 加好友验证方式
	AllowTypeNeedConfirm = enum.AllowTypeNeedConfirm // 需要经过自己确认对方才能添加自己为好友
	AllowTypeAllowAny    = enum.AllowTypeAllowAny    // 允许任何人添加自己为好友
	AllowTypeDenyAny     = enum.AllowTypeDenyAny     // 不允许任何人添加自己为好友

	// 管理员禁止加好友标识类型
	AdminForbidTypeNone    = enum.AdminForbidTypeNone    // 默认值，允许加好友
	AdminForbidTypeSendOut = enum.AdminForbidTypeSendOut // 禁止该用户发起加好友请求

	// 标准资料字段
	StandardAttrNickname        = enum.StandardAttrNickname        // 昵称
	StandardAttrGender          = enum.StandardAttrGender          // 性别
	StandardAttrBirthday        = enum.StandardAttrBirthday        // 生日
	StandardAttrLocation        = enum.StandardAttrLocation        // 所在地
	StandardAttrSignature       = enum.StandardAttrSignature       // 个性签名
	StandardAttrAllowType       = enum.StandardAttrAllowType       // 加好友验证方式
	StandardAttrLanguage        = enum.StandardAttrLanguage        // 语言
	StandardAttrAvatar          = enum.StandardAttrAvatar          // 头像URL
	StandardAttrMsgSettings     = enum.StandardAttrMsgSettings     // 消息设置
	StandardAttrAdminForbidType = enum.StandardAttrAdminForbidType // 管理员禁止加好友标识
	StandardAttrLevel           = enum.StandardAttrLevel           // 等级
	StandardAttrRole            = enum.StandardAttrRole            // 角色
)
