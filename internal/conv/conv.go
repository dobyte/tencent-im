/**
 * @Author: fuxiao
 * @Email: 576101059@qq.com
 * @Date: 2021/8/31 11:16 下午
 * @Desc: TODO
 */

package conv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type stringInterface interface {
	String() string
}

type errorInterface interface {
	Error() string
}

func String(any interface{}) string {
	switch v := any.(type) {
	case nil:
		return ""
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.Itoa(int(v))
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case []byte:
		return string(v)
	case time.Time:
		return v.String()
	case *time.Time:
		if v == nil {
			return ""
		}
		return v.String()
	default:
		if v == nil {
			return ""
		}
		
		if i, ok := v.(stringInterface); ok {
			return i.String()
		}
		
		if i, ok := v.(errorInterface); ok {
			return i.Error()
		}
		
		var (
			rv   = reflect.ValueOf(v)
			kind = rv.Kind()
		)
		
		switch kind {
		case reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return ""
			}
		case reflect.String:
			return rv.String()
		}
		
		if kind == reflect.Ptr {
			return String(rv.Elem().Interface())
		}
		
		if b, e := json.Marshal(v); e != nil {
			return fmt.Sprint(v)
		} else {
			return string(b)
		}
	}
}
