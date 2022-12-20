package foldertest

import (
	"context"

	"github.com/grafana/grafana/pkg/services/folder"
)

type FakeService struct {
	ExpectedFolders []*folder.Folder
	ExpectedFolder  *folder.Folder
	ExpectedError   error
}

func NewFakeService() *FakeService {
	return &FakeService{}
}

var _ folder.Service = (*FakeService)(nil)

func (s *FakeService) GetChildren(ctx context.Context, cmd *folder.GetChildrenQuery) ([]*folder.Folder, error) {
	return s.ExpectedFolders, s.ExpectedError
}
func (s *FakeService) Create(ctx context.Context, cmd *folder.CreateFolderCommand) (*folder.Folder, error) {
	return s.ExpectedFolder, s.ExpectedError
}
func (s *FakeService) Get(ctx context.Context, cmd *folder.GetFolderQuery) (*folder.Folder, error) {
	return s.ExpectedFolder, s.ExpectedError
}
func (s *FakeService) Update(ctx context.Context, cmd *folder.UpdateFolderCommand) (*folder.Folder, error) {
	return s.ExpectedFolder, s.ExpectedError
}
func (s *FakeService) DeleteFolder(ctx context.Context, cmd *folder.DeleteFolderCommand) error {
	return s.ExpectedError
}
func (s *FakeService) MakeUserAdmin(ctx context.Context, orgID int64, userID, folderID int64, setViewAndEditPermissions bool) error {
	return s.ExpectedError
}

func (s *FakeService) Move(ctx context.Context, cmd *folder.MoveFolderCommand) (*folder.Folder, error) {
	return s.ExpectedFolder, s.ExpectedError
}

func (s *FakeService) GetParents(ctx context.Context, orgID int64, folderUID string) ([]*folder.Folder, error) {
	return s.ExpectedFolders, s.ExpectedError
}

func (s *FakeService) GetTree(ctx context.Context, orgID int64, folderUID string, depth int64) (map[string][]*folder.Folder, error) {
	ret := make(map[string][]*folder.Folder)
	ret[folderUID] = s.ExpectedFolders
	return ret, s.ExpectedError
}
