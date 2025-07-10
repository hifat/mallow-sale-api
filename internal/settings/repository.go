package settingModule

type Repository interface {
	Get() (*Response, error)
	Update(costPercentage float32) error
}
