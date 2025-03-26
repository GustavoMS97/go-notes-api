package note

import (
	"errors"
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
	log.Println("[NoteService] Creating note")
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

func (s *NoteService) UpdateNote(noteID string, userID string, title *string, content *string) (Note, error) {
	log.Println("[NoteService] Updating note:", noteID, userID)
	updates := make(map[string]interface{})

	if title != nil {
		updates["title"] = *title
	}

	if content != nil {
		updates["content"] = *content
	}

	if len(updates) == 0 {
		return Note{}, errors.New("no fields to update")
	}

	return s.repo.UpdateByID(noteID, userID, updates)
}

func (s *NoteService) DeleteNote(noteID string, userID string) error {
	log.Println("[NoteService] Deliting note:", noteID, userID)
	return s.repo.DeleteByID(noteID, userID)
}
