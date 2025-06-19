package impl

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"plugin"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/internal/registry"
	"ziniki.org/deployer/deployer/internal/repo"
	"ziniki.org/deployer/deployer/internal/runtime"
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/errorsink"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/utils"
)

type DriverImpl struct {
	tools        *pluggable.CoreTools
	userErrorsTo io.StringWriter
	srcdir       string
	input        []string
}

func (d *DriverImpl) ObtainCoreTools() *pluggable.CoreTools {
	return d.tools
}

func (d *DriverImpl) ReadScriptsFrom(indir string) error {
	d.srcdir = indir
	input, err := utils.FindFiles(indir, ".dply")
	if err != nil {
		fmt.Printf("could not read dir %s: %v\n", indir, err)
		return err
	}
	d.input = append(d.input, input...)
	return nil
}

func (d *DriverImpl) FindAndReadEnvs(dirs []string, file string) bool {
	found := false
	for _, d := range dirs {
		envs, err := utils.ReadEnvs(filepath.Join(d, file))
		if err == nil {
			found = true
			utils.SetEnvs(envs)
		}
	}
	return found
}

func (d *DriverImpl) UseModule(mod string) error {
	p, err := plugin.Open(mod)
	if err != nil {
		return err
	}
	init, err := p.Lookup("RegisterWithDriver")
	if err != nil {
		log.Printf("ignoring module " + mod + " as it does not have RegisterWithDriver")
		return nil
	}
	return init.(func(deployer.Driver) error)(d)
}

func (d *DriverImpl) DoStuff() error {
	for _, f := range d.input {
		from := filepath.Join(d.srcdir, f)
		d.tools.Repository.ReadingFile(f)
		parser.Parse(d.tools, f, from)
	}
	if d.tools.Reporter.HasErrors() {
		return fmt.Errorf("errors during parsing")
	}
	d.tools.Repository.ResolveAll(d.tools)
	if d.tools.Reporter.HasErrors() {
		return fmt.Errorf("errors during resolving")
	}
	return nil
}

// Mainly support for the test harness, but do with them as you will (not used in $cmd$)
func (d *DriverImpl) AddSymbolListener(lsnr pluggable.SymbolListener) {
	d.tools.Repository.AddSymbolListener(lsnr)
}

func (d *DriverImpl) Traverse(lsnr pluggable.RepositoryTraverser) {
	d.tools.Repository.Traverse(lsnr)
}

func (d *DriverImpl) UserError(msg string) {
	d.userErrorsTo.WriteString(msg)
}

func NewDriver(sink errorsink.ErrorSink, userErrorsTo io.StringWriter) *DriverImpl {
	reg := registry.NewRegistry()
	reporter := errorsink.NewErrorReporter(sink)
	repo := repo.NewRepository()
	storage := runtime.NewRuntimeStorage(reg, repo, sink)
	tools := pluggable.NewTools(reporter, reg, reg, repo, storage)
	reg.BindTools(tools)
	return &DriverImpl{tools: tools, userErrorsTo: userErrorsTo}
}
