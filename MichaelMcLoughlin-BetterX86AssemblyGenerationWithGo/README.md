# Better x86 Assembly Generation from Go

Michael McLoughlin ([@mmcloughlin](https://github.com/mmcloughlin))

* [Slides](slides.pdf)
* [`avo` Project](https://github.com/mmcloughlin/avo)

This talk will equip you with the tools to safely write lightning fast assembly functions for Go. We will introduce Go assembly and the role it plays in the Go ecosystem, promote best practices for assembly development and demonstrate how code generation tools can manage complexity in large assembly projects.

In the Go standard library and beyond, assembly language is used to provide access to architecture and OS-specific features, to accelerate hot loops in scientific code, and provide high-performance cryptographic routines. For these specific applications correctness is paramount, yet assembly functions can be painstaking to write and nearly impossible to review; the subtle crypto/elliptic bug in 6 lines of a 2000+ line assembly file led to CVE-2017-8932 and should serve as a warning to those entering this game. Due to this, code generation tools have gained traction in other languages.

This talk will demonstrate the use of code generation for Go assembly, with a particular focus on the avo library for x86 assembly generation. Such techniques simplify assembly functions by enabling the use of Go control structures, offering automatic register allocation and handling interaction with Go types. Using examples, we will demonstrate how these methods help produce maintainable high-performance functions. Intermediate to advanced Gophers will leave with greater confidence in the world of Go assembly.
