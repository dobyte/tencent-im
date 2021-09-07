/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/9/7 17:26
 * @Desc: TODO
 */

package group

const (
    GroupTypePublic     GroupType = "Public"     // Public（陌生人社交群）
    GroupTypePrivate    GroupType = "Private"    // Private（即 Work，好友工作群）
    GroupTypeChatRoom   GroupType = "ChatRoom"   // ChatRoom（即 Meeting，会议群）
    GroupTypeAVChatRoom GroupType = "AVChatRoom" // AVChatRoom（直播群）
    
    ApplyJoinOptionFreeAccess     ApplyJoinOption = "FreeAccess"     // 自由加入
    ApplyJoinOptionNeedPermission ApplyJoinOption = "NeedPermission" // 需要验证
    ApplyJoinOptionDisableApply   ApplyJoinOption = "DisableApply"   // 禁止加群
)
