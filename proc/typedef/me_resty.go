package typedef

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/orglang/go-sdk/adt/typesem"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Incept(typeQN string) (typesem.SemRef, error) {
	return typesem.SemRef{}, nil
}

func (sdk *RestySDK) Create(spec DefSpec) (DefSnap, error) {
	var dto DefSnap
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/types")
	if err != nil {
		return DefSnap{}, err
	}
	if res.IsError() {
		return DefSnap{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}

func (sdk *RestySDK) Modify(snap DefSnap) (DefSnap, error) {
	return DefSnap{}, nil
}

func (sdk *RestySDK) Retrieve(defID string) (DefSnap, error) {
	return DefSnap{}, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]typesem.SemRef, error) {
	return []typesem.SemRef{}, nil
}
