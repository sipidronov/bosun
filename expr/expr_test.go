package expr

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestExprSimple(t *testing.T) {
	var exprTests = []struct {
		input  string
		output Value
	}{
		{"!1", 0},
		{"-2", -2},
		{"1.444-010+2*3e2-4/5+0xff", 847.644},
		{"1>2", 0},
		{"3>2", 1},
		{"1==1", 1},
		{"1==2", 0},
		{"1!=01", 0},
		{"1!=2", 1},
		{"1<2", 1},
		{"2<1", 0},
		{"1||0", 1},
		{"0||0", 0},
		{"1&&0", 0},
		{"1&&2", 1},
		{"1<=0", 0},
		{"1<=1", 1},
		{"1<=2", 1},
		{"1>=0", 1},
		{"1>=1", 1},
		{"1>=2", 0},
	}

	for _, et := range exprTests {
		e, err := New(et.input)
		if err != nil {
			t.Error(err)
			break
		}
		r, err := e.Execute("")
		if err != nil {
			t.Error(err)
			break
		} else if len(r) != 1 {
			t.Error("bad r len", len(r))
			break
		} else if len(r[0].Group) != 0 {
			t.Error("bad group len", r[0].Group)
			break
		} else if r[0].Value != et.output {
			t.Errorf("expected %v, got %v: %v", et.output, r[0].Value, et.input)
		}
	}
}

func TestExprQuery(t *testing.T) {
	e, err := New(`avg([avg:proc.stat.cpu{host=*}], "5m") > 4e7`)
	if err != nil {
		t.Fatal(err)
	}
	r, err := e.Execute("")
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.MarshalIndent(&r, "", "  ")
	fmt.Println(string(b))
}
