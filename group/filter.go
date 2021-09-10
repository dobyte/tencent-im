/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/9/8 10:17
 * @Desc: 群组响应过滤器
 */

package group

type (
    // BaseInfoField 群基础信息字段
    BaseInfoField string
    
    // MemberInfoField 群成员信息字段
    MemberInfoField string
)

const (
    BaseFieldGroupId         BaseInfoField = "GroupId"         // 群组的唯一标识
    BaseFieldType            BaseInfoField = "Type"            // 群组类型
    BaseFieldName            BaseInfoField = "Name"            // 群组名称
    BaseFieldIntroduction    BaseInfoField = "Introduction"    // 群组简介
    BaseFieldNotification    BaseInfoField = "Notification"    // 群组公告
    BaseFieldAvatar          BaseInfoField = "FaceUrl"         // 群组头像URL
    BaseFieldOwner           BaseInfoField = "Owner_Account"   // 群主ID
    BaseFieldCreateTime      BaseInfoField = "CreateTime"      // 群组的创建时间
    BaseFieldInfoSeq         BaseInfoField = "InfoSeq"         // 群资料变更次数
    BaseFieldLastInfoTime    BaseInfoField = "LastInfoTime"    // 群组最后一次信息变更时间
    BaseFieldLastMsgTime     BaseInfoField = "LastMsgTime"     // 群组内最后发消息的时间
    BaseFieldNextMsgSeq      BaseInfoField = "NextMsgSeq"      // 群内下一条消息的Seq
    BaseFieldMemberNum       BaseInfoField = "MemberNum"       // 当前成员数量
    BaseFieldMaxMemberNum    BaseInfoField = "MaxMemberNum"    // 最大成员数量
    BaseFieldApplyJoinOption BaseInfoField = "ApplyJoinOption" // 申请加群选项
    
    MemberFieldUserId          MemberInfoField = "Member_Account"  // 群成员ID
    MemberFieldRole            MemberInfoField = "Role"            // 群内身份
    MemberFieldJoinTime        MemberInfoField = "JoinTime"        // 入群时间
    MemberFieldMsgSeq          MemberInfoField = "MsgSeq"          // 该成员当前已读消息Seq
    MemberFieldMsgFlag         MemberInfoField = "MsgFlag"         // 消息接收选项
    MemberFieldLastSendMsgTime MemberInfoField = "LastSendMsgTime" // 最后发送消息的时间
    MemberFieldNameCard        MemberInfoField = "NameCard"        // 群名片
)

type Filter struct {
    baseInfo         map[string]bool
    memberInfo       map[string]bool
    memberRole       map[string]bool
    groupCustomData  map[string]bool
    memberCustomData map[string]bool
}

// AddBaseInfoFilter 添加基础信息过滤器
func (f *Filter) AddBaseInfoFilter(field BaseInfoField) {
    if f.baseInfo == nil {
        f.baseInfo = make(map[string]bool)
    }
    
    f.baseInfo[string(field)] = true
}

// RemBaseInfoFilter 移除基础信息过滤器
func (f *Filter) RemBaseInfoFilter(field BaseInfoField) {
    if f.baseInfo == nil {
        return
    }
    
    delete(f.baseInfo, string(field))
}

// GetAllBaseInfoFilterFields 获取所有基础信息过滤器字段
func (f *Filter) GetAllBaseInfoFilterFields() (filters []string) {
    if f.baseInfo == nil {
        return
    }
    
    filters = make([]string, 0, len(f.baseInfo))
    for k, _ := range f.baseInfo {
        filters = append(filters, k)
    }
    
    return
}

// AddMemberInfoFilter 添加成员信息过滤器
func (f *Filter) AddMemberInfoFilter(field MemberInfoField) {
    if f.memberInfo == nil {
        f.memberInfo = make(map[string]bool)
    }
    
    f.memberInfo[string(field)] = true
}

// RemMemberInfoFilter 移除成员信息过滤器
func (f *Filter) RemMemberInfoFilter(field MemberInfoField) {
    if f.memberInfo == nil {
        return
    }
    
    delete(f.memberInfo, string(field))
}

// GetAllMemberInfoFilterFields 获取所有成员信息过滤器字段
func (f *Filter) GetAllMemberInfoFilterFields() (filters []string) {
    if f.memberInfo == nil {
        return
    }
    
    filters = make([]string, 0, len(f.memberInfo))
    for k, _ := range f.memberInfo {
        filters = append(filters, k)
    }
    
    return
}

// AddMemberRoleFilter 添加群成员角色过滤器
func (f *Filter) AddMemberRoleFilter(field string) {
    if f.memberRole == nil {
        f.memberRole = make(map[string]bool)
    }
    
    f.memberRole[field] = true
}

// RemMemberRoleFilter 移除群成员角色过滤器
func (f *Filter) RemMemberRoleFilter(field string) {
    if f.memberRole == nil {
        return
    }
    
    delete(f.memberRole, field)
}

// GetAllMemberRoleFilterValues 获取所有群成员角色过滤器值
func (f *Filter) GetAllMemberRoleFilterValues() (filters []string) {
    if f.memberRole == nil {
        return
    }
    
    filters = make([]string, 0, len(f.memberRole))
    for k, _ := range f.memberRole {
        filters = append(filters, k)
    }
    
    return
}

// AddGroupCustomDataFilter 添加群自定义数据过滤器
func (f *Filter) AddGroupCustomDataFilter(field string) {
    if f.groupCustomData == nil {
        f.groupCustomData = make(map[string]bool)
    }
    
    f.groupCustomData[field] = true
}

// RemGroupCustomDataFilter 移除群自定义数据过滤器
func (f *Filter) RemGroupCustomDataFilter(field string) {
    if f.groupCustomData == nil {
        return
    }
    
    delete(f.groupCustomData, field)
}

// GetAllGroupCustomDataFilterFields 获取所有群自定义数据过滤器字段
func (f *Filter) GetAllGroupCustomDataFilterFields() (filters []string) {
    if f.groupCustomData == nil {
        return
    }
    
    filters = make([]string, 0, len(f.groupCustomData))
    for k, _ := range f.groupCustomData {
        filters = append(filters, k)
    }
    
    return
}

// AddMemberCustomDataFilter 添加群成员自定义数据过滤器
func (f *Filter) AddMemberCustomDataFilter(field string) {
    if f.memberCustomData == nil {
        f.memberCustomData = make(map[string]bool)
    }
    
    f.memberCustomData[field] = true
}

// RemMemberCustomDataFilter 移除群成员自定义数据过滤器
func (f *Filter) RemMemberCustomDataFilter(field string) {
    if f.memberCustomData == nil {
        return
    }
    
    delete(f.memberCustomData, field)
}

// GetAllMemberCustomDataFilterFields 获取所有群成员自定义数据过滤器字段
func (f *Filter) GetAllMemberCustomDataFilterFields() (filters []string) {
    if f.memberCustomData == nil {
        return
    }
    
    filters = make([]string, 0, len(f.memberCustomData))
    for k, _ := range f.memberCustomData {
        filters = append(filters, k)
    }
    
    return
}
