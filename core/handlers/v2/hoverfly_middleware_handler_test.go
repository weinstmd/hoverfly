package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"testing"
)

type HoverflyMiddlewareStub struct {
	Middleware string
}

func (this HoverflyMiddlewareStub) GetMiddleware() string {
	return this.Middleware
}

func (this *HoverflyMiddlewareStub) SetMiddleware(middleware string) error {
	this.Middleware = middleware
	if middleware == "error" {
		return fmt.Errorf("error")
	}

	return nil
}

func TestHoverflyMiddlewareHandlerGetReturnsTheCorrectMiddleware(t *testing.T) {
	RegisterTestingT(t)

	stubHoverfly := &HoverflyMiddlewareStub{Middleware: "test-middleware"}
	unit := HoverflyMiddlewareHandler{Hoverfly: stubHoverfly}

	request, err := http.NewRequest("GET", "", nil)
	Expect(err).To(BeNil())

	response := makeRequestOnHandler(unit.Get, request)

	Expect(response.Code).To(Equal(http.StatusOK))

	middlewareView, err := unmarshalMiddlewareView(response.Body)
	Expect(err).To(BeNil())
	Expect(middlewareView.Middleware).To(Equal("test-middleware"))
}

func TestHoverflyMiddlewareHandlerPutSetsTheNewMiddlewarendReplacesTheTestMiddleware(t *testing.T) {
	RegisterTestingT(t)

	stubHoverfly := &HoverflyMiddlewareStub{Middleware: "test-middleware"}
	unit := HoverflyMiddlewareHandler{Hoverfly: stubHoverfly}

	middlewareView := &MiddlewareView{Middleware: "new-middleware"}

	bodyBytes, err := json.Marshal(middlewareView)
	Expect(err).To(BeNil())

	request, err := http.NewRequest("PUT", "", ioutil.NopCloser(bytes.NewBuffer(bodyBytes)))
	Expect(err).To(BeNil())

	response := makeRequestOnHandler(unit.Put, request)
	Expect(response.Code).To(Equal(http.StatusOK))
	Expect(stubHoverfly.Middleware).To(Equal("new-middleware"))

	middlewareViewResponse, err := unmarshalMiddlewareView(response.Body)
	Expect(err).To(BeNil())

	Expect(middlewareViewResponse.Middleware).To(Equal("new-middleware"))
}

func TestHoverflyMiddlewareHandlerPutWill422ErrorIfHoverflyErrors(t *testing.T) {
	RegisterTestingT(t)

	var stubHoverfly HoverflyMiddlewareStub
	unit := HoverflyMiddlewareHandler{Hoverfly: &stubHoverfly}

	middlewareView := &MiddlewareView{Middleware: "error"}

	bodyBytes, err := json.Marshal(middlewareView)
	Expect(err).To(BeNil())

	request, err := http.NewRequest("PUT", "", ioutil.NopCloser(bytes.NewBuffer(bodyBytes)))
	Expect(err).To(BeNil())

	response := makeRequestOnHandler(unit.Put, request)
	Expect(response.Code).To(Equal(http.StatusUnprocessableEntity))

	errorViewResponse, err := unmarshalErrorView(response.Body)
	Expect(err).To(BeNil())

	Expect(errorViewResponse.Error).To(Equal("Invalid middleware: error"))
}

func TestHoverflyMiddlewareHandlerPutWill400ErrorIfJsonIsBad(t *testing.T) {
	RegisterTestingT(t)

	var stubHoverfly HoverflyMiddlewareStub
	unit := HoverflyMiddlewareHandler{Hoverfly: &stubHoverfly}

	bodyBytes := []byte("{{}{}}")

	request, err := http.NewRequest("PUT", "/api/v2/hoverfly/mode", ioutil.NopCloser(bytes.NewBuffer(bodyBytes)))
	Expect(err).To(BeNil())

	response := makeRequestOnHandler(unit.Put, request)
	Expect(response.Code).To(Equal(http.StatusBadRequest))

	errorViewResponse, err := unmarshalErrorView(response.Body)
	Expect(err).To(BeNil())

	Expect(errorViewResponse.Error).To(Equal("Malformed JSON"))
}

func unmarshalMiddlewareView(buffer *bytes.Buffer) (MiddlewareView, error) {
	body, err := ioutil.ReadAll(buffer)
	if err != nil {
		return MiddlewareView{}, err
	}

	var middlewareView MiddlewareView

	err = json.Unmarshal(body, &middlewareView)
	if err != nil {
		return MiddlewareView{}, err
	}

	return middlewareView, nil
}
