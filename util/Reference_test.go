/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package util

import (
	"testing"

	"github.com/tursom/GoCollections/lang"
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
