package jhttp

import (
	"net/http"
	"strings"
)

type Mux struct {
	*http.ServeMux
}

func NewServeMux() *Mux {
	return &Mux{http.NewServeMux()}
}

func (m *Mux) HandleConnect(pattern string, handler http.Handler) {
	m.handle(http.MethodConnect, pattern, handler)
}

func (m *Mux) HandleDelete(pattern string, handler http.Handler) {
	m.handle(http.MethodDelete, pattern, handler)
}

func (m *Mux) HandleGet(pattern string, handler http.Handler) {
	m.handle(http.MethodGet, pattern, handler)
}

func (m *Mux) HandleHead(pattern string, handler http.Handler) {
	m.handle(http.MethodHead, pattern, handler)
}

func (m *Mux) HandleOptions(pattern string, handler http.Handler) {
	m.handle(http.MethodOptions, pattern, handler)
}

func (m *Mux) HandlePost(pattern string, handler http.Handler) {
	m.handle(http.MethodPost, pattern, handler)
}

func (m *Mux) HandlePut(pattern string, handler http.Handler) {
	m.handle(http.MethodPut, pattern, handler)
}

func (m *Mux) HandlePatch(pattern string, handler http.Handler) {
	m.handle(http.MethodPatch, pattern, handler)
}

func (m *Mux) HandleTrace(pattern string, handler http.Handler) {
	m.handle(http.MethodTrace, pattern, handler)
}

func (m *Mux) handle(method string, pattern string, handler http.Handler) {
	m.Handle(strings.Join([]string{method, pattern}, " "), handler)
}

// HandlerFunc

func (m *Mux) HandleFuncConnect(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodConnect, pattern, handler)
}

func (m *Mux) HandleFuncDelete(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodDelete, pattern, handler)
}

func (m *Mux) HandleFuncGet(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodGet, pattern, handler)
}

func (m *Mux) HandleFuncHead(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodHead, pattern, handler)
}

func (m *Mux) HandleFuncOptions(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodOptions, pattern, handler)
}

func (m *Mux) HandleFuncPost(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodPost, pattern, handler)
}

func (m *Mux) HandleFuncPut(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodPut, pattern, handler)
}

func (m *Mux) HandleFuncPatch(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodPatch, pattern, handler)
}

func (m *Mux) HandleFuncTrace(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodTrace, pattern, handler)
}

func (m *Mux) handleFunc(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.HandleFunc(strings.Join([]string{method, pattern}, " "), handler)
}
