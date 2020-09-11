package main

//create routes
func (s *Server) routes() {
	s.router.HandleFunc("/asset", s.handlegetasset()).Methods("POST")
	s.router.HandleFunc("/assets", s.handlegetassets()).Methods("POST")
}
