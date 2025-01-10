package str

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TEST1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "TEST2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "TEST3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "TEST4",
			args: args{
				s: "ababac",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TEST1",
			args: args{
				s1: "listen",
				s2: "silent",
			},
			want: true,
		},
		{
			name: "TEST2",
			args: args{
				s1: "triangle",
				s2: "integral",
			},
			want: true,
		},
		{
			name: "TEST3",
			args: args{
				s1: "hello",
				s2: "world",
			},
			want: false,
		},
		{
			name: "TEST4",
			args: args{
				s1: "race",
				s2: "care",
			},
			want: true,
		},
		{
			name: "TEST5",
			args: args{
				s1: "anagram",
				s2: "nagaram",
			},
			want: true,
		},
		{
			name: "TEST6",
			args: args{
				s1: "rat",
				s2: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
