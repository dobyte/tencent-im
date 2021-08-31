/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/27 1:12 下午
 * @Desc: TODO
 */

package core

type Error interface {
	error
	Code() int
	Message() string
}

type respError struct {
	code    int
	message string
}

func NewError(code int, message string) Error {
	return &respError{
		code:    code,
		message: message,
	}
}

func (e *respError) Error() string {
	return e.message
}

func (e *respError) Code() int {
	return e.code
}

func (e *respError) Message() string {
	return e.message
}
