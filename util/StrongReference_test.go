/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package util

import (
	"reflect"
	"testing"

	"github.com/tursom/GoCollections/lang"
)

func TestStrongReference_GetReference(t *testing.T) {
	type fields struct {
		reference lang.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   lang.Int
	}{
		{"1", fields{reference: 1}, 1},
		{"2", fields{reference: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewStrongReference(tt.fields.reference)
			if got := r.GetReference(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrongReference_SetReference(t *testing.T) {
	type fields struct {
		reference lang.Int
	}
	type args struct {
		reference lang.Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   lang.Int
	}{
		{"1", fields{reference: 1}, args{2}, 2},
		{"2", fields{reference: 2}, args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewStrongReference(tt.fields.reference)
			r.SetReference(tt.args.reference)
			if got := r.GetReference(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReference() = %v, want %v", got, tt.want)
			}
		})
	}
}
