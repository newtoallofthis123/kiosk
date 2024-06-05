package main

import "log"

func main() {
	env := NewEnv()
	server, err := NewServer(env)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	err = server.Start()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	defer server.Db.Db.Close()
}
