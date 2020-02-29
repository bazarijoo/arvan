package RequestBody

type UpdateUserBalanceRequestBody struct {
	PhoneNumber string `json:"phone_number"`
	Credit      int    `json:"credit"`
}
