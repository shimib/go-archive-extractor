package archive_extractor

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"strings"
)

func TestTarUnexpectedEofArchiver(t *testing.T) {
	za := &TarArchvier{}
	funcParams := params()
	if err := za.ExtractArchive("./fixtures/test.deb", processingFunc, funcParams); err != nil {
		fmt.Print(err.Error() + "\n")
		assert.Equal(t, "archive/tar: invalid tar header", strings.Trim(err.Error(), ""))
	}
}

func TestTarArchiver(t *testing.T) {
	za := &TarArchvier{}
	funcParams := params()
	if err := za.ExtractArchive("./fixtures/test.tar.gz", processingFunc, funcParams); err != nil {
		fmt.Print(err.Error())
		t.Fatal(err)
	}
	ad := funcParams["archiveData"].(*ArchiveData)
	assert.Equal(t, ad.Name, "logRotator-1.0/log_rotator.go")
	assert.Equal(t, ad.ModTime, int64(1531307652))
	assert.Equal(t, ad.IsFolder, false)
	assert.Equal(t, ad.Size, int64(3685))
	assert.Equal(t, ad.IsSparse, false)
}
func TestTarArchiverSparse(t *testing.T) {
	za := &TarArchvier{}
	funcParams := params()
	if err := za.ExtractArchive("./fixtures/test-gnu-sparse-big.tar", processingFunc, funcParams); err != nil {
		fmt.Print(err.Error())
		t.Fatal(err)
	}
	ad := funcParams["archiveData"].(*ArchiveData)
	assert.Equal(t, ad.Name, "gnu-sparse")
	assert.Equal(t, ad.ModTime, int64(0))
	assert.Equal(t, ad.IsFolder, false)
	assert.Equal(t, ad.Size, int64(60000000000))
	assert.Equal(t, true, ad.IsSparse)
}
