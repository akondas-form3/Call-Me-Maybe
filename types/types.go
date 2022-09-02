package types

// Human is our type that we'll serialize to/from
type Human struct {
	ID         int
	FirstName  string
	LastName   string
	Age        int
	LikesPizza bool
}
