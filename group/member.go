/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/9/7 18:06
 * @Desc: TODO
 */

package group

import "time"

type Member struct {
    userId   string // 成员ID
    role     string // 群内身份
    joinTime int64  // 加入时间
    nameCard string // 群名片
}

func NewMember() *Member {
    return &Member{}
}

// SetUserId 设置群成员ID
func (m *Member) SetUserId(userId string) {
    m.userId = userId
}

// GetUserId 获取群成员ID
func (m *Member) GetUserId() string {
    return m.userId
}

// SetRole 设置群内身份
func (m *Member) SetRole(role string) {
    m.role = role
}

// GetRole 获取群内身份
func (m *Member) GetRole() string {
    return m.role
}

// SetJoinTime 设置加入时间
func (m *Member) SetJoinTime(joinTime time.Time) {
    m.joinTime = joinTime.Unix()
}

// GetJoinTime 获取加入时间
func (m *Member) GetJoinTime() time.Time {
    return time.Unix(m.joinTime, 0)
}

// SetNameCard 设置群名片
func (m *Member) SetNameCard(nameCard string) {
    m.nameCard = nameCard
}

// GetNameCard 获取群名片
func (m *Member) GetNameCard() string {
    return m.nameCard
}

func (m *Member) CheckError() (err error) {
    return nil
}
