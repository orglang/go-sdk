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

func (sdk *RestySDK) Incept(decQN string) (typesem.SemRef, error) {
	return typesem.SemRef{}, nil
}

func (sdk *RestySDK) Create(spec DecSpec) (DecSnap, error) {
	var dto DecSnap
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/procs/decs")
	if err != nil {
		return DecSnap{}, err
	}
	if res.IsError() {
		return DecSnap{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}

func (sdk *RestySDK) RetrieveSnap(decID string) (DecSnap, error) {
	var dto DecSnap
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetPathParam("id", decID).
		Get("/decs/{id}")
	if err != nil {
		return DecSnap{}, err
	}
	if res.IsError() {
		return DecSnap{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]typesem.SemRef, error) {
	dtos := []typesem.SemRef{}
	return dtos, nil
}
