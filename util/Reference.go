/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package util

import "github.com/tursom/GoCollections/lang"

type (
	Reference[R any] interface {
		GetReference() R
	}
	MutableReference[R any] interface {
		SetReference(reference R)
		GetReference() R
	}
	StrongReference[R any] struct {
		lang.BaseObject
		reference R
	}
	NoReference[R any] struct {
		lang.BaseObject
	}
)

func NewStrongReference[R any](reference R) *StrongReference[R] {
	return &StrongReference[R]{
		reference: reference,
	}
}

func NewNoReference[R any]() *NoReference[R] {
	return &NoReference[R]{}
}

func (r *StrongReference[R]) Equals(o lang.Object) bool {
	if r == o {
		return true
	}
	if objectReference, ok := o.(Reference[R]); ok {
		if reference, ok := any(r.reference).(lang.Equable); ok {
			if referenceObject, ok := any(objectReference.GetReference()).(lang.Object); ok {
				return reference.Equals(referenceObject)
			}
		}
	}
	return r == o
}

func (r *StrongReference[R]) SetReference(reference R) {
	r.reference = reference
}

func (r *StrongReference[R]) GetReference() R {
	return r.reference
}

func (r *NoReference[R]) SetReference(_ R) {
}

func (r *NoReference[R]) GetReference() R {
	return lang.Nil[R]()
}
