package mockRules

import "github.com/hifat/goroger-core/rules"

//go:generate mockgen -source=./rules.go -destination=./rules_mock.go -package=mockRules
type validator interface {
	rules.Validator
}
