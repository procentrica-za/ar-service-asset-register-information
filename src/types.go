package main

import "github.com/gorilla/mux"

//router service struct
type Server struct {
	router *mux.Router
}

//config struct
type Config struct {
	CRUDHost                     string
	CRUDPort                     string
	ASSETREGISTERINFORMATIONPort string
}

type AssetID struct {
	AssetID string `json:"id"`
}

type AssetRegisterResponse struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	SerialNo           string `json:"serialno"`
	Size               string `json:"size"`
	Type               string `json:"type"`
	Class              string `json:"class"`
	Dimension1Val      string `json:"dimension1val"`
	Dimension2Val      string `json:"dimension2val"`
	Dimension3Val      string `json:"dimension3val`
	Dimension4Val      string `json:"dimension4val"`
	Dimension5Val      string `json:"dimension5val"`
	Dimension6Val      string `json:"dimension6val"`
	Extent             string `json:"extent"`
	ExtentConfidence   string `json:"extentconfidence"`
	TakeOnDate         string `json:"takeondate"`
	DeRecognitionvalue string `json:"derecognitionvalue"`
	Latitude           string `json:"latitude"`
	Longitude          string `json:"longitude"`
}

type AssetList struct {
	Assets []AssetRegisterResponse `json:"assets"`
}

type FunclocDetails struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Geom        string `json:"geom"`
}

type FunclocAssets struct {
	AssetID                     string `json:"assetid"`
	Name                        string `json:"name"`
	DerecognitionDate           string `json:"derecognitiondate"`
	Derecognitionvalue          string `json:"derecognitionvalue"`
	Description                 string `json:"description"`
	Dimension1Value             string `json:"dimension1value"`
	Dimension2Value             string `json:"dimension2value"`
	Dimension3Value             string `json:"dimension3value"`
	Dimension4Value             string `json:"dimension4value"`
	Dimension5Value             string `json:"dimension5value"`
	Extent                      string `json:"extent"`
	ExtentConfidence            string `json:"extentconfidence"`
	ManufactureDate             string `json:"manufacturedate"`
	ManufactureDateConfidence   string `json:"manufacturedateconfidence"`
	Takeondate                  string `json:"takeondate"`
	SerialNo                    string `json:"serialno"`
	Lat                         string `json:"lat"`
	Lon                         string `json:"lon"`
	CuName                      string `json:"cuname"`
	CuDescription               string `json:"cudescription"`
	EulYears                    string `json:"eulyears"`
	ResidualValFactor           string `json:"residualvalfactor"`
	Size                        string `json:"size"`
	SizeUnit                    string `json:"sizeunit"`
	Type                        string `json:"type"`
	Class                       string `json:"class"`
	IsActive                    string `json:"isactive"`
	Age                         string `json:"age"`
	CarryingValueClosingBalance string `json:"carryingvalueclosingbalance"`
	CarryingValueOpeningBalance string `json:"carryingvalueopeningbalance"`
	CostClosingBalance          string `json:"costclosingbalance"`
	CostOpeningBalance          string `json:"costopeningbalance"`
	CRC                         string `json:"crc"`
	DepreciationClosingBalance  string `json:"depreciationclosingbalance"`
	DepreciationOpeningBalance  string `json:"depreciationopeningbalance"`
	ImpairmentClosingBalance    string `json:"impairmentclosingbalance"`
	ImpairmentOpeningBalance    string `json:"impairmentopeningbalance"`
	ResidualValue               string `json:"residualvalue"`
	RulYears                    string `json:"rulyears"`
	DRC                         string `json:"drc"`
	FY                          string `json:"fy"`
}

type FuncLocAssetList struct {
	ID          string          `json:"id"`
	Description string          `json:"description"`
	Name        string          `json:"name"`
	Latitude    string          `json:"latitude"`
	Longitude   string          `json:"longitude"`
	Assets      []FunclocAssets `json:"funclocassets"`
}

type ShadowLocation struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type FuncLocsList struct {
	Locations []ShadowLocation `json:"funclocs"`
}
