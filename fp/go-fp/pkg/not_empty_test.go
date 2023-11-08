package pkg_test

import (
	"example/pkg"
	"reflect"
	"testing"
)

func TestNotEmpty(t *testing.T) {
	t.Run("to array", func(t *testing.T) {
		for _, tt := range []struct {
			name string
			args pkg.NotEmpty[string]
			want []string
		}{
			{
				name: "two elements",
				args: pkg.NotEmpty[string]{
					First: "a",
					Left:  []string{"b"},
				},
				want: []string{"a", "b"},
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				got := pkg.ToArray(tt.args)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ToArray() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("from array", func(t *testing.T) {
		for _, tt := range []struct {
			name   string
			args   []string
			want   pkg.NotEmpty[string]
			wantOk bool
		}{
			{
				name: "two elements",
				args: []string{"a", "b"},
				want: pkg.NotEmpty[string]{
					First: "a",
					Left:  []string{"b"},
				},
				wantOk: true,
			},
			{
				name:   "no elements",
				args:   []string{},
				want:   pkg.NotEmpty[string]{},
				wantOk: false,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				got, ok := pkg.FromArrayToNotEmpty(tt.args)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromArrayToNotEmpty() got = %v, want %v", got, tt.want)
				}
				if ok != tt.wantOk {
					t.Errorf("FromArrayToNotEmpty() ok = %v, want %v", ok, tt.wantOk)
				}
			})
		}
	})
}
