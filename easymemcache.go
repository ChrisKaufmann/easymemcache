package easymemcache

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"strings"
)

func New(server string) Client {
	var c Client
	c.client = memcache.New(server)
	c.keys = make(map[string]int)
	c.prefix = ""
	return c
}

type Client struct {
	client *memcache.Client
	keys map[string]int
	prefix string
}

func (c Client) Set(key string, i interface{}) error {
	var timeout int32 = 86400
	err := c.SetTime(key, i, timeout)
	return err
}
func (c Client) SetTime(key string, i interface{}, t int32) (err error) {
	key = strings.Replace(key, " ", "_", -1)
	b, err := json.Marshal(i)
	//add it to the list of keys
	c.keys[key]=1
	if err != nil {
		return err
	}
	err = c.client.Set(&memcache.Item{Key: c.prefix+key, Value: []byte(b), Expiration: t})
	return err
}
func (c Client) Get(key string, i interface{}) error {
	key = strings.Replace(key, " ", "_", -1)
	thing, err := c.client.Get(c.prefix+key)
	if err != nil {
		return err
	}
	err = json.Unmarshal(thing.Value, &i)
	return err
}
func (c Client) Delete(key string) (err error) {
	err = c.client.Delete(key)
	delete(c.keys, key)
	return err
}
func (c Client) Keys() []string {
	var keys []string
	for k := range c.keys {
		keys = append(keys, k)
	}
	return keys
}
func (c Client) Find(s string) (rl []string) {
	for i,_ := range c.keys {
		if strings.Contains(i,s) {
			rl = append(rl,i)
		}
	}
	return rl
}
func (c Client) StartsWith(s string) (rl []string) {
	for i,_ := range c.keys {
		if strings.HasPrefix(i,s) {
			rl = append(rl,i)
		}
	}
	return rl
}
