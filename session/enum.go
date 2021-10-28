/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/10/28 17:22
 * @Desc: TODO
 */

package session

// SessionType 会话类型
type SessionType int

const (
	SessionTypeC2C SessionType = 1 // C2C 会话
	SessionTypeG2C SessionType = 2 // G2C 会话
)
