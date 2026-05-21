package storageModule

type UploadRequest struct {
	ID          string `json:"id"`
	File        []byte `json:"file"`
	ServiceCode string `json:"serviceCode"`

	ContentType string `json:"-"`
	Filename    string `json:"-"`
}
