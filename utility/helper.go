package utility

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func JsonEncode(json string) string {
	json = strings.Replace(json, "'", "\\'", -1)
	json = strings.Replace(json, "\"", "\\\"", -1)
	json = strings.Replace(json, "\r\n", "\\u000d\\u000a", -1)
	json = strings.Replace(json, "\n", "\\u000a", -1)
	return json
}

func JsonDecode(json string) string {
	json = strings.Replace(json, "\\'", "'", -1)
	json = strings.Replace(json, "\\\"", "\"", -1)
	json = strings.Replace(json, "\\u000d\\u000a", "\r\n", -1)
	json = strings.Replace(json, "\\u000a", "\n", -1)
	return json
}

func ToStr(obj interface{}) string {
	switch obj.(type) {
	case string:
		return obj.(string)
	case int:
		return strconv.Itoa(obj.(int))
	case int64:
		return strconv.FormatInt(obj.(int64), 10)
	case float32:
		return strconv.FormatFloat(float64(obj.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(obj.(float64), 'f', -1, 64)
	case []byte:
		return string(obj.([]byte))
	}
	return ""
}

func ToInt(obj interface{}) int {
	switch obj.(type) {
	case string:
		i, err := strconv.Atoi(obj.(string))
		if err == nil {
			return i
		} else {
			return 0
		}
	case int:
		return obj.(int)
	case int64:
		return int(obj.(int64))
	case float32:
		return int(float64(obj.(float32)))
	case float64:
		return int(obj.(float64))
	}
	return 0
}

func MD5(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func SHA1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func NewGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return MD5(base64.URLEncoding.EncodeToString(b))
}

func GetRequestStr(r *http.Request, key string) string {
	return FilterSQL(r.FormValue(key))
}

func GetRequestInt(r *http.Request, key string) int {
	return ToInt(r.FormValue(key))
}

//防sql注入处理
func FilterSQL(sql string) string {
	return regexp.MustCompile("(?i)(and|or|exec|insert|select|delete|update|chr|truncate|char|declare|join|mid|cmd|xp_|sp_|0x|\"|;|@|%|#|&|<|>)").ReplaceAllStringFunc(sql, ToSBC)
}

//半角转全角
func ToSBC(input string) string {
	rs := []rune(input)
	str := ""
	for _, r := range rs {
		if r == 32 {
			r = 12288
			continue
		}
		if r < 127 {
			r = r + 65248
		}
		str += string(r)
	}

	return string(str)
}

//全角转半角
func ToDBC(input string) string {
	rs := []rune(input)
	str := ""
	for _, r := range rs {
		if r == 12288 {
			r = 32
			continue
		}
		if r > 65280 && r < 65375 {
			r = r - 65248
		}
		str += string(r)
	}

	return string(str)
}

func NoHTML(htmlstring string) string {
	if htmlstring == "" {
		return ""
	}

	return regexp.MustCompile("(?i)(\\r\\n|<script.*?</script>|<style.*?</style>|<.*?>|<(.[^>]*)>|[\\s]+|-->|<!--.*|&(nbsp|#160);|&#(\\d+);|<|>)").ReplaceAllString(htmlstring, "")
}
