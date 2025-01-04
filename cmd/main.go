package main

func main() {
	server := api.NewAPIServer(":8080", nil)
	server.run().Error()
	// This is the entry point of the program
}
