package repository

import (
	"fmt"
	"github.com/ktruedat/notionBackendGo/model"
)

type DocumentRepository interface {
	CreateDocument(doc *model.Document) (*model.Document, error)
	GetDocumentByID(id string) (*model.Document, error)
}

type InMemoryDocumentRepository struct {
	iNodeTable map[string]*model.Document
}

func NewInMemoryDocumentRepository(iNodeTable map[string]*model.Document) *InMemoryDocumentRepository {
	return &InMemoryDocumentRepository{iNodeTable: iNodeTable}
}

func (memRepo InMemoryDocumentRepository) CreateDocument(doc *model.Document) (*model.Document, error) {
	memRepo.iNodeTable[string(doc.ID)] = doc
	return doc, nil
}

func (memRepo InMemoryDocumentRepository) GetDocumentByID(id string) (*model.Document, error) {
	if document, exists := memRepo.iNodeTable[id]; exists {
		return document, nil
	}
	return nil, fmt.Errorf("document with id %v does not exists", id)
}
