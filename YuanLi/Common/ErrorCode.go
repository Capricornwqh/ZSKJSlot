package Common

// 錯誤代碼
const (
	ERROR_CODE_OK = iota
	ERROR_CODE_UNMARSHAL
	ERROR_CODE_NO_SERVICE
	ERROR_CODE_NO_RESPONSE
	ERROR_CODE_NO_BALANCE
	ERROR_CODE_NO_TOKEN     // 找不到token
	ERROR_CODE_REPEAT_LOGIN // 重複登入
	ERROR_CODE_KICKED       // 後踢前或平台踢出
	ERROR_CODE_NO_GAME      // 找不到遊戲
	ERROR_CODE_NO_PLATFORM  // 無此平台
	ERROR_CODE_MAX_ODDS     // 超過最大贏分倍數
)

const (
	ERROR_CODE_PLATFORM_API = iota + 100 // 平台api錯誤
)

// 錯誤代碼 (遊戲相關)
const (
	ERROR_CODE_INVALID_RTP          = iota + 1000 // 無效的 RTP
	ERROR_CODE_INVALID_BUY_TYPE                   // 無效的購買類型
	ERROR_CODE_SETTING_NOT_FOUND                  // 未找到設定
	ERROR_CODE_ROULETTE_SPIN_FAILED               // 機率輪盤擲骰失敗
)

// 錯誤代碼 (測試相關)
const (
	ERROR_CODE_INVALID_DEBUG_CODE       = iota + 1100 // 無效的測試代碼
	ERROR_CODE_INVALID_DEBUG_SYMBOL                   // 無效的盤面
	ERROR_CODE_INVALID_DEBUG_FLAG                     // 無效的特別旗標
	ERROR_CODE_INVALID_DEBUG_REEL_INDEX               // 無效的停輪 index
	ERROR_CODE_INVALID_DEBUG_MULTIPLIER               // 無效的乘倍
)
