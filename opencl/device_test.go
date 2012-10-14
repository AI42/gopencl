package opencl

import (
  "testing"
)

func TestGetPlatforms(t *testing.T) {
  ps, err := getPlatforms()

  if err != nil {
    t.Error(err)
  } else {
    for i, p := range ps {
      t.Logf("%d name:       %s", i, p.Name())
      t.Logf("%d profile:    %s", i, p.Profile())
      t.Logf("%d vendor:     %s", i, p.Vendor())
      t.Logf("%d version:    %s", i, p.Version())
      t.Logf("%d extensions: %v", i, p.Extensions())

      for j, d := range p.Devices() {
        t.Logf("    %d name:               %s", j, d.Name())
        t.Logf("    %d version:            %s", j, d.Version())
        t.Logf("    %d opencl C version:   %s", j, d.OpenCLCVersion())
        t.Logf("    %d vendor:             %s", j, d.Vendor())
        t.Logf("    %d extensions:         %v", j, d.Extensions())
        t.Logf("    %d profile:            %s", j, d.Profile())
        t.Logf("    %d available:          %v", j, d.IsAvailable())
        t.Logf("    %d default:            %v", j, d.IsDefault())
        t.Logf("    %d CPU:                %v", j, d.IsCPU())
        t.Logf("    %d GPU:                %v", j, d.IsGPU())
        t.Logf("    %d vendor id:          %x", j, d.VendorID())
        t.Logf("    %d compiler:           %v", j, d.IsCompilerAvailable())
        t.Logf("    %d max units:          %d", j, d.MaxComputeUnits())
        t.Logf("    %d max constant args:  %d", j, d.MaxConstantArgs())
        t.Logf("    %d global memory size: %d bytes", j, d.GlobalMemorySize())
        t.Logf("    %d local memory size:  %d bytes", j, d.LocalMemorySize())
        t.Logf("    %d image support:      %v", j, d.ImageSupport())
        t.Logf("    %d error correction:   %v", j, d.ErrorCorrection())
        t.Logf("    %d unified memory:     %v", j, d.UnifiedMemory())
        t.Logf("    %d little endian:      %v", j, d.IsLittleEndian())
      }
    }
  }
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// fill-column: 70
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
