package termdec

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/typesem"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec DecSpec) (typesem.SemRef, error) {
	var dto typesem.SemRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/pools/decs")
	if err != nil {
		return typesem.SemRef{}, err
	}
	if res.IsError() {
		return typesem.SemRef{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}
