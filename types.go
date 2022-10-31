package qweather

import "time"

type QWeather struct {
	mode string
	host string
	key  string
}

type responseType interface {
	aqiResponse | aqiDailyResponse | indiceResponse | pDailyResponse | pHourlyResponse | rainResponse | realtimeResponse | warningResponse

	GetCode() string
	isOk() bool
}

type baseResponse struct {
	Code       string    `json:"code"`
	UpdateTime time.Time `json:"update_time"`
}

func (b baseResponse) GetCode() string {
	return b.Code
}

func (b baseResponse) isOk() bool {
	return b.Code == "200"
}

type realtimeNow struct {
	ObsTime   string `json:"obsTime"`
	Temp      string `json:"temp"`
	FeelsLike string `json:"feelsLike"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Precip    string `json:"precip"`
	Pressure  string `json:"pressure"`
	Vis       string `json:"vis"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}

type realtimeResponse struct {
	baseResponse
	Now realtimeNow `json:"now"`
}

type pHourlyItem struct {
	FxTime    string `json:"fxTime"`
	Temp      string `json:"temp"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Pop       string `json:"pop"`
	Precip    string `json:"precip"`
	Pressure  string `json:"pressure"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}

type pHourlyResponse struct {
	baseResponse
	Hourly []*pHourlyItem `json:"hourly"`
}

type rainItem struct {
	FxTime string `json:"fxTime"`
	Precip string `json:"precip"`
	Type   string `json:"type"`
}

type rainResult struct {
	Summary  string      `json:"summary"`
	Minutely []*rainItem `json:"minutely"`
}

type rainResponse struct {
	baseResponse
	rainResult
}

type pDailyItem struct {
	FxDate         string `json:"fxDate"`
	Sunrise        string `json:"sunrise"`
	Sunset         string `json:"sunset"`
	Moonrise       string `json:"moonrise"`
	Moonset        string `json:"moonset"`
	MoonPhase      string `json:"moonPhase"`
	MoonPhaseIcon  string `json:"moonPhaseIcon"`
	TempMax        string `json:"tempMax"`
	TempMin        string `json:"tempMin"`
	IconDay        string `json:"iconDay"`
	TextDay        string `json:"textDay"`
	IconNight      string `json:"iconNight"`
	TextNight      string `json:"textNight"`
	Wind360Day     string `json:"wind360Day"`
	WindDirDay     string `json:"windDirDay"`
	WindScaleDay   string `json:"windScaleDay"`
	WindSpeedDay   string `json:"windSpeedDay"`
	Wind360Night   string `json:"wind360Night"`
	WindDirNight   string `json:"windDirNight"`
	WindScaleNight string `json:"windScaleNight"`
	WindSpeedNight string `json:"windSpeedNight"`
	Humidity       string `json:"humidity"`
	Precip         string `json:"precip"`
	Pressure       string `json:"pressure"`
	Vis            string `json:"vis"`
	Cloud          string `json:"cloud"`
	UvIndex        string `json:"uvIndex"`
}

type pDailyResponse struct {
	baseResponse
	Daily []*pDailyItem `json:"daily"`
}

type warningItem struct {
	Id            string `json:"id"`
	Sender        string `json:"sender"`
	PubTime       string `json:"pubTime"`
	Title         string `json:"title"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	Status        string `json:"status"`
	Level         string `json:"level"`
	Severity      string `json:"severity"`
	SeverityColor string `json:"severityColor"`
	Type          string `json:"type"`
	TypeName      string `json:"typeName"`
	Urgency       string `json:"urgency"`
	Certainty     string `json:"certainty"`
	Text          string `json:"text"`
	Related       string `json:"related"`
}

type warningResponse struct {
	baseResponse
	Warning []*warningItem `json:"warning"`
}

type indiceItem struct {
	Date     string `json:"date"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	Category string `json:"category"`
	Text     string `json:"text"`
}

type indiceResponse struct {
	baseResponse
	Daily []*indiceItem `json:"daily"`
}

type aqiItem struct {
	PubTime  string `json:"pubTime"`
	Aqi      string `json:"aqi"`
	Level    string `json:"level"`
	Category string `json:"category"`
	Primary  string `json:"primary"`
	Pm10     string `json:"pm10"`
	Pm2P5    string `json:"pm2p5"`
	No2      string `json:"no2"`
	So2      string `json:"so2"`
	Co       string `json:"co"`
	O3       string `json:"o3"`
}

type aqiResponse struct {
	baseResponse
	Now *aqiItem `json:"now"`
}

type aqiDailyItem struct {
	FxDate   string `json:"fxDate"`
	Aqi      string `json:"aqi"`
	Level    string `json:"level"`
	Category string `json:"category"`
	Primary  string `json:"primary"`
}

type aqiDailyResponse struct {
	baseResponse
	Daily []*aqiDailyItem `json:"daily"`
}
