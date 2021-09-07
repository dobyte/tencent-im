/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/29 17:38
 * @Desc: 资料管理结构体定义
 */

package profile

import "github.com/dobyte/tencent-im/internal/types"

type (
    // 设置资料（请求）
    setProfileReq struct {
        UserId  string          `json:"From_Account"` // （必填）需要设置该 UserID 的资料
        Profile []types.TagPair `json:"ProfileItem"`  // （必填）待设置的用户的资料对象数组
    }
    
    // 获取资料（请求）
    getProfileReq struct {
        UserIds []string `json:"To_Account"` // （必填）需要拉取这些UserID的资料
        TagList []string `json:"TagList"`    // （必填）指定要拉取的资料字段的 Tag，支持的字段有
    }
    
    // 获取资料（响应）
    getProfileResp struct {
        types.ActionBaseResp
        ErrorDisplay string        `json:"ErrorDisplay"`    // 详细的客户端展示信息
        UserProfiles []UserProfile `json:"UserProfileItem"` // 用户资料结构化信息
    }
    
    // UserProfile 用户资料
    UserProfile struct {
        UserId     string          `json:"To_Account"`  // 用户的UserID
        Profile    []types.TagPair `json:"ProfileItem"` // 用户的资料对象数组
        ResultCode int             `json:"ResultCode"`  // 处理结果，0表示成功，非0表示失败
        ResultInfo string          `json:"ResultInfo"`  // 错误描述信息，成功时该字段为空
    }
    
    // GenderType 性别类型
    GenderType = types.GenderType
    
    // AllowType 加好友验证方式
    AllowType = types.AllowType
    
    // AdminForbidType 管理员禁止加好友标识类型
    AdminForbidType = types.AdminForbidType
)
