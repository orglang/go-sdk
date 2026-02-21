package uniqref

type Msg struct {
	DescID string `json:"id" param:"id"`
	DescRN int64  `json:"rn" query:"rn"`
}
