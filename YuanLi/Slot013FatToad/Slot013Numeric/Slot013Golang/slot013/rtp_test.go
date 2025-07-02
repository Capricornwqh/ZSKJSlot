package slot013

import (
	"testing"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestRTP(t *testing.T) {
	type args struct {
		rtp     int
		buyType int
	}
	tests := []struct {
		name string
		rv   *RTPVerifier
		args args
	}{
		// Test cases
		{
			"RTP0 上線數值 一般投注",
			&RTPVerifier{TotalCount: 100000000},
			args{rtp: 0, buyType: BUY_NONE},
		},
		{
			"RTP0 上線數值 購買免費",
			&RTPVerifier{TotalCount: 10000000},
			args{rtp: 0, buyType: BUY_FREE_SPINS},
		},
		{
			"RTP0 上線數值 購買超級",
			&RTPVerifier{TotalCount: 10000000},
			args{rtp: 0, buyType: BUY_SUPER_FREE_SPINS},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startTime := time.Now()
			tt.rv.Init()
			tt.rv.Run(tt.args.rtp, tt.args.buyType)

			elapsed := time.Since(startTime)
			hours := int(elapsed.Hours())
			minutes := int(elapsed.Minutes()) % 60
			seconds := int(elapsed.Seconds()) % 60
			p := message.NewPrinter(language.Chinese)
			p.Printf("Test %s %d次 耗时: %d时%d分%d秒\n",
				tt.name, tt.rv.TotalCount, hours, minutes, seconds)
			tt.rv.Dump(true)
			p.Println("========================================")
		})
	}
}
