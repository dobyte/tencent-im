/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/9/7 14:40
 * @Desc: TODO
 */

package sns

import "github.com/dobyte/tencent-im/internal/enum"

const (
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
    
    // 好友属性
    FriendAttrAddSource  = "Tag_SNS_IM_AddSource"  // 添加源
    FriendAttrRemark     = "Tag_SNS_IM_Remark"     // 备注
    FriendAttrGroup      = "Tag_SNS_IM_Group"      // 分组
    FriendAttrAddWording = "Tag_SNS_IM_AddWording" // 附言信息
    FriendAttrAddTime    = "Tag_SNS_IM_AddTime"    // 添加时间
    FriendAttrRemarkTime = "Tag_SNS_IM_RemarkTime" // 备注时间
    
    // 添加类型
    AddTypeSingle AddType = "Add_Type_Single" // 单向添加
    AddTypeBoth   AddType = "Add_Type_Both"   // 双向添加
    
    // 删除类型
    DeleteTypeSingle DeleteType = "Delete_Type_Single" // 单向删除
    DeleteTypeBoth   DeleteType = "Delete_Type_Both"   // 双向删除
    
    // 校验模式
    CheckTypeSingle CheckType = "CheckResult_Type_Single" // 单向校验好友关系
    CheckTypeBoth   CheckType = "CheckResult_Type_Both"   // 双向校验好友关系
    
    // 黑名单校验模式
    BlacklistCheckTypeSingle BlacklistCheckType = "BlackCheckResult_Type_Single" // 单向校验黑名单关系
    BlacklistCheckTypeBoth   BlacklistCheckType = "BlackCheckResult_Type_Both"   // 双向校验黑名单关系
    
    // 好友关系结果
    CheckResultTypeNoRelation = "CheckResult_Type_NoRelation" // From_Account 的好友表中没有 To_Account，但无法确定 To_Account 的好友表中是否有 From_Account
    CheckResultTypeAWithB     = "CheckResult_Type_AWithB"     // From_Account 的好友表中有 To_Account，但无法确定 To_Account 的好友表中是否有 From_Account
    CheckResultTypeBWithA     = "CheckResult_Type_BWithA"     // From_Account 的好友表中没有 To_Account，但 To_Account 的好友表中有 From_Account
    CheckResultTypeBothWay    = "CheckResult_Type_BothWay"    // From_Account 的好友表中有 To_Account，To_Account 的好友表中也有 From_Account
    
    // 黑名单关系结果
    BlackCheckResultTypeNO      = "BlackCheckResult_Type_NO"      // From_Account 的黑名单中没有 To_Account，但无法确定 To_Account 的黑名单中是否有 From_Account
    BlackCheckResultTypeAWithB  = "BlackCheckResult_Type_AWithB"  // From_Account 的黑名单中有 To_Account，但无法确定 To_Account 的黑名单中是否有 From_Account
    BlackCheckResultTypeBWithA  = "BlackCheckResult_Type_BWithA"  // From_Account 的黑名单中没有 To_Account，但 To_Account 的黑名单中有 From_Account
    BlackCheckResultTypeBothWay = "BlackCheckResult_Type_BothWay" // From_Account 的黑名单中有 To_Account，To_Account 的黑名单中也有 From_Account
)
