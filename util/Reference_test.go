package util

import (
	"github.com/tursom/GoCollections/lang"
	"testing"
)

func TestStrongReference_Equals(t *testing.T) {
	type fields struct {
		reference int
	}
	type args struct {
		o lang.Object
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"1", fields{1}, args{NewStrongReference[int](1)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &StrongReference[int]{
				reference: tt.fields.reference,
			}
			if got := r.Equals(tt.args.o); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
