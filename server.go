package main

import (
	Model "Rentmatics_App/Model"
	Service "Rentmatics_App/Services"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	fs3 := http.FileServer(http.Dir("./HTML/"))
	//http.Handle("/", fs)
	router.PathPrefix("/HTML/").Handler(http.StripPrefix("/HTML/", fs3))
	fs4 := http.FileServer(http.Dir("./rentmatics_final/"))
	//http.Handle("/", fs)
	router.PathPrefix("/rentmatics_final/").Handler(http.StripPrefix("/rentmatics_final/", fs4))

	fs5 := http.FileServer(http.Dir("./Rent/"))

	router.PathPrefix("/Rent/").Handler(http.StripPrefix("/Rent/", fs5))

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

	router.HandleFunc("/RentUser", Service.RentUserdata)
	router.HandleFunc("/RentLogin", Service.RentLogin)
	router.HandleFunc("/RentUserlogout", Service.RentUserlogout)

	router.HandleFunc("/Login", Service.Login)
	router.HandleFunc("/Logout", Service.Userlogout)
	router.HandleFunc("/Changepassword", Service.Changepassword)
	router.HandleFunc("/Forgotpassword", Service.Forgot)

	router.HandleFunc("/RentLogin", Service.RentLogin)

	//Notify
	router.HandleFunc("/Notifyme", Service.Notification)

	//payment

	router.HandleFunc("/paytm", Service.Payment)
	router.HandleFunc("/PaymentResponse/{loginid}/{homeid}/{payment_id}/{payment_request_id}/", Service.PaymentResponse)

	//Booking
	router.HandleFunc("/Booknow", Service.Booknow)

	//Subscriber
	router.HandleFunc("/Postsubscriber", Service.Postsubscriber)

	//Myprofile
	router.HandleFunc("/Myprofile", Service.Profileinsert)
	router.HandleFunc("/GetProfile", Service.Getprofile)

	//Tenants
	router.HandleFunc("/GetTenant", Service.GetTenant)
	router.HandleFunc("/GetIndiviualTenant", Service.GetSingleTenant)

	router.HandleFunc("/Sentotp", Service.SendOTP)
	router.HandleFunc("/OTPAUTH", Service.OTPAUTH1)
	router.HandleFunc("/Resend", Service.ResendOtp)

	//Complaint
	router.HandleFunc("/GetComplaint", Service.GetComplaint)
	router.HandleFunc("/GetAllpendingComplaints", Service.Getpendingstatus)
	router.HandleFunc("/GetSingleComplaint", Service.GetsingleComplaint)
	router.HandleFunc("/InsertComplaint", Service.InsertComplaints)
	router.HandleFunc("/UpdateComplaintstaus", Service.Upadtecomplaintstatus)

	//Vacate
	router.HandleFunc("/Getvacatedetails", Service.Getvacatedetails)
	router.HandleFunc("/GetSinglevacate", Service.Getsinglevacate)
	router.HandleFunc("/Insertvacate", Service.InsertVacate)
	router.HandleFunc("/Updatevacate", Service.UpdatetVacate)

	//Activites
	router.HandleFunc("/GetActivitiesdetails", Service.GetActivity)
	router.HandleFunc("/GetAllActivitiesPending", Service.GetActivityPending)
	router.HandleFunc("/GetSingleActivity", Service.GetsingleActivity)
	router.HandleFunc("/InsertAcivity", Service.InsertActivity)
	router.HandleFunc("/UpdateActivity", Service.UpdateActivity)

	//Tenant Request
	router.HandleFunc("/GetRequestdetails", Service.GetRequest)
	router.HandleFunc("/GetRequestPendingdetails", Service.GetPendingrequest)
	router.HandleFunc("/GetSingleRequest", Service.GetsingleRequest)
	router.HandleFunc("/InsertRequest", Service.InsertRequest)
	router.HandleFunc("/Updaterequeststaus", Service.Updaterequeststatus)

	//Owner Request
	router.HandleFunc("/GetOwnerdetails", Service.GetOwners)
	router.HandleFunc("/GetSingleOwner", Service.GetsingleOwner)
	router.HandleFunc("/InsertOwner", Service.Insertowner)

	router.HandleFunc("/Ownertoexecutive", Service.Ownertoexecutive)

	//Executive request
	router.HandleFunc("/GetExecutivedetails", Service.GetExecutive)
	router.HandleFunc("/GetSingleExecutive", Service.GetsingleExecutive)
	router.HandleFunc("/InsertExecutive", Service.Insertexecutive)

	//Owner complaints
	router.HandleFunc("/GetAllownersComplaint", Service.GetOwnercomplaint)
	router.HandleFunc("/Getsingleownercomplaint", Service.GetSingleOwnercomplaint)
	router.HandleFunc("/Insertownercomplaints", Service.Insertownercomplaints)

	//Ownerdashboard

	router.HandleFunc("/Getownerhomedetils", Service.Getmyhomedetails)
	router.HandleFunc("/GetOwnertenantdetails", Service.GetOwnerTenantDetails)
	router.HandleFunc("/Insertownerrequest", Service.InsertOwnerRequestDetails)
	router.HandleFunc("/InsertOwnervisit", Service.InsertOwnerVisitDetails)
	router.HandleFunc("/OwnerMessage", Service.InsertOwnerMessage)

	//GetOwnerid
	router.HandleFunc("/GetOwnerId", Service.GetsingleId)
	router.HandleFunc("/GetOwnercomplaintAll", Service.GetOwnercomplaintId)
	router.HandleFunc("/GetOwnerRequestAll", Service.GetOwnerResquestId)
	router.HandleFunc("/GetSingleOwnerResquest", Service.GetSingleOwnerResquest)

	//GetComplaints and request Details
	router.HandleFunc("/GetcomplaintsRequestData", Service.GetComplaintRequest)
	router.HandleFunc("/GetOwnercomplaintsRequestData", Service.GetOwnerComplaintRequest)
	router.HandleFunc("/GetOwnermessage", Service.GetOwnermessage)
	router.HandleFunc("/GetOwnermessagedetail", Service.GetOwnermessagedetails)

	//Tenant Api
	router.HandleFunc("/GetTenantid", Service.GetAdminTenant)
	router.HandleFunc("/Postfeedback", Service.Feedback)
	router.HandleFunc("/Getfeedback", Service.GetFeedback)
	router.HandleFunc("/OwnerReferandearn", Service.Referandearn)

	//top homes
	router.HandleFunc("/Tophome", Service.Tophome)

	//Refer and earn
	router.HandleFunc("/Getrefercode", Service.GetRefercode)

	//For HTTPS

	//Contactform
	router.HandleFunc("/insertcontact", Service.Insertcontact)

	//Browsecategories
	router.HandleFunc("/Browsecategories", Service.Browswcat)

	// Default server - non-trusted for debugging

	serverhttp := func() {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*", "http://develop.rentmatics.com", "http://api.msg91.com"},
			AllowCredentials: true,
		})
		handler := c.Handler(router)

		fmt.Println("Server should be available at http", config.Port)
		fmt.Println(http.ListenAndServe(config.Port, handler))
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
