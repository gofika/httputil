package httputil

const HTTPBIN_ENDPOINT = "https://httpbin.org"

func endpoint(path string) string {
	return HTTPBIN_ENDPOINT + path
}
