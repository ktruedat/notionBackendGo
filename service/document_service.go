package service

import "github.com/ktruedat/notionBackendGo/model"
import "github.com/ktruedat/notionBackendGo/repository"

type IDocumentService interface {
	CreateDocument(doc *model.Document) *model.Document
}

type DocumentService struct {
	Repository repository.DocumentRepository
}

func NewDocumentService(repository repository.DocumentRepository) *DocumentService {
	return &DocumentService{Repository: repository}
}

func (dSvc DocumentService) CreateDocument(doc *model.Document) *model.Document {
	document, _ := dSvc.Repository.CreateDocument(doc)
	return document
}
