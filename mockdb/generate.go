package mockdb

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// generate mock data
// TODO: pseudorandom generation
func generateFakeUsers() []User {
	var users = []User{
		User{ID: "123456789", Name: "Beatrice", Email: "beatrice@gmail.com"},
		User{ID: "987654321", Name: "Amy", Email: "amy@amy.net"},
	}
	return users
}

func FetchAllUsers() []User {
	users := generateFakeUsers()
	return users
}

