package opencl

/*
#cgo CFLAGS: -I/usr/include/CL
#include <opencl.h>
*/
import "C"

import (
  "errors"
  "fmt"
)

var errorAsString = []string{
  "success",
  "device not found",
  "device not available",
  "compiler not available",
  "mem object allocation failure",
  "out of resources",
  "out of host memory",
  "profiling info not available",
  "mem copy overlap",
  "image format mismatch",
  "image format not supported",
  "build program failure",
  "map failure",
  "misaligned sub buffer offset",
  "exec status error for events in wait list",
  "", // 15
  "", // 16
  "", // 17
  "", // 18
  "", // 19
  "", // 20
  "", // 21
  "", // 22
  "", // 23
  "", // 24
  "", // 25
  "", // 26
  "", // 27
  "", // 28
  "", // 29
  "invalid value",
  "invalid device type",
  "invalid platform",
  "invalid device",
  "invalid context",
  "invalid queue properties",
  "invalid command queue",
  "invalid host ptr",
  "invalid mem object",
  "invalid image format descriptor",
  "invalid image size",
  "invalid sampler",
  "invalid binary",
  "invalid build options",
  "invalid program",
  "invalid program executable",
  "invalid kernel name",
  "invalid kernel definition",
  "invalid kernel",
  "invalid arg index",
  "invalid arg value",
  "invalid arg size",
  "invalid kernel args",
  "invalid work dimension",
  "invalid work group size",
  "invalid work item size",
  "invalid global offset",
  "invalid event wait list",
  "invalid event",
  "invalid operation",
  "invalid gl object",
  "invalid buffer size",
  "invalid mip level",
  "invalid global work size",
  "invalid property",
}

func newError(code_ C.cl_int) (err error) {
  code := -int(code_)

  if code > 0 {
    if code < len(errorAsString) && errorAsString[code] != "" {
      err = errors.New(errorAsString[code])
    } else {
      err = fmt.Errorf("unknown error %d", -code)
    }
  }

  return
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// fill-column: 70
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
