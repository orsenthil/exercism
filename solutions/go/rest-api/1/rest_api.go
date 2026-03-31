package restapi

import (
	"slices"
	"strings"
)

// Define the Rest API interface. You should not modify the code in this block.

type User struct {
	Name    string
	Owes    map[string]float64
	OwedBy  map[string]float64
	Balance float64
}

type GetUsersRequest struct {
	Users []string
}

type GetUsersResponse struct {
	Users []User
}

type AddUserRequest struct {
	User string
}

type AddUserResponse struct {
	User User
}

type AddIouRequest struct {
	Lender   string
	Borrower string
	Amount   float64
}

type AddIouResponse struct {
	Users []User
}

type RestApi interface {
	GetUsers(GetUsersRequest) GetUsersResponse
	AddUser(AddUserRequest) AddUserResponse
	AddIou(AddIouRequest) AddIouResponse
}

// Your code goes below here. Implement the RestApi interface.

type Api struct {
	database []User
}

func NewApi(database []User) RestApi {
	return &Api{database: database}
}

func (a *Api) GetUsers(req GetUsersRequest) GetUsersResponse {
	if len(req.Users) == 0 {
		slices.SortFunc(a.database, func(a, b User) int {
			return strings.Compare(a.Name, b.Name)
		})
		return GetUsersResponse{Users: a.database}
	}
	users := []User{}
	for _, user := range a.database {
		if slices.Contains(req.Users, user.Name) {
			users = append(users, user)
		}
	}
	slices.SortFunc(users, func(a, b User) int {
		return strings.Compare(a.Name, b.Name)
	})

	return GetUsersResponse{Users: users}
}

func (a *Api) AddUser(req AddUserRequest) AddUserResponse {
	for _, user := range a.database {
		if user.Name == req.User {
			return AddUserResponse{User: user}
		}
	}
	newUser := User{Name: req.User, Owes: map[string]float64{}, OwedBy: map[string]float64{}, Balance: 0}

	a.database = append(a.database, newUser)

	return AddUserResponse{User: newUser}
}

func (a *Api) AddIou(req AddIouRequest) AddIouResponse {
	lender := req.Lender
	borrower := req.Borrower
	amount := req.Amount
	lenderIdx := -1
	borrowerIdx := -1

	for i := range a.database {
		if a.database[i].Name == lender {
			if a.database[i].Owes[borrower] > 0 {
				a.database[i].Owes[borrower] -= amount
				if a.database[i].Owes[borrower] > 0 {
					// nothing
				} else if a.database[i].Owes[borrower] == 0 {
					delete(a.database[i].Owes, borrower)
				} else {
						a.database[i].OwedBy[borrower] = -a.database[i].Owes[borrower]
						delete(a.database[i].Owes, borrower)
				}
			} else {
				a.database[i].OwedBy[borrower] += amount
			}
			a.database[i].Balance += amount
			lenderIdx = i
		}
		if a.database[i].Name == borrower {
			if a.database[i].OwedBy[lender] > 0 {
				a.database[i].OwedBy[lender] -= amount
				if a.database[i].OwedBy[lender] > 0 {
					// nothing
				} else if a.database[i].OwedBy[lender] == 0 {
					delete(a.database[i].OwedBy, lender)
				} else {
					a.database[i].Owes[lender] = -a.database[i].OwedBy[lender]
					delete(a.database[i].OwedBy, lender)
				}
			} else {
				a.database[i].Owes[lender] += amount
			}
			a.database[i].Balance -= amount
			borrowerIdx = i
		}
	}

	sortedUsers := []User{a.database[lenderIdx], a.database[borrowerIdx]}
	slices.SortFunc(sortedUsers, func(a, b User) int {
		return strings.Compare(a.Name, b.Name)
	})

	return AddIouResponse{Users: sortedUsers}

}
