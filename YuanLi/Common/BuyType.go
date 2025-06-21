package Common

// Buy Type 購買類型 (注意：若要新增，請往後新增，不要從中間新增)
const (
	BUY_NONE             = iota // 未購買，即一般投注
	BUY_EXTRA_BET               // 額外投注
	BUY_FREE_SPINS              // 購買免費遊戲
	BUY_SUPER_FREE_SPINS        // 購買超級免費遊戲
)
