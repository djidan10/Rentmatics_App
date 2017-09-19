package Services

import (
	MobilePay "Rentmatics_App/MobilePay"
	"net/http"
)

//For Instomoja mobile payment
func MobilePayService(w http.ResponseWriter, r *http.Request) {

	MobilePay.SetCredentials("PqwSFo2t29uV0NtYfv6NpQvcg2L7M8bMf312Gome", "dIGyHd2wCyuod7TPM7erVHZV6zfoMFWTScVz4vzGgj27cJDZ9jZMLDkN4043ktzDBPdRDUUvh0zpn0pUGn11zDpHrpDRZdTPWtWNNhZLDgdcAIqNjW8zrrtaJbSM1STy", "vthnWJGYlvelDImii3HMBxOHFVaRdAXRb5oU9X4R", "7Agkfi7An5oUFFzKd5XOkkYYPH4vA211k6n7mH7MuSYWdy0Kw7IY4vpLZGtDPDRpWvxFNgCU58YIE9ImY8XgeV7hi6nVEEEfDsRhQKShQiw1xSIci1rIUve5kgdtfuzY")
	b, err := MobilePay.CreateOrderTokens("")
	if err != nil {
		log.Info("Error on Mobile Pay service", err)
	}
	w.Write(b)

}
