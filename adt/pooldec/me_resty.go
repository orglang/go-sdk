package pooldec

import (
	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/descsem"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec DecSpec) (descsem.SemRef, error) {
	var dto descsem.SemRef
	_, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/pools/decs")
	if err != nil {
		return descsem.SemRef{}, err
	}
	return dto, nil
}
