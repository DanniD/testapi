package storage

// InitUsers returns initdata for users
func InitUsers() []User {
	return []User{
		{
			ID:             1,
			OrganisationID: 1,
			Name:           "Nina Greger",
			Email:          "Nina.Greger@propia.com",
			Password:       "password",
			Admin:          true,
		},
		{
			ID:             2,
			OrganisationID: 2,
			Name:           "Chegu vara",
			Email:          "Chegu.vara@cybercom.com",
			Password:       "password",
			Admin:          true,
		},
		{
			ID:             3,
			OrganisationID: 2,
			Name:           "rep repsson",
			Email:          "rep.repsson@cybercom.com",
			Password:       "",
			Admin:          false,
		},
	}
}

// InitOrganisations returns initdata for Organisations
func InitOrganisations() []Organisation {
	return []Organisation{
		{
			ID:   1,
			Name: "Propia",
			Division: []Division{
				{Name: "Konsult", Users: []User{
					{
						ID:       1,
						Name:     "Nina Greger",
						Email:    "Nina.Greger@propia.com",
						Password: "password",
						Admin:    true,
					},
				}},
			},
		},
		{
			ID:   2,
			Name: "Cybercom",
			Division: []Division{
				{Name: "Inköp", Users: []User{
					{
						ID:       2,
						Name:     "Chegu vara",
						Email:    "Chegu.vara@cybercom.com",
						Password: "password",
						Admin:    true,
					},
				},
				},
				{Name: "Ledning", Users: []User{
					{
						ID:       2,
						Name:     "Chegu vara",
						Email:    "Chegu.vara@cybercom.com",
						Password: "password",
						Admin:    true},
					{
						ID:       3,
						Name:     "rep repsson",
						Email:    "rep.repsson@cybercom.com",
						Password: "",
						Admin:    false,
					},
				},
				},
			},
		},
	}
}
