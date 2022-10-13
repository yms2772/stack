# Installation
`go get github.com/yms2772/stack`

# Example
```Go
package main

import (
	"github.com/yms2772/stack"
	"github.com/yms2772/stack/comparator"
)

func main() {
	s := stack.New[int]()           // Create a new stack
	                                // Available type:
	                                // ~int | ~int8 | ~int16 | ~int32 | ~int64 
					//~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr 
					//~float32 | ~float64 
					//~string
	
	s.Push(1)                       // Push data to the stack
	
	s.Peek()                        // Peek data from the stack (not erase the last index)
	                                // Output: 1
				                    
	s.Empty()                       // Check if the stack is empty
	                                // Output: false
				                    
	s.Search(1)                     // Search data index number from the stack
	                                // Output: 0
				                    
	s.Size()                        // Get the size of the stack
	                                // Output: 1
	
	s.Sort()                        // Sort the stack in ascending order		
	s.Sort(comparator.ASC)          // Sort the stack in ascending order		
	s.Sort(comparator.DESC)         // Sort the stack in descending order
	
	s.Index(0)                      // Get the stack data of the index number

	s.Pop()                         // Pop from the stack (erase the last index)
	                                // Output: 1
									
	s.Push(4)
	s.Push(5)
	s.Push(6)
	
	s.Remove(1)                     // Remove stack from the stack
	                                // Result: []int{4, 6}
					                
	s.Clean()                       // Clean stack
}
```
