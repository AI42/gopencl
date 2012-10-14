package opencl

/*
#cgo CFLAGS: -I/usr/include/CL
#include <opencl.h>
*/
import "C"

import (
  "unsafe"
)

const (
  deviceTypeDefault     = (1 << 0)
  deviceTypeCPU         = (1 << 1)
  deviceTypeGPU         = (1 << 2)
  deviceTypeAccelerator = (1 << 3)
)

type Device struct {
  available         bool
  compilerAvailable bool
  endianLittle      bool
  errorCorrection   bool
  extensions        []string
  globalMemorySize  uint64
  id                C.cl_device_id
  images            bool
  image2DMaxWidth   int
  image2DMaxHeight  int
  image3DMaxWidth   int
  image3DMaxHeight  int
  image3DMaxDepth   int
  localMemorySize   uint64
  maxComputeUnits   int
  maxConstantArgs   int
  name              string
  nativeKernels     bool
  openclCVersion    string
  profile           string
  typeMask          int
  unifiedMemory     bool
  vendor            string
  vendorID          int
  version           string
}

var getDeviceInfo = makeGetInfo(func(id interface{}, info interface{}, size C.size_t, ptr unsafe.Pointer, retSize *C.size_t) C.cl_int {
  return C.clGetDeviceInfo(id.(C.cl_device_id), info.(C.cl_device_info), size, ptr, retSize)
})

func getDeviceByID(id C.cl_device_id) (d *Device, err error) {
  var cl_uint C.cl_uint
  var cl_ulong C.cl_ulong
  var cl_device_type C.cl_device_type
  var cl_bool C.cl_bool
  var str [4096]byte

  d = &Device{
    id: id,
  }

  getMap := []struct {
    info C.cl_device_info
    dst  interface{}
    hold interface{}
  }{
    {C.CL_DEVICE_AVAILABLE, &d.available, &cl_bool},
    {C.CL_DEVICE_COMPILER_AVAILABLE, &d.compilerAvailable, &cl_bool},
    {C.CL_DEVICE_ENDIAN_LITTLE, &d.endianLittle, &cl_bool},
    {C.CL_DEVICE_EXTENSIONS, &d.extensions, &str},
    {C.CL_DEVICE_ERROR_CORRECTION_SUPPORT, &d.errorCorrection, &cl_bool},
    {C.CL_DEVICE_GLOBAL_MEM_SIZE, &d.globalMemorySize, &cl_ulong},
    {C.CL_DEVICE_HOST_UNIFIED_MEMORY, &d.unifiedMemory, &cl_bool},
    {C.CL_DEVICE_IMAGE_SUPPORT, &d.images, &cl_bool},
    {C.CL_DEVICE_LOCAL_MEM_SIZE, &d.localMemorySize, &cl_ulong},
    {C.CL_DEVICE_MAX_COMPUTE_UNITS, &d.maxComputeUnits, &cl_uint},
    {C.CL_DEVICE_MAX_CONSTANT_ARGS, &d.maxConstantArgs, &cl_uint},
    {C.CL_DEVICE_NAME, &d.name, &str},
    {C.CL_DEVICE_OPENCL_C_VERSION, &d.openclCVersion, &str},
    {C.CL_DEVICE_PROFILE, &d.profile, &str},
    {C.CL_DEVICE_TYPE, &d.typeMask, &cl_device_type},
    {C.CL_DEVICE_VENDOR, &d.vendor, &str},
    {C.CL_DEVICE_VENDOR_ID, &d.vendorID, &cl_uint},
    {C.CL_DEVICE_VERSION, &d.version, &str},
  }

  for _, v := range getMap {
    if _, err = getDeviceInfo(id, v.info, v.hold, v.dst); err != nil {
      return
    }
  }

  return
}

func (d *Device) ErrorCorrection() bool {
  return d.errorCorrection
}

func (d *Device) Extensions() []string {
  return d.extensions
}

func (d *Device) GlobalMemorySize() uint64 {
  return d.globalMemorySize
}

func (d *Device) IsAccelerator() bool {
  return d.typeMask&deviceTypeAccelerator != 0
}

func (d *Device) IsAvailable() bool {
  return d.available
}

func (d *Device) IsDefault() bool {
  return d.typeMask&deviceTypeDefault != 0
}

func (d *Device) IsCPU() bool {
  return d.typeMask&deviceTypeCPU != 0
}

func (d *Device) IsCompilerAvailable() bool {
  return d.compilerAvailable
}

func (d *Device) IsGPU() bool {
  return d.typeMask&deviceTypeGPU != 0
}

func (d *Device) IsLittleEndian() bool {
  return d.endianLittle
}

func (d *Device) ImageSupport() bool {
  return d.images
}

func (d *Device) LocalMemorySize() uint64 {
  return d.localMemorySize
}

func (d *Device) MaxComputeUnits() int {
  return d.maxComputeUnits
}

func (d *Device) MaxConstantArgs() int {
  return d.maxConstantArgs
}

func (d *Device) MaxImageDimensions2D() (int, int) {
  return d.image2DMaxWidth, d.image2DMaxHeight
}

func (d *Device) MaxImageDimensions3D() (int, int, int) {
  return d.image3DMaxWidth, d.image3DMaxHeight, d.image3DMaxDepth
}

func (d *Device) Name() string {
  return d.name
}

func (d *Device) NativeKernelSupport() bool {
  return d.nativeKernels
}

func (d *Device) OpenCLCVersion() string {
  return d.openclCVersion
}

func (d *Device) Profile() string {
  return d.profile
}

func (d *Device) UnifiedMemory() bool {
  return d.unifiedMemory
}

func (d *Device) Vendor() string {
  return d.vendor
}

func (d *Device) VendorID() int {
  return d.vendorID
}

func (d *Device) Version() string {
  return d.version
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// fill-column: 70
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
