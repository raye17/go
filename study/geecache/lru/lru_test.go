package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestNew(t *testing.T) {
	type args struct {
		maxBytes  int64
		onEvicted func(string, Value)
	}
	tests := []struct {
		name string
		args args
		want *Cache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.maxBytes, tt.args.onEvicted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Add(t *testing.T) {
	type args struct {
		key   string
		value Value
	}
	tests := []struct {
		name string
		c    *Cache
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Add(tt.args.key, tt.args.value)
		})
	}
}

func TestCache_RemoveOldest(t *testing.T) {
	tests := []struct {
		name string
		c    *Cache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.RemoveOldest()
		})
	}
}

func TestCache_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		c         *Cache
		args      args
		wantValue Value
		wantOk    bool
	}{
		// TODO: Add test cases.
		{
			name:      "get",
			c:         New(0, nil),
			args:      args{key: "key"},
			wantValue: String("value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotOk := tt.c.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Cache.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Cache.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
