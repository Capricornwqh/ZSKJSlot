package Slot005GemBonanza

import (
	"Force/GameServer/Common"
	"testing"
)

func TestRTPVerifierRun(t *testing.T) {
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
			"RTP0 上線數值 一般投注 整合測試 97% 10億手",
			&RTPVerifier{TotalCount: 1000000000},
			args{rtp: 0, buyType: Common.BUY_NONE},
		},
		{
			"RTP0 上線數值 額外投注 整合測試 97% 10億手",
			&RTPVerifier{TotalCount: 1000000000},
			args{rtp: 0, buyType: Common.BUY_EXTRA_BET},
		},
		{
			"RTP0 上線數值 購買免費 整合測試 97% 1億手",
			&RTPVerifier{TotalCount: 100000000},
			args{rtp: 0, buyType: Common.BUY_FREE_SPINS},
		},
		{
			"RTP0 上線數值 購買超級 整合測試 97% 1億手",
			&RTPVerifier{TotalCount: 100000000},
			args{rtp: 0, buyType: Common.BUY_SUPER_FREE_SPINS},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rv.Init()
			tt.rv.Run(tt.args.rtp, tt.args.buyType)
			tt.rv.Dump(true)
		})
	}
}
