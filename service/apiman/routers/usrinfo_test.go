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
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

var getUserMock = apitest.NewMock().
	Get("http://127.0.0.1:5000/api/v1/usr/16701").
	RespondWith().
	Body(`{"name": "Mir", "id": "16701"}`).
	Status(http.StatusOK).
	End()

func TestAPI_GetUserInfo(t *testing.T) {
	req := gofight.New()
	server := apiman.New()
	req.GET("/api/v1/usr/16701").
		SetHeader(gofight.H{
			"Contest-Type": "aapplication/json; charset=utf-8",
		}).
		Run(server.GetRouter(), func(res gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(res.Body.String())
			usrname, _ := jsonparser.GetString(data, "name")
			banchmarkname := "Mir"
			assert.Equal(t, banchmarkname, usrname)
			assert.Equal(t, http.StatusOK, res.Code)
			assert.Equal(t, "application/json; charset=utf-8", res.HeaderMap.Get("Content-Type"))
		})
}

func TestAPI_GetUserInfo_Success_report(t *testing.T) {
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
		Report(apitest.SequenceDiagram("../../../test/report/routers/usr/usrinfo/success")).
		Get("http://127.0.0.1:5000/api/v1/usr/16701").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestAPI_GetUserInfo_NotFound_report(t *testing.T) {
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
		Report(apitest.SequenceDiagram("../../../test/report/routers/usr/usrinfo/notfound")).
		Get("http://127.0.0.1:5000/api/v1/usr/16").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestAPI_GetUserInfo_CookieMatching_report(t *testing.T) {
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
		Report(apitest.SequenceDiagram("../../../test/report/routers/usr/usrinfo/cookmatch")).
		Get("http://127.0.0.1:5000/api/v1/usr/16701").
		Expect(t).
		Cookies(apitest.NewCookie("TomsFavouriteDrink").
			Value("Beer").
			Path("/")).
		Status(http.StatusOK).
		End()
}

func TestAPI_GetUserInfo_JSONPath_report(t *testing.T) {
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
		Report(apitest.SequenceDiagram("../../../test/report/routers/usr/usrinfo/jsonpath")).
		Mocks(getUserMock).
		Get("http://127.0.0.1:5000/api/v1/usr/16701").
		Expect(t).
		Assert(jsonpath.Equal(`$.id`, "16701")).
		Status(http.StatusOK).
		End()
}
