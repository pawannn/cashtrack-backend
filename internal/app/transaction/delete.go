package transaction

import (
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (tA *TransactionApp) Delete(tx *models.Transaction) utils.CashTrackError {
	err := tA.databaseRepo.DeleteTransaction(*tx)
	if err != utils.NoErr {
		return err
	}
	return utils.NoErr
}
