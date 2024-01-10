package linkedlist

func ExampleNew() {
	linkedList := New()
	linkedList.Append(1).Append(6).Append("dasoianfaoifnw").PrintAll()
}

func ExampleNewWithDeletion() {
	linkedList := New()
	linkedList.Append(1).Append(6).Append("dasoianfaoifnw").DeleteWithValue(6).PrintAll()
}

func ExampleNewWithDeletionOfHead() {
	linkedList := New()
	linkedList.Append(1).Append(6).Append("dasoianfaoifnw").DeleteWithValue(1).PrintAll()
}

func ExampleNewWithDeletionOfTail() {
	linkedList := New()
	linkedList.Append(1).Append(6).Append("dasoianfaoifnw").DeleteWithValue("dasoianfaoifnw").PrintAll()
}
