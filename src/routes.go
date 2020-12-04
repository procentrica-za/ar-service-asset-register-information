package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/asset", s.handlegetasset()).Methods("GET")
	s.router.HandleFunc("/assets", s.handlegetassets()).Methods("GET")

	//Demo Routes
	s.router.HandleFunc("/funclocdetails", s.handlegetfunclocCurrentDetails()).Methods("GET")
	//Demo shadow Routes
	s.router.HandleFunc("/funclocshadowdetails", s.handlegetfunclocShadowDetails()).Methods("GET")
	//Demo Location Routes
	s.router.HandleFunc("/funclocs", s.handleshadowlocations()).Methods("GET")
	//Get Node func locs
	s.router.HandleFunc("/nodefunclocs", s.handleGetNodeFuncLocs()).Methods("GET")
	//Get Node assetlocs
	s.router.HandleFunc("/nodeassets", s.handleGetNodeAssets()).Methods("GET")
}
