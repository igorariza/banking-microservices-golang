package utilsvalidate

import (
	"fmt"
	"strings"

	protoTransaction "banking-system/transaction-service/rpc/types/transaction/v1alpha1"
)

//var regexName string = `^[a-z][-_0-9a-z]+`

func ValidateDataTransaction(request *protoTransaction.Transaction) error {

	stringArray := []string{}

	if len(stringArray) > 0 {
		justString := strings.Join(stringArray, " ")
		return fmt.Errorf(justString)
	}
	return nil

}