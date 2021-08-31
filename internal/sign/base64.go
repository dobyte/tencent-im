/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/27 19:15
 * @Desc: BASE64
 */

package sign

import (
    "encoding/base64"
    "strings"
)

// base64Encode base64 encode a string.
func base64Encode(data []byte) string {
    str := base64.StdEncoding.EncodeToString(data)
    str = strings.Replace(str, "+", "*", -1)
    str = strings.Replace(str, "/", "-", -1)
    str = strings.Replace(str, "=", "_", -1)
    return str
}

// base64Decode base64 decode a string.
func base64Decode(str string) ([]byte, error) {
    str = strings.Replace(str, "_", "=", -1)
    str = strings.Replace(str, "-", "/", -1)
    str = strings.Replace(str, "*", "+", -1)
    return base64.StdEncoding.DecodeString(str)
}
