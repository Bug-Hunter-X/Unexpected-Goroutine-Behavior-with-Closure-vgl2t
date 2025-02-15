# Unexpected Goroutine Behavior with Closure in Go

This example demonstrates a common issue in Go involving closures and goroutines.  The program intends to print numbers 0-9, but due to how Go handles variable scoping in goroutines, it actually prints 10 for every iteration.  This is because the goroutine accesses the variable `i` from its surrounding scope, by reference, not value.  By the time the goroutines actually run, the loop has already completed and `i` is 10.

## How to fix it
The solution below shows two common ways to correctly pass the value of `i` to the goroutine, either by passing a copy or using a helper function.