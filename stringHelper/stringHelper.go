package stringHelper

import (
	"bytes"
	"golangutils/pconst"

	"regexp"
	"strconv"
	"strings"


)

// Concat concatenates string to buffer.
// According to 'Efficient String Concatenation in Go(http://herman.asia/efficient-string-concatenation-in-go)',
// bytes.Buffer is best choice for heavy-duty case.
// You should call buffer.String() to get a concatenated string after all concaternating finished.
func Concat(buffer *bytes.Buffer, str string) {
	buffer.WriteString(str)
}

// ConcatExist concatenates string to string array.
// According to 'Efficient String Concatenation in Go(http://herman.asia/efficient-string-concatenation-in-go)',
// When str is already exist, it's faster than buffer concatenation.
// You should call strings.Join(strs, "") to get a concatenated string after all concaternating finished.
func ConcatExist(strs []string, str string) []string {
	return append(strs, str)
}

func CheckValueInSliceUInt32(id uint32, s []uint32) bool {
	for _, v := range s {
		if id == v {
			return true
		}
	}
	return false
}

func CheckValueInSliceUInt64(id uint64, s []uint64) bool {
	for _, v := range s {
		if id == v {
			return true
		}
	}
	return false
}

func CheckValueInSliceString(str string, s []string) bool {
	for _, v := range s {
		if str == v {
			return true
		}
	}
	return false
}

func StrToUni(str string) string {
	rs := []rune(str)
	ret := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			ret += string(r)
		} else {
			ret += `\u` + strconv.FormatInt(int64(rint), 16)
		}
	}
	return ret
}

func CheckStringExisted(strs []string, str string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}

	return false
}

func CheckUInt64Existed(ids []uint64, id uint64) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}

	return false
}

func CheckStringLen(str string, mixLen, maxLen uint32) (bool, string) {
	status := true
	str = strings.Trim(str, " ")
	strLen := uint32(len(str))

	if strLen < mixLen && strLen > maxLen {
		status = false
	}

	return status, str
}

func StringToUInt16s(str string, space string) []uint16 {
	idsSlice := strings.Split(str, space)
	var Ints []uint16
	for _, v := range idsSlice {
		id, _ := strconv.ParseUint(v, 10, 16)
		nid := uint16(id)
		if nid == uint16(0) {
			Ints = []uint16{}
			break
		}
		Ints = append(Ints, nid)
	}

	return Ints
}


func GetPageNum(limitNum int, defaultNum int) int {
	if limitNum > pconst.COMMON_PAGE_LIMIT_NUM_MAX {
		return pconst.COMMON_PAGE_LIMIT_NUM_MAX
	} else if limitNum <= 0 {
		return defaultNum
	} else {
		return limitNum
	}
}

func CheckPicUrls(picUrls []string) bool {
	if len(picUrls) == 0 {
		return true
	}
	for _, v := range picUrls {
		if !CheckPicUrl(v) {
			return false
			break
		}
	}
	return true
}

//http://askpic-10003009.image.myqcloud.com/87a0e7a6-6053-4c15-8097-590e9eb3f44e
//http://dynamicpic-10003009.image.myqcloud.com/95f0dcc3-266b-4344-abae-d3e6815b3fc8
func CheckPicUrl(picUrl string) bool {
	b, err := regexp.MatchString(`^http://[\-\w]+\.image\.myqcloud\.com/[\-\w]{36}$`, picUrl)
	if err != nil {
		return false
	}
	return b
}
