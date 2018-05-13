package usecases

import (
	"collectionneur-cli/src/domain/interfaces"
	"time"
)

type LoadAndSendSpendInfoUseCase struct {
	api interfaces.IAPI
	dao interfaces.ISpendInfoDAO
}

func (u *LoadAndSendSpendInfoUseCase) Execute() (int, error) {
	var startDt time.Time
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return 0, err
	}

	info, err := u.api.GetLastSpendInfo()
	if err == nil {
		startDt = info.Date
	} else if err == interfaces.ErrApiNoInfo {
		startDt = time.Date(2000, 1, 1, 0, 0, 0, 0, loc)
	} else {
		return 0, err
	}

	newInfos, err := u.dao.GetSpendInfoBefore(startDt)
	if err != nil {
		return 0, err
	}

	count, err := u.api.SendSpendInfos(newInfos)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func NewLoadAndSendSpendInfoUseCase(api interfaces.IAPI, dao interfaces.ISpendInfoDAO) *LoadAndSendSpendInfoUseCase {
	return &LoadAndSendSpendInfoUseCase{
		api: api,
		dao: dao,
	}
}
