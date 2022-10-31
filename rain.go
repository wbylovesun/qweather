package qweather

import (
	"fmt"
	"net/url"
)

func (q *QWeather) Rain(lat, lng float64) (*rainResult, error) {
	values := url.Values{}
	values.Add("location", fmt.Sprintf("%f,%f", lng, lat))
	respBytes, err := q.request("/v7/minutely/5m", values)
	if err != nil {
		return nil, err
	}
	var r rainResponse
	err = decode(respBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r.rainResult, nil
}
