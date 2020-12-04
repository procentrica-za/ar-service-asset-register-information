package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) handlegetasset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset Has Been Called...")
		//Get Asset ID from URL
		assetid := r.URL.Query().Get("assetid")

		//Check if Asset ID provided is null
		if assetid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Asset ID not properly provided in URL")
			fmt.Println("Asset ID not proplery provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/asset?assetid=" + assetid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for export asset")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(500)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to asset export")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed..."+bodyString)
			fmt.Println("Request to DB can't be completed..." + bodyString)
			return
		}

		//close the request.
		defer req.Body.Close()

		//create new response struct
		var assetResponse AssetRegisterResponse

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&assetResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding asset response")
			return
		}
		js, jserr := json.Marshal(assetResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the response to export asset")
			return
		}

		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetassets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Assets Has Been Called...")
		//Get Asset ID from URL
		assettypeid := r.URL.Query().Get("assettypeid")

		//Check if Asset ID provided is null
		if assettypeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Asset ID not properly provided in URL")
			fmt.Println("Asset ID not proplery provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/assets?assettypeid=" + assettypeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve advertisement information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get asset data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get asset data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := AssetList{}
		assetsList.Assets = []AssetRegisterResponse{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get Messages response ")
			return
		}
		//convert struct back to JSON.
		js, jserr := json.Marshal(assetsList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetfunclocCurrentDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle get func loc for current db details has been called")
		//Get Functional Location ID from URL

		funclocid := r.URL.Query().Get("funclocid")

		//Check if no Email address was provided in the URL

		if funclocid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/funclocdetails?funclocid=" + funclocid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve get functional location information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get functional location information"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get functional location information" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var funclocdetails FunclocDetails

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&funclocdetails)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding functional location information ")
			return
		}

		//send CRUD response to email service
		req1, respErr1 := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/funclocassets?funclocid=" + funclocid)

		fmt.Println("Sent to text messaging service")
		//check for response error of 500
		if respErr1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr1.Error())
			fmt.Println("Error received from text service->" + respErr1.Error())
			return
		}
		if req1.StatusCode != 200 {
			w.WriteHeader(req1.StatusCode)
			fmt.Fprint(w, "Request to CRUD to get asset information can't be completed...")
			fmt.Println("Unable to process asset functional location information")
			return
		}
		if req1.StatusCode == 500 {
			w.WriteHeader(500)

			bodyBytes1, err1 := ioutil.ReadAll(req1.Body)
			if err1 != nil {
				log.Fatal(err1)
			}
			bodyString := string(bodyBytes1)
			fmt.Fprintf(w, "Request to CRUD to get asset information can't be completed..."+bodyString)
			fmt.Println("Unable to process asset functional location information..." + bodyString)
			return
		}

		//close the request
		defer req1.Body.Close()

		//create new response struct for JSON list
		assetsList := FuncLocAssetList{}
		assetsList.ID = funclocdetails.ID
		assetsList.Description = funclocdetails.Description
		assetsList.Name = funclocdetails.Name
		assetsList.Latitude = funclocdetails.Latitude
		assetsList.Longitude = funclocdetails.Longitude
		assetsList.Assets = []FunclocAssets{}

		//decode request into decoder which converts to the struct
		decoder1 := json.NewDecoder(req1.Body)
		err1 := decoder1.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get Messages response ")
			return
		}
		//convert struct back to JSON.
		js, jserr := json.Marshal(assetsList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handlegetfunclocShadowDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle get func loc for shadow db details has been called")
		//Get Functional Location ID from URL

		funclocid := r.URL.Query().Get("funclocid")

		//Check if no Email address was provided in the URL

		if funclocid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/funclocshadowdetails?funclocid=" + funclocid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve get functional location information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get functional location information"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get functional location information" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var funclocdetails FunclocDetails

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&funclocdetails)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding functional location information ")
			return
		}

		//send CRUD response to email service
		req1, respErr1 := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/funclocshadowassets?funclocid=" + funclocid)

		fmt.Println("Sent to text messaging service")
		//check for response error of 500
		if respErr1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr1.Error())
			fmt.Println("Error received from text service->" + respErr1.Error())
			return
		}
		if req1.StatusCode != 200 {
			w.WriteHeader(req1.StatusCode)
			fmt.Fprint(w, "Request to CRUD to get asset information can't be completed...")
			fmt.Println("Unable to process asset functional location information")
			return
		}
		if req1.StatusCode == 500 {
			w.WriteHeader(500)

			bodyBytes1, err1 := ioutil.ReadAll(req1.Body)
			if err1 != nil {
				log.Fatal(err1)
			}
			bodyString := string(bodyBytes1)
			fmt.Fprintf(w, "Request to CRUD to get asset information can't be completed..."+bodyString)
			fmt.Println("Unable to process asset functional location information..." + bodyString)
			return
		}

		//close the request
		defer req1.Body.Close()

		//create new response struct for JSON list
		assetsList := FuncLocAssetList{}
		assetsList.ID = funclocdetails.ID
		assetsList.Description = funclocdetails.Description
		assetsList.Name = funclocdetails.Name
		assetsList.Latitude = funclocdetails.Latitude
		assetsList.Longitude = funclocdetails.Longitude
		assetsList.Assets = []FunclocAssets{}

		//decode request into decoder which converts to the struct
		decoder1 := json.NewDecoder(req1.Body)
		err1 := decoder1.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get Messages response ")
			return
		}
		//convert struct back to JSON.
		js, jserr := json.Marshal(assetsList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleshadowlocations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Assets Has Been Called...")

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/funclocs")

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to location information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get location data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get location data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		locationsList := FuncLocsList{}
		locationsList.Locations = []ShadowLocation{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&locationsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get Messages response ")
			return
		}
		//convert struct back to JSON.
		js, jserr := json.Marshal(locationsList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetNodeFuncLocs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Node Func Locs Has Been Called...")

		nodeid := r.URL.Query().Get("nodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/nodefunclocs?nodeid=" + nodeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to location information")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(req.StatusCode)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Request to DB can't be completed...")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)
			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "An internal error has occured whilst trying to get location data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get location data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		nodesList := NodeFuncLocsList{}
		nodesList.NodeFuncLocs = []NodeFuncLocs{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&nodesList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get location response ")
			return
		}
		//convert struct back to JSON.
		js, jserr := json.Marshal(nodesList)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the decoded response into specified JSON format!")
			return
		}

		//return success back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
