package easymemcache
import (
	"testing"
	"time"
)
var mc Client

func TestAll(t *testing.T) {
	testServer := "localhost:11211"
	mc = New(testServer)
	var err error
	var ts string

	//set by itself
	err = mc.Set("test","blah blah")
	if err != nil { t.Errorf("Set", err) }

	//set with get
	err = mc.Set("setwithget","setwithgetval")
	if err != nil { t.Errorf("setwithget", err) }
	err = mc.Get("setwithget",&ts)
	if err != nil { t.Errorf("setwithget", err) }
	if ts != "setwithgetval" { t.Errorf("setwithget", err) }

	//set time with get (should hit)
	err = mc.SetTime("setwithgettime", "setwithgettimeval", 30)
	if err != nil { t.Errorf("setwithgettime", err) }
	err = mc.Get("setwithgettime", &ts)
	if err != nil { t.Errorf("setwithgettime", err) }
	if ts != "setwithgettimeval" { t.Errorf("setwithgettime", err) }

	//set time with get (should miss)
	err = mc.SetTime("setwithgettime2", "setwithgettime2val", 1)
	if err != nil { t.Errorf("setwithgettime2", err) }
	time.Sleep(2*time.Second)
	err = mc.Get("setwithgettime2", &ts)
	if err == nil {t.Errorf("setwithgettime2 should have been nil", err) }

	//test delete
	mc.Delete("setwithgettime2")
	err = mc.Get("setwithgettime2", &ts)
	if err == nil {t.Errorf("Delete didn't work, got key for setwithgettime2", err) }

	//At this point, the list of keys should have just the first 3 as values (last was deleted)
	tk := []string{"test", "setwithget", "setwithgettime"}
	avkeys := allvals(mc.Keys())
	avtk := allvals(tk)
	if len(avkeys) != len(avtk) { t.Error("Length of keys is wrong", err) }

	//Test find a couple of times
	tfl := mc.Find("test")
	if len(tfl) != 1 { t.Error("Wrong number of keys for find", err) }
	tswg := mc.Find("setwith")
	if len(tswg) != 2 {t.Error("Wrong number of keys for find setwith", err) }
	twf := mc.Find("shouldn'texist")
	if len(twf) != 0 {t.Error("Wrong number of keys for find shouldn't exist", err) }

	//test StartsWith
	sfl := mc.StartsWith("test")
	if len(sfl) != 1 { t.Error("Wrong number of keys for startswith", err) }
	sswg := mc.StartsWith("setwith")
	if len(sswg) != 2 {t.Error("Wrong number of keys for startswith setwith", err) }
	swf := mc.StartsWith("shouldn'texist")
	if len(swf) != 0 {t.Error("Wrong number of keys for startswith shouldn't exist", err) }

}

func allvals(s []string) []string {
	var r []string
	for _,i := range(s) {
		var v string
		err := mc.Get(i, &v)
		if err == nil {
			//normally we'd test for !- nil, but we expect some errors here and just want the passes
			r = append(r, v)
		}
	}
	return r
}

func ps( s []string) {
	for _,i := range(s) {
		print(i+"\n")
	}
	print("\n")
}
