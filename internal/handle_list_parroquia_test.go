/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"vigo360.es/fuentes/internal/fuente"
)

func TestHandleListParroquiaValid(t *testing.T) {
	var (
		fr = fuente.NewMockRepository(false)
		s  = NewServer(fr)
		r  = httptest.NewRequest("GET", "/parroquias/valid", nil)
		w  = httptest.NewRecorder()
	)

	s.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200 received %d\n", w.Code)
	}
}

func TestHandleListParroquiaInvalid(t *testing.T) {
	var (
		fr = fuente.NewMockRepository(false)
		s  = NewServer(fr)
		r  = httptest.NewRequest("GET", "/parroquias/invalid", nil)
		w  = httptest.NewRecorder()
	)

	s.ServeHTTP(w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404 received %d\n", w.Code)
	}
}

func TestHandleListParroquiaDataError(t *testing.T) {
	var (
		fr = fuente.NewMockRepository(true)
		s  = NewServer(fr)
		r  = httptest.NewRequest("GET", "/parroquias/valid", nil)
		w  = httptest.NewRecorder()
	)

	s.ServeHTTP(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500 received %d\n", w.Code)
	}
}
