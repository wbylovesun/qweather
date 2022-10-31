package qweather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (q *QWeather) request(api string, values url.Values) ([]byte, error) {
	values.Add("key", q.key)
	u := "https://" + q.host + api + "?" + values.Encode()
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func NewQWeather(mode string, key string) (*QWeather, error) {
	qWeather := new(QWeather)
	if _, ok := hosts[mode]; !ok {
		return nil, fmt.Errorf("invalid mode")
	}
	qWeather.mode = mode
	qWeather.host = hosts[mode]
	qWeather.key = key
	return qWeather, nil
}
