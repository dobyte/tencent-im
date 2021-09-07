/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/28 1:14 上午
 * @Desc: TODO
 */

package enum

import (
    "github.com/dobyte/tencent-im/internal/types"
)

const (
    SuccessActionStatus = "OK"
    FailActionStatus    = "FAIL"
    SuccessCode         = 0
    
    // 消息类型
    MsgText     = "TIMTextElem"      // 消息元素
    MsgLocation = "TIMLocationElem"  // 地理位置消息元素
    MsgFace     = "TIMFaceElem"      // 表情消息元素
    MsgCustom   = "TIMCustomElem"    // 自定义消息元素
    MsgSound    = "TIMSoundElem"     // 语音消息元素
    MsgImage    = "TIMImageElem"     // 图像消息元素
    MsgFile     = "TIMFileElem"      // 文件消息元素
    MsgVideo    = "TIMVideoFileElem" // 视频消息元素
    
    // 图片格式
    ImageFormatJPG   = 1   // JPG格式
    ImageFormatGIF   = 2   // GIF格式
    ImageFormatPNG   = 3   // PNG格式
    ImageFormatBMP   = 4   // BMP格式
    ImageFormatOTHER = 255 // 其他格式
    
    // 图片类型
    ImageTypeOriginal = 1 // 原图
    ImageTypePic      = 2 // 大图
    ImageTypeThumb    = 3 // 缩略图
    
    // 标准资料字段
    StandardAttrNickname        = "Tag_Profile_IM_Nick"            // 昵称
    StandardAttrGender          = "Tag_Profile_IM_Gender"          // 性别
    StandardAttrBirthday        = "Tag_Profile_IM_BirthDay"        // 生日
    StandardAttrLocation        = "Tag_Profile_IM_Location"        // 所在地
    StandardAttrSignature       = "Tag_Profile_IM_SelfSignature"   // 个性签名
    StandardAttrAllowType       = "Tag_Profile_IM_AllowType"       // 加好友验证方式
    StandardAttrLanguage        = "Tag_Profile_IM_Language"        // 语言
    StandardAttrAvatar          = "Tag_Profile_IM_Image"           // 头像URL
    StandardAttrMsgSettings     = "Tag_Profile_IM_MsgSettings"     // 消息设置
    StandardAttrAdminForbidType = "Tag_Profile_IM_AdminForbidType" // 管理员禁止加好友标识
    StandardAttrLevel           = "Tag_Profile_IM_Level"           // 等级
    StandardAttrRole            = "Tag_Profile_IM_Role"            // 角色
    
    // 自定义属性前缀
    CustomAttrPrefix = "Tag_Profile_Custom" // 自定义属性前缀
    
    // 性别类型
    GenderTypeUnknown types.GenderType = "Gender_Type_Unknown" // 没设置性别
    GenderTypeFemale  types.GenderType = "Gender_Type_Female"  // 女性
    GenderTypeMale    types.GenderType = "Gender_Type_Male"    // 男性
    
    // 加好友验证方式
    AllowTypeNeedConfirm types.AllowType = "AllowType_Type_NeedConfirm" // 需要经过自己确认对方才能添加自己为好友
    AllowTypeAllowAny    types.AllowType = "AllowType_Type_AllowAny"    // 允许任何人添加自己为好友
    AllowTypeDenyAny     types.AllowType = "AllowType_Type_DenyAny"     // 不允许任何人添加自己为好友
    
    // 管理员禁止加好友标识类型
    AdminForbidTypeNone    types.AdminForbidType = "AdminForbid_Type_None"    // 默认值，允许加好友
    AdminForbidTypeSendOut types.AdminForbidType = "AdminForbid_Type_SendOut" // 禁止该用户发起加好友请求
    
    // 同步至其他设备
    SyncOtherMachineYes types.SyncOtherMachine = 1 // 把消息同步到From_Account在线终端和漫游上
    SyncOtherMachineNo  types.SyncOtherMachine = 2 // 消息不同步至From_Account
    
    // 推送标识
    PushFlagYes types.PushFlag = 0 // 正常推送
    PushFlagNo  types.PushFlag = 1 // 不离线推送
    
    // 华为推送通知消息分类
    HuaWeiImportanceLow    types.HuaWeiImportance = "LOW"    // LOW类消息
    HuaWeiImportanceNormal types.HuaWeiImportance = "NORMAL" // NORMAL类消息
    
    // 华为推送为“打开应用内指定页面”的前提下透传参数行为
    HuaweiIntentParamAction types.HuaweiIntentParam = 0 // 将透传内容Ext作为Action参数
    HuaweiIntentParamIntent types.HuaweiIntentParam = 1 // 将透传内容Ext作为Intent参数
    
    // VIVO手机推送消息分类
    VivoClassificationOperation types.VivoClassification = 0 // 运营类消息
    VivoClassificationSystem    types.VivoClassification = 1 // 系统类消息
    
    // IOS徽章计数模式
    BadgeModeNormal types.BadgeMode = 0 // 本条消息需要计数
    BadgeModeIgnore types.BadgeMode = 1 // 本条消息不需要计数
    
    // iOS10的推送扩展开关
    MutableContentNormal types.MutableContent = 0 // 关闭iOS10的推送扩展
    MutableContentEnable types.MutableContent = 1 // 开启iOS10的推送扩展
)
