package service

import "github.com/ktruedat/notionBackendGo/model"
import "github.com/ktruedat/notionBackendGo/repository"

type IDocumentService interface {
	CreateDocument(doc *model.Document) *model.Document
	GetDocumentByID(id string) (*model.Document, error)
}

type DocumentService struct {
	Repository repository.DocumentRepository
}

func NewDocumentService(repository repository.DocumentRepository) *DocumentService {
	return &DocumentService{Repository: repository}
}

func (docSvc DocumentService) CreateDocument(doc *model.Document) (*model.Document, error) {
	document, err := docSvc.Repository.CreateDocument(doc)
	return document, err
}

func (docSvc DocumentService) GetDocumentByID(id string) (*model.Document, error) {
	return docSvc.Repository.GetDocumentByID(id)
}
