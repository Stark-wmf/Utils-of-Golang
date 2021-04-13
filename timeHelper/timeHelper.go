package timeHelper

import (
	"golangutils/pconst"
	"math"
	"strconv"
	"time"
)

func FewDaysLater(day int) time.Time {
	return FewDurationLater(time.Duration(day) * 24 * time.Hour)
}

func TwentyFourHoursLater() time.Time {
	return FewDurationLater(time.Duration(24) * time.Hour)
}

func SixHoursLater() time.Time {
	return FewDurationLater(time.Duration(6) * time.Hour)
}
// Add 返回的=原来时间+duration时间
func FewDurationLater(duration time.Duration) time.Time {
	// When Save time should considering UTC
	baseTime := time.Now()
	fewDurationLater := baseTime.Add(duration)
	return fewDurationLater
}

func IsExpired(expirationTime time.Time) bool {
	after := time.Now().After(expirationTime)
	return after
}

//几种日期格式
func GetCommonDateYdmHis(t time.Time) string {
	return t.Format(pconst.TIME_FORMAT_Y_M_D_H_I_S)
}

func GetCommonDateYdmHis2(t time.Time) string {
	return t.Format(pconst.TIME_FORMAT_Y_M_D_H_I_S_2)
}

func GetCommonDateYdm(t time.Time) string {
	return t.Format(pconst.TIME_FORMAT_Y_M_D)
}

func GetCommonDateYdm_(t time.Time) string {
	return t.Format(pconst.TIME_FORMAT_Y_M_D_)
}

func GetMonthUnix() (result int64) {
	now := time.Now()
	year := now.Year()
	month := now.Month().String()
	nowMonth := strconv.Itoa(year) + "-" + month + "-01"
	location, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(pconst.TIME_FORMAT_Y_MS_D_, nowMonth, location)
	if err != nil {
		return result
	} else {
		return t.Unix()
	}
}

func GetPRCTodayTime() (time.Time, error) {
	todayStr := GetCommonDateYdm_(time.Now())
	location, _ := time.LoadLocation("Asia/Shanghai")
	dateTime, err := time.ParseInLocation(pconst.TIME_FORMAT_Y_M_D_, todayStr, location)
	return dateTime, err
}
// 判断是多少时间之前，时间大，可能很多天
func GetBigFormatTime(ts time.Time) (ret string) {
	t := time.Now().Unix() - ts.Unix()
	if t < pconst.TIME_ONE_MINUTE {
		ret = "刚刚"
	} else if t < pconst.TIME_ONE_HOUR {
		minute := int(math.Floor(float64(t) / float64(pconst.TIME_ONE_MINUTE)))
		minuteStr := strconv.Itoa(minute)
		ret = minuteStr + "分钟前"
	} else if t < pconst.TIME_ONE_DAY {
		hour := int(math.Floor(float64(t) / float64(pconst.TIME_ONE_HOUR)))
		hourStr := strconv.Itoa(hour)
		ret = hourStr + "小时前"
	} else if t < pconst.TIME_ONE_WEEK {
		day := int(math.Floor(float64(t) / float64(pconst.TIME_ONE_DAY)))
		dayStr := strconv.Itoa(day)
		ret = dayStr + "天前"
	} else {
		ret = GetCommonDateYdm(ts)
	}

	return ret
}
// 判断是多少时间之前，时间小，分钟计时
func GetMiddleFormatTime(ts time.Time) (ret string) {
	t := time.Now().Unix() - ts.Unix()
	if t < pconst.TIME_ONE_MINUTE {
		ret = "刚刚"
	} else if t < pconst.TIME_ONE_HOUR {
		minute := int(math.Floor(float64(t) / float64(pconst.TIME_ONE_MINUTE)))
		minuteStr := strconv.Itoa(minute)
		ret = minuteStr + "分钟前"
	} else {
		ret = ts.Format(pconst.TIME_FORMAT_M_D_H_I)
	}

	return ret
}
//剩多少时间，这里用uint32的原因还不太清楚，可能是负数？
func GetCocEventLeftTime(ts uint32) (ret string) {
	if ts == 0 {
		return ""
	}
	left := int64(ts) - time.Now().Unix()
	if left >= pconst.TIME_ONE_DAY {
		day := int(math.Floor(float64(left) / float64(pconst.TIME_ONE_DAY)))
		dayStr := strconv.Itoa(day)
		ret = "剩" + dayStr + "天"
	} else if left >= pconst.TIME_ONE_HOUR {
		hour := int(math.Floor(float64(left) / float64(pconst.TIME_ONE_HOUR)))
		hourStr := strconv.Itoa(hour)
		ret = "剩" + hourStr + "小时"
	} else if left > 0 {
		minute := int(math.Floor(float64(left) / float64(pconst.TIME_ONE_MINUTE)))
		second := int(left % pconst.TIME_ONE_MINUTE)
		minuteStr := strconv.Itoa(minute)
		secondStr := strconv.Itoa(second)
		ret = "剩" + minuteStr + "分" + secondStr + "秒"
	} else {
		ret = "活动结束"
	}

	return ret
}

func GetCocEventLeftTimeSeconds(ts uint32) (ret uint32) {
	if ts == 0 {
		return 0
	}
	left := int64(ts) - time.Now().Unix()
	if left <= 0 {
		left = 0
	}

	return uint32(left)
}
