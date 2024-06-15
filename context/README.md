# context

**What is the purpose of the context package in Go?**

The context package provides a way to pass cancellation signals, deadlines, and other request-scoped values across API boundaries.

It is used to manage the lifecycle of operations and facilitate cancellation and timeout handling.

The main purpose of using Context is to manage long-running processes (like an HTTP request, a database call, or a network call) in a way that doesn’t waste resources.

**Benifits of context**

- Cancel long running process, We can send signal through context to cancel process across API boundries
- Set deadlines for processes to complete
- pass common scoped data with the application

**USE Case**

- Pass request Id for function calls and go routines that belongs to same HTTP call.
- Errors encountered when fetching data from a database, we can send signal database down through context.


https://www.sohamkamani.com/golang/context/
