package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	res := "id,name,address,plan-name,date,amount\n1900,Shilpa,Kolkata,SuperSaver,01-03-2022,86"
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "/1")
		// Return response to be tested
		rw.Write([]byte(res))
	}))

	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	client := InternetProviderClient{
		Client: *server.Client(),
		Url:    server.URL,
	}
	body, _ := client.GetBillById(1)

	assert.Equal(t, res, string(body))
}
