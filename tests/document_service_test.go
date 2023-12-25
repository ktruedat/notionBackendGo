package tests

import (
	"github.com/google/go-cmp/cmp"
	"github.com/ktruedat/notionBackendGo/model"
	"github.com/ktruedat/notionBackendGo/repository"
	"github.com/ktruedat/notionBackendGo/service"
	"testing"
)

func TestCreateDocument(t *testing.T) {
	inMemDocRepo := repository.NewInMemoryDocumentRepository()
	docService := service.NewDocumentService(inMemDocRepo)
	wantedDocument := &model.Document{ID: "1234",
		Name:      "Some Document",
		Documents: []model.ID{},
	}
	gotDocument := docService.CreateDocument(wantedDocument)

	if !cmp.Equal(wantedDocument, gotDocument) {
		t.Fatal(cmp.Diff(wantedDocument, gotDocument))
	}
}
