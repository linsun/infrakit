// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/docker/client (interfaces: APIClient)

package client

import (
	types "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	events "github.com/docker/docker/api/types/events"
	filters "github.com/docker/docker/api/types/filters"
	network "github.com/docker/docker/api/types/network"
	registry "github.com/docker/docker/api/types/registry"
	swarm "github.com/docker/docker/api/types/swarm"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	io "io"
	time "time"
)

// Mock of APIClient interface
type MockAPIClient struct {
	ctrl     *gomock.Controller
	recorder *_MockAPIClientRecorder
}

// Recorder for MockAPIClient (not exported)
type _MockAPIClientRecorder struct {
	mock *MockAPIClient
}

func NewMockAPIClient(ctrl *gomock.Controller) *MockAPIClient {
	mock := &MockAPIClient{ctrl: ctrl}
	mock.recorder = &_MockAPIClientRecorder{mock}
	return mock
}

func (_m *MockAPIClient) EXPECT() *_MockAPIClientRecorder {
	return _m.recorder
}

func (_m *MockAPIClient) ClientVersion() string {
	ret := _m.ctrl.Call(_m, "ClientVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ClientVersion() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ClientVersion")
}

func (_m *MockAPIClient) ContainerAttach(_param0 context.Context, _param1 string, _param2 types.ContainerAttachOptions) (types.HijackedResponse, error) {
	ret := _m.ctrl.Call(_m, "ContainerAttach", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.HijackedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerAttach(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerAttach", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerCommit(_param0 context.Context, _param1 string, _param2 types.ContainerCommitOptions) (types.ContainerCommitResponse, error) {
	ret := _m.ctrl.Call(_m, "ContainerCommit", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ContainerCommitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerCommit(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerCommit", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerCreate(_param0 context.Context, _param1 *container.Config, _param2 *container.HostConfig, _param3 *network.NetworkingConfig, _param4 string) (types.ContainerCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "ContainerCreate", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(types.ContainerCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerCreate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerCreate", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockAPIClient) ContainerDiff(_param0 context.Context, _param1 string) ([]types.ContainerChange, error) {
	ret := _m.ctrl.Call(_m, "ContainerDiff", _param0, _param1)
	ret0, _ := ret[0].([]types.ContainerChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerDiff(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerDiff", arg0, arg1)
}

func (_m *MockAPIClient) ContainerExecAttach(_param0 context.Context, _param1 string, _param2 types.ExecConfig) (types.HijackedResponse, error) {
	ret := _m.ctrl.Call(_m, "ContainerExecAttach", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.HijackedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerExecAttach(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerExecAttach", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerExecCreate(_param0 context.Context, _param1 string, _param2 types.ExecConfig) (types.ContainerExecCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "ContainerExecCreate", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ContainerExecCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerExecCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerExecCreate", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerExecInspect(_param0 context.Context, _param1 string) (types.ContainerExecInspect, error) {
	ret := _m.ctrl.Call(_m, "ContainerExecInspect", _param0, _param1)
	ret0, _ := ret[0].(types.ContainerExecInspect)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerExecInspect(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerExecInspect", arg0, arg1)
}

func (_m *MockAPIClient) ContainerExecResize(_param0 context.Context, _param1 string, _param2 types.ResizeOptions) error {
	ret := _m.ctrl.Call(_m, "ContainerExecResize", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerExecResize(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerExecResize", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerExecStart(_param0 context.Context, _param1 string, _param2 types.ExecStartCheck) error {
	ret := _m.ctrl.Call(_m, "ContainerExecStart", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerExecStart(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerExecStart", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerExport(_param0 context.Context, _param1 string) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "ContainerExport", _param0, _param1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerExport(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerExport", arg0, arg1)
}

func (_m *MockAPIClient) ContainerInspect(_param0 context.Context, _param1 string) (types.ContainerJSON, error) {
	ret := _m.ctrl.Call(_m, "ContainerInspect", _param0, _param1)
	ret0, _ := ret[0].(types.ContainerJSON)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerInspect(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerInspect", arg0, arg1)
}

func (_m *MockAPIClient) ContainerInspectWithRaw(_param0 context.Context, _param1 string, _param2 bool) (types.ContainerJSON, []byte, error) {
	ret := _m.ctrl.Call(_m, "ContainerInspectWithRaw", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ContainerJSON)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAPIClientRecorder) ContainerInspectWithRaw(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerInspectWithRaw", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerKill(_param0 context.Context, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "ContainerKill", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerKill(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerKill", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerList(_param0 context.Context, _param1 types.ContainerListOptions) ([]types.Container, error) {
	ret := _m.ctrl.Call(_m, "ContainerList", _param0, _param1)
	ret0, _ := ret[0].([]types.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerList", arg0, arg1)
}

func (_m *MockAPIClient) ContainerLogs(_param0 context.Context, _param1 string, _param2 types.ContainerLogsOptions) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "ContainerLogs", _param0, _param1, _param2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerLogs(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerLogs", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerPause(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "ContainerPause", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerPause(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerPause", arg0, arg1)
}

func (_m *MockAPIClient) ContainerRemove(_param0 context.Context, _param1 string, _param2 types.ContainerRemoveOptions) error {
	ret := _m.ctrl.Call(_m, "ContainerRemove", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerRemove(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerRemove", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerRename(_param0 context.Context, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "ContainerRename", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerRename(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerRename", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerResize(_param0 context.Context, _param1 string, _param2 types.ResizeOptions) error {
	ret := _m.ctrl.Call(_m, "ContainerResize", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerResize(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerResize", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerRestart(_param0 context.Context, _param1 string, _param2 *time.Duration) error {
	ret := _m.ctrl.Call(_m, "ContainerRestart", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerRestart(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerRestart", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerStart(_param0 context.Context, _param1 string, _param2 types.ContainerStartOptions) error {
	ret := _m.ctrl.Call(_m, "ContainerStart", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerStart(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerStart", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerStatPath(_param0 context.Context, _param1 string, _param2 string) (types.ContainerPathStat, error) {
	ret := _m.ctrl.Call(_m, "ContainerStatPath", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ContainerPathStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerStatPath(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerStatPath", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerStats(_param0 context.Context, _param1 string, _param2 bool) (types.ContainerStats, error) {
	ret := _m.ctrl.Call(_m, "ContainerStats", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ContainerStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerStats(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerStats", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerStop(_param0 context.Context, _param1 string, _param2 *time.Duration) error {
	ret := _m.ctrl.Call(_m, "ContainerStop", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerStop(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerStop", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerTop(_param0 context.Context, _param1 string, _param2 []string) (types.ContainerProcessList, error) {
	ret := _m.ctrl.Call(_m, "ContainerTop", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ContainerProcessList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerTop(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerTop", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerUnpause(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "ContainerUnpause", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ContainerUnpause(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerUnpause", arg0, arg1)
}

func (_m *MockAPIClient) ContainerUpdate(_param0 context.Context, _param1 string, _param2 container.UpdateConfig) (types.ContainerUpdateResponse, error) {
	ret := _m.ctrl.Call(_m, "ContainerUpdate", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ContainerUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerUpdate", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ContainerWait(_param0 context.Context, _param1 string) (int, error) {
	ret := _m.ctrl.Call(_m, "ContainerWait", _param0, _param1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ContainerWait(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ContainerWait", arg0, arg1)
}

func (_m *MockAPIClient) CopyFromContainer(_param0 context.Context, _param1 string, _param2 string) (io.ReadCloser, types.ContainerPathStat, error) {
	ret := _m.ctrl.Call(_m, "CopyFromContainer", _param0, _param1, _param2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(types.ContainerPathStat)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAPIClientRecorder) CopyFromContainer(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CopyFromContainer", arg0, arg1, arg2)
}

func (_m *MockAPIClient) CopyToContainer(_param0 context.Context, _param1 string, _param2 string, _param3 io.Reader, _param4 types.CopyToContainerOptions) error {
	ret := _m.ctrl.Call(_m, "CopyToContainer", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) CopyToContainer(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CopyToContainer", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockAPIClient) Events(_param0 context.Context, _param1 types.EventsOptions) (<-chan events.Message, <-chan error) {
	ret := _m.ctrl.Call(_m, "Events", _param0, _param1)
	ret0, _ := ret[0].(<-chan events.Message)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) Events(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Events", arg0, arg1)
}

func (_m *MockAPIClient) ImageBuild(_param0 context.Context, _param1 io.Reader, _param2 types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	ret := _m.ctrl.Call(_m, "ImageBuild", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ImageBuildResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImageBuild(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageBuild", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ImageCreate(_param0 context.Context, _param1 string, _param2 types.ImageCreateOptions) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "ImageCreate", _param0, _param1, _param2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImageCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageCreate", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ImageHistory(_param0 context.Context, _param1 string) ([]types.ImageHistory, error) {
	ret := _m.ctrl.Call(_m, "ImageHistory", _param0, _param1)
	ret0, _ := ret[0].([]types.ImageHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImageHistory(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageHistory", arg0, arg1)
}

func (_m *MockAPIClient) ImageImport(_param0 context.Context, _param1 types.ImageImportSource, _param2 string, _param3 types.ImageImportOptions) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "ImageImport", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImageImport(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageImport", arg0, arg1, arg2, arg3)
}

func (_m *MockAPIClient) ImageInspectWithRaw(_param0 context.Context, _param1 string) (types.ImageInspect, []byte, error) {
	ret := _m.ctrl.Call(_m, "ImageInspectWithRaw", _param0, _param1)
	ret0, _ := ret[0].(types.ImageInspect)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAPIClientRecorder) ImageInspectWithRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageInspectWithRaw", arg0, arg1)
}

func (_m *MockAPIClient) ImageList(_param0 context.Context, _param1 types.ImageListOptions) ([]types.Image, error) {
	ret := _m.ctrl.Call(_m, "ImageList", _param0, _param1)
	ret0, _ := ret[0].([]types.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImageList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageList", arg0, arg1)
}

func (_m *MockAPIClient) ImageLoad(_param0 context.Context, _param1 io.Reader, _param2 bool) (types.ImageLoadResponse, error) {
	ret := _m.ctrl.Call(_m, "ImageLoad", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ImageLoadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImageLoad(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageLoad", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ImagePull(_param0 context.Context, _param1 string, _param2 types.ImagePullOptions) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "ImagePull", _param0, _param1, _param2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImagePull(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImagePull", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ImagePush(_param0 context.Context, _param1 string, _param2 types.ImagePushOptions) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "ImagePush", _param0, _param1, _param2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImagePush(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImagePush", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ImageRemove(_param0 context.Context, _param1 string, _param2 types.ImageRemoveOptions) ([]types.ImageDelete, error) {
	ret := _m.ctrl.Call(_m, "ImageRemove", _param0, _param1, _param2)
	ret0, _ := ret[0].([]types.ImageDelete)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImageRemove(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageRemove", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ImageSave(_param0 context.Context, _param1 []string) (io.ReadCloser, error) {
	ret := _m.ctrl.Call(_m, "ImageSave", _param0, _param1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImageSave(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageSave", arg0, arg1)
}

func (_m *MockAPIClient) ImageSearch(_param0 context.Context, _param1 string, _param2 types.ImageSearchOptions) ([]registry.SearchResult, error) {
	ret := _m.ctrl.Call(_m, "ImageSearch", _param0, _param1, _param2)
	ret0, _ := ret[0].([]registry.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ImageSearch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageSearch", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ImageTag(_param0 context.Context, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "ImageTag", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ImageTag(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImageTag", arg0, arg1, arg2)
}

func (_m *MockAPIClient) Info(_param0 context.Context) (types.Info, error) {
	ret := _m.ctrl.Call(_m, "Info", _param0)
	ret0, _ := ret[0].(types.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) Info(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Info", arg0)
}

func (_m *MockAPIClient) NetworkConnect(_param0 context.Context, _param1 string, _param2 string, _param3 *network.EndpointSettings) error {
	ret := _m.ctrl.Call(_m, "NetworkConnect", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) NetworkConnect(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkConnect", arg0, arg1, arg2, arg3)
}

func (_m *MockAPIClient) NetworkCreate(_param0 context.Context, _param1 string, _param2 types.NetworkCreate) (types.NetworkCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "NetworkCreate", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.NetworkCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) NetworkCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkCreate", arg0, arg1, arg2)
}

func (_m *MockAPIClient) NetworkDisconnect(_param0 context.Context, _param1 string, _param2 string, _param3 bool) error {
	ret := _m.ctrl.Call(_m, "NetworkDisconnect", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) NetworkDisconnect(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkDisconnect", arg0, arg1, arg2, arg3)
}

func (_m *MockAPIClient) NetworkInspect(_param0 context.Context, _param1 string) (types.NetworkResource, error) {
	ret := _m.ctrl.Call(_m, "NetworkInspect", _param0, _param1)
	ret0, _ := ret[0].(types.NetworkResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) NetworkInspect(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkInspect", arg0, arg1)
}

func (_m *MockAPIClient) NetworkInspectWithRaw(_param0 context.Context, _param1 string) (types.NetworkResource, []byte, error) {
	ret := _m.ctrl.Call(_m, "NetworkInspectWithRaw", _param0, _param1)
	ret0, _ := ret[0].(types.NetworkResource)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAPIClientRecorder) NetworkInspectWithRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkInspectWithRaw", arg0, arg1)
}

func (_m *MockAPIClient) NetworkList(_param0 context.Context, _param1 types.NetworkListOptions) ([]types.NetworkResource, error) {
	ret := _m.ctrl.Call(_m, "NetworkList", _param0, _param1)
	ret0, _ := ret[0].([]types.NetworkResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) NetworkList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkList", arg0, arg1)
}

func (_m *MockAPIClient) NetworkRemove(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "NetworkRemove", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) NetworkRemove(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkRemove", arg0, arg1)
}

func (_m *MockAPIClient) NodeInspectWithRaw(_param0 context.Context, _param1 string) (swarm.Node, []byte, error) {
	ret := _m.ctrl.Call(_m, "NodeInspectWithRaw", _param0, _param1)
	ret0, _ := ret[0].(swarm.Node)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAPIClientRecorder) NodeInspectWithRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeInspectWithRaw", arg0, arg1)
}

func (_m *MockAPIClient) NodeList(_param0 context.Context, _param1 types.NodeListOptions) ([]swarm.Node, error) {
	ret := _m.ctrl.Call(_m, "NodeList", _param0, _param1)
	ret0, _ := ret[0].([]swarm.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) NodeList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeList", arg0, arg1)
}

func (_m *MockAPIClient) NodeRemove(_param0 context.Context, _param1 string, _param2 types.NodeRemoveOptions) error {
	ret := _m.ctrl.Call(_m, "NodeRemove", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) NodeRemove(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeRemove", arg0, arg1, arg2)
}

func (_m *MockAPIClient) NodeUpdate(_param0 context.Context, _param1 string, _param2 swarm.Version, _param3 swarm.NodeSpec) error {
	ret := _m.ctrl.Call(_m, "NodeUpdate", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) NodeUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NodeUpdate", arg0, arg1, arg2, arg3)
}

func (_m *MockAPIClient) RegistryLogin(_param0 context.Context, _param1 types.AuthConfig) (types.AuthResponse, error) {
	ret := _m.ctrl.Call(_m, "RegistryLogin", _param0, _param1)
	ret0, _ := ret[0].(types.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) RegistryLogin(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegistryLogin", arg0, arg1)
}

func (_m *MockAPIClient) ServerVersion(_param0 context.Context) (types.Version, error) {
	ret := _m.ctrl.Call(_m, "ServerVersion", _param0)
	ret0, _ := ret[0].(types.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ServerVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServerVersion", arg0)
}

func (_m *MockAPIClient) ServiceCreate(_param0 context.Context, _param1 swarm.ServiceSpec, _param2 types.ServiceCreateOptions) (types.ServiceCreateResponse, error) {
	ret := _m.ctrl.Call(_m, "ServiceCreate", _param0, _param1, _param2)
	ret0, _ := ret[0].(types.ServiceCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ServiceCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceCreate", arg0, arg1, arg2)
}

func (_m *MockAPIClient) ServiceInspectWithRaw(_param0 context.Context, _param1 string) (swarm.Service, []byte, error) {
	ret := _m.ctrl.Call(_m, "ServiceInspectWithRaw", _param0, _param1)
	ret0, _ := ret[0].(swarm.Service)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAPIClientRecorder) ServiceInspectWithRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceInspectWithRaw", arg0, arg1)
}

func (_m *MockAPIClient) ServiceList(_param0 context.Context, _param1 types.ServiceListOptions) ([]swarm.Service, error) {
	ret := _m.ctrl.Call(_m, "ServiceList", _param0, _param1)
	ret0, _ := ret[0].([]swarm.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) ServiceList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceList", arg0, arg1)
}

func (_m *MockAPIClient) ServiceRemove(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "ServiceRemove", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ServiceRemove(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceRemove", arg0, arg1)
}

func (_m *MockAPIClient) ServiceUpdate(_param0 context.Context, _param1 string, _param2 swarm.Version, _param3 swarm.ServiceSpec, _param4 types.ServiceUpdateOptions) error {
	ret := _m.ctrl.Call(_m, "ServiceUpdate", _param0, _param1, _param2, _param3, _param4)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) ServiceUpdate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ServiceUpdate", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockAPIClient) SwarmInit(_param0 context.Context, _param1 swarm.InitRequest) (string, error) {
	ret := _m.ctrl.Call(_m, "SwarmInit", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) SwarmInit(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SwarmInit", arg0, arg1)
}

func (_m *MockAPIClient) SwarmInspect(_param0 context.Context) (swarm.Swarm, error) {
	ret := _m.ctrl.Call(_m, "SwarmInspect", _param0)
	ret0, _ := ret[0].(swarm.Swarm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) SwarmInspect(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SwarmInspect", arg0)
}

func (_m *MockAPIClient) SwarmJoin(_param0 context.Context, _param1 swarm.JoinRequest) error {
	ret := _m.ctrl.Call(_m, "SwarmJoin", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) SwarmJoin(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SwarmJoin", arg0, arg1)
}

func (_m *MockAPIClient) SwarmLeave(_param0 context.Context, _param1 bool) error {
	ret := _m.ctrl.Call(_m, "SwarmLeave", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) SwarmLeave(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SwarmLeave", arg0, arg1)
}

func (_m *MockAPIClient) SwarmUpdate(_param0 context.Context, _param1 swarm.Version, _param2 swarm.Spec, _param3 swarm.UpdateFlags) error {
	ret := _m.ctrl.Call(_m, "SwarmUpdate", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) SwarmUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SwarmUpdate", arg0, arg1, arg2, arg3)
}

func (_m *MockAPIClient) TaskInspectWithRaw(_param0 context.Context, _param1 string) (swarm.Task, []byte, error) {
	ret := _m.ctrl.Call(_m, "TaskInspectWithRaw", _param0, _param1)
	ret0, _ := ret[0].(swarm.Task)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAPIClientRecorder) TaskInspectWithRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskInspectWithRaw", arg0, arg1)
}

func (_m *MockAPIClient) TaskList(_param0 context.Context, _param1 types.TaskListOptions) ([]swarm.Task, error) {
	ret := _m.ctrl.Call(_m, "TaskList", _param0, _param1)
	ret0, _ := ret[0].([]swarm.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) TaskList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskList", arg0, arg1)
}

func (_m *MockAPIClient) UpdateClientVersion(_param0 string) {
	_m.ctrl.Call(_m, "UpdateClientVersion", _param0)
}

func (_mr *_MockAPIClientRecorder) UpdateClientVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateClientVersion", arg0)
}

func (_m *MockAPIClient) VolumeCreate(_param0 context.Context, _param1 types.VolumeCreateRequest) (types.Volume, error) {
	ret := _m.ctrl.Call(_m, "VolumeCreate", _param0, _param1)
	ret0, _ := ret[0].(types.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) VolumeCreate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VolumeCreate", arg0, arg1)
}

func (_m *MockAPIClient) VolumeInspect(_param0 context.Context, _param1 string) (types.Volume, error) {
	ret := _m.ctrl.Call(_m, "VolumeInspect", _param0, _param1)
	ret0, _ := ret[0].(types.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) VolumeInspect(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VolumeInspect", arg0, arg1)
}

func (_m *MockAPIClient) VolumeInspectWithRaw(_param0 context.Context, _param1 string) (types.Volume, []byte, error) {
	ret := _m.ctrl.Call(_m, "VolumeInspectWithRaw", _param0, _param1)
	ret0, _ := ret[0].(types.Volume)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockAPIClientRecorder) VolumeInspectWithRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VolumeInspectWithRaw", arg0, arg1)
}

func (_m *MockAPIClient) VolumeList(_param0 context.Context, _param1 filters.Args) (types.VolumesListResponse, error) {
	ret := _m.ctrl.Call(_m, "VolumeList", _param0, _param1)
	ret0, _ := ret[0].(types.VolumesListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAPIClientRecorder) VolumeList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VolumeList", arg0, arg1)
}

func (_m *MockAPIClient) VolumeRemove(_param0 context.Context, _param1 string, _param2 bool) error {
	ret := _m.ctrl.Call(_m, "VolumeRemove", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAPIClientRecorder) VolumeRemove(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "VolumeRemove", arg0, arg1, arg2)
}
