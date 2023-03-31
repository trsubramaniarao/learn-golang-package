package pkg

import "testing"

func TestAdder(t *testing.T) {
	type args struct {
		a int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "The usual happy path.",
			args: args{
				a: 1,
				b: []int{1},
			},
			want: 2,
		},
		{
			name: "One argument.",
			args: args{
				a: 1,
			},
			want: 1,
		},
		{
			name: "More than two arguments.",
			args: args{
				a: 1,
				b: []int{2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Adder(tt.args.a, tt.args.b...); got != tt.want {
				t.Errorf("Adder() = %v, want %v", got, tt.want)
			}
		})
	}
}
