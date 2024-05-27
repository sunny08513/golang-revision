# Question 2

Write a Go program that simulates a simple producer-consumer problem:

1. **Producer**: Creates a goroutine that generates a sequence of numbers from 1 to 20 and sends them to a channel.
2. **Consumer**: Creates two separate goroutines that receive numbers from the producer's channel, process them (for example, multiply by 2), and send the results to another channel.
3. **Printer**: Creates a final goroutine that receives the processed numbers from the consumers' channel and prints them to the console.

Your main function should ensure all goroutines complete their tasks before exiting. Also, ensure the producer does not produce more numbers until at least one consumer is ready to process.

