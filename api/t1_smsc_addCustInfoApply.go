package api

import "github.com/codingeasygo/util/xmap"

func (c *Config) T1SmscAddCustInfoApplyRequest() (data xmap.M, err error) {
	method := "t1.smsc.addCustInfoApply"
	version := "1.2"
	url := proUrlPrefix + t1SmscAddCustInfoApplyUrl
	if IsDev {
		url = devUrlPrefix + t1SmscAddCustInfoApplyUrl
	}
	_, data, err = c.Request(url, method, version, "bizContent") // TODO bizContent
	return
}
