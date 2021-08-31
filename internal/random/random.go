/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/5/12 10:35
 * @Desc: 随机数类库
 */

package random

import (
	"math/rand"
	"time"
)

const (
	AlphaStr         = iota // 字母字
	AlphaLowerStr           // 小写字母
	AlphaUpperStr           // 大写字母
	NumericStr              // 数字
	NoZeroNumericStr        // 无0数字
)

// GenStr 生成指定长度的字符串
func GenStr(mode, length int) string {
	var (
		pos     int
		lastStr string
		seedStr string
	)

	rand.Seed(time.Now().UnixNano())
	switch mode {
	case AlphaStr:
		seedStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case AlphaLowerStr:
		seedStr = "abcdefghijklmnopqrstuvwxyz"
	case AlphaUpperStr:
		seedStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case NumericStr:
		seedStr = "0123456789"
	case NoZeroNumericStr:
		seedStr = "123456789"
	}

	seedLen := len(seedStr)
	for i := 0; i < length; i++ {
		pos = rand.Intn(seedLen)
		lastStr += seedStr[pos : pos+1]
	}

	return lastStr
}

// GenNumeric 生成指定范围的数字
func GenNumeric(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	if min < max {
		return rand.Intn(max-min) + min
	} else {
		return rand.Intn(min-max) + max
	}
}
