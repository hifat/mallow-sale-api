package settingModule

type Repository interface {
	Get() (*Entity, error)
	Update(costPercentage float32) error
}
