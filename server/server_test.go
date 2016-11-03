package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

// Note that these tests are neither for unit-testing
// nor for extensive API testing. The coverage is also low.
//
// The purpose of these tests is to just to highlight
// the API usage, like a swagger doc and to catch
// any high-level regressions.

func TestTodosAPIs(t *testing.T) {

	client := &http.Client{}

	t.Log("Getting the JWT from the server")
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth("admin", "admin")
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.Status)
	}

	t.Log("Extracting the JWT from the HTTP response")
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	jwt := string(b)
	t.Log("JWT string is:\n", jwt)

	//------------------------------------------
	t.Log("Add a new todo item")
	todo := []byte(`{"Task": "Buy milk"}`)
	req, err = http.NewRequest("POST", "http://127.0.0.1:8080/todos", bytes.NewBuffer(todo))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+jwt)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatal(resp)
	}
	//------------------------------------------

	//------------------------------------------
	t.Log("Get a list of todos from the server")
	req, err = http.NewRequest("GET", "http://127.0.0.1:8080/todos", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+jwt)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp)
	}
	//------------------------------------------

	//------------------------------------------
	t.Log("Update a todo item")
	todo = []byte(`{"Task": "Buy milk", "Completed": true}`)
	req, err = http.NewRequest("PUT", "http://127.0.0.1:8080/todo/0", bytes.NewBuffer(todo))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+jwt)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp)
	}
	//------------------------------------------

	//------------------------------------------
	t.Log("Delete a todo item")
	req, err = http.NewRequest("DELETE", "http://127.0.0.1:8080/todo/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer "+jwt)
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp)
	}
	//------------------------------------------
}
