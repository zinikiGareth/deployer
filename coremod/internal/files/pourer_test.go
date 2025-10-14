package files_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"slices"
	"testing"

	"ziniki.org/deployer/coremod/internal/files"
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

var loc errorsink.Location

func TestCanFindADirectoryPourerForTest1Dir(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get cwd: %v", err)
	}
	cwd = path.Dir(path.Dir(cwd))
	dp, err := files.NewDirModel(&loc, []any{cwd, "test/testdata", "filesystem", "test1/dir"})
	if err != nil {
		t.Fatalf("error obtaining model: %v", err)
	}
	testPourAll(t, dp, []string{"xx.txt", "yy.txt"})
}

func TestErrorAskingAFileToBeADir(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get cwd: %v", err)
	}
	cwd = path.Dir(path.Dir(cwd))
	_, err = files.NewDirModel(&loc, []any{cwd, "test/testdata", "filesystem", "test1/dir", "xx.txt"})
	if err == nil {
		t.Fatalf("no error obtaining model")
	}
}

func TestErrorAskingNonExistentPathToBeADir(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get cwd: %v", err)
	}
	cwd = path.Dir(path.Dir(cwd))
	_, err = files.NewDirModel(&loc, []any{cwd, "test/testdata", "filesystem", "test1/dir", "zz.txt"})
	if err == nil {
		t.Fatalf("no error obtaining model")
	}
}

func TestCanFindAFilePourerForTest1Dir(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get cwd: %v", err)
	}
	cwd = path.Dir(path.Dir(cwd))
	dp, err := files.NewFileModel(&loc, []any{cwd, "test/testdata", "filesystem", "test1/dir", "xx.txt"})
	if err != nil {
		t.Fatalf("error obtaining model: %v", err)
	}
	testPourAll(t, dp, []string{"xx.txt"})
}

func TestErrorAskingADirToBeAFile(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get cwd: %v", err)
	}
	cwd = path.Dir(path.Dir(cwd))
	_, err = files.NewFileModel(&loc, []any{cwd, "test/testdata", "filesystem", "test1/dir"})
	if err == nil {
		t.Fatalf("no error obtaining model")
	}
}

func testPourAll(t *testing.T, dp *files.DirModel, exp []string) {
	dest := &DestDumper{expect: exp}
	err := dp.PourAll(dest)
	if err != nil {
		t.Fatalf("error during pouration: %v\n", err)
	}
	if len(dest.expect) > 0 {
		t.Fatalf("did not receive %v", dest.expect)
	}

}

type DestDumper struct {
	expect []string
}

// PourInto implements corebottom.FileDest.
func (d *DestDumper) PourInto(name string, contents io.Reader) error {
	if !slices.Contains(d.expect, name) {
		log.Printf("did not expect %s", name)
		return fmt.Errorf("did not expect %s", name)
	}
	d.expect = slices.DeleteFunc(d.expect, func(n string) bool {
		return n == name
	})
	return nil
}

// Relative implements corebottom.FileDest.
func (d *DestDumper) Relative(s string) (corebottom.FileDest, error) {
	panic("unimplemented")
}
