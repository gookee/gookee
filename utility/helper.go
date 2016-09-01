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
	"time"
)

func JsonEncode(json string) string {
	json = strings.Replace(json, "\\", "\\\\", -1)
	json = strings.Replace(json, "\"", "\\\"", -1)
	json = strings.Replace(json, "\r\n", "\\u000d\\u000a", -1)
	json = strings.Replace(json, "\n", "\\u000a", -1)
	json = strings.Replace(json, "\t", "\\u0009", -1)
	return json
}

func JsonDecode(json string) string {
	json = strings.Replace(json, "\\\\", "\\", -1)
	json = strings.Replace(json, "\\\"", "\"", -1)
	json = strings.Replace(json, "\\u000d\\u000a", "\r\n", -1)
	json = strings.Replace(json, "\\u000a", "\n", -1)
	json = strings.Replace(json, "\\u0009", "\t", -1)
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
	case time.Time:
		tmp := obj.(time.Time).Format("2006-01-02 15:04:05")
		if tmp == "0001-01-01 00:00:00" {
			return ""
		} else {
			return tmp
		}
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

func SubString(str string, begin, length int) (substr string) {
	rs := []rune(str)
	lth := len(rs)

	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := 0
	if length == -1 {
		end = lth
	} else {
		end = begin + length
		if end > lth {
			end = lth
		}
	}

	return string(rs[begin:end])
}

func If(b bool, t, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}

func AjaxPageList(count, pageCount, takeNum, pageIndex int) string {
	html := ""

	if count > 0 {
		midhtml := ""
		s0 := "首页"
		s1 := "上页"
		s2 := "下页"
		s3 := "末页"
		s4 := "GO"
		pageButtonText := "共 {count} 条记录 {pageCount} 页 每页{takeNum}条 当前第 {pageIndex} 页"

		html = strings.Replace(pageButtonText, "{count}", ToStr(count), -1)
		html = strings.Replace(html, "{pageCount}", ToStr(pageCount), -1)
		html = strings.Replace(html, "{pageIndex}", ToStr(pageIndex), -1)
		html = strings.Replace(html, "{takeNum}",
			"<select onchange=\"bindData("+ToStr(pageIndex)+",$(this).val());\">"+
				"<option"+ToStr(If(takeNum == 5, " selected=\"selected\"", ""))+">5</option>"+
				"<option"+ToStr(If(takeNum == 10, " selected=\"selected\"", ""))+">10</option>"+
				"<option"+ToStr(If(takeNum == 15, " selected=\"selected\"", ""))+">15</option>"+
				"<option"+ToStr(If(takeNum == 20, " selected=\"selected\"", ""))+">20</option>"+
				"<option"+ToStr(If(takeNum == 25, " selected=\"selected\"", ""))+">25</option>"+
				"<option"+ToStr(If(takeNum == 30, " selected=\"selected\"", ""))+">30</option>"+
				"<option"+ToStr(If(takeNum == 35, " selected=\"selected\"", ""))+">35</option>"+
				"<option"+ToStr(If(takeNum == 40, " selected=\"selected\"", ""))+">40</option>"+
				"<option"+ToStr(If(takeNum == 45, " selected=\"selected\"", ""))+">45</option>"+
				"<option"+ToStr(If(takeNum == 50, " selected=\"selected\"", ""))+">50</option>"+
				"<option"+ToStr(If(takeNum == 100, " selected=\"selected\"", ""))+">100</option>"+
				"<option"+ToStr(If(takeNum == 200, " selected=\"selected\"", ""))+">200</option>"+
				"<option"+ToStr(If(takeNum == 500, " selected=\"selected\"", ""))+">500</option>"+
				"</select>", -1)
		PageNum := 5
		k := PageNum / 2
		j := pageIndex - k
		if j+PageNum > pageCount {
			j = pageCount - PageNum + 1
		}
		if j <= 0 {
			j = 1
		}
		for i := j; i < PageNum+j; i++ {
			if i > pageCount {
				break
			}
			midhtml += " <a href=\"javascript:void(0)\" onclick=\"bindData(" + ToStr(i) + "," + ToStr(takeNum) + ");\" style=\"padding:0px 3px;\"" + ToStr(If(pageIndex == i, " disabled", "")) + ">" + ToStr(i) + "</a> "
		}

		html += "<span><a href=\"javascript:void(0)\" onclick=\"bindData(1," +
			ToStr(takeNum) +
			");\" style=\"padding:0px 3px;\"" +
			ToStr(If(pageIndex <= 1, " disabled", "")) +
			">" +
			s0 +
			"</a> <a href=\"javascript:void(0)\" onclick=\"bindData(" +
			ToStr(If(pageIndex <= 1, 1, pageIndex-1)) +
			"," +
			ToStr(takeNum) +
			");\" style=\"padding:0px 3px;\"" +
			ToStr(If(pageIndex <= 1, " disabled", "")) +
			">" +
			s1 +
			"</a> " +
			midhtml +
			" <a href=\"javascript:void(0)\" onclick=\"bindData(" +
			ToStr(If(pageIndex >= pageCount, pageIndex, pageIndex+1)) +
			"," +
			ToStr(takeNum) +
			");\" style=\"padding:0px 3px;\"" +
			ToStr(If(pageIndex >= pageCount, " disabled", "")) +
			">" +
			s2 +
			"</a> <a href=\"javascript:void(0)\" onclick=\"bindData(" +
			ToStr(pageCount) +
			"," +
			ToStr(takeNum) +
			");\" style=\"padding:0px 3px;\"" +
			ToStr(If(pageIndex >= pageCount, " disabled", "")) +
			">" +
			s3 +
			"</a> <input type=\"text\" style=\"margin:0px 3px; font-size:12px; width:30px; height:15px;\" attr=\"num\" value=\"" +
			ToStr(pageIndex) +
			"\"> <a href=\"javascript:void(0)\" onclick=\"bindData($(this).parent().find('input').first().val()," +
			ToStr(takeNum) + ");\">" + s4 + "</a></span>"
	}

	return JsonEncode(html)
}
