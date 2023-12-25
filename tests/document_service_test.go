package tests

import (
	"github.com/google/go-cmp/cmp"
	"github.com/ktruedat/notionBackendGo/model"
	"github.com/ktruedat/notionBackendGo/repository"
	"github.com/ktruedat/notionBackendGo/service"
	"testing"
)

func TestCreateDocument(t *testing.T) {
	var iNodeTable map[string]*model.Document
	inMemDocRepo := repository.NewInMemoryDocumentRepository(iNodeTable)
	docService := service.NewDocumentService(inMemDocRepo)
	wantedDocument := &model.Document{ID: "1234",
		Name:      "Some Document",
		Documents: []model.ID{},
	}
	gotDocument, _ := docService.CreateDocument(wantedDocument)

	if !cmp.Equal(wantedDocument, gotDocument) {
		t.Fatal(cmp.Diff(wantedDocument, gotDocument))
	}
}

func TestGetDocumentByID(t *testing.T) {
	iNodeTable := make(map[string]*model.Document)
	inMemDocRepo := repository.NewInMemoryDocumentRepository(iNodeTable)
	docService := service.NewDocumentService(inMemDocRepo)
	wantedDocument := &model.Document{ID: "1234",
		Name:      "Some Document",
		Documents: []model.ID{},
	}
	docService.CreateDocument(wantedDocument)
	gotDocumentID, _ := docService.GetDocumentByID(string(wantedDocument.ID))

	if !cmp.Equal(wantedDocument, gotDocumentID) {
		t.Fatal(cmp.Diff(wantedDocument, gotDocumentID))
	}
}
