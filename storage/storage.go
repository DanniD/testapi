package storage

// User data structure
type User struct {
	ID             int    `json:"id"`
	OrganisationID int    `json:"organisationId,omitempty"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Admin          bool   `json:"admin"`
}

// Organisation data structure
type Organisation struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Division []Division `json:"´division"`
}

// Division data structure
type Division struct {
	OrganisationID int    `json:"-"`
	Name           string `json:"name"`
	Users          []User `json:"users"`
}

// Storage .db is an alias for *gorm.DB
type Storage struct {
}

// New creates a new storage
func New() *Storage {
	return &Storage{}
}

// Users is dummyData of users
var Users = []User{}

// Organisations is dummyData of Organisations
var Organisations = []Organisation{}

// Init Storeage ...
func (s *Storage) Init() {
	Users = InitUsers()
	Organisations = InitOrganisations()
}

// GetUsers return all users
func (s *Storage) GetUsers() []User {
	return Users
}

// GetUser return one user
func (s *Storage) GetUser(id int) User {
	return Users[id]
}

// GetOrganisations returns all organisation ...
func (s *Storage) GetOrganisations() []Organisation {
	return Organisations
}

// GetOrganisation returns one organisation
func (s *Storage) GetOrganisation(id int) Organisation {
	return Organisations[id]
}
