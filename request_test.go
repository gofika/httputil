package httputil

import (
	"bytes"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostJSON(t *testing.T) {
	resp, err := PostJSON(endpoint("/post"), map[string]string{"foo": "bar"})
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	type PostJSONResp struct {
		JSON map[string]string `json:"json"`
	}
	res, err := ReadJSON[PostJSONResp](resp)
	assert.Nil(t, err)
	assert.Equal(t, "bar", res.JSON["foo"])
}

func TestPostForm(t *testing.T) {
	resp, err := PostForm(endpoint("/post"), url.Values{"foo": {"bar"}})
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	type PostFormResp struct {
		Form map[string]string `json:"form"`
	}
	res, err := ReadJSON[PostFormResp](resp)
	assert.Nil(t, err)
	assert.Equal(t, "bar", res.Form["foo"])
}

func TestPostFormFiles(t *testing.T) {
	resp, err := PostFormFiles(endpoint("/post"), url.Values{"foo": {"bar"}}, []*UploadFile{
		{
			FieldName: "file",
			FileName:  "test.txt",
			Body:      bytes.NewBufferString("hello world"),
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	type PostFormFilesResp struct {
		Files map[string]string `json:"files"`
		Form  map[string]string `json:"form"`
	}
	res, err := ReadJSON[PostFormFilesResp](resp)
	assert.Nil(t, err)
	assert.Equal(t, "bar", res.Form["foo"])
	assert.Equal(t, "hello world", res.Files["file"])
}
