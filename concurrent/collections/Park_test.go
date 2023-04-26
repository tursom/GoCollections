package collections

import (
	"testing"

	"github.com/tursom/GoCollections/util/time"
)

func TestPark_Park(t *testing.T) {
	var p Park
	t1 := time.Now()
	go func() {
		<-time.After(time.Second)

		p.Unpark()
	}()

	p.Park()
	t2 := time.Now()

	sub := t2.Sub(t1)
	if sub > time.Duration(float64(time.Second)*1.01) ||
		sub < time.Duration(float64(time.Second)*0.09) {
		t.Fatal(sub)
	}
}

func TestPark_ParkT(t *testing.T) {
	var p Park
	t1 := time.Now()
	p.ParkT(time.Second)
	t2 := time.Now()

	sub := t2.Sub(t1)
	if sub > time.Duration(float64(time.Second)*1.01) ||
		sub < time.Duration(float64(time.Second)*0.09) {
		t.Fatal(sub)
	}
}
