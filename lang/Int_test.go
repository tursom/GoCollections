package lang

import "testing"

func TestInt_Compare(t *testing.T) {
	type args struct {
		t Int
	}
	tests := []struct {
		name string
		i    Int
		args args
		want int
	}{
		{"", 1, args{1}, 0},
		{"", 1, args{0}, -1},
		{"", 1, args{2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Compare(tt.args.t); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_Equals(t *testing.T) {
	type args struct {
		e Object
	}
	tests := []struct {
		name string
		i    Int
		args args
		want bool
	}{
		{"", 1, args{Int(1)}, true},
		{"", 1, args{Int(0)}, false},
		{"", 1, args{Int(2)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Equals(tt.args.e); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt_String(t *testing.T) {
	tests := []struct {
		name string
		i    Int
		want string
	}{
		{"", 1, "1"},
		{"", 2, "2"},
		{"", 3, "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
