// Code generated by mockery v2.6.0. DO NOT EDIT.

package marvel

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

// GetCharacterById provides a mock function with given fields: id
func (_m *MockInterface) GetCharacterById(id string) (*http.Response, error) {
	ret := _m.Called(id)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string) *http.Response); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCharacters provides a mock function with given fields: offset, limit
func (_m *MockInterface) GetCharacters(offset int, limit int) (*http.Response, error) {
	ret := _m.Called(offset, limit)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(int, int) *http.Response); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}