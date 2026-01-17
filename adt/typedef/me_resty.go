package typedef

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Incept(typeQN string) (DefRefME, error) {
	return DefRefME{}, nil
}

func (sdk *RestySDK) Create(spec DefSpecME) (DefSnapME, error) {
	var res DefSnapME
	resp, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		Post("/types")
	if err != nil {
		return DefSnapME{}, err
	}
	if resp.IsError() {
		return DefSnapME{}, fmt.Errorf("received: %v", string(resp.Body()))
	}
	return res, nil
}

func (sdk *RestySDK) Modify(snap DefSnapME) (DefSnapME, error) {
	return DefSnapME{}, nil
}

func (sdk *RestySDK) Retrieve(defID string) (DefSnapME, error) {
	return DefSnapME{}, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]DefRefME, error) {
	return []DefRefME{}, nil
}
