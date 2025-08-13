package impl

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"plugin"

	"ziniki.org/deployer/driver/internal/parser"
	"ziniki.org/deployer/driver/internal/registry"
	"ziniki.org/deployer/driver/internal/repo"
	"ziniki.org/deployer/driver/internal/runtime"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
	"ziniki.org/deployer/driver/pkg/utils"
)

type DriverImpl struct {
	tools        *driverbottom.CoreTools
	userErrorsTo io.StringWriter
	srcdir       string
	input        []string
}

func (d *DriverImpl) ObtainCoreTools() *driverbottom.CoreTools {
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
		envfile := filepath.Join(d, file)
		envs, err := utils.ReadEnvs(envfile)
		if err != nil {
			log.Printf("Failed to read %s: %v\n", envfile, err)
			return false
		}
		if envs != nil {
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
		log.Printf("ignoring module %s as it does not have RegisterWithDriver", mod)
		return nil
	}
	return init.(func(driverbottom.Driver) error)(d)
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
	hasErrors := d.tools.Repository.ResolveAll(d.tools)
	if hasErrors || d.tools.Reporter.HasErrors() {
		return fmt.Errorf("errors during resolving")
	}
	return nil
}

// Mainly support for the test harness, but do with them as you will (not used in $cmd$)
func (d *DriverImpl) AddSymbolListener(lsnr driverbottom.SymbolListener) {
	d.tools.Repository.AddSymbolListener(lsnr)
}

func (d *DriverImpl) Traverse(lsnr driverbottom.RepositoryTraverser) {
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
	tools := driverbottom.NewTools(reporter, reg, reg, repo, storage)
	reg.BindTools(tools)
	RegisterBasicFunctions(tools)
	ret := &DriverImpl{tools: tools, userErrorsTo: userErrorsTo}
	return ret
}
