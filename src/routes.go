package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/asset", s.handlegetasset()).Methods("GET")
	s.router.HandleFunc("/assets", s.handlegetassets()).Methods("GET")

	//Demo Routes
	s.router.HandleFunc("/funclocdetails", s.handlegetfunclocCurrentDetails()).Methods("GET")
}
