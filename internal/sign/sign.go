/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 19:11
 * @Desc: Transmission signature.
 */

package sign

import (
	"bytes"
	"compress/zlib"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"time"
)

// GenUserSig gen a user sign.
func GenUserSig(sdkAppId int, key string, userid string, expire int) (string, error) {
	return genUserSig(sdkAppId, key, userid, expire, nil)
}

// GenPrivateMapKey gen a private map.
func GenPrivateMapKey(sdkAppId int, key string, userid string, expire int, roomId uint32, privilegeMap uint32) (string, error) {
	var userBuf []byte = genUserBuf(userid, sdkAppId, roomId, expire, privilegeMap, 0, "")
	return genUserSig(sdkAppId, key, userid, expire, userBuf)
}

// GenPrivateMapKeyWithRoomId gen a private map with room id.
func GenPrivateMapKeyWithRoomId(sdkAppId int, key string, userid string, expire int, roomId string, privilegeMap uint32) (string, error) {
	var userBuf []byte = genUserBuf(userid, sdkAppId, 0, expire, privilegeMap, 0, roomId)
	return genUserSig(sdkAppId, key, userid, expire, userBuf)
}

// genUserBuf gen a user buffer.
func genUserBuf(account string, dwSdkappid int, dwAuthID uint32,
	dwExpTime int, dwPrivilegeMap uint32, dwAccountType uint32, roomStr string) []byte {

	offset := 0
	length := 1 + 2 + len(account) + 20 + len(roomStr)
	if len(roomStr) > 0 {
		length = length + 2
	}

	userBuf := make([]byte, length)

	// ver
	if len(roomStr) > 0 {
		userBuf[offset] = 1
	} else {
		userBuf[offset] = 0
	}

	offset++
	userBuf[offset] = (byte)((len(account) & 0xFF00) >> 8)
	offset++
	userBuf[offset] = (byte)(len(account) & 0x00FF)
	offset++

	for ; offset < len(account)+3; offset++ {
		userBuf[offset] = account[offset-3]
	}

	// dwSdkAppid
	userBuf[offset] = (byte)((int64(dwSdkappid) & 0xFF000000) >> 24)
	offset++
	userBuf[offset] = (byte)((dwSdkappid & 0x00FF0000) >> 16)
	offset++
	userBuf[offset] = (byte)((dwSdkappid & 0x0000FF00) >> 8)
	offset++
	userBuf[offset] = (byte)(dwSdkappid & 0x000000FF)
	offset++

	// dwAuthId
	userBuf[offset] = (byte)((dwAuthID & 0xFF000000) >> 24)
	offset++
	userBuf[offset] = (byte)((dwAuthID & 0x00FF0000) >> 16)
	offset++
	userBuf[offset] = (byte)((dwAuthID & 0x0000FF00) >> 8)
	offset++
	userBuf[offset] = (byte)(dwAuthID & 0x000000FF)
	offset++

	// dwExpTime now+300;
	currTime := time.Now().Unix()
	var expire = currTime + int64(dwExpTime)
	userBuf[offset] = (byte)((expire & 0xFF000000) >> 24)
	offset++
	userBuf[offset] = (byte)((expire & 0x00FF0000) >> 16)
	offset++
	userBuf[offset] = (byte)((expire & 0x0000FF00) >> 8)
	offset++
	userBuf[offset] = (byte)(expire & 0x000000FF)
	offset++

	// dwPrivilegeMap
	userBuf[offset] = (byte)((dwPrivilegeMap & 0xFF000000) >> 24)
	offset++
	userBuf[offset] = (byte)((dwPrivilegeMap & 0x00FF0000) >> 16)
	offset++
	userBuf[offset] = (byte)((dwPrivilegeMap & 0x0000FF00) >> 8)
	offset++
	userBuf[offset] = (byte)(dwPrivilegeMap & 0x000000FF)
	offset++

	// dwAccountType
	userBuf[offset] = (byte)((dwAccountType & 0xFF000000) >> 24)
	offset++
	userBuf[offset] = (byte)((dwAccountType & 0x00FF0000) >> 16)
	offset++
	userBuf[offset] = (byte)((dwAccountType & 0x0000FF00) >> 8)
	offset++
	userBuf[offset] = (byte)(dwAccountType & 0x000000FF)
	offset++

	if len(roomStr) > 0 {
		userBuf[offset] = (byte)((len(roomStr) & 0xFF00) >> 8)
		offset++
		userBuf[offset] = (byte)(len(roomStr) & 0x00FF)
		offset++

		for ; offset < length; offset++ {
			userBuf[offset] = roomStr[offset-(length-len(roomStr))]
		}
	}

	return userBuf
}

// hmacSha256 encrypt with HMAC SHA256.
func hmacSha256(sdkAppId int, key string, identifier string, currTime int64, expire int, base64UserBuf *string) string {
	var contentToBeSigned string
	contentToBeSigned = "TLS.identifier:" + identifier + "\n"
	contentToBeSigned += "TLS.sdkappid:" + strconv.Itoa(sdkAppId) + "\n"
	contentToBeSigned += "TLS.time:" + strconv.FormatInt(currTime, 10) + "\n"
	contentToBeSigned += "TLS.expire:" + strconv.Itoa(expire) + "\n"
	if nil != base64UserBuf {
		contentToBeSigned += "TLS.userbuf:" + *base64UserBuf + "\n"
	}

	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(contentToBeSigned))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// genUserSig gen a sign
func genUserSig(sdkAppId int, key string, identifier string, expire int, userBuf []byte) (string, error) {
	currTime := time.Now().Unix()
	var sigDoc map[string]interface{}
	sigDoc = make(map[string]interface{})
	sigDoc["TLS.ver"] = "2.0"
	sigDoc["TLS.identifier"] = identifier
	sigDoc["TLS.sdkappid"] = sdkAppId
	sigDoc["TLS.expire"] = expire
	sigDoc["TLS.time"] = currTime
	var base64UserBuf string
	if nil != userBuf {
		base64UserBuf = base64.StdEncoding.EncodeToString(userBuf)
		sigDoc["TLS.userbuf"] = base64UserBuf
		sigDoc["TLS.sig"] = hmacSha256(sdkAppId, key, identifier, currTime, expire, &base64UserBuf)
	} else {
		sigDoc["TLS.sig"] = hmacSha256(sdkAppId, key, identifier, currTime, expire, nil)
	}

	data, err := json.Marshal(sigDoc)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, _ = w.Write(data)
	_ = w.Close()
	return base64Encode(b.Bytes()), nil
}
