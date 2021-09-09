/**
 * @Author: wanglin
 * @Author: wanglin@vspn.com
 * @Date: 2021/9/3 18:15
 * @Desc: 推送消息实体
 */

package push

import (
	"errors"

	"github.com/dobyte/tencent-im/internal/entity"
)

var errInvalidPushCondition = errors.New("attrs and tags condition cannot be set at the same time")

type Message struct {
	entity.Message
	condition *condition
}

func NewMessage() *Message {
	return &Message{}
}

// SetConditionTagsOr 设置标签的或条件（设置会冲掉之前的标签或条件）
func (m *Message) SetConditionTagsOr(tags ...string) {
	if m.condition != nil && m.condition.TagsOr != nil {
		m.condition.TagsOr = m.condition.TagsOr[0:0]
	}

	m.AddConditionTagsOr(tags...)
}

// AddConditionTagsOr 添加标签的或条件（添加会累加之前的条件或条件）
func (m *Message) AddConditionTagsOr(tags ...string) {
	if m.condition == nil {
		m.condition = &condition{}
	}
	if m.condition.TagsOr == nil {
		m.condition.TagsOr = make([]string, 0)
	}
	m.condition.TagsOr = append(m.condition.TagsOr, tags...)
}

// SetConditionTagsAnd 设置标签的与条件（设置会冲掉之前的标签与条件）
func (m *Message) SetConditionTagsAnd(tags ...string) {
	if m.condition != nil && m.condition.TagsAnd != nil {
		m.condition.TagsAnd = m.condition.TagsAnd[0:0]
	}

	m.AddConditionTagsAnd(tags...)
}

// AddConditionTagsAnd 添加标签的与条件（添加会累加之前的标签与条件）
func (m *Message) AddConditionTagsAnd(tags ...string) {
	if m.condition == nil {
		m.condition = &condition{}
	}
	if m.condition.TagsAnd == nil {
		m.condition.TagsAnd = make([]string, 0)
	}
	m.condition.TagsAnd = append(m.condition.TagsAnd, tags...)
}

// SetConditionAttrsOr 设置属性的或条件（设置会冲掉之前的属性或条件）
func (m *Message) SetConditionAttrsOr(attrs map[string]interface{}) {
	if m.condition != nil && m.condition.AttrsOr != nil {
		m.condition.AttrsOr = make(map[string]interface{})
	}

	m.AddConditionAttrsOr(attrs)
}

// AddConditionAttrsOr 添加属性的或条件（添加会累加之前的属性或条件）
func (m *Message) AddConditionAttrsOr(attrs map[string]interface{}) {
	if m.condition == nil {
		m.condition = &condition{}
	}
	if m.condition.AttrsOr == nil {
		m.condition.AttrsOr = make(map[string]interface{})
	}
	for k, v := range attrs {
		m.condition.AttrsOr[k] = v
	}
}

// SetConditionAttrsAnd 设置属性的与条件（设置会冲掉之前的属性与条件）
func (m *Message) SetConditionAttrsAnd(attrs map[string]interface{}) {
	if m.condition != nil && m.condition.AttrsAnd != nil {
		m.condition.AttrsAnd = make(map[string]interface{})
	}

	m.AddConditionAttrsAnd(attrs)
}

// AddConditionAttrsAnd 添加属性的与条件（添加会累加之前的属性与条件）
func (m *Message) AddConditionAttrsAnd(attrs map[string]interface{}) {
	if m.condition == nil {
		m.condition = &condition{}
	}
	if m.condition.AttrsAnd == nil {
		m.condition.AttrsAnd = make(map[string]interface{})
	}
	for k, v := range attrs {
		m.condition.AttrsAnd[k] = v
	}
}

// GetCondition 获取推送条件
func (m *Message) GetCondition() *condition {
	return m.condition
}

// checkError 检测错误
func (m *Message) checkError() (err error) {
	if err = m.CheckLifeTimeArgError(); err != nil {
		return
	}

	if err = m.CheckBodyArgError(); err != nil {
		return
	}

	if err = m.checkConditionArgError(); err != nil {
		return
	}

	return
}

// checkConditionArgError 检测条件参数错误
func (m *Message) checkConditionArgError() error {
	hasAttrs, hasTags := false, false

	if m.condition != nil {
		if m.condition.AttrsAnd != nil || m.condition.AttrsOr != nil {
			hasAttrs = true
		}

		if m.condition.TagsAnd != nil || m.condition.TagsOr != nil {
			hasTags = true
		}
	}

	if hasAttrs && hasTags {
		return errInvalidPushCondition
	}

	return nil
}
