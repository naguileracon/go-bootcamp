package main

import (
	"errors"
	"fmt"
)

type Client struct {
	File        string
	Name        string
	ID          int
	PhoneNumber string
	Home        string
}

type StorageClientsSlice struct {
	Clients []Client
}

func (s *StorageClientsSlice) GetItem(id int) Client {
	return s.Clients[id]
}

func (s *StorageClientsSlice) AddClient(client *Client) {

	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println("Several errors were detected at runtime")
		}
		fmt.Println("End of execution")

	}()

	// validate if client exists
	s.validateIfClientExists(client)
	// validating if the client has zero values
	_, err := s.validateZeroValuesInClient(client)
	// handling error
	if err != nil {
		panic(err.Error())
	}

	// add client to slice
	s.Clients = append(s.Clients, *client)
	fmt.Println("Client added successfully")
}

func (s *StorageClientsSlice) validateIfClientExists(client *Client) {
	for _, c := range s.Clients {
		if c.ID == client.ID {
			panic("Error: the client already exists")
		}
	}
	return
}

func (s *StorageClientsSlice) validateZeroValuesInClient(client *Client) (bool, error) {
	if client.ID == 0 {
		return false, errors.New("error: the ID is zero")
	}
	if client.Name == "" {
		return false, errors.New("error: the name is empty")
	}
	if client.File == "" {
		return false, errors.New("error: the file is empty")
	}
	if client.Home == "" {
		return false, errors.New("error: the home is empty")
	}
	if client.PhoneNumber == "" {
		return false, errors.New("error: the phone number is empty")
	}
	return true, nil
}
