package api

type BucketEntry struct {
	Id       string            `json:"id"`
	BucketId string            `json:"bucket_id"`
	Method   string            `json:"method"`
	Path     string            `json:"path"`
	IP       string            `json:"ip"`
	Time     string            `json:"time"`
	Headers  map[string]string `json:"headers"`
	Body     string            `json:"data"`
}
