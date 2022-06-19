package common

type RequestOptsHeader struct {
	Key   string
	Value string
}

type RequestOpts struct {
	Credentials string              `json:"credentials"`
	Headers     []RequestOptsHeader `json:"headers"`
	Method      string              `json:"method"`
	Mode        string              `json:"mode"`
}
