package main

import "easymemcache"
import "strconv"

type mything struct {
	number    int
	letters   string
	stringary []string
	intary    []int
}
func printmything(t mything) {
	print(strconv.Itoa(t.number)+"\n")
	print(t.letters+"\n")
	for _,s := range(t.stringary) {
		print("\t"+s+"\n")
	}
	for _,i := range(t.intary) {
		print("\t"+strconv.Itoa(i)+"\n")
	}
}

func main() {
	mc := easymemcache.New("localhost:11211")
	mc.Set("Aletter","a")
	var anum string
	mc.Get("Aletter",&anum)
	print("A letter:"+anum+"\n")
	var set_a=1
	mc.Set("A Number",set_a)
	var a int
	mc.Get("A Number",&a)
	print("A Number"+strconv.Itoa(a)+"\n")


	var mt mything
	mt.number=33
	mt.letters="things letters"
	mt.stringary = []string{"Test", "another"}
	mt.intary = []int{3,2,1}
	print("Printing mt\n")
	printmything(mt)
	err := mc.Set("mything",mt)
	if err != nil { panic(err) }

	var newt mything
	err = mc.Get("mything",&newt)
	if err != nil { panic(err) }
	print("Print newt\n")
	printmything(newt)

}
