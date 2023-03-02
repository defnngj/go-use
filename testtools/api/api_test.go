package api

import (
	"net/http"
	"testing"
	"time"

	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func Client() http.Client {
	cli := &http.Client{
		Timeout: time.Second * 10,
	}
	return *cli
}

func TestGetSample(t *testing.T) {
	cli := Client()
	apitest.New().EnableNetworking(&cli).
		Get("http://httpbin.org/get").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetParams(t *testing.T) {
	cli := Client()
	apitest.New().
		EnableNetworking(&cli).
		Intercept(func(req *http.Request) {
			req.URL.RawQuery = "id=1&name=jack"
		}).
		Get("http://httpbin.org/get").
		Expect(t).
		Assert(
			jsonpath.Contains(`$.args.id`, "1")).
		Assert(
			jsonpath.Equal(`$.args.name`, "jack")).
		End()
}

func TestPostFormData(t *testing.T) {
	cli := Client()
	apitest.New().
		EnableNetworking(&cli).
		Post("http://httpbin.org/post").
		FormData("key1", "value1").
		FormData("key2", "value2").
		Expect(t).
		Assert(
			jsonpath.Chain().
				Equal(`$.form.key1`, "value1").
				Equal(`$.form.key2`, "value2").
				End()).
		End()
}

func TestPostJson(t *testing.T) {
	cli := Client()
	apitest.New().
		EnableNetworking(&cli).
		Post("http://httpbin.org/post").
		JSON(`{"message": "hi"}`).
		Expect(t).
		Assert(
			jsonpath.Chain().
				Contains(`$.data`, "message").
				Contains(`$.data`, "hi").
				End()).
		End()
}
