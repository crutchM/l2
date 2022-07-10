package main

import "testing"

type TestCase struct {
	inp string
	exp string
}

func TestDecode(t *testing.T) {
	cases := []TestCase{
		{
			inp: "a4bc2d5e",
			exp: "aaaabccddddde",
		},
		{
			inp: "asrs6/22",
			exp: "asrssssss22",
		},
		{
			inp: "11",
			exp: "incorrect row",
		},
		{
			inp: "",
			exp: "incorrect row",
		},
		{
			inp: "/22",
			exp: "22",
		},
	}

	for _, val := range cases {
		res := unpack(val.inp)
		if res != val.exp {
			t.Errorf("unpac %s: Expected %s, got %s", val.inp, val.exp, res)
		}

	}
}
