# question 3

**Write a Go program that simulates a web server log processing system:**

 **Log Generator**: 
Create a goroutine that generates log entries. Each log entry should be a simple string like "Log Entry X", where X is a number from 1 to 50. Send each log entry to a channel.

**Log Processor**:
 Create three separate goroutines that receive log entries from the log generator's channel. Each processor should simulate processing the log entry by printing it to the console along with the processor's ID (e.g., "Processor 1: Log Entry 1").

**Coordinator**: Ensure that the main function waits for all log entries to be processed before exiting.

**Requirements**

- Use channels to communicate between the log generator and log processors.
- Use a WaitGroup to ensure the main function waits for all goroutines to complete.
- Ensure that the log generator waits if all processors are busy.
