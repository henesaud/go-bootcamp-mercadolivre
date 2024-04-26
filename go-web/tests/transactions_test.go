package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/cmd/server/handler"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/internal/transactions"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/pkg/store"
	"github.com/henesaud/go-bootcamp-mercadolivre/go-web/pkg/web"
	"github.com/stretchr/testify/assert"
)

func extractIdFromResponse(data interface{}) uint64 {
	id := data.(map[string]interface{})["id"]
	return uint64(id.(float64))
}
func Test_GetTransaction(t *testing.T) {
	t.Run("Test_PatchTransaction_OK", func(t *testing.T) {
		var response web.Response

		r := createServer()

		req, rr := createRequestTest(http.MethodPost, "/transactions/", `{
			"code": "codeTest",
			"currency": "currencyTest",
			"amount": 100.0,
			"emiter": "emiterTest",
			"receiver": "receiverTest",
			"date": "01/01/2023"
		}`)
		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code)
		err := json.Unmarshal(rr.Body.Bytes(), &response.Data)
		assert.Nil(t, err)

		id := extractIdFromResponse(response.Data)
		patchUrl := "/transactions/" + strconv.FormatUint(uint64(id), 10)
		req, rr = createRequestTest(http.MethodPatch, patchUrl, `{
			"amount": 400.0
		}`)

		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)

		err = json.Unmarshal(rr.Body.Bytes(), &response.Data)
		assert.Nil(t, err)

	})

	t.Run("Test_DeleteTransaction", func(t *testing.T) {
		r := createServer()
		var response web.Response

		req, rr := createRequestTest(http.MethodDelete, "/transactions/54", "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Code)

		req, rr = createRequestTest(http.MethodPost, "/transactions/", `{
			"code": "codeTest",
			"currency": "currencyTest",
			"amount": 100.0,
			"emiter": "emiterTest",
			"receiver": "receiverTest",
			"date": "01/01/2023"
		}`)
		r.ServeHTTP(rr, req)
		err := json.Unmarshal(rr.Body.Bytes(), &response.Data)
		assert.Nil(t, err)
		id := extractIdFromResponse(response.Data)
		deleteUrl := "/transactions/" + strconv.FormatUint(uint64(id), 10)

		req, rr = createRequestTest(http.MethodDelete, deleteUrl, "")
		r.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.NewFileStore(store.FileType, "./transactions.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)
	r := gin.Default()

	pr := r.Group("/transactions")
	pr.POST("/", t.Store)
	pr.GET("/", t.All)
	pr.PATCH("/:id", t.UpdateAmount)
	pr.DELETE("/:id", t.Delete)
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}
