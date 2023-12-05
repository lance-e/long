package timer

import "time"

// GetNowTime 获取当前时间
func GetNowTime() time.Time {
	return time.Now()
}

// GetCalculateTime 计算时间
func GetCalculateTime(currentTimer time.Time, duration string) (time.Time, error) {
	d, err := time.ParseDuration(duration)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(d), nil
}
