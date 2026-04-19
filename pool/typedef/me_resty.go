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

func (sdk *RestySDK) Incept(xactQN string) (typesem.SemRef, error) {
	return typesem.SemRef{}, nil
}

func (sdk *RestySDK) Create(spec DefSpec) (typesem.SemRef, error) {
	var dto typesem.SemRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/xacts/defs")
	if err != nil {
		return typesem.SemRef{}, err
	}
	if res.IsError() {
		return typesem.SemRef{}, fmt.Errorf("received: %v", string(res.Body()))
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
