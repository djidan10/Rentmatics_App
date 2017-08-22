package main

import (
	"crypto/tls"

	"encoding/json"

	"fmt"

	"net/http"

	Model "Rentmatics_App/Model"
	Service "Rentmatics_App/Services"

	"github.com/gorilla/mux"
)

const (
	contenttypeJSON = "application/json; charset=utf-8"
)

//Error provides error message in JSON format to client (utility)
func Error(w http.ResponseWriter, err int, msg string) {

	e := Model.APIError{}
	e.Error.Code = err
	e.Error.Message = msg
	w.Header().Set("Content-Type", contenttypeJSON)
	w.WriteHeader(err)
	json.NewEncoder(w).Encode(e)
}

func Serve() bool {

	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./rentmatics_theme/"))
	//http.Handle("/", fs)
	router.PathPrefix("/rentmatics_theme/").Handler(http.StripPrefix("/rentmatics_theme/", fs))

	fs1 := http.FileServer(http.Dir("./Houseimage/"))
	//http.Handle("/", fs)
	router.PathPrefix("/Houseimage/").Handler(http.StripPrefix("/Houseimage/", fs1))

	fs2 := http.FileServer(http.Dir("./rentmatics/"))
	//http.Handle("/", fs)
	router.PathPrefix("/rentmatics/").Handler(http.StripPrefix("/rentmatics/", fs2))

	fs3 := http.FileServer(http.Dir("./rentmatics/"))
	//http.Handle("/", fs)
	router.PathPrefix("/rentmatics/").Handler(http.StripPrefix("/rentmatics/", fs3))

	//Handlers

	//Rent Homes

	router.HandleFunc("/GetAllhomedetails", Service.Getallhomedetails)
	router.HandleFunc("/Gethomedetails", Service.Gethomedetails)
	router.HandleFunc("/GetIndivual_home", Service.GetIndivual_home)
	router.HandleFunc("/GetFilter", Service.GetFilter)

	//Insert Home Details
	router.HandleFunc("/GetImage", Service.InsertHomeDetails)

	//User authentication
	router.HandleFunc("/GetLogin", Service.Auth)

	//Favourites
	router.HandleFunc("/InsertFavourites", Service.InsertFavourite)
	router.HandleFunc("/GetFavourites", Service.Getfavourites)

	//Schedule Visit
	//router.HandleFunc("/GetSchedule", Service.GetSchedule)

	//Get Cities
	router.HandleFunc("/GetCities", Service.GetCities)
	//router.HandleFunc("/Scheduleotp", Service.Scheduleotp)
	router.HandleFunc("/ScheduleVisit", Service.ScheduleVisit)

	router.HandleFunc("/Wishlist", Service.Wishlist)
	router.HandleFunc("/Insertwishlist", Service.InsertWishlist)
	router.HandleFunc("/Deletewishlist", Service.DeleteWishlist)

	router.HandleFunc("/User", Service.Userdata)
	router.HandleFunc("/Login", Service.Login)
	router.HandleFunc("/Logout", Service.Userlogout)
	//router.HandleFunc("/Changepassword", Service.change)
	//Notify

	router.HandleFunc("/Notifyme", Service.Notification)

	//payment

	router.HandleFunc("/paytm", Service.Payment)
	router.HandleFunc("/PaymentResponse/{loginid}/{homeid}/{payment_id}/{payment_request_id}/", Service.PaymentResponse)

	//Subscriber
	router.HandleFunc("/Postsubscriber", Service.Postsubscriber)

	//Myprofile
	router.HandleFunc("/Myprofile", Service.Profileinsert)
	router.HandleFunc("/GetProfile", Service.Getprofile)

	//For HTTPS

	//
	// Default server - non-trusted for debugging
	serverhttp := func() {
		fmt.Println("Server should be available at http", config.Port)
		fmt.Println(http.ListenAndServe(config.Port, router))
	}

	// Setup TLS parameters for trusted server implementation
	if config.SSL && config.Key != "" && config.Cert != "" {
		// Setup TLS parameters
		tlsConfig := &tls.Config{
			ClientAuth:   tls.NoClientCert,
			MinVersion:   tls.VersionTLS12,
			Certificates: make([]tls.Certificate, 1),
		}

		var err error
		// Setup API server private key and certificate
		tlsConfig.Certificates[0], err = tls.X509KeyPair([]byte(config.Cert), []byte(config.Key))
		if err != nil {
			fmt.Println("Error during decoding service key and certificate:", err)
			return false
		}

		tlsConfig.BuildNameToCertificate()

		https := &http.Server{
			Addr:      config.Https_port,
			TLSConfig: tlsConfig,
			Handler:   router,
		}

		// Trusted server implementation
		server := func() {
			fmt.Println("Server should be available at https", config.Https_port)
			fmt.Println(https.ListenAndServeTLS("", ""))
		}
		go server()
	}

	// Schedule API server
	go serverhttp()

	return true
}
