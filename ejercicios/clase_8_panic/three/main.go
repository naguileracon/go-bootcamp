package main

func main() {
	// Create a new store
	store := &StorageClientsSlice{
		Clients: []Client{
			{ID: 1, Name: "Juan", File: "juan.txt", Home: "Calle 123", PhoneNumber: "123456789"},
			{ID: 2, Name: "Juan", File: "juan.txt", Home: "Calle 123", PhoneNumber: "123456789"},
		},
	}

	//store.AddClient(&Client{ID: 4, Name: "Juan", File: "juan.txt", Home: "Calle 123", PhoneNumber: "123456789"})
	store.AddClient(&Client{ID: 2, Name: "", File: "juan.txt", Home: "Calle 123", PhoneNumber: "123456789"})

}
