package character

// Character represents a player character with a name, job, and attributes.
// It includes JSON struct tags for serialization.
// Attributes is represented using the AttributesMap type.
type Character struct {
	Name       string        `json:"name"`
	Job        string        `json:"job"`
	Attributes AttributesMap `json:"attributes"`
}

// NewCharacter creates and returns a new Character instance.
// It takes the character's name, job, and attributes as parameters.
// It returns a pointer to the newly created Character.
func NewCharacter(name string, job string, attributes AttributesMap) *Character {
	return &Character{
		Name:       name,
		Job:        job,
		Attributes: attributes,
	}
}

// String returns a string representation of the Character.
func (c *Character) String() string {
	return c.Name + " the " + c.Job + " [" + c.Attributes.String() + "]"
}
