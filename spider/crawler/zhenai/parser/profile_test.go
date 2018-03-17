package parser

import (
	"testing"
	"io/ioutil"
	"hans/learn/spider/crawler/model"
)


func TestParseProfile(t *testing.T) {
	contents,err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents,"幽诺")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element,but was %v ",result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:"幽诺",
		Age:0,
		Marriage:"离异",
	}

	if profile != expected {
		t.Errorf("expected %v ,but was %v ",expected,profile)
	}


}