package routers_test

import (
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/appleboy/gofight/v2"
	"github.com/buger/jsonparser"
	"github.com/go-playground/assert/v2"
	"github.com/or0986113303/remotedebug/service/apiman"
	"github.com/steinfletcher/apitest"
)

func Test_test(t *testing.T) {
	req := gofight.New()
	server := apiman.New()
	req.GET("/api/v1/test/message").
		SetHeader(gofight.H{
			"Contest-Type": "aapplication/json; charset=utf-8",
		}).
		Run(server.GetRouter(), func(res gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(res.Body.String())

			message, _ := jsonparser.GetString(data, "message")
			benchmark := "Hello world"
			assert.Equal(t, benchmark, message)
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "application/json; charset=utf-8", res.HeaderMap.Get("Content-Type"))
		})
}

func Test_test_report(t *testing.T) {
	cookieJar, _ := cookiejar.New(nil)
	cli := &http.Client{
		Timeout: time.Second * 1,
		Jar:     cookieJar,
	}

	server := apiman.New()
	server.StartServer()
	time.Sleep(1 * time.Second)
	defer server.StopServer()

	apitest.New().
		EnableNetworking(cli).
		Report(apitest.SequenceDiagram("../../../test/report/routers/test/message")).
		Get("http://127.0.0.1:5000/api/v1/test/message").
		Expect(t).
		Body(`{"message": "Hello world"}`).
		Status(http.StatusOK).
		End()

}
