package xactdef

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Incept(xactQN string) (DefRef, error) {
	return DefRef{}, nil
}

func (sdk *RestySDK) Create(spec DefSpec) (DefSnap, error) {
	var res DefSnap
	resp, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		Post("/xacts")
	if err != nil {
		return DefSnap{}, err
	}
	if resp.IsError() {
		return DefSnap{}, fmt.Errorf("received: %v", string(resp.Body()))
	}
	return res, nil
}

func (sdk *RestySDK) Modify(snap DefSnap) (DefSnap, error) {
	return DefSnap{}, nil
}

func (sdk *RestySDK) Retrieve(defID string) (DefSnap, error) {
	return DefSnap{}, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]DefRef, error) {
	return []DefRef{}, nil
}
