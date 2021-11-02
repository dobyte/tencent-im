/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/9/7 10:28
 * @Desc: 运营管理枚举参数
 */

package operation

type (
	// ChatType 聊天类型
	ChatType string

	// FieldType 运营数据字段类型
	FieldType string
)

const (
	ChatTypeC2C   ChatType = "C2C"   // 单聊消息
	ChatTypeGroup ChatType = "Group" // 群聊消息

	FieldTypeAppName               FieldType = "AppName"              // 应用名称
	FieldTypeAppId                 FieldType = "AppId"                // 应用 SDKAppID
	FieldTypeCompany               FieldType = "Company"              // 所属客户名称
	FieldTypeActiveUserNum         FieldType = "ActiveUserNum"        // 活跃用户数
	FieldTypeRegisterUserNumOneDay FieldType = "RegistUserNumOneDay"  // 新增注册人数
	FieldTypeRegisterUserNumTotal  FieldType = "RegistUserNumTotal"   // 累计注册人数
	FieldTypeLoginTimes            FieldType = "LoginTimes"           // 登录次数
	FieldTypeLoginUserNum          FieldType = "LoginUserNum"         // 登录人数
	FieldTypeUpMsgNum              FieldType = "UpMsgNum"             // 上行消息数
	FieldTypeSendMsgUserNum        FieldType = "SendMsgUserNum"       // 发消息人数
	FieldTypeAPNSMsgNum            FieldType = "APNSMsgNum"           // APNs 推送数
	FieldTypeC2CUpMsgNum           FieldType = "C2CUpMsgNum"          // 上行消息数（C2C）
	FieldTypeC2CSendMsgUserNum     FieldType = "C2CSendMsgUserNum"    // 发消息人数（C2C）
	FieldTypeC2CAPNSMsgNum         FieldType = "C2CAPNSMsgNum"        // APNs 推送数（C2C）
	FieldTypeMaxOnlineNum          FieldType = "MaxOnlineNum"         // 最高在线人数
	FieldTypeChainIncrease         FieldType = "ChainIncrease"        // 关系链对数增加量
	FieldTypeChainDecrease         FieldType = "ChainDecrease"        // 关系链对数删除量
	FieldTypeGroupUpMsgNum         FieldType = "GroupUpMsgNum"        // 上行消息数（群）
	FieldTypeGroupSendMsgUserNum   FieldType = "GroupSendMsgUserNum"  // 发消息人数（群）
	FieldTypeGroupAPNSMsgNum       FieldType = "GroupAPNSMsgNum"      // APNs 推送数（群）
	FieldTypeGroupSendMsgGroupNum  FieldType = "GroupSendMsgGroupNum" // 发消息群组数
	FieldTypeGroupJoinGroupTimes   FieldType = "GroupJoinGroupTimes"  // 入群总数
	FieldTypeGroupQuitGroupTimes   FieldType = "GroupQuitGroupTimes"  // 退群总数
	FieldTypeGroupNewGroupNum      FieldType = "GroupNewGroupNum"     // 新增群组数
	FieldTypeGroupAllGroupNum      FieldType = "GroupAllGroupNum"     // 累计群组数
	FieldTypeGroupDestroyGroupNum  FieldType = "GroupDestroyGroupNum" // 解散群个数
	FieldTypeCallBackReq           FieldType = "CallBackReq"          // 回调请求数
	FieldTypeCallBackRsp           FieldType = "CallBackRsp"          // 回调应答数
	FieldTypeDate                  FieldType = "Date"                 // 日期
)
