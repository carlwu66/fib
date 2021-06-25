package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	resp, e := http.Get("http://127.0.0.1:8001/fibonacci/46")
	CheckError(e)
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	CheckError(e)
	assert.Equal(t, "1836311903", string(body))
}

func TestOrder(t *testing.T) {
	resp, e := http.Get("http://127.0.0.1:8001/order/1836311903")
	CheckError(e)
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	CheckError(e)
	assert.Equal(t, "46", string(body))
}
