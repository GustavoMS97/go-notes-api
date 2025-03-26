package note

type NoteRepository interface {
	Create(note Note) (Note, error)
	FindAllByUserID(userID string, search string) ([]Note, error)
	UpdateByID(noteID string, userID string, updates map[string]interface{}) (Note, error)
}
