package webreq

type H map[string]string

type Headers struct {
	List []H
}

func NewHeaders() *Headers {
	return &Headers{}
}

func (h *Headers) Add(key string, value string) {
	h.List = append(h.List, H{key: value})
}
