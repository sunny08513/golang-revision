# Question 1

Write a Go program that creates three goroutines:

1. The first goroutine generates a sequence of numbers from 1 to 10 and sends them to a channel.
2. The second goroutine receives numbers from the first goroutine's channel, squares each number, and sends the squared numbers to another channel.
3. The third goroutine receives the squared numbers from the second goroutine's channel and prints each squared number to the console.
Your main function should wait for all goroutines to complete their tasks before exiting.

**Requirements**
Use channels to communicate between goroutines.
Ensure that your program does not exit before all goroutines have completed their work.


# Bufferd Channel
Buffered channels in Go allow you to specify the capacity of the channel at the time of creation. Unlike unbuffered channels, where sends and receives block until the other side is ready, buffered channels allow sends to proceed without blocking if there's room in the buffer.

**Benefits and Use Cases**

1. **Decoupling Goroutines**: Buffered channels can decouple the sending and receiving goroutines, allowing the sender to proceed without waiting for the receiver.
2. **Throughput Improvement**: In some cases, buffered channels can improve throughput by reducing the time spent waiting for sends and receives to synchronize.
3. **Burst Handling**: Buffered channels are useful for handling bursts of messages where the rate of production temporarily exceeds the rate of consumption.

**Considerations**

**Size Limit**: The buffer size should be chosen carefully based on the application's requirements. Too small a buffer might not provide the desired benefits, while too large a buffer can consume unnecessary memory.

**Complexity**: Using buffered channels introduces more complexity into the program. You need to consider buffer size and potential deadlocks if the buffer size is not sufficient to handle the message bursts.