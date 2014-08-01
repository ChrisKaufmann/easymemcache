package easymemcache

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"strings"
)

func New(server string) Client {
	var c Client
	c.MClient = memcache.New(server)
	c.Prefix = ""
	return c
}

type Client struct {
	MClient *memcache.Client
	Prefix string
}

func (c Client) Print() {
	print("Prefix:" + c.Prefix + "\n")
}
func (c Client) KeyList(kp ...string) (l map[string]int) {
	var mykey = "easymemcache_keys"
	if len(l)<1 {l=make(map[string]int)}
	//if l == nil {l=make(map[string]int)}
	if len(kp) > 0 {
		var k string
		k = kp[0]
		l = c.KeyList()
		l[k]=1
		b, err := json.Marshal(l)
		err = c.MClient.Set(&memcache.Item{Key: mykey, Value:  []byte(b), Expiration: 86400})
		if err != nil {print(err.Error())}
	}
	c.Get(mykey, &l)
	return l
}
func (c Client) KeyListDelete(kp ...string) (err error) {
	var mykey = "easymemcache_keys"
	for _,k := range kp {
		l := c.KeyList()
		delete(l,k)
		b, err := json.Marshal(l)
		err = c.MClient.Set(&memcache.Item{Key: mykey, Value:  []byte(b), Expiration: 86400})
		if err != nil {print(err.Error())}
	}
	return err
}
func (c Client) Set(key string, i interface{}) error {
	var timeout int32 = 86400
	err := c.SetTime(key, i, timeout)
	return err
}
func (c Client) SetTime(key string, i interface{}, t int32) (err error) {
	key = strings.Replace(key, " ", "_", -1)
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	//add it to the list of keys
	c.KeyList(key)
	err = c.MClient.Set(&memcache.Item{Key: c.Prefix + key, Value: []byte(b), Expiration: t})
	return err
}
func (c Client) Increment(key string, incr uint64) {
	c.MClient.Increment(c.Prefix+key,incr)
}
func (c Client) Decrement(key string, decr uint64) {
	c.MClient.Decrement(c.Prefix+key,decr)
}
func (c Client) Get(key string, i interface{}) error {
	key = strings.Replace(key, " ", "_", -1)
	thing, err := c.MClient.Get(c.Prefix + key)
	if err != nil {
		return err
	}
	err = json.Unmarshal(thing.Value, &i)
	return err
}
func (c Client) GetOr(key string, i interface{}, f func()) {
	err := c.Get(key, i)
	if err != nil {
		f()
		c.Set(key,i)
	}
}
func (c Client) Gets(key string) (s string, err error) {
	err = c.Get(key, &s)
	return s, err
}
func (c Client) Geti(key string) (i int, err error) {
	err = c.Get(key, &i)
	return i, err
}
func (c Client) Count() (i int) {
	return len(c.KeyList())
}
func (c Client) Delete(otherkeys ...string) (err error) {
	for _,k := range otherkeys {
		k = strings.Replace(k, " ", "_", -1)
		err = c.MClient.Delete(c.Prefix+k)
		c.KeyListDelete(k)
		if err != nil {return err}
	}
	return err
}
func (c Client) DeleteLike(s string) (err error) {
	kl := c.Find(s)
	for _,k := range kl {
		err = c.Delete(k)
		if err != nil {
			return err
		}
		c.KeyListDelete(k)
	}
	return err
}
func (c Client) DeleteAll() (err error) {
	for k := range c.KeyList() {
		err = c.Delete(k)
		if err != nil {
			return err
		}
	}
	err=c.MClient.DeleteAll()
	return err
}
func (c Client) Keys() []string {
	var keys []string
	for k := range c.KeyList() {
		keys = append(keys, k)
	}
	return keys
}
func (c Client) Find(s string) (rl []string) {
	for i, _ := range c.KeyList() {
		if strings.Contains(i, s) {
			rl = append(rl, i)
		}
	}
	return rl
}
func (c Client) StartsWith(s string) (rl []string) {
	for i, _ := range c.KeyList() {
		if strings.HasPrefix(i, s) {
			rl = append(rl, i)
		}
	}
	return rl
}
