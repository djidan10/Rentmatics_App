package Services

import (
	"fmt"
	//	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"bytes"
	"encoding/json"
	//	"fmt"
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

func Payment(w http.ResponseWriter, r *http.Request) {
	var payment Model.Payment

	var paymentinsto Model.Paymentinsto
	err := json.NewDecoder(r.Body).Decode(&payment)
	fmt.Println("getpayment", payment)

	if err != nil {
		fmt.Println("errpayment", err)
	}

	paymentinsto.Purpose = payment.Purpose
	paymentinsto.Amount = payment.Amount
	paymentinsto.Buyer_name = payment.Buyer_name
	paymentinsto.Email = payment.Email
	paymentinsto.Phone = payment.Phone
	paymentinsto.Send_email = payment.Send_email
	paymentinsto.Send_sms = payment.Send_sms
	//paymentinsto.Redirect_url = "http://localhost:8083/PaymentResponse/" + payment.Loginid + "/" + strconv.Itoa(payment.Homeid)
	paymentinsto.Redirect_url = "http://localhost:8083/rentmatics/SampleLoading.html?loginid=" + payment.Loginid + "&homeid=" + strconv.Itoa(payment.Homeid)
	fmt.Println("rediecturl", paymentinsto.Redirect_url)

	jsonvalue, err := json.Marshal(paymentinsto)
	fmt.Println("#######error marshal", err, string(jsonvalue))

	//paymentvalue := "{ 'purpose' :'" + paymentinsto.Purpose + "`," + "'amount' :'" + paymentinsto.Amount + "'," + "'buyer_name' :'" + paymentinsto.Buyer_name + "'," + "'email' :'" + paymentinsto.Email + "'," + "'phone' :'" + paymentinsto.Phone + ",'" + "'send_email' :'" + paymentinsto.Send_email + "'," + "'send_sms' :'" + paymentinsto.Send_sms + "'," + "'redirect_url' :'" + paymentinsto.Redirect_url + "' }"
	//	paymentvalue:=`'{    "purpose"      	:" ' + paymentinsto.Purpose +		 '`,'
	//	+ ' "amount" 	:" ' + paymentinsto.Amount + 	'",'
	//+ ' "buyer_name" 	:" ' + paymentinsto.Buyer_name + 	'",'
	//+ ' "email"		 :" ' + paymentinsto.Email + 		'",'
	//+ ' "phone" 	:" ' + paymentinsto.Phone + 		',"'
	//+ ' "send_email" 	:" ' + paymentinsto.Send_email + 	'",'
	//+ ' "send_sms" 	:" ' + paymentinsto.Send_sms + 	'",'
	//+ ' "redirect_url" 	:" ' + paymentinsto.Redirect_url + 	' " }'`
	//	fmt.Println(paymentvalue)

	req, err := http.NewRequest("POST", "https://www.instamojo.com/api/1.1/payment-requests/", bytes.NewBuffer(jsonvalue))
	//	req, err := http.NewRequest("POST", "https://www.instamojo.com/api/1.1/payment-requests/", bytes.NewBuffer(jsonvalue))

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("X-Api-Key", "5aeafa0841d6946200d8d08fd31def29")
	req.Header.Set("X-Auth-Token", "98b16c2818b1d48807a9e7d7b91e082b")

	Httpclient := &http.Client{}
	resp, err := Httpclient.Do(req)

	//	Senddata, err := json.Marshal(resp.Body.)

	//	if err != nil {
	//		fmt.Println("err", err)
	//	}
	fmt.Println("**********resoponse", resp, err)
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set("Access-Control-Allow-Orgin", "*")

	if resp.Body != nil {
		io.Copy(w, resp.Body)
	}

}
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
