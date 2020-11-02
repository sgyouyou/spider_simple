package parser

import (
	"fmt"
	"io/ioutil"
	"spider_simple/engine"
	"testing"
)

func TestParseDetail(t *testing.T) {
	// 造页面
	byte, err := ioutil.ReadFile("detail_test.html")
	if err != nil {
		panic(err)
	}
	//reader := bytes.NewReader(byte)

	// 造数据
	expected := engine.Details{
		Rank: "SUV",
		Engine: "185kW(3.0L自然吸气)",
		PowerType: "汽油机",
		Gearbox: "5挡AT",
		Size: "4695×1815×1825",
		BodyStructure: "5门 7座 SUV",
		ListedTime: "2015",
		OilWear: 10.8,
	}

	//把拿到的和自己造的对比
	res := ParseDetail(byte)
	if err != nil {
		panic(err)
	}
	got  := res.Items[0]
	if expected != got {
		t.Errorf("Expected: %v, but got %v\n", expected, got)
	}
	fmt.Printf("Result: %v\n", res)
}
