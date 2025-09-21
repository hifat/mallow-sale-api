package settingModule

type IRepository interface {
	Get() (*Response, error)
	Update(costPercentage float32) error
}
