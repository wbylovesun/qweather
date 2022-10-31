package qweather

const (
	FreeMode     = "free"
	StandardMode = "standard"
	PremiumMode  = "premium"

	FreeModeHost     = "devapi.qweather.com"
	StandardModeHost = "api.qweather.com"
	PremiumModeHost  = StandardModeHost
)

var hosts = map[string]string{
	FreeMode:     FreeModeHost,
	StandardMode: StandardModeHost,
	PremiumMode:  PremiumModeHost,
}
