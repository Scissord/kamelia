package client

type Client struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type GetClientsInput struct {
	Limit  int
	Offset int
}

type GetClientsResponse struct {
	Items []Client
	Total int64
}

type ClientCreateInput struct {
	ID        int64 // для update/delete
	FirstName string
	LastName  string
	Email     string
}
