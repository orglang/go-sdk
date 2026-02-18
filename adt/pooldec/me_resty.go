package pooldec

import (
	"github.com/go-resty/resty/v2"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec DecSpec) (PoolRef, error) {
	var res PoolRef
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		Post("/pools/decs")
	if err != nil {
		return PoolRef{}, err
	}
	return res, nil
}
