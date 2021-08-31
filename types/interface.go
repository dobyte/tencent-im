/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/27 12:54 下午
 * @Desc: TODO
 */

package types

type BaseRespInterface interface {
	GetErrorCode() int
	GetErrorInfo() string
}

func (r *BaseResp) GetErrorCode() int {
	return r.ErrorCode
}

func (r *BaseResp) GetErrorInfo() string {
	return r.ErrorInfo
}

type ActionBaseRespInterface interface {
	BaseRespInterface
	GetActionStatus() string
}

func (r *ActionBaseResp) GetActionStatus() string {
	return r.ActionStatus
}