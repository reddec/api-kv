package client

import (
	"context"
	"fmt"
)

func ExampleDial() {
	ctx := context.Background()

	// errors omitted for demonstration purpose
	kv, _ := Dial("https://example.com", "my-app", "deadbeaf")
	// set item
	_ = kv.Set(ctx, "foo", []byte("bar"))
	// get item
	value, _ := kv.Get(ctx, "foo")
	fmt.Println("Got", string(value))
	// delete item
	_ = kv.Delete(ctx, "foo")
}

func ExampleKV_Keys() {
	ctx := context.Background()

	// errors omitted for demonstration purpose
	kv, _ := Dial("https://example.com", "my-app", "deadbeaf")
	// set item
	_ = kv.Set(ctx, "foo", []byte("bar"))
	_ = kv.Set(ctx, "bar", []byte("baz"))
	// iterate
	it := kv.Keys()
	for it.Next(ctx) {
		for _, key := range it.Keys() {
			fmt.Println(key)
		}
	}
	// check errors
	if it.Error() != nil {
		panic(it.Error())
	}
}
