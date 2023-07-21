/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/27 11:31 上午
 * @Desc: TODO
 */

package core

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dobyte/http"

	"github.com/dobyte/tencent-im/internal/enum"
	"github.com/dobyte/tencent-im/internal/sign"
	"github.com/dobyte/tencent-im/internal/types"
)

const (
	defaultBaseUrl     = "https://adminapiger.im.qcloud.com"
	defaultVersion     = "v4"
	defaultContentType = "json"
	defaultExpiration  = 3600
)

var invalidResponse = NewError(enum.InvalidResponseCode, "invalid response")

type Client interface {
	// Get GET请求
	Get(serviceName string, command string, data interface{}, resp interface{}) error
	// Post POST请求
	Post(serviceName string, command string, data interface{}, resp interface{}) error
	// Put PUT请求
	Put(serviceName string, command string, data interface{}, resp interface{}) error
	// Patch PATCH请求
	Patch(serviceName string, command string, data interface{}, resp interface{}) error
	// Delete DELETE请求
	Delete(serviceName string, command string, data interface{}, resp interface{}) error
}

type client struct {
	client          *http.Client
	opt             *Options
	userSig         string
	userSigExpireAt int64
}

type Options struct {
	AppId      int    // 应用SDKAppID，可在即时通信 IM 控制台 的应用卡片中获取。
	AppSecret  string // 密钥信息，可在即时通信 IM 控制台 的应用详情页面中获取，具体操作请参见 获取密钥
	UserId     string // 用户ID
	Expiration int    // UserSig过期时间
	RequestUrl string // 请求地址, 因为不同的区域对应不同的请求地址
}

func NewClient(opt *Options) Client {
	rand.Seed(time.Now().UnixNano())
	c := new(client)
	c.opt = opt
	c.client = http.NewClient()
	c.client.SetContentType(http.ContentTypeJson)
	if opt.RequestUrl != "" {
		c.client.SetBaseUrl(opt.RequestUrl)
	} else {
		c.client.SetBaseUrl(defaultBaseUrl)
	}

	return c
}

// Get GET请求
func (c *client) Get(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodGet, serviceName, command, data, resp)
}

// Post POST请求
func (c *client) Post(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodPost, serviceName, command, data, resp)
}

// Put PUT请求
func (c *client) Put(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodPut, serviceName, command, data, resp)
}

// Patch PATCH请求
func (c *client) Patch(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodPatch, serviceName, command, data, resp)
}

// Delete DELETE请求
func (c *client) Delete(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodDelete, serviceName, command, data, resp)
}

// request Request请求
func (c *client) request(method, serviceName, command string, data, resp interface{}) error {
	res, err := c.client.Request(method, c.buildUrl(serviceName, command), data)
	if err != nil {
		return err
	}

	if err = res.Scan(resp); err != nil {
		return err
	}

	if r, ok := resp.(types.ActionBaseRespInterface); ok {
		if r.GetActionStatus() == enum.FailActionStatus {
			return NewError(r.GetErrorCode(), r.GetErrorInfo())
		}

		if r.GetErrorCode() != enum.SuccessCode {
			return NewError(r.GetErrorCode(), r.GetErrorInfo())
		}
	} else if r, ok := resp.(types.BaseRespInterface); ok {
		if r.GetErrorCode() != enum.SuccessCode {
			return NewError(r.GetErrorCode(), r.GetErrorInfo())
		}
	} else {
		return invalidResponse
	}

	return nil
}

// buildUrl 构建一个请求URL
func (c *client) buildUrl(serviceName string, command string) string {
	format := "/%s/%s/%s?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=%s"
	random := rand.Int31()
	userSig := c.getUserSig()
	return fmt.Sprintf(format, defaultVersion, serviceName, command, c.opt.AppId, c.opt.UserId, userSig, random, defaultContentType)
}

// getUserSig 获取签名
func (c *client) getUserSig() string {
	now, expiration := time.Now(), c.opt.Expiration

	if expiration <= 0 {
		expiration = defaultExpiration
	}

	if c.userSig == "" || c.userSigExpireAt <= now.Unix() {
		c.userSig, _ = sign.GenUserSig(c.opt.AppId, c.opt.AppSecret, c.opt.UserId, expiration)
		c.userSigExpireAt = now.Add(time.Duration(expiration) * time.Second).Unix()
	}

	return c.userSig
}
