package uniqref

type Msg struct {
	ID string `json:"id" param:"id"`
	RN int64  `json:"rn" query:"rn"`
}
