package types

// Human is our type that we'll serialize to/from
type Human struct {
	ID         int    // The ID of the human
	FirstName  string // The name
	LastName   string // The last name
	Age        int    // The age
	LikesPizza bool   // The only question you need to ask a human, ever
}
