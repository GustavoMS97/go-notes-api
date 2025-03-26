package note

type NoteService struct {
	repo NoteRepository
}

func NewNoteService(repo NoteRepository) *NoteService {
	return &NoteService{repo}
}

func (s *NoteService) CreateNote(title, content, userID string) (Note, error) {
	note := Note{
		Title:   title,
		Content: content,
		UserID:  userID,
	}

	return s.repo.Create(note)
}
