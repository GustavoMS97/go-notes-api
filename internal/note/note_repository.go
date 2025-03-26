package note

type NoteRepository interface {
	Create(note Note) (Note, error)
	FindAllByUserID(userID string, search string) ([]Note, error)
}
