package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"io/ioutil"
	"net/http"
	"strings"
)

var host = "http://localhost:3333"
var res *http.Response

func aRequestIsSentToTheEndpoint(method, endpoint string) error {
	var reader = strings.NewReader("")
	var request, err = http.NewRequest(method, host+endpoint, reader)
	if err != nil {
		return fmt.Errorf("could not create request: %s", err.Error())
	}
	res, err = http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("could not send request: %s", err.Error())
	}
	return nil
}

func theHTTPresponseCodeShouldBe(expectedCode int) error {
	if res.StatusCode != expectedCode {
		return fmt.Errorf("expected status code was %d but got %d", expectedCode, res.StatusCode)
	}
	return nil
}

func theResponseContentShouldBe(expectedContent string) error {
	body, _ := ioutil.ReadAll(res.Body)
	if expectedContent != string(body) {
		return fmt.Errorf("expected content was %s but got %s instead", expectedContent, body)
	}
	return nil
}

func FeatureContext(ctx *godog.ScenarioContext) {
	ctx.Step(`^a "([^"]*)" request is sent to the endpoint "([^"]*)"$`, aRequestIsSentToTheEndpoint)
	ctx.Step(`^the HTTP-response code should be "([^"]*)"$`, theHTTPresponseCodeShouldBe)
	ctx.Step(`^the response content should be "([^"]*)"$`, theResponseContentShouldBe)
}
