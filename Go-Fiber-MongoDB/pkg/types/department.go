package types

type CreateDepartmentRequest struct {
	Name				string 	`json:"name"`
	Description string	`json:"description"`
	ManagerID 	string	`json:"manager_id"`
}

type CreateDepartmentResponse struct {
	ID 					string `json:"id"`
	Name 				string `json:"name"`
	Description string `json:"description"`
	ManagerID 	string `json:"manager_id"`
}