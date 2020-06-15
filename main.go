package main

import (
	"context"
	"fmt"
	clientexternal "github.com/insolar/testdoc/external_generated"
	serverexternal "github.com/insolar/testdoc/external_server"
	clientinternal "github.com/insolar/testdoc/internal_generated"
	serverinternal "github.com/insolar/testdoc/internal_server"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Server struct{}

// Set password
// (POST /auth/set-password)
func (s Server) SetPassword(ctx echo.Context, params serverexternal.SetPasswordParams) error {
	return nil
}

// Get token
// (GET /auth/token)
func (s Server) Token(ctx echo.Context, params serverexternal.TokenParams) error {
	return ctx.JSON(http.StatusOK, serverexternal.ResponsesToken{
		AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJpbnNvbGFyIiwic3ViIjoiZXhjaGFuZ2UtbmFtZSIsImlhdCI6MTUxNjIzOTAyMiwiZXhwIjoxNTE2MjQ1MDAwfQ.kT9y2S3FrHT_x25J8vx8n_WEkd0zbIHiqmzDW0H1ViI",
		Expiration:  1590656252,
	})
}

// Get clients
// (GET /clients)
func (s Server) ClientArray(ctx echo.Context) error {
	clients := &[]serverinternal.ResponsesClientYaml{
		{
			Active:     true,
			Login:      "unique_client_login",
			ModifiedAt: 1590669972,
		}, {
			Active:     false,
			Login:      "unique_client_login_1",
			ModifiedAt: 1590669972,
		},
	}

	return ctx.JSON(http.StatusOK, serverinternal.ResponsesClientsYaml{Clients: clients})
}

// Add client
// (POST /clients)
func (s Server) AddClient(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, serverinternal.ResponsesNewClient{
		Login:      "unique_client_login",
		RegisterAt: "https://auth.insolar.io/auth/set-password?code=asdfEGXDKBOI346sdfg",
	})
}

// Activate or deactivate client
// (PUT /clients)
func (s Server) ActivateClient(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, serverinternal.ResponsesClientYaml{
		Active:     true,
		Login:      "unique_client_login",
		ModifiedAt: 1590669972,
	})
}

// Get client
// (GET /clients/{login})
func (s Server) GetClient(ctx echo.Context, login string) error {
	return ctx.JSON(http.StatusOK, serverinternal.ResponsesClientYaml{
		Active:     true,
		Login:      "unique_client_login",
		ModifiedAt: 1590669972,
	})
}

// NewServer returns instance of API server
func NewServer() *Server {
	return &Server{}
}

func localSrv() {
	e := echo.New()
	apiServer := NewServer()
	serverexternal.RegisterHandlers(e, apiServer)
	serverinternal.RegisterHandlers(e, apiServer)
	e.Logger.Fatal(e.Start(":8080"))
}

func main() {
	go localSrv()
	time.Sleep(1 * time.Second)
	cfg := &clientexternal.Configuration{
		BasePath: "",
		Host:     "localhost:8080",
		Scheme:   "http",
		Debug:    true,
	}
	c := clientexternal.NewAPIClient(cfg)
	resp, err := c.DefaultApi.SetPassword(context.Background(), "", clientexternal.ParametersSetPassword{
		Login:    "abc",
		Password: "zxc",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp external: %v\n", resp)
	cfg2 := &clientinternal.Configuration{
		BasePath: "",
		Host:     "localhost:8080",
		Scheme:   "http",
		Debug:    true,
	}
	ci := clientinternal.NewAPIClient(cfg2)
	_, resp, err = ci.DefaultApi.ClientArray(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp internal: %v\n", resp)
}
