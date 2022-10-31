package qweather

import (
	"net/url"
	"strings"
)

func (q *QWeather) Indice1d(location string, typ []string) ([]*indiceItem, error) {
	var types = "0"
	if len(typ) > 0 {
		types = strings.Join(typ, ",")
	}
	values := url.Values{}
	values.Add("location", location)
	values.Add("type", types)
	respBytes, err := q.request("/v7/indices/1d", values)
	if err != nil {
		return nil, err
	}
	var r indiceResponse
	err = decode(respBytes, &r)
	if err != nil {
		return nil, err
	}
	return r.Daily, nil
}
