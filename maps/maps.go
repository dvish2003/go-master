package main


func main(){
	studentAges := map[string]int{}

	studentAges["Alice"] = 21
	studentAges["Bob"] = 22
	studentAges["Charlie"] = 20

	// Update age
	studentAges["Alice"] = 23

	// Delete an entry
	delete(studentAges, "Bob")

	// Iterate over the map
	for name, age := range studentAges {
		println(name, age)
	}
}