package main

import (
	"bytes"
	"compress/gzip"
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
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(assetResponse)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
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
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(assetsList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
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

		fmt.Println("Sent tocrud service for funclocassets service")
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
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(assetsList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
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

		fmt.Println("Sent to crud service for shadowassets funcloc")
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
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(assetsList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
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
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(locationsList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
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
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(nodesList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetNodeAssets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Node Assets Has Been Called...")

		nodeid := r.URL.Query().Get("nodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/nodeassets?nodeid=" + nodeid)

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
		assetsList := NodeAssetsList{}
		assetsList.NodeAssets = []NodeAssets{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(assetsList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetAssetDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle get asset detail has been called")
		//Get Asset ID from URL

		id := r.URL.Query().Get("id")

		//Check if no Email address was provided in the URL

		if id == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Asset ID not properly provided in URL")
			fmt.Println("Asset ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/assetdetails?id=" + id)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve get asset information")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get asset information"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get asset information" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var assetdetails Assetdetails

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&assetdetails)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding asset information ")
			return
		}

		//send CRUD response to flexval service
		req1, respErr1 := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/assetflexval?id=" + id)

		fmt.Println("Sent to crud to get flex val service")
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
			fmt.Println("Unable to process asset asset information")
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
			fmt.Println("Unable to process asset asset information..." + bodyString)
			return
		}

		//close the request
		defer req1.Body.Close()

		//create new response struct for JSON list
		assetsList := AssetDetail{}
		assetsList.ID = assetdetails.ID
		assetsList.Name = assetdetails.Name
		assetsList.Type = assetdetails.Type
		assetsList.TypeFriendly = assetdetails.TypeFriendly
		assetsList.Description = assetdetails.Description
		assetsList.ManufactureDate = assetdetails.ManufactureDate
		assetsList.TakeOnDate = assetdetails.TakeOnDate
		assetsList.SerialNo = assetdetails.SerialNo
		assetsList.DerecognitionDate = assetdetails.DerecognitionDate
		assetsList.DerecognitionValue = assetdetails.DerecognitionValue
		assetsList.CompatibleUnitID = assetdetails.CompatibleUnitID
		assetsList.CompatibleUnitName = assetdetails.CompatibleUnitName
		assetsList.Dimension1Name = assetdetails.Dimension1Name
		assetsList.Dimension1Description = assetdetails.Dimension1Description
		assetsList.Dimension1Unit = assetdetails.Dimension1Unit
		assetsList.Dimension2Name = assetdetails.Dimension2Name
		assetsList.Dimension2Description = assetdetails.Dimension2Description
		assetsList.Dimension2Unit = assetdetails.Dimension2Unit
		assetsList.Dimension3Name = assetdetails.Dimension3Name
		assetsList.Dimension3Description = assetdetails.Dimension3Description
		assetsList.Dimension3Unit = assetdetails.Dimension3Unit
		assetsList.Dimension4Name = assetdetails.Dimension4Name
		assetsList.Dimension4Description = assetdetails.Dimension4Description
		assetsList.Dimension4Unit = assetdetails.Dimension4Unit
		assetsList.Dimension5Name = assetdetails.Dimension5Name
		assetsList.Dimension5Description = assetdetails.Dimension5Description
		assetsList.Dimension5Unit = assetdetails.Dimension5Unit
		assetsList.Dimension1Value = assetdetails.Dimension1Value
		assetsList.Dimension2Value = assetdetails.Dimension2Value
		assetsList.Dimension3Value = assetdetails.Dimension3Value
		assetsList.Dimension4Value = assetdetails.Dimension4Value
		assetsList.Dimension5Value = assetdetails.Dimension5Value
		assetsList.Extent = assetdetails.Extent
		assetsList.Rulyears = assetdetails.Rulyears
		assetsList.Crc = assetdetails.Crc
		assetsList.Drc = assetdetails.Drc
		assetsList.Cost = assetdetails.Cost
		assetsList.CarryingValue = assetdetails.CarryingValue
		assetsList.Size = assetdetails.Size
		assetsList.Flexvals = []FlexVals{}
		assetsList.ALevels = []Levels{}

		//decode request into decoder which converts to the struct
		decoder1 := json.NewDecoder(req1.Body)
		err1 := decoder1.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get flex vals response ")
			return
		}

		//send CRUD response to flexval service
		req3, respErr2 := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/assetlevel?id=" + id)

		fmt.Println("Sent to crud to get asset level service")
		//check for response error of 500
		if respErr2 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr2.Error())
			fmt.Println("Error received from text service->" + respErr2.Error())
			return
		}
		if req3.StatusCode != 200 {
			w.WriteHeader(req3.StatusCode)
			fmt.Fprint(w, "Request to CRUD to get asset information can't be completed...")
			fmt.Println("Unable to process asset asset information")
			return
		}
		if req3.StatusCode == 500 {
			w.WriteHeader(500)

			bodyBytes2, err2 := ioutil.ReadAll(req3.Body)
			if err2 != nil {
				log.Fatal(err2)
			}
			bodyString := string(bodyBytes2)
			fmt.Fprintf(w, "Request to CRUD to get asset information can't be completed..."+bodyString)
			fmt.Println("Unable to process asset asset information..." + bodyString)
			return
		}

		//decode request into decoder which converts to the struct
		decoder3 := json.NewDecoder(req3.Body)
		err3 := decoder3.Decode(&assetsList)
		if err3 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err3.Error())
			fmt.Println("Error occured in decoding get levels response ")
			return
		}
		//close the request
		defer req3.Body.Close()
		//convert struct back to JSON.
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err4 := json.NewEncoder(gz).Encode(assetsList)
		if err4 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handlegetFuncLocAssets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Assets Has Been Called...")

		//Get Asset ID from URL

		funclocid := r.URL.Query().Get("funclocid")

		//Check if no Email address was provided in the URL

		if funclocid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Asset ID not properly provided in URL")
			fmt.Println("Asset ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/funclocassets?funclocid=" + funclocid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to funclocassets information")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get funclocassets data"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get funclocassets data" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := FunclocationAssetsList{}
		assetsList.Funclocassets = []FunclocationAssets{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get FuncLocAssets response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(assetsList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetFuncLoc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get funcloc Has Been Called...")

		funclocnodeid := r.URL.Query().Get("funclocnodeid")
		id := r.URL.Query().Get("id")

		//Check if no Node ID was provided in the URL

		if funclocnodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location node ID not properly provided in URL")
			fmt.Println("Functional Location node ID not properly provided in URL")
			return
		}
		if id == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/funcloc?funclocnodeid=" + funclocnodeid + "&id=" + id)

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
		funcslist := FuncLocList{}
		funcslist.Funclocs = []FuncLoc{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&funcslist)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(funcslist)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetFuncLocDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset Has Been Called...")
		//Get Asset ID from URL
		id := r.URL.Query().Get("id")

		//Check if Asset ID provided is null
		if id == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Funcloc ID not properly provided in URL")
			fmt.Println("Funcloc ID not proplery provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/funclocdetail?id=" + id)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for export funcloc")
			return
		}
		if req.StatusCode != 200 {
			w.WriteHeader(500)
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to asset funcloc")
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
		var locationResponse FuncLocDetail

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&locationResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding asset response")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(locationResponse)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetFuncLocSpatial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get funcloc spatial Has Been Called...")

		id := r.URL.Query().Get("id")

		//Check if no Node ID was provided in the URL

		if id == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/funclocspatial?id=" + id)

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
		funcslist := FuncLocSpatialList{}
		funcslist.FuncLocSpatial = []FuncLocSpatial{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&funcslist)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get assets response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(funcslist)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetNodeFuncLocSpatial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Node Func Locs spatial Has Been Called...")

		nodeid := r.URL.Query().Get("funclocnodeid")

		//Check if no Node ID was provided in the URL

		if nodeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Functional Location ID not properly provided in URL")
			fmt.Println("Functional Location ID not properly provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/nodefunclocspatial?funclocnodeid=" + nodeid)

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
		nodesList := NodeFuncLocsSpatialList{}
		nodesList.NodeFuncLocsSpatial = []NodeFuncLocsSpatial{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&nodesList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get location response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(nodesList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetNodeHierarchyFlattened() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Node Hierarchy Flattened Has Been Called...")

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/nodehierarchyflattened")

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
		nodesList := FlattenedHierarchyList{}
		nodesList.FlattenedHierarchy = []FlattenedHierarchy{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&nodesList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding get location response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(nodesList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetNodeHierarchyFlattenedFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Node Hierarchy Flattened Filtered Has Been Called...")
		///get JSON payload
		filterparams := FlattenedHierarchyFilter{}
		err := json.NewDecoder(r.Body).Decode(&filterparams)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error with hierarchy filter parameters")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(filterparams)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/nodehierarchyfiltered", "application/json", bytes.NewBuffer(requestByte))
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieved flattened node hierarchy.")
			return
		}

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get flattened node hierarchy filtered")
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
		nodesList := FlattenedHierarchyList{}
		nodesList.FlattenedHierarchy = []FlattenedHierarchy{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&nodesList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get location response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(nodesList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetNodeAssetsFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Node Assets Filtered Has Been Called...")
		///get JSON payload
		filterparams := FlattenedHierarchyFilter{}
		err := json.NewDecoder(r.Body).Decode(&filterparams)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error with hierarchy filter parameters")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(filterparams)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/nodeassetsfiltered", "application/json", bytes.NewBuffer(requestByte))
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve node assets filtered.")
			return
		}

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get nodeassetsfiltered")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get node assets filtered"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get node assets filtered" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		nodesList := NodeAssetsList{}
		nodesList.NodeAssets = []NodeAssets{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&nodesList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get location response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(nodesList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handlegetFuncLocAssetsFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Funcloc Assets Filtered Has Been Called...")
		///get JSON payload
		filterparams := FlattenedHierarchyFilter{}
		err := json.NewDecoder(r.Body).Decode(&filterparams)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error with hierarchy filter parameters")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(filterparams)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/funclocassetsfiltered", "application/json", bytes.NewBuffer(requestByte))
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve funclocassets filtered.")
			return
		}

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get funclocassets filtered")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get funcloc assets filtered"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get funcloc assets filtered" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		assetsList := FunclocationAssetsList{}
		assetsList.Funclocassets = []FunclocationAssets{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&assetsList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get location response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(assetsList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetNodeFuncLocsFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Node FuncLocs Filtered Has Been Called...")
		///get JSON payload
		filterparams := FlattenedHierarchyFilter{}
		err := json.NewDecoder(r.Body).Decode(&filterparams)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error with hierarchy filter parameters")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(filterparams)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/nodefunclocsfiltered", "application/json", bytes.NewBuffer(requestByte))
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve funclocs filtered.")
			return
		}

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get funclocs filtered")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get funclocs filtered"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get funclocs filtered" + bodyString)
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
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(nodesList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}

func (s *Server) handleGetNodeFuncLocSpatialFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Node FuncLocs Spatial Filtered Has Been Called...")
		///get JSON payload
		filterparams := FlattenedHierarchyFilter{}
		err := json.NewDecoder(r.Body).Decode(&filterparams)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error with hierarchy filter parameters")
			return
		}
		//create byte array from JSON payload
		requestByte, _ := json.Marshal(filterparams)

		//post to crud service
		req, respErr := http.Post("http://"+config.CRUDHost+":"+config.CRUDPort+"/nodefunclocspatialfiltered", "application/json", bytes.NewBuffer(requestByte))
		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to retrieve nodefunclocs spatial filtered.")
			return
		}

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for request to get nodefunclocs spatial filtered")
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
			fmt.Fprintf(w, "An internal error has occured whilst trying to get nodefunclocs spatial filtered"+bodyString)
			fmt.Println("An internal error has occured whilst trying to get nodefunclocs spatial filtered" + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct for JSON list
		nodesList := NodeFuncLocsSpatialList{}
		nodesList.NodeFuncLocsSpatial = []NodeFuncLocsSpatial{}

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)
		err1 := decoder.Decode(&nodesList)
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err1.Error())
			fmt.Println("Error occured in decoding get location response ")
			return
		}
		w.Header().Add("Accept-Charset", "utf-8")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		err2 := json.NewEncoder(gz).Encode(nodesList)
		if err2 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Error processing action"}`))
			return
		}
		gz.Close()
		return
	}
}
