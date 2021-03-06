package linkedList

/*
2.1 Remove Dups: Remove duplicates from an unsorted linked list without a temp buffer
*/
// Hint: Hash map
func (ll *DoubleLinkedList) removeDups() {
	seen := make(map[int]struct{})
	node := ll.head
	count := ll.count

	for i := 0; i < count; i++ {
		if _, ok := seen[node.value]; ok {
			ll.deleteNode(node)
		} else {
			seen[node.value] = struct{}{}
		}
		// As the linked list gets shorter,
		// it's important to check node.next
		if node.next == nil {
			break
		}
		node = node.next
	}
}
