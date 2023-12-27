### Behind the Scenes

Whether an operation executes concurrently or not, it has to either to with
some kind of **computation/calculus** or **memory access**. In the computers' era,
things are controlled and coordinated by the CPU, whether it's memory or computation.

#### CPU Registers

Alongside many things a CPU has, one of the most important things
are its **registers**.

What is a Register?
CPU registers are high-speed storage locations inside the microprocessor used
for **Memory Addressing**, **Data Operation**, and **Processing**.

A register is a **small** **high-speed memory** inside the **CPU**. It is used to **store** **temporary results**.
Registers are designed to be assessed at a much **higher speed** than conventional memory.
Some registers are **general purpose**, while others are **classified** according to the function they perform.

Some of these registers are:

- `Memory Address Register` (MAR)
- `Stack Control Register` (ACR)
- `Memory Data Register` (MDR)
- `Flag Register` (FR)
- `Accumulator Register` (AC)
- `Instruction Pointer Register` (IPR)
- `Data Register`
- `Address or Segment Register` (AR or SR)
- `Memory Buffer Register` (MBR)
- `Index Register`
- `Program Counter` (PC)

Shortly every **memory/computation** operation has to deal with **registers**.
A register can either be **used** or **not used**. What that means for concurrent code,
is that every process will get its turn to work with the register and there's
no need for any kind of **memory access synchronization** or other types of
measures we usually take in a high level language which executes concurrent code.

If we go back to **Atomic(s)**, these are **Go ASM (Assembly)** routines that execute directly on CPU
thus, **not needing** any kind of **memory access synchronisation**, and they are **blazing fast**
if benchmarked against other types in the sync package.

For more on what kind of registers are available check out the [CPU Registers](./registers.md)

#### Cross Compilation Workflow

Along with many other reasons why people choose Go, it also provides a beautiful and super
elegant code **cross compilation model**, which works beautifully pretty much on **any
operating system** and **architecture** available nowadays.

When you think of architectures, there's plenty of them. These are the ones Go supports:

- `386`
- `amd64`
- `arm`
- `arm64`
- `mips`
- `mipsle`
- `mips64`
- `mipsle64`
- `ppc64`
- `ppc64le`
- `riscv64`
- `s390x`
- `wasm`

Here's a small overview of how things really happen under the hood, from **compiling** your
Go source code, till you have a **generated binary**, that actually gets to be executed on
your **specific OS** and **Architecture**:

1. Go source Code (obviously :D)
2. The `$GOOS` and `$GOARCH` variables, accepting a bunch of OSes and Architectures
3. `Go Compiler` which compiles your code given the Go Code and `$GOOS` & `$GOARCH`
4. Once the compiler is done, it'll generate a `main.s`, which are `Go` pseudo `ASM` (Assembly) instructions
5. That is then picked up by the Go ASM (Assembler) tool and goes through the `obj` library
6. The `obj` library maps the **pseudo instructions** to **real architecture instructions**
   generating an object file i.e `main.o` (this process is called assembling)
7. Then the `Go Linker` tool will pick the object file and **generate a binary** out of it (this process is called linking)
8. The generated binary is nothing but 0s and 1s, which is the only thing the CPU gets to eat

#### `go build` Workflow

When running the `go build` command, pretty much the exact same process described before happens.

1. Apply the **Cross Compilation** phase for the `main` package usually the `main.go` containing the `main()` func
2. Look for any **Go ASM files**, usually the ones that end with `.s` for the provided specific architecture.
3. **Create object files** for those ASM files
4. **Link** all **objects** and **generate** the final **binary**.

If your regular good old Go code, contains `functions` with `missing body`, the code
will expect some **external ASM declarations**, pretty much like is done in most C/C++ code.
This is crucial to making the **link between Go code and ASM code**.

Just like regular go files (`.go`) the ASM ones (`.s`) the same
`name_$GOOS_GOARCH.go` and respectively `name_$GOOS_GOARCH.s` rules apply.
Naming your files using the `$GOOS_$GOARCH` combination, will make the
build tool pick only the ones for the provided **OS** and respective **Architecture**.

For a more detailed example, have a look [here](https://github.com/golang-basics/concurrency/blob/master/atomics/asm/main.go)

#### Go ASM

Most statically typed languages only have 3 phases: `Compilation`->`Assembling`->`Linking`

Go has obviously more phases, to easier **accommodate multiple architectures** and make **cross
compilation** easy & make **builds faster** by eliminating some heavy work on
the assembly process.

**Go** has its own **Assembler**, which as I mentioned before generates **ASM like instructions**
which are pretty **architecture agnostic** and are pretty much pseudo instructions
meant to be used internally by the `obj` library, which then **maps** against **real
architecture instructions** and generates real **ASM instructions** for your
specific **OS** and **Architecture**.

Here are some ASM symbols and abbreviations Go uses internally

- `FP` Frame pointer: arguments and locals.
- `PC` Program counter: jumps and branches.
- `SB` Static base pointer: global symbols.
- `SP` Stack pointer: top of stack.

For more checkout out

- `$GOROOT/src/cmd/asm/internal/arch`
- `$GOROOT/src/cmd/internal/obj`

#### 32bit vs 64bit

Most Operating Systems out there run on either **32bit** or **64bit** architecture.
The question is, which one is better?

The answer is pretty simple, obviously 64bit.

There are 2 significant differences.

1. **Maximum number of ALU operations** performed: `32` and `64` respectively.

Which means, the CPU will be able to process that many **arithmetic** and **logical**
operations in a 32bit and respectively 64bit system.

2. **Max RAM memory available**: `3.25 GB` and respectively `16 EB` (16B GB)

Which of course means **larger CPU registers** 32bit vs 64bit values.
A 64-bit register can theoretically reference 18,446,744,073,709,551,616 bytes,
or 17,179,869,184 gigabytes (16 exabytes) of memory.
This is several million times more than an average workstation would need to access.

### Tips

```shell script
# cd into asm directory
cd asm

# generate object file from Go's pseudo assembly instructions
go tool asm sqrt_amd64.s

# compiles the go program and generates an object file
# with pseudo assembly instructions
go tool compile -S main.go

# compiles the go program and generates an object file
# with pseudo assembly instructions
# and saves the instructions in the specified file
go tool compile -S main.go > main.s

# generate a binary from the specified object file
go tool link main.o

# compiles and generates a binary from main.go
go build -o exec main.go
# dumps the assembly instructions from the generated binary
go tool objdump -s main.main exec
```
