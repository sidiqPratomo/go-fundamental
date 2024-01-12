package main

import "fmt"

type Queue struct {
	items []int
}

func (q Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q Queue) Size() int {
	return len(q.items)
}

func main() {
	queue := Queue{}

	// Enqueue elements
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Dequeue elements
	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		fmt.Println("Dequeued:", item)
	}

	// Check the size of the queue
	fmt.Println("Queue size:", queue.Size())

	array()

	hashTable()
}

// =========================================
// datastructure static array

func array() {
	// Declaring an array of integers with a fixed size of 5
	var numbers [5]int

	// Initializing elements of the array
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// Accessing elements of the array
	fmt.Println("Element at index 0:", numbers[0])
	fmt.Println("Element at index 3:", numbers[3])

	// Length of the array (number of elements)
	fmt.Println("Length of the array:", len(numbers))

	// Iterating over the array using a for loop
	fmt.Println("Array elements:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

//================================================
// Hash table biasa digunakan untuk fitur search dan lebih cepat

const arraySize = 10

// Node represents a key-value pair in the linked list
type Node struct {
	key   string
	value interface{}
	next  *Node
}

// HashTable represents the hash table data structure
type HashTable struct {
	array [arraySize]*Node
}

// hashFunction takes a key and returns its hash value (a simple hash function)
func hashFunction(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % arraySize
}

// Insert inserts a key-value pair into the hash table
func (ht HashTable) Insert(key string, value interface{}) {
	index := hashFunction(key)
	newNode := &Node{key: key, value: value}

	if ht.array[index] == nil {
		ht.array[index] = newNode
	} else {
		currentNode := ht.array[index]
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
}

// Find searches for a key in the hash table and returns its value (if found)
func (ht HashTable) Find(key string) (interface{}, bool) {
	index := hashFunction(key)

	currentNode := ht.array[index]
	for currentNode != nil {
		if currentNode.key == key {
			return currentNode.value, true
		}
		currentNode = currentNode.next
	}

	return nil, false
}

func hashTable() {

	ht := HashTable{}

	// Insert key-value pairs
	ht.Insert("apple", 5)
	ht.Insert("banana", 10)
	ht.Insert("orange", 15)

	// Find values by key
	value, found := ht.Find("apple")
	if found {
		fmt.Println("Value for 'apple':", value)
	} else {
		fmt.Println("'apple' not found.")
	}

	value, found = ht.Find("grape")
	if found {
		fmt.Println("Value for 'grape':", value)
	} else {
		fmt.Println("'grape' not found.")
	}

}
