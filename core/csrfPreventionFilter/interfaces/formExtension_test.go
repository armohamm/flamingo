package interfaces

import (
	"testing"

	"github.com/stretchr/testify/suite"

	applicationMocks "flamingo.me/flamingo/v3/core/csrfPreventionFilter/application/mocks"
	"flamingo.me/flamingo/v3/core/form2/domain"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	CsrfFormExtensionTestSuite struct {
		suite.Suite

		formExtension *CrsfTokenFormExtension
		service       *applicationMocks.Service

		webRequest *web.Request
	}
)

func TestCsrfFormExtensionTestSuite(t *testing.T) {
	suite.Run(t, &CsrfFormExtensionTestSuite{})
}

func (t *CsrfFormExtensionTestSuite) SetupSuite() {
	t.webRequest = web.RequestFromRequest(nil, nil)
}

func (t *CsrfFormExtensionTestSuite) SetupTest() {
	t.service = &applicationMocks.Service{}

	t.formExtension = &CrsfTokenFormExtension{}
	t.formExtension.Inject(t.service)
}

func (t *CsrfFormExtensionTestSuite) TearDown() {
	t.service.AssertExpectations(t.T())
	t.service = nil
}

func (t *CsrfFormExtensionTestSuite) TestValidate_WrongToken() {
	t.service.On("IsValid", t.webRequest).Return(false).Once()

	validationInfo, err := t.formExtension.Validate(nil, t.webRequest, nil, nil)

	t.NoError(err)
	t.True(validationInfo.HasGeneralErrors())
	t.Equal([]domain.Error{
		{
			MessageKey:   "formError.crsfToken.invalid",
			DefaultLabel: "Invalid crsf token.",
		},
	}, validationInfo.GetGeneralErrors())
}

func (t *CsrfFormExtensionTestSuite) TestFilter_Success() {
	t.service.On("IsValid", t.webRequest).Return(true).Once()

	validationInfo, err := t.formExtension.Validate(nil, t.webRequest, nil, nil)

	t.NoError(err)
	t.False(validationInfo.HasGeneralErrors())
}
