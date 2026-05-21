package storageHelper

import (
	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	utilsModule "github.com/hifat/mallow-sale-api/internal/utils"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"github.com/hifat/mallow-sale-api/pkg/define"
	"github.com/hifat/mallow-sale-api/pkg/handling"
)

type helper struct {
	cfg *config.Config
}

func New(cfg *config.Config) storageModule.IHelper {
	return &helper{cfg: cfg}
}

func (h *helper) GetDirName(serviceCode string) (string, error) {
	tail := ""
	if h.cfg.App.Mode == utilsModule.EnumDevMode {
		tail = "-dev"
	}

	dirSet := map[string]string{
		"PURCHASE_EVIDENCE": "purchase-evidences",
	}
	if _, ok := dirSet[serviceCode]; !ok {
		return "", handling.ThrowErrByCode(define.CodeInvalidServiceCode)
	}

	return dirSet[serviceCode] + tail, nil
}
