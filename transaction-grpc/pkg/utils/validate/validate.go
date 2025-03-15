package utilsvalidate

import (
	"fmt"
	"strings"

	protoTransaction "banking-system/transaction-service/rpc/types/transaction/v1alpha1"
)

//var regexName string = `^[a-z][-_0-9a-z]+`

func ValidateDataDeletePayment(request *protoTransaction.Transaction) error {

	stringArray := []string{}

	// organizationId := request.GetOrganizationId()
	// cardId := request.GetCardId()

	// if organizationId == "" {
	// 	stringArray = append(stringArray, "organizationId is required")
	// }

	// if cardId == "" {
	// 	stringArray = append(stringArray, "cardId is required")
	// }

	if len(stringArray) > 0 {
		justString := strings.Join(stringArray, " ")
		return fmt.Errorf(justString)
	}
	return nil

}