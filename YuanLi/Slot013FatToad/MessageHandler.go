package Slot013FatToad

import (
	"Force/GameServer/Common"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var RELEASE_MODE = false

var slot SlotProb

func init() {
	slot.Init()
}

func MessageHandler(data []byte) []byte {
	cmd := SpinCmd{RTP: DEFAULT_RTP}
	traceId := strings.ReplaceAll(uuid.New().String(), "-", "")
	slotResult := SlotResult{Code: Common.ERROR_CODE_OK, TraceId: traceId}
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		fmt.Printf("[ERROR] %v, data: %v\n", err, string(data))
		slotResult.Code = Common.ERROR_CODE_UNMARSHAL
		result, _ := json.Marshal(slotResult)
		return result
	}
	fmt.Printf("[MessageHandler] Receive: %#v\n", cmd)

	// 檢查 RTP 是否有效
	if cmd.RTP < 0 || cmd.RTP >= RTP_TOTAL {
		fmt.Printf("[ERROR][MessageHandler] Invalid RTP: %d\n", cmd.RTP)
		slotResult.Code = Common.ERROR_CODE_INVALID_RTP
		result, _ := json.Marshal(slotResult)
		return result
	}

	// 計算單線押注，並檢查 BuyType 是否有效
	lineBet := cmd.TotalBet / PAYLINE_TOTAL // TODO: 檢查投注額合理範圍
	betRatio, ok := BetRatio[cmd.BuyType]
	if !ok {
		fmt.Printf("[ERROR][MessageHandler] Invalid BuyType: %d\n", cmd.BuyType)
		slotResult.Code = Common.ERROR_CODE_INVALID_BUY_TYPE
		result, _ := json.Marshal(slotResult)
		return result
	}
	lineBet = int(float64(lineBet) / betRatio)

	// 測試指令
	debugCmdList := cmd.DebugCmdList
	if RELEASE_MODE {
		debugCmdList = nil
	}

	// 取得遊戲結果
	slotResult.BuyType = cmd.BuyType
	_ = slot.Run(cmd.RTP, lineBet, &slotResult, debugCmdList)
	result, err := json.Marshal(slotResult)
	if err == nil {
		prettyJSON, err := json.MarshalIndent(slotResult, "", "  ")
		if err == nil {
			fmt.Printf("[MessageHandler] Send Result:\n%s\n", string(prettyJSON))
		} else {
			fmt.Printf("[MessageHandler] Send: %#v\n", slotResult)
		}
	}
	return result
}
