# Atomic(s) - `sync/atomic`

**High Level languages** are cool, but nothing beats code that executes directly on the **CPU Level**.
In Go those are **_atomic(s)_**. Running concurrent code can both be: **Slow** and **Cumbersome**
(due to the **synchronisation** involved). However, the `atomic.Value` type is the best
of both worlds: **blazing fast** and **less cumbersome** compared to other **sync types**.

When you think of **Atomic(s)**, it's nothing fancy or anything Go related,
it's really something which is as **close** as possible to the **CPU Level**,
which works directly with **CPU registers**, hence the blazing fast speed.

When you think of **Atoms**, these are the **smallest** **Indivisible** units. The same approach
pretty much is valid when it comes to Concurrency.

These are some examples of the simplest **Atomic Operations**: `+`, `-`, `*`, `/`, `=`
that respect all principles of **Atomicity**.

So why is it called **Atomic**, and what exactly does the term **ATOMIC** mean?
Something is Atomic, if it **HAPPENS** from point A to point Z in its **ENTIRETY**
**WITHOUT BEING INTERRUPTED** by other processes in the same **Context**.

### Atomicity

An operation is considered **Atomic**, if within the **context** it is operating it is `Indivisible` or `Uninterruptible`.
The word that's important here is **Context**. Something may be **atomic** in one **context** but not in another.

Operations that are **Atomic** in the **Context** of your **Process**, may not be atomic in the context of the **Operating System**.
Operations that are **Atomic** in the **Context** of your **Operating System**, may not be **Atomic** in the **Context** of your **Machine**,
and **Operations** that are **Atomic** in the **Context** of your **Machine**, may not be **Atomic** in the **Context** of your **Application**.

**Indivisible** and **Uninterruptible** means, that within the **Context** you've defined something that is **Atomic**
`WILL HAPPEN` in its `ENTIRETY`, without anything else `HAPPENING SIMULTANEOUSLY` in the same context.

Let's take for example the statement `i++`. This may look like one Atomic operation, but in reality this happens:

1. Retrieve the value of `i`
2. Increment the value of `i`
3. Store the value of `i`

While **each** of the operations above are **atomic**, **the combination** of these in a certain **Context may not be**,
which also means, **combining several atomic operations** does not necessarily produce a **bigger atomic operation**.

### Memory Access Synchronization

**Memory Access Synchronization** comes hand in hand with **Atomicity**. In order to make a specific operation Atomic,
we MUST allow some kind of Memory Access Synchronization. If **multiple go routines** try to **read/write** from/to the **same
memory space**, they need a way to **communicate**, that **only 1 go routine at a time** can **read or write** from that **memory space**.

This kind of memory access synchronization is done through a process called **Mutual Exclusion**, which provides a **Locking
Mechanism**, so that when 1 concurrent process (go routine) tries to access some kind of memory space, that memory space
can be **guarded by a Mutex** which **holds a Lock** on that **memory space**.

The **Lock** can only be **acquired** by **only 1 go routine** at a time, thus making the rest of concurrent processes (go routines)
**wait** their turn to acquire the Lock, in order to read from or write to that memory space. This way a certain **memory
space** in the **context** of an operation is considered to be **Atomic**, resulting in **deterministic** and **correct** results when
**multiple concurrent operations** are involved in the game.

### `sync/atomic`

Most functionalities provided by the `sync/atomic` package are there to work with **numbers**.
These types are:

- `int32`
- `int64`
- `uint32`
- `uint64`
- `uintptr`
- `Pointer`

Respectively available functions for each of those types:

- `Load__`
- `Add__`
- `Store__`
- `Swap__`
- `CompareAndSwap__`

Besides, working mostly with numbers (except `Pointer`) the `sync/atomic` package
also provides the `atomic.Value` type which has the handy methods:

- `Store`
- `Load`

which facilitate Concurrency for a more complex/hybrid type.

**Note:** A more complex type `sync.Map` provided by the sync package
also makes use of `atomic.Value`

### Atomic(s) Downsides

- Mostly works with **Numbers**
- `atomic.Value` is generic (**type assertion** required)
- `atomic.Value` has **Limited Context** (cumbersome/impossible to use on sequential data)
- `atomic.Value` is **not always Atomic** (if not careful, we may run into race conditions)
