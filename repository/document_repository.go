package repository

import "github.com/ktruedat/notionBackendGo/model"

type DocumentRepository interface {
	CreateDocument(doc *model.Document) (*model.Document, error)
}

type InMemoryDocumentRepository struct {
}

func NewInMemoryDocumentRepository() *InMemoryDocumentRepository {
	return &InMemoryDocumentRepository{}
}

func (memRepo InMemoryDocumentRepository) CreateDocument(doc *model.Document) (*model.Document, error) {

	return doc, nil
}
