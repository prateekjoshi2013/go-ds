### Built-in Data Structures

- Go has **four basic types** of data structures: **arrays, slices, maps, and structs**

- Arrays: 
  
  - An ordered sequence of elements of the **same type**. 
  
  - Array **sizes are static**; they are set when the array is defined and cannot be changed subsequently. 
  
  - In Go, **arrays are values**. This is an important difference, because in some languages an array is a pointer to the first item in the array. 
    
    - This means **if you pass an array to a function you will be passing a copy of the array, and this could be expensive**.

- Slices:
  
  - Represent **an ordered sequence of elements built on top of arrays**
  
  - Are **used much more often than arrays because of their flexibility**
  
  - Slices have **no fixed length**.
  
  - **A slice is a struct that consists of a pointer to an array, the length of the segment of the array, and the capacity of the underlying array**.
  
- Maps

- **Data structures that associate the values of one type (called the key) with values of another type (called the value)**.

- Internally, **a map is a pointer to runtime.hmap structure**.

### Collections

- Unlike many other programming languages, Go does not provide other basic data structures (although there is the container package in the standard library it has very few implementations). 

-  We wiil create some other data structures that are not provided either in the language or in the standard library, and it’s not very difficult to re-create them.

  - Queue
  - Stack
  - Set
  - Linked list
  - Heap
  - Graph
  
- **None of the data structures are concurrency-safe**. 

- This is because **fundamentally they are built on the three basic Go data structures—arrays, slices, and maps—and these are not concurrency-safe**. 

- **To avoid race conditions, you can add a mutex to the data structure and use it to lock the data structure before any reads or writes**. 

- **The sync package has a RWMutex that you can use for this purpose**.