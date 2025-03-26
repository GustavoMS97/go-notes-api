package note

type NoteRepository interface {
	Create(note Note) (Note, error)
}
