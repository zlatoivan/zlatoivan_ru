package color

import (
	"strconv"
	"testing"
)

func TestStatus(t *testing.T) {
	type args struct {
		status int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "< 200",
			args: args{
				status: 100,
			},
			want: Blue(strconv.Itoa(100)),
		},
		{
			name: "< 300",
			args: args{
				status: 200,
			},
			want: Green(strconv.Itoa(200)),
		},
		{
			name: "< 400",
			args: args{
				status: 301,
			},
			want: Cyan(strconv.Itoa(301)),
		},
		{
			name: "< 500",
			args: args{
				status: 404,
			},
			want: Yellow(strconv.Itoa(404)),
		},
		{
			name: ">= 500",
			args: args{
				status: 505,
			},
			want: Red(strconv.Itoa(505)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Status(tt.args.status); got != tt.want {
				t.Errorf("Status() = %v, want %v", got, tt.want)
			}
		})
	}
}
