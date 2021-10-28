/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/27 11:31 上午
 * @Desc: TODO
 */

package core

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dobyte/http"

	"github.com/dobyte/tencent-im/internal/enum"
	"github.com/dobyte/tencent-im/internal/sign"
	"github.com/dobyte/tencent-im/internal/types"
)

const (
	defaultBaseUrl     = "https://console.tim.qq.com"
	defaultVersion     = "v4"
	defaultContentType = "json"
	defaultExpire      = 3600
)

var invalidResponse = errors.New("invalid response")

type Client interface {
	// Get send an http request use get method.
	Get(serviceName string, command string, data interface{}, resp interface{}) error
	// Post send an http request use post method.
	Post(serviceName string, command string, data interface{}, resp interface{}) error
	// Put send an http request use put method.
	Put(serviceName string, command string, data interface{}, resp interface{}) error
	// Patch send an http request use patch method.
	Patch(serviceName string, command string, data interface{}, resp interface{}) error
	// Delete send an http request use patch method.
	Delete(serviceName string, command string, data interface{}, resp interface{}) error
}

type client struct {
	client          *http.Client
	opt             *Options
	userSig         string
	userSigExpireAt int64
}

type Options struct {
	AppId     int    // 应用SDKAppID，可在即时通信 IM 控制台 的应用卡片中获取。
	AppSecret string // 密钥信息，可在即时通信 IM 控制台 的应用详情页面中获取，具体操作请参见 获取密钥
	UserId    string // 用户ID
	Expire    int    // UserSig过期时间
}

func NewClient(opt *Options) Client {
	rand.Seed(time.Now().UnixNano())
	c := new(client)
	c.opt = opt
	c.client = http.NewClient()
	c.client.SetContentType(http.ContentTypeJson)
	c.client.SetBaseUrl(defaultBaseUrl)

	return c
}

// Get send an http request use get method.
func (c *client) Get(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodGet, serviceName, command, data, resp)
}

// Post send an http request use post method.
func (c *client) Post(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodPost, serviceName, command, data, resp)
}

// Put send an http request use put method.
func (c *client) Put(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodPut, serviceName, command, data, resp)
}

// Patch send an http request use patch method.
func (c *client) Patch(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodPatch, serviceName, command, data, resp)
}

// Delete send an http request use patch method.
func (c *client) Delete(serviceName string, command string, data interface{}, resp interface{}) error {
	return c.request(http.MethodDelete, serviceName, command, data, resp)
}

// request send an http request.
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

// buildUrl build a request url.
func (c *client) buildUrl(serviceName string, command string) string {
	format := "/%s/%s/%s?sdkappid=%d&identifier=%s&usersig=%s&random=%d&contenttype=%s"
	random := rand.Int31()
	userSig := c.getUserSig()
	return fmt.Sprintf(format, defaultVersion, serviceName, command, c.opt.AppId, c.opt.UserId, userSig, random, defaultContentType)
}

// getUserSig get a userSig
func (c *client) getUserSig() string {
	now := time.Now()
	expire := c.opt.Expire

	if expire <= 0 {
		expire = defaultExpire
	}

	if c.userSig == "" || c.userSigExpireAt <= now.Unix() {
		c.userSig, _ = sign.GenUserSig(c.opt.AppId, c.opt.AppSecret, c.opt.UserId, expire)
		c.userSigExpireAt = now.Add(time.Duration(expire) * time.Second).Unix()
	}

	return c.userSig
}
