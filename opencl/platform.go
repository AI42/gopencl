package opencl

/*
#cgo CFLAGS: -I/usr/include/CL
#include <opencl.h>
*/
import "C"

import (
  "unsafe"
)

type Platform struct {
  extensions []string
  id         C.cl_platform_id
  name       string
  profile    string
  vendor     string
  version    string
  devices    []*Device
}

const (
  maxPlatforms       = 16
  maxPlatformDevices = 64
)

var getPlatformInfo = makeGetInfo(func(id interface{}, info interface{}, size C.size_t, ptr unsafe.Pointer, retSize *C.size_t) C.cl_int {
  return C.clGetPlatformInfo(id.(C.cl_platform_id), info.(C.cl_platform_info), size, ptr, retSize)
})

func getPlatformByID(id C.cl_platform_id) (p *Platform, err error) {
  var numDevices C.cl_uint
  deviceIDs := [maxPlatformDevices]C.cl_device_id{}
  deviceIDsPtr := (*C.cl_device_id)(unsafe.Pointer(&deviceIDs))

  ret := C.clGetDeviceIDs(id, C.CL_DEVICE_TYPE_ALL, C.cl_uint(maxPlatformDevices), deviceIDsPtr, &numDevices)

  if ret != 0 {
    err = newError(ret)
    return
  }

  p = &Platform{
    id:      id,
    devices: make([]*Device, int(numDevices)),
  }

  strMap := []struct {
    dst  *string
    info C.cl_platform_info
  }{
    {&p.name, C.CL_PLATFORM_NAME},
    {&p.profile, C.CL_PLATFORM_PROFILE},
    {&p.vendor, C.CL_PLATFORM_VENDOR},
    {&p.version, C.CL_PLATFORM_VERSION},
  }

  for _, s := range strMap {
    var str [256]byte

    if _, err = getPlatformInfo(id, s.info, &str, s.dst); err != nil {
      return
    }
  }

  for i := 0; i < int(numDevices); i++ {
    if p.devices[i], err = getDeviceByID(deviceIDs[i]); err != nil {
      break
    }
  }

  return
}

func getPlatforms() (ps []*Platform, err error) {
  var numPlatforms C.cl_uint
  platformIDs := [maxPlatforms]C.cl_platform_id{}
  platformIDsPtr := (*C.cl_platform_id)(unsafe.Pointer(&platformIDs))

  ret := C.clGetPlatformIDs(maxPlatforms, platformIDsPtr, &numPlatforms)

  if ret != 0 {
    err = newError(ret)
    return
  }

  ps = make([]*Platform, int(numPlatforms))

  for i := 0; i < int(numPlatforms); i++ {
    if ps[i], err = getPlatformByID(platformIDs[i]); err != nil {
      break
    }
  }

  return
}

func (p *Platform) Devices() []*Device {
  return p.devices
}

func (p *Platform) Name() string {
  return p.name
}

func (p *Platform) Profile() string {
  return p.profile
}

func (p *Platform) Vendor() string {
  return p.vendor
}

func (p *Platform) Version() string {
  return p.version
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// fill-column: 70
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
