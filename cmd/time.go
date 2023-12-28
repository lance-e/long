package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"long/internal/timer"
	"strconv"
	"strings"
	"time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowtime := timer.GetNowTime()
		log.Printf("输出结果为:%s,%d", nowtime.Format("2006-01-02 15:04:05"), nowtime.Unix())
	},
}

var duration string        //持续时间
var calculationTime string //需要计算的时间
var calculateTimeCmd = &cobra.Command{
	Use:   "calculate",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculationTime == "" {
			currentTimer = timer.GetNowTime()

		} else {
			var err error
			count := strings.Count(calculationTime, " ")
			if count == 0 {
				layout = "2006-01-02"
			}
			if count == 1 {
				layout = "2006-01-02 15:04:05"
			}
			currentTimer, err = time.Parse(layout, calculationTime)
			if err != nil {
				t, _ := strconv.Atoi(calculationTime)
				currentTimer = time.Unix(int64(t), 0)
			}

		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer calculate err : %v", err)
		}
		log.Printf("输出结果:%s,%d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)       //增加nowTime子命令
	timeCmd.AddCommand(calculateTimeCmd) //增加calculateTime子命令
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间,有效单位是ns,us,(或n,u),ms,s,m,h")
	calculateTimeCmd.Flags().StringVarP(&calculationTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或者格式化后的时间")
}
