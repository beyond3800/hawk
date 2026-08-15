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

func (b *Builder) Paginate(page, perPage int, dest any) (*Pagination, error) {

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 15
	}

	countBuilder := b.Table(b.table)
	countBuilder.wheres = append([]string{}, b.wheres...)
	countBuilder.bindings = append([]any{}, b.bindings...)

	total, err := countBuilder.Count()
	if err != nil {
		return nil, err
	}

	b.page = page
	b.perPage = perPage
	b.limit = perPage
	b.offset = (page - 1) * perPage

	if err := b.Get(dest); err != nil {
		return nil, err
	}

	lastPage := total / perPage
	if total%perPage != 0 {
		lastPage++
	}

	return &Pagination{
		Data: dest,
		Meta: Meta{
			Page: page,
			PerPage: perPage,
			Total: total,
			LastPage: lastPage,
		},
	}, nil
}