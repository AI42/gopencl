package opencl

/*
#cgo CFLAGS: -I/usr/include/CL
#include <opencl.h>
*/
import "C"

import (
  "reflect"
  "strings"
  "unsafe"
)

type getInfoFunc func(interface{}, interface{}, C.size_t, unsafe.Pointer, *C.size_t) C.cl_int
type getInfo func(interface{}, interface{}, interface{}, interface{}) (int, error)

var trimSpace = reflect.ValueOf(strings.TrimSpace)
var fields = reflect.ValueOf(strings.Fields)

func makeGetInfo(get getInfoFunc) getInfo {
  return func(id interface{}, info interface{}, holdPtr interface{}, dst interface{}) (size int, err error) {
    holdPtrV := reflect.ValueOf(holdPtr)
    holdType := holdPtrV.Elem().Type()
    holdSize := C.size_t(holdType.Size())
    holdPtrC := unsafe.Pointer(holdPtrV.Pointer())
    var retSize C.size_t

    if ret := get(id, info, C.size_t(holdSize), holdPtrC, &retSize); ret != 0 {
      err = newError(ret)
    } else {
      size = int(retSize)

      if dst != holdPtr {
        dptr := reflect.ValueOf(dst)
        var d reflect.Value

        switch dst.(type) {
        case *bool:
          isFalse := reflect.DeepEqual(reflect.Zero(holdType).Interface(), holdPtrV.Elem().Interface())
          d = reflect.ValueOf(!isFalse)

        case *string:
          d = holdPtrV.Elem().Slice(0, size-1).Convert(dptr.Elem().Type())
          ret := trimSpace.Call([]reflect.Value{d})
          d = ret[0]

        case *[]string:
          d = holdPtrV.Elem().Slice(0, size-1).Convert(reflect.TypeOf(""))
          ret := fields.Call([]reflect.Value{d})
          d = ret[0]

        default:
          d = holdPtrV.Elem().Convert(dptr.Elem().Type())
        }

        dptr.Elem().Set(d)
      }
    }

    return
  }
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// fill-column: 70
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
