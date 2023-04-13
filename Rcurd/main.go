package main

func main() {
	// Run data migration to create necessary database tables and columns.
	DataMigration()

	// Set up API endpoint routing and start the server.
	HandlerRouting()
}
