package codec

import (
	"reflect"
	"testing"
)

func TestGobCodec(t *testing.T) {

	println(reflect.TypeOf(C).String())
	println(reflect.ValueOf(C).String())

}
