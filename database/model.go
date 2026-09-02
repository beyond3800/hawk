package database


type ModelQuery struct {
	model   any
	builder *Builder
}


func (m *ModelQuery) All() any{
	m.builder.Get(&m.model)
	return m.model
}