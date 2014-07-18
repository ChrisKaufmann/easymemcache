## About
Super easy to use memcached client in go, built on github.com/bradfitz/gomemcache/memcache
(http://golang.org/).

## Installing

### *go get*

	$ go get github.com/ChrisKaufmann/easymemcache

Src will be in:

	$GOROOT/src/pkg/github.com/ChrisKaufmann/easymemcache

Update with 'go get -u github.com/ChrisKaufmann/easymemcache'

## Example

	import "github.com/ChrisKaufmann/easymemcache"
	func main() {
		mc := easymemcache.New("localhost:11211")
		mc.Set("my_key","my value")
		mystr := mc.Gets("my_key")
		mc.Set("My_int_key": 32)
		myint := mc.Geti("My_int_key")

		// For structs, or really any data:
		var mything structtype
		mc.Set("struct_key",mything)
		mc.Get("struct_key", &mything)

		//Delete
		mc.Delete("my_key")

		//Get all keys used with this:
		keys := mc.Keys()

		//Find keys
		fkeys := mc.Find("ruct")
	}

Full Example in example/example.go

## Docs

godoc github.com/ChrisKaufmann/easymemcache

