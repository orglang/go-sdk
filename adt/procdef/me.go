package procdef

type stepKind string

const (
	Msg = stepKind("msg")
	Svc = stepKind("svc")
)

type DefRec struct {
	ID string   `json:"id"`
	K  stepKind `json:"kind"`
}
