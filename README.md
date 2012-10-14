# gopencl

Go to OpenCL compiler WIP.

## TODO

 * a tool to parse Go source code and translate kernels written in Go to OpenCL
   C language, then compile them to IR and generate another Go source file with
   all kernels wrapped as normal Go functions
 * the generated Go code should check (in `init()`) if it's possible to use
   compiled kernels and fall back to native Go functions if not
 * the tool should also have lots of debug options, like dumping IR, ASM, LLVM,
   etc.
 * should support nvidia, radeon and intel OpenCL >= 1.1
