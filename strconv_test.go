package sconv

import "testing"

func TestStr19DecToU64(t *testing.T) {
	type args struct {
		s string
	}
	type wants struct {
		v  uint64
		ok bool
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{name: "", args: args{s: ""}, wants: wants{v: 0, ok: false}},
		{name: "", args: args{s: "0"}, wants: wants{v: 0, ok: true}},
		{name: "", args: args{s: "1"}, wants: wants{v: 1, ok: true}},
		{name: "", args: args{s: "12"}, wants: wants{v: 12, ok: true}},
		{name: "", args: args{s: "123"}, wants: wants{v: 123, ok: true}},
		{name: "", args: args{s: "1234"}, wants: wants{v: 1234, ok: true}},
		{name: "", args: args{s: "12345"}, wants: wants{v: 12345, ok: true}},
		{name: "", args: args{s: "123456"}, wants: wants{v: 123456, ok: true}},
		{name: "", args: args{s: "1234567"}, wants: wants{v: 1234567, ok: true}},
		{name: "", args: args{s: "12345678"}, wants: wants{v: 12345678, ok: true}},
		{name: "", args: args{s: "123456789"}, wants: wants{v: 123456789, ok: true}},
		{name: "", args: args{s: "1234567890"}, wants: wants{v: 1234567890, ok: true}},
		{name: "", args: args{s: "12345678901"}, wants: wants{v: 12345678901, ok: true}},
		{name: "", args: args{s: "123456789012"}, wants: wants{v: 123456789012, ok: true}},
		{name: "", args: args{s: "1234567890123"}, wants: wants{v: 1234567890123, ok: true}},
		{name: "", args: args{s: "12345678901234"}, wants: wants{v: 12345678901234, ok: true}},
		{name: "", args: args{s: "123456789012345"}, wants: wants{v: 123456789012345, ok: true}},
		{name: "", args: args{s: "1234567890123456"}, wants: wants{v: 1234567890123456, ok: true}},
		{name: "", args: args{s: "12345678901234567"}, wants: wants{v: 12345678901234567, ok: true}},
		{name: "", args: args{s: "123456789012345678"}, wants: wants{v: 123456789012345678, ok: true}},
		{name: "", args: args{s: "1234567890123456789"}, wants: wants{v: 1234567890123456789, ok: true}},
		{name: "", args: args{s: "12345678901234567890"}, wants: wants{v: 2345678901234567890, ok: false}},
		{name: "", args: args{s: "1234567890*"}, wants: wants{v: 0, ok: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Str19DecToU64(tt.args.s)
			if got != tt.wants.v {
				t.Errorf("Str19DecToU64() got = %v, want %v", got, tt.wants.v)
			}
			if got1 != tt.wants.ok {
				t.Errorf("Str19DecToU64() got1 = %v, want %v", got1, tt.wants.ok)
			}
		})
	}
}
