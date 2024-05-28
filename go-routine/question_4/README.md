# question 4

Write a Go program that simulates a task scheduling system:

1. Task Generator:
 Create a goroutine that generates a series of tasks. Each task should be a string like "Task X", where X is a number from 1 to 100. Send each task to a channel.

2. Workers: 
Create five separate goroutines that act as workers. Each worker should receive tasks from the task generator's channel, process them by printing the task along with the worker's ID, and then mark the task as done.
3. Task Completion: 
Create a final goroutine that monitors the completion of all tasks and prints a "All tasks completed" message once all tasks are done.

**Requirements**
Use channels to communicate between the task generator and workers.
Use a WaitGroup to ensure the main function waits for all goroutines to complete.
Implement a mechanism to signal when all tasks have been completed.