package testmod

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/testmod/internal/blob"
	"ziniki.org/deployer/testmod/internal/testDB"
	"ziniki.org/deployer/testmod/internal/testS3"
	"ziniki.org/deployer/testmod/internal/testenv"
)

var testRunner driverbottom.TestRunner

func ProvideTestRunner(runner driverbottom.TestRunner) error {
	testRunner = runner
	return nil
}

func RegisterWithDriver(driver driverbottom.Driver) error {
	if testRunner != nil {
		eh := testRunner.ErrorHandlerFor("log")
		eh.WriteMsg("Installing things from testmod\n")
	}
	toolsI := driver.ObtainCoreTools().RetrieveOther("coremod")
	if toolsI == nil {
		panic("must load coremod first")
	}
	tools := toolsI.(*corebottom.Tools)

	tools.Register.ProvideDriver("testS3.TestAwsEnv", &testenv.TestAwsEnv{})

	tools.Register.Register("target", "test.assertBucketHas", testS3.NewAssertBucketHandler(tools))
	tools.Register.Register("target", "blob", blob.NewBlobCommandHandler(tools))

	tools.Register.Register("blank", "test.S3.Bucket", &testS3.BucketBlank{})
	tools.Register.Register("blank", "test.S3.Configuration", &testS3.ConfigurationCoin{})

	tools.Register.Register("blank", "test.DB.Table", &testDB.TableBlank{})

	tools.Register.Register("prop-interpreter", "test.FieldInterpreter", driverbottom.CreateInterpreter(testDB.CreateFieldInterpreter))
	return nil
}
