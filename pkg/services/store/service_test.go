package store

import (
	"bytes"
	"context"
	acmock "github.com/grafana/grafana/pkg/services/accesscontrol/mock"
	"path"
	"path/filepath"
	"testing"

	"github.com/grafana/grafana-plugin-sdk-go/experimental"
	"github.com/grafana/grafana/pkg/tsdb/testdatasource"
	"github.com/stretchr/testify/require"
)

func TestListFiles(t *testing.T) {
	publicRoot, err := filepath.Abs("../../../public")
	require.NoError(t, err)
	res := &nestedTree{
		roots: []storageRuntime{
			newDiskStorage("public", "Public static files", &StorageLocalDiskConfig{
				Path: publicRoot,
				Roots: []string{
					"/testdata/",
					"/img/icons/",
					"/img/bg/",
					"/gazetteer/",
					"/maps/",
					"/upload/",
				},
			}).setReadOnly(true).setBuiltin(true),
		},
	}

	store := newStandardStorageService(res, &nestedTree{}, NewStorageAuthService(&acmock.Mock{}, acmock.NewPermissionsServicesMock()))
	frame, err := store.List(context.Background(), nil, "res/public/testdata")
	require.NoError(t, err)

	err = experimental.CheckGoldenFrame(path.Join("testdata", "public_testdata.golden.txt"), frame, true)
	require.NoError(t, err)

	file, err := store.Read(context.Background(), nil, "res/public/testdata/js_libraries.csv")
	require.NoError(t, err)
	require.NotNil(t, file)

	frame, err = testdatasource.LoadCsvContent(bytes.NewReader(file.Contents), file.Name)
	require.NoError(t, err)
	err = experimental.CheckGoldenFrame(path.Join("testdata", "public_testdata_js_libraries.golden.txt"), frame, true)
	require.NoError(t, err)
}
