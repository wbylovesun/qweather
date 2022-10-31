package qweather

import (
	"net/url"
)

func (q *QWeather) Warning(location string) ([]*warningItem, error) {
	values := url.Values{}
	values.Add("location", location)
	respBytes, err := q.request("/v7/warning/now", values)
	if err != nil {
		return nil, err
	}
	var r warningResponse
	err = decode(respBytes, &r)
	if err != nil {
		return nil, err
	}
	return r.Warning, nil
}
