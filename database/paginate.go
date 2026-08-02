package database


type Pagination struct {
    Data any    `json:"data"`
    Meta Meta   `json:"meta"`
}

type Meta struct {
    Page int `json:"page"`

    PerPage int `json:"per_page"`

    Total int `json:"total"`

    LastPage int `json:"last_page"`
}