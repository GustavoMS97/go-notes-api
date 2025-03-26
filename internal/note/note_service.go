package note

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteService struct {
	repo NoteRepository
}

func NewNoteService(repo NoteRepository) *NoteService {
	return &NoteService{repo}
}

func (s *NoteService) CreateNote(title, content, userID string) (Note, error) {
	log.Println("[NoteService] Creating note:", title, content, userID)
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("[NoteService] Error creating note", err)
		return Note{}, err
	}

	note := Note{
		Title:   title,
		Content: content,
		UserID:  oid,
	}

	createdNote, err := s.repo.Create(note)
	if err != nil {
		log.Println("[NoteService] Failed to create note:", err)
		return Note{}, err
	}

	log.Println("[NoteService] Note created successfully:", createdNote.ObjectID)
	return createdNote, nil
}

func (s *NoteService) GetNotesByUser(userID string, search string) ([]Note, error) {
	log.Println("[NoteService] Fetching notes:", userID, search)
	return s.repo.FindAllByUserID(userID, search)
}
