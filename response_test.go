package httputil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAll(t *testing.T) {
	resp, err := Get(endpoint("/get"))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	body, err := ReadAll(resp)
	assert.Nil(t, err)
	assert.NotEmpty(t, body)
}

func TestReadString(t *testing.T) {
	resp, err := Get(endpoint("/get"))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	body, err := ReadString(resp)
	assert.Nil(t, err)
	assert.NotEmpty(t, body)
	assert.Contains(t, body, "/get")
}

func TestReadAnyJSON(t *testing.T) {
	resp, err := Get(endpoint("/get"))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	type GetResp struct {
		URL string `json:"url"`
	}
	var v GetResp
	err = ReadAnyJSON(resp, &v)
	assert.Nil(t, err)
	assert.Contains(t, v.URL, "/get")
}

func TestReadJSON(t *testing.T) {
	resp, err := Get(endpoint("/get"))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	type GetResp struct {
		URL string `json:"url"`
	}
	v, err := ReadJSON[GetResp](resp)
	assert.Nil(t, err)
	assert.Contains(t, v.URL, "/get")
}

func TestSaveFile(t *testing.T) {
	resp, err := Get(endpoint("/get"))
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	written, err := SaveFile(resp, "test.txt")
	assert.Nil(t, err)
	assert.NotZero(t, written)
	data, err := os.ReadFile("test.txt")
	assert.Nil(t, err)
	assert.NotEmpty(t, data)
	assert.Contains(t, string(data), "/get")
	err = os.Remove("test.txt")
	assert.Nil(t, err)
}
