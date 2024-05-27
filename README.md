# golang-revision

In this repo we will add golang concepts one by one

1. GO-ROTINE
Goroutines are lightweight compared to threads and are managed by the Go runtime.
A goroutine is a lightweight thread of execution in Go. Goroutines are managed by the Go runtime and are distinct from operating system threads, which are typically heavier in terms of resources.

Here's a more detailed explanation:
```
Concurrency Model:
    Goroutines are a key component of Go's concurrency model, allowing functions to be executed concurrently. Multiple goroutines can run simultaneously within the same address space, enabling efficient use of multi-core processors.

Low Overhead:
   Goroutines have a lower overhead compared to threads because they are multiplexed onto a smaller number of OS threads. This means that you can create thousands of goroutines without much additional memory overhead.

Go Runtime: 
   Goroutines are managed by the Go runtime, which handles the scheduling of goroutines onto OS threads. This makes it easier to write concurrent code in Go compared to languages where developers need to manage threads explicitly.

Communication: 
   Goroutines communicate with each other using channels, which are built-in synchronization primitives in Go. Channels allow goroutines to safely send and receive data, enabling them to coordinate their actions.
```
In summary, goroutines are a lightweight and efficient way to achieve concurrency in Go, allowing developers to write highly concurrent programs with ease.

Link to revise golang functionality

https://www.sohamkamani.com/golang/