package pooldec

import (
	"github.com/go-resty/resty/v2"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec DecSpec) (DecRef, error) {
	var res DecRef
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		Post("/pools/decs")
	if err != nil {
		return DecRef{}, err
	}
	return res, nil
}
