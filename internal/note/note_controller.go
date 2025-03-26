package note

type NoteController struct {
	service *NoteService
}

func NewNoteController(service *NoteService) *NoteController {
	return &NoteController{service: service}
}
