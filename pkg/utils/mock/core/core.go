package mockCore

import core "github.com/hifat/goroger-core"

//go:generate mockgen -source=./core.go -destination=./core_mock.go -package=mockCore
type helper interface {
	core.Helper
}

type logger interface {
	core.Logger
}
