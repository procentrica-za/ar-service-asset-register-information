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
	ID          string          `json:"id,omitempty"`
	Description string          `json:"description,omitempty"`
	Name        string          `json:"name,omitempty"`
	Latitude    string          `json:"latitude,omitempty"`
	Longitude   string          `json:"longitude,omitempty"`
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

type NodeFuncLocs struct {
	Id              string `json:"id,omitempty"`
	FuncLocNodeId   string `json:"funclocnodeid,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	InstallDate     string `json:"installdate,omitempty"`
	Status          string `json:"status,omitempty"`
	FuncLocNodeName string `json:"funclocnodename,omitempty"`
}

type NodeFuncLocsList struct {
	NodeFuncLocs []NodeFuncLocs `json:"nodefunclocs"`
}

type NodeAssets struct {
	Id            string  `json:"id,omitempty"`
	FuncLocNodeId string  `json:"funclocnodeid,omitempty"`
	FuncLocID     string  `json:"funclocid,omitempty"`
	Name          string  `json:"name,omitempty"`
	Description   string  `json:"description,omitempty"`
	Lat           float32 `json:"lat,omitempty"`
	Lon           float32 `json:"lon,omitempty"`
	Cuname        string  `json:"cuname,omitempty"`
	Typename      string  `json:"typename,omitempty"`
	SerialNo      string  `json:"serialno,omitempty"`
	Extent        float32 `json:"extent,omitempty"`
	Crc           float32 `json:"crc,omitempty"`
	Drc           float32 `json:"drc,omitempty"`
	Cost          float32 `json:"cost,omitempty"`
	CarryingValue float32 `json:"carryingvalue,omitempty"`
	TakeOnDate    string  `json:"takeondate,omitempty"`
	Rulyears      float32 `json:"rulyears,omitempty"`
	TypeFriendly  string  `json:"typefriendly,omitempty"`
	Size          float32 `json:"size,omitempty"`
}

type NodeAssetsList struct {
	NodeAssets []NodeAssets `json:"nodeassets"`
}

type Assetdetails struct {
	ID                    string  `json:"id,omitempty"`
	Name                  string  `json:"name,omitempty"`
	Type                  string  `json:"type,omitempty"`
	TypeFriendly          string  `json:"typefriendly,omitempty"`
	Description           string  `json:"description,omitempty"`
	ManufactureDate       string  `json:"manufacturedate,omitempty"`
	TakeOnDate            string  `json:"takeondate,omitempty"`
	SerialNo              string  `json:"serialno,omitempty"`
	DerecognitionDate     string  `json:"derecognitiondate,omitempty"`
	DerecognitionValue    string  `json:"derecognitionvalue,omitempty"`
	CompatibleUnitID      string  `json:"compatibleunitid,omitempty"`
	CompatibleUnitName    string  `json:"compatibleunitname,omitempty"`
	Dimension1Name        string  `json:"dimension1name,omitempty"`
	Dimension1Description string  `json:"dimension1description,omitempty"`
	Dimension1Unit        string  `json:"dimension1unit,omitempty"`
	Dimension2Name        string  `json:"dimension2name,omitempty"`
	Dimension2Description string  `json:"dimension2description,omitempty"`
	Dimension2Unit        string  `json:"dimension2unit,omitempty"`
	Dimension3Name        string  `json:"dimension3name,omitempty"`
	Dimension3Description string  `json:"dimension3description,omitempty"`
	Dimension3Unit        string  `json:"dimension3unit,omitempty"`
	Dimension4Name        string  `json:"dimension4name,omitempty"`
	Dimension4Description string  `json:"dimension4description,omitempty"`
	Dimension4Unit        string  `json:"dimension4unit,omitempty"`
	Dimension5Name        string  `json:"dimension5name,omitempty"`
	Dimension5Description string  `json:"dimension5description,omitempty"`
	Dimension5Unit        string  `json:"dimension5unit,omitempty"`
	Dimension1Value       float32 `json:"dimension1value,omitempty"`
	Dimension2Value       float32 `json:"dimension2value,omitempty"`
	Dimension3Value       float32 `json:"dimension3value,omitempty"`
	Dimension4Value       float32 `json:"dimension4value,omitempty"`
	Dimension5Value       float32 `json:"dimension5value,omitempty"`
	Extent                float32 `json:"extent,omitempty"`
	Rulyears              float32 `json:"rulyears,omitempty"`
	Crc                   float32 `json:"crc,omitempty"`
	Drc                   float32 `json:"drc,omitempty"`
	Cost                  float32 `json:"cost,omitempty"`
	CarryingValue         float32 `json:"carryingvalue,omitempty"`
	Size                  float32 `json:"size,omitempty"`
}

type FlexVals struct {
	Category     string `json:"category,omitempty"`
	Name         string `json:"name,omitempty"`
	Value        string `json:"value,omitempty"`
	DisplayOrder string `json:"displayorder,omitempty"`
	Flddefname   string `json:"flddefname,omitempty"`
	Datatype     string `json:"datatype,omitempty"`
	Controltype  string `json:"controltype,omitempty"`
	Isunique     bool   `json:"isunique,omitempty"`
	Unit         string `json:"unit,omitempty"`
	Lookupvals   string `json:"lookupvals,omitempty"`
	DateAdded    string `json:"timestamp,omitempty"`
}

type Levels struct {
	TypeLevelName string `json:"typelevelname,omitempty"`
	Level         int    `json:"level,omitempty"`
	Name          string `json:"name,omitempty"`
}

type AssetDetail struct {
	ID                    string     `json:"id,omitempty"`
	Name                  string     `json:"name,omitempty"`
	Type                  string     `json:"type,omitempty"`
	TypeFriendly          string     `json:"typefriendly,omitempty"`
	Description           string     `json:"description,omitempty"`
	ManufactureDate       string     `json:"manufacturedate,omitempty"`
	TakeOnDate            string     `json:"takeondate,omitempty"`
	SerialNo              string     `json:"serialno,omitempty"`
	DerecognitionDate     string     `json:"derecognitiondate,omitempty"`
	DerecognitionValue    string     `json:"derecognitionvalue,omitempty"`
	CompatibleUnitID      string     `json:"compatibleunitid,omitempty"`
	CompatibleUnitName    string     `json:"compatibleunitname,omitempty"`
	Dimension1Name        string     `json:"dimension1name,omitempty"`
	Dimension1Description string     `json:"dimension1description,omitempty"`
	Dimension1Unit        string     `json:"dimension1unit,omitempty"`
	Dimension2Name        string     `json:"dimension2name,omitempty"`
	Dimension2Description string     `json:"dimension2description,omitempty"`
	Dimension2Unit        string     `json:"dimension2unit,omitempty"`
	Dimension3Name        string     `json:"dimension3name,omitempty"`
	Dimension3Description string     `json:"dimension3description,omitempty"`
	Dimension3Unit        string     `json:"dimension3unit,omitempty"`
	Dimension4Name        string     `json:"dimension4name,omitempty"`
	Dimension4Description string     `json:"dimension4description,omitempty"`
	Dimension4Unit        string     `json:"dimension4unit,omitempty"`
	Dimension5Name        string     `json:"dimension5name,omitempty"`
	Dimension5Description string     `json:"dimension5description,omitempty"`
	Dimension5Unit        string     `json:"dimension5unit,omitempty"`
	Dimension1Value       float32    `json:"dimension1value,omitempty"`
	Dimension2Value       float32    `json:"dimension2value,omitempty"`
	Dimension3Value       float32    `json:"dimension3value,omitempty"`
	Dimension4Value       float32    `json:"dimension4value,omitempty"`
	Dimension5Value       float32    `json:"dimension5value,omitempty"`
	Extent                float32    `json:"extent,omitempty"`
	Rulyears              float32    `json:"rulyears,omitempty"`
	Crc                   float32    `json:"crc,omitempty"`
	Drc                   float32    `json:"drc,omitempty"`
	Cost                  float32    `json:"cost,omitempty"`
	CarryingValue         float32    `json:"carryingvalue,omitempty"`
	Size                  float32    `json:"size,omitempty"`
	Flexvals              []FlexVals `json:"flexvalues"`
	ALevels               []Levels   `json:"levels"`
}

type FunclocationAssets struct {
	ID            string  `json:"id,omitempty"`
	FuncLocId     string  `json:"funclocId,omitempty"`
	Name          string  `json:"name,omitempty"`
	Description   string  `json:"description,omitempty"`
	Lat           float32 `json:"lat,omitempty"`
	Lon           float32 `json:"lon,omitempty"`
	Cuname        string  `json:"cuname,omitempty"`
	Typename      string  `json:"typename,omitempty"`
	SerialNo      string  `json:"serialno,omitempty"`
	Extent        float32 `json:"extent,omitempty"`
	Crc           float32 `json:"crc,omitempty"`
	Drc           float32 `json:"drc,omitempty"`
	Cost          float32 `json:"cost,omitempty"`
	CarryingValue float32 `json:"carryingvalue,omitempty"`
	TakeOnDate    string  `json:"takeondate,omitempty"`
	Rulyears      float32 `json:"rulyears,omitempty"`
	TypeFriendly  string  `json:"typefriendly,omitempty"`
	Size          float32 `json:"size,omitempty"`
}

type FunclocationAssetsList struct {
	Funclocassets []FunclocationAssets `json:"funclocassets"`
}

type FuncLoc struct {
	Id              string `json:"id,omitempty"`
	FuncLocNodeId   string `json:"funclocnodeid,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	InstallDate     string `json:"installdate,omitempty"`
	Status          string `json:"status,omitempty"`
	FuncLocNodeName string `json:"funclocnodename,omitempty"`
}

type FuncLocList struct {
	Funclocs []FuncLoc `json:"funcloc"`
}

type FuncLocDetail struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type FuncLocSpatial struct {
	Name string  `json:"name,omitempty"`
	Lat  float32 `json:"lat,omitempty"`
	Lon  float32 `json:"lon,omitempty"`
	Id   string  `json:"id,omitempty"`
}

type FuncLocSpatialList struct {
	FuncLocSpatial []FuncLocSpatial `json:"funclocspatial"`
}

type NodeFuncLocsSpatial struct {
	Name string  `json:"name,omitempty"`
	Lat  float32 `json:"lat,omitempty"`
	Lon  float32 `json:"lon,omitempty"`
	Id   string  `json:"id,omitempty"`
}

type NodeFuncLocsSpatialList struct {
	NodeFuncLocsSpatial []NodeFuncLocsSpatial `json:"nodefunclocspatial"`
}

type FlattenedHierarchy struct {
	ParentId string `json:"parentid,omitempty"`
	Id       string `json:"Id,omitempty"`
	Name     string `json:"name,omitempty"`
	Nodetype string `json:"nodetype,omitempty"`
	IsLeaf   bool   `json:"isleaf,omitempty"`
}

type FlattenedHierarchyFilter struct {
	NodeID      string `json:"nodeid,omitempty"`
	Likelyhood  string `json:"likelyhood,omitempty"`
	Consequence string `json:"consequence,omitempty"`
	AssettypeID string `json:"assettypeid,omitempty"`
	Rulyears    int    `json:"rulyears,omitempty"`
}

type FlattenedHierarchyList struct {
	FlattenedHierarchy []FlattenedHierarchy `json:"nodehierarchyflattened"`
}
