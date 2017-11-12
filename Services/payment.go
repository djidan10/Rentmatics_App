package Services

import (
	Model "Rentmatics_App/Model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var paylo = `{
 "purpose": "purpose",
  "amount": "10",
  "buyer_name": "sandy",
  "email": "sandhiyabalakrishnan6@gmail.com",
  "phone": "7418924864",
  "send_email": "false",
  "send_sms": "false",
	"redirect_url":"http://localhost:8083/rentmatics/"
  }`
var Authkey = "150219AJLZP0ww58ff3960"
var OtpAuth int

func Payment(w http.ResponseWriter, r *http.Request) {
	var payment Model.Payment
	var paymentinsto Model.Paymentinsto
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		log.Error("Error on Getting Payment value", err)
	}

	paymentinsto.Purpose = payment.Purpose
	paymentinsto.Amount = payment.Amount
	paymentinsto.Buyer_name = payment.Buyer_name
	paymentinsto.Email = payment.Email
	paymentinsto.Phone = payment.Phone
	paymentinsto.Send_email = payment.Send_email
	paymentinsto.Send_sms = payment.Send_sms
	paymentinsto.Redirect_url = "http://localhost:8083/rentmatics/SampleLoading.html?loginid=" + payment.Loginid + "&homeid=" + strconv.Itoa(payment.Homeid)
	log.Info("rediecturl", paymentinsto.Redirect_url)
	jsonvalue, err := json.Marshal(paymentinsto)
	log.Error("Error on payment", err, string(jsonvalue))
	req, err := http.NewRequest("POST", "https://www.instamojo.com/api/1.1/payment-requests/", bytes.NewBuffer(jsonvalue))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("X-Api-Key", "5aeafa0841d6946200d8d08fd31def29")
	req.Header.Set("X-Auth-Token", "98b16c2818b1d48807a9e7d7b91e082b")
	Httpclient := &http.Client{}
	resp, err := Httpclient.Do(req)
	if resp.Body != nil {
		io.Copy(w, resp.Body)
	}

}

//Payment response store to database
func PaymentResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside PaymentResponse")
	vars := mux.Vars(r)
	loginid := vars["loginid"]
	homeid := vars["homeid"]
	payment_id := vars["payment_id"]
	payment_request_id := vars["payment_request_id"]
	fmt.Println(loginid)
	fmt.Println(homeid)
	fmt.Println(payment_id)
	fmt.Println(payment_request_id)
}

type SentOtp struct {
	OTP         string
	Phonenumber string
}

func SendOTP(w http.ResponseWriter, r *http.Request) {
	var Getsentotp SentOtp
	err := json.NewDecoder(r.Body).Decode(&Getsentotp)
	if err != nil {
		log.Error("Struct Mismatch on get profile", err)
	}

	resp, err := http.Get(Getsentotp.OTP)
	if err != nil {
		log.Error("Otp error", err)
	}
	log.Info(resp)

}

type GetOtpAuth struct {
	OTp string
}

func OTPAUTH1(w http.ResponseWriter, r *http.Request) {

	var Getotp GetOtpAuth
	err := json.NewDecoder(r.Body).Decode(&Getotp)
	if err != nil {
		log.Error("Struct Mismatch on get profile", err)
	}
	fmt.Println("insideauth", Getotp.OTp)
	resp, err := http.Get(Getotp.OTp)
	fmt.Println(resp)
	fmt.Println("value", &resp)
	fmt.Println("aaa", resp.Body)

	if err != nil {
		log.Error("Otp error", err)
	}

	if resp.Body != nil {
		io.Copy(w, resp.Body)
	}

}
func ResendOtp(w http.ResponseWriter, r *http.Request) {
	var Getsentotp GetOtpAuth
	err := json.NewDecoder(r.Body).Decode(&Getsentotp)
	if err != nil {
		log.Error("Struct Mismatch on get profile", err)
	}

	resp, err := http.Get(Getsentotp.OTp)
	if err != nil {
		log.Error("Otp error", err)
	}
	log.Info(resp)

}
