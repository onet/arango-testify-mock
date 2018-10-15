// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// DatabaseGraphs is an autogenerated mock type for the DatabaseGraphs type
type DatabaseGraphs struct {
	mock.Mock
}

// CreateGraph provides a mock function with given fields: ctx, name, options
func (_m *DatabaseGraphs) CreateGraph(ctx context.Context, name string, options *driver.CreateGraphOptions) (driver.Graph, error) {
	ret := _m.Called(ctx, name, options)

	var r0 driver.Graph
	if rf, ok := ret.Get(0).(func(context.Context, string, *driver.CreateGraphOptions) driver.Graph); ok {
		r0 = rf(ctx, name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Graph)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *driver.CreateGraphOptions) error); ok {
		r1 = rf(ctx, name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Graph provides a mock function with given fields: ctx, name
func (_m *DatabaseGraphs) Graph(ctx context.Context, name string) (driver.Graph, error) {
	ret := _m.Called(ctx, name)

	var r0 driver.Graph
	if rf, ok := ret.Get(0).(func(context.Context, string) driver.Graph); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Graph)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GraphExists provides a mock function with given fields: ctx, name
func (_m *DatabaseGraphs) GraphExists(ctx context.Context, name string) (bool, error) {
	ret := _m.Called(ctx, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Graphs provides a mock function with given fields: ctx
func (_m *DatabaseGraphs) Graphs(ctx context.Context) ([]driver.Graph, error) {
	ret := _m.Called(ctx)

	var r0 []driver.Graph
	if rf, ok := ret.Get(0).(func(context.Context) []driver.Graph); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]driver.Graph)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
