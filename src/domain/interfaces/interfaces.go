package interfaces

import (
	"collectionneur-cli/src/domain/entities"
	"time"
)

type IAPI interface {
	GetLastSpendInfo() (*entities.SpendInfo, error)
	SendSpendInfos([]*entities.SpendInfo) (int, error)
}

type ISpendInfoDAO interface {
	GetSpendInfoBefore(dt time.Time) ([]*entities.SpendInfo, error)
}
