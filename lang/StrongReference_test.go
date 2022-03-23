package lang

import (
	"fmt"
	"github.com/tursom/GoCollections/util"
	"reflect"
	"testing"
)

func TestStrongReference_ToString(t *testing.T) {
	reference := util.NewStrongReference[Int](1)
	fmt.Println(reference.ToString())
	fmt.Println(reference)
}

func TestStrongReference_GetReference(t *testing.T) {
	type fields struct {
		reference Int
	}
	tests := []struct {
		name   string
		fields fields
		want   Int
	}{
		{"1", fields{reference: 1}, 1},
		{"2", fields{reference: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := util.NewStrongReference(tt.fields.reference)
			if got := r.GetReference(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrongReference_SetReference(t *testing.T) {
	type fields struct {
		reference Int
	}
	type args struct {
		reference Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Int
	}{
		{"1", fields{reference: 1}, args{2}, 2},
		{"2", fields{reference: 2}, args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := util.NewStrongReference(tt.fields.reference)
			r.SetReference(tt.args.reference)
			if got := r.GetReference(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReference() = %v, want %v", got, tt.want)
			}
		})
	}
}
