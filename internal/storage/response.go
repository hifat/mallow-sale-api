package storageModule

type UploadResponse struct {
	Filename  string `json:"filename"`
	ObjectKey string `json:"objectKey"`
	Url       string `json:"url"`
}
