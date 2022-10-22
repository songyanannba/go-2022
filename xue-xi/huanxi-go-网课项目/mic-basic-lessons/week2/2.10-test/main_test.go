package main

import "testing"

func TestShow(t *testing.T) {
	res := Show()
	want := "main"
	if res == want {
		t.Logf("show = %v wang = %v", res, want)
	} else {
		t.Errorf("show =%v want = %v \n", res, want)
	}
}

func TestShowWithTable(t *testing.T) {
	test := []struct {
		name string
		want string
	}{
		{
			"面向对象",
			"面向对象",
		},
		{
			"jie口",
			"接口",
		},
		{
			"111",
			"111",
		},
		{
			"222",
			"main",
		},
	}
	for _, item := range test {
		t.Run(item.name, func(t *testing.T) {
			if got := Show(); got != item.want {
				t.Errorf("show =%v want = %v \n", got, item.want)
			}
		})
	}
}
