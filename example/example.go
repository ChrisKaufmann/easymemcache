package main

import "github.com/ChrisKaufmann/easymemcache"
import "fmt"
import "os/exec"
import "strconv"

type mything struct {
	Number    int
	Letters   string
	Stringary []string
	Intary    []int
}

func printmything(t mything) {
	print(strconv.Itoa(t.Number) + "\n")
	print(t.Letters + "\n")
	for _, s := range t.Stringary {
		print("\t" + s + "\n")
	}
	for _, i := range t.Intary {
		print("\t" + strconv.Itoa(i) + "\n")
	}
}

func main() {
	mc := easymemcache.New("localhost:11211")

	//Generic set and get
	mc.Set("Some_key", "some value")
	var return_val string
	mc.Get("Some_key", &return_val)

	//Letter and number variables
	mc.Set("Aletter", "a")
	anum, err := mc.Gets("Aletter")
	if err != nil {
		panic(err)
	}
	print("A letter:" + anum + "\n")

	mc.Set("A_Number", 32)
	a, err := mc.Geti("A_Number")
	if err != nil {
		panic(err)
	}
	print("A Number" + strconv.Itoa(a) + "\n")

	//Increment and decrement a key
	a,_ = mc.Geti("A_Number")
	print("Before Increment: "+strconv.Itoa(a)+"\n")
	mc.Increment("A_Number",1)
	a,_ = mc.Geti("A_Number")
	print("After Increment: "+strconv.Itoa(a)+"\n")
	mc.Decrement("A_Number",11)
	a,_ = mc.Geti("A_Number")
	print("After Decrement: "+strconv.Itoa(a)+"\n")

	//Setting and retrieving a struct
	mt := mything{
		Number:    33,
		Letters:   "my letters",
		Stringary: []string{"Test", "another"},
		Intary:    []int{3, 2, 3},
	}
	err = mc.Set("my_thing", mt)
	if err != nil {
		panic(err)
	}

	var newt mything
	err = mc.Get("my_thing", &newt)
	if err != nil {
		panic(err)
	}
	print("Print newt\n")
	printmything(newt)

	//Deleting one
	mc.Delete("my_thing")

	//Getting a list
	print("Print a list of keys\n")
	el := mc.Keys()
	for _, i := range el {
		print("\t" + i + "\n")
	}

	//Get keys with "letter" in the name
	print("List of keys with 'letter' in the name\n")
	el = mc.Find("letter")
	for _, i := range el {
		print("\t" + i + "\n")
	}

	//Get keys that start with "A"
	print("List of keys that start with 'A'\n")
	el = mc.StartsWith("A")
	for _, i := range el {
		print("\t" + i + "\n")
	}

	//Deleting all
	print("Deleting all keys\n")
	print("Length before: "+strconv.Itoa(len(mc.Keys()))+"\n")
	err = mc.DeleteAll()
	if err != nil {
		panic(err)
	}
	el = mc.Keys()
	for _, i := range el {
		print("\t" + i + "\n")
	}
	print("Length after: "+strconv.Itoa(len(mc.Keys()))+"\n")


	//Use a prefix in front of all your keys for separation
	mc.Prefix="easymemcachego-example"
	mc.Print()
	//it will be transparent
	mc.Set("with_prefix",42)
	cmd := "echo 'get easymemcachego-examplewith_prefix' | nc 127.0.0.1 11211"
	out,err := exec.Command("sh","-c",cmd).Output()
	fmt.Printf("Output from nc: %q\n", out)

}
