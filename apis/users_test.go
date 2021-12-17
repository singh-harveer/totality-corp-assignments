package apis_test

import (
	"net/http"
	"testing"
	"totality/users/apis"
	"totality/users/mocks"

	"github.com/gin-gonic/gin"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	var client, err = mocks.NewUserClient()
	if err != nil {
		t.Fatalf("failed to create mock client :%v", err)
	}

	var handler = apis.NewHandler(client)
	_ = handler

	var testcases = []struct {
		name   string
		ginCTX *gin.Context
		status int
		want   string
	}{
		{
			name: "Ok",

			ginCTX: &gin.Context{
				Params: gin.Params{{
					Key:   "id",
					Value: "1",
				}},
			},
		},
		{
			name: "notFound",

			ginCTX: &gin.Context{
				Params: gin.Params{{
					Key:   "id",
					Value: "11",
				}},
			},
		},
		{
			name: "invalidParamValue",

			ginCTX: &gin.Context{
				Params: gin.Params{{
					Key:   "id",
					Value: "ewuyre",
				}},
			},
			status: http.StatusBadRequest,
		},
		{
			name: "missingParam",

			ginCTX: &gin.Context{},
			status: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// TODO- in progress..
		})

	}
}
