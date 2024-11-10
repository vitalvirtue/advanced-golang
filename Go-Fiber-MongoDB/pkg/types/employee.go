package types

type CreateEmployeeRequest struct {
    FirstName string  `json:"first_name"`
    LastName  string  `json:"last_name"`
    Position  string  `json:"position"`
    Salary    float64 `json:"salary"`
}

type CreateEmployeeResponse struct {
    ID        string  `json:"id"`
    FirstName string  `json:"first_name"`
    LastName  string  `json:"last_name"`
    Position  string  `json:"position"`
    Salary    float64 `json:"salary"`
}
