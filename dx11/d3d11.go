package dx11

import (
	"reflect"
	"syscall"
	"unsafe"

	"github.com/gonutz/w32"
	"github.com/lxn/win"
	"github.com/whiteboxsolutions/go-ole"
	"golang.org/x/sys/windows"
)

var (
	d3d11DLL = windows.NewLazySystemDLL("d3d11.dll")
)

const D3D11_SDK_VERSION = 7

type D3D11_CREATE_DEVICE_FLAG uint32

const (
	D3D11_CREATE_DEVICE_SINGLETHREADED                                D3D11_CREATE_DEVICE_FLAG = 0x1
	D3D11_CREATE_DEVICE_DEBUG                                         D3D11_CREATE_DEVICE_FLAG = 0x2
	D3D11_CREATE_DEVICE_SWITCH_TO_REF                                 D3D11_CREATE_DEVICE_FLAG = 0x4
	D3D11_CREATE_DEVICE_PREVENT_INTERNAL_THREADING_OPTIMIZATIONS      D3D11_CREATE_DEVICE_FLAG = 0x8
	D3D11_CREATE_DEVICE_BGRA_SUPPORT                                  D3D11_CREATE_DEVICE_FLAG = 0x20
	D3D11_CREATE_DEVICE_DEBUGGABLE                                    D3D11_CREATE_DEVICE_FLAG = 0x40
	D3D11_CREATE_DEVICE_PREVENT_ALTERING_LAYER_SETTINGS_FROM_REGISTRY D3D11_CREATE_DEVICE_FLAG = 0x80
	D3D11_CREATE_DEVICE_DISABLE_GPU_TIMEOUT                           D3D11_CREATE_DEVICE_FLAG = 0x100
	D3D11_CREATE_DEVICE_VIDEO_SUPPORT                                 D3D11_CREATE_DEVICE_FLAG = 0x800
)

var ID3D11DeviceID = ole.NewGUID("{db6f6ddb-ac77-4e88-8253-819df9bbf140}")

type ID3D11Device struct {
	ole.IUnknown
}

type ID3D11DeviceVtbl struct {
	ole.IUnknownVtbl
	CreateBuffer                         uintptr
	CreateTexture1D                      uintptr
	CreateTexture2D                      uintptr
	CreateTexture3D                      uintptr
	CreateShaderResourceView             uintptr
	CreateUnorderedAccessView            uintptr
	CreateRenderTargetView               uintptr
	CreateDepthStencilView               uintptr
	CreateInputLayout                    uintptr
	CreateVertexShader                   uintptr
	CreateGeometryShader                 uintptr
	CreateGeometryShaderWithStreamOutput uintptr
	CreatePixelShader                    uintptr
	CreateHullShader                     uintptr
	CreateDomainShader                   uintptr
	CreateComputeShader                  uintptr
	CreateClassLinkage                   uintptr
	CreateBlendState                     uintptr
	CreateDepthStencilState              uintptr
	CreateRasterizerState                uintptr
	CreateSamplerState                   uintptr
	CreateQuery                          uintptr
	CreatePredicate                      uintptr
	CreateCounter                        uintptr
	CreateDeferredContext                uintptr
	OpenSharedResource                   uintptr
	CheckFormatSupport                   uintptr
	CheckMultisampleQualityLevels        uintptr
	CheckCounterInfo                     uintptr
	CheckCounter                         uintptr
	CheckFeatureSupport                  uintptr
	GetPrivateData                       uintptr
	SetPrivateData                       uintptr
	SetPrivateDataInterface              uintptr
	GetFeatureLevel                      uintptr
	GetCreationFlags                     uintptr
	GetDeviceRemovedReason               uintptr
	GetImmediateContext                  uintptr
	SetExceptionMode                     uintptr
	GetExceptionMode                     uintptr
}

type _D3D11_BOX struct {
	Left, Top, Front, Right, Bottom, Back uint32
}

type _DXGI_SAMPLE_DESC struct {
	Count   uint32
	Quality uint32
}

type D3D11_TEXTURE2D_DESC struct {
	Width          uint32
	Height         uint32
	MipLevels      uint32
	ArraySize      uint32
	Format         uint32
	SampleDesc     _DXGI_SAMPLE_DESC
	Usage          uint32
	BindFlags      uint32
	CPUAccessFlags uint32
	MiscFlags      uint32
}

type iD3D11DebugVtbl struct {
	iUnknownVtbl

	SetFeatureMask             uintptr
	GetFeatureMask             uintptr
	SetPresentPerRenderOpDelay uintptr
	GetPresentPerRenderOpDelay uintptr
	SetSwapChain               uintptr
	GetSwapChain               uintptr
	ValidateContext            uintptr
	ReportLiveDeviceObjects    uintptr
	ValidateContextForDispatch uintptr
}

type iUnknownVtbl struct {
	// every COM object starts with these three
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	// _QueryInterface2 uintptr
}

type iD3D11DeviceChildVtbl struct {
	iUnknownVtbl

	GetDevice               uintptr
	GetPrivateData          uintptr
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
}

type iD3D11InfoQueueVtbl struct {
	iUnknownVtbl

	AddApplicationMessage                        uintptr
	AddMessage                                   uintptr
	AddRetrievalFilterEntries                    uintptr
	AddStorageFilterEntries                      uintptr
	ClearRetrievalFilter                         uintptr
	ClearStorageFilter                           uintptr
	ClearStoredMessages                          uintptr
	GetBreakOnCategory                           uintptr
	GetBreakOnID                                 uintptr
	GetBreakOnSeverity                           uintptr
	GetMessage                                   uintptr
	GetMessageCountLimit                         uintptr
	GetMuteDebugOutput                           uintptr
	GetNumMessagesAllowedByStorageFilter         uintptr
	GetNumMessagesDeniedByStorageFilter          uintptr
	GetNumMessagesDiscardedByMessageCountLimit   uintptr
	GetNumStoredMessages                         uintptr
	GetNumStoredMessagesAllowedByRetrievalFilter uintptr
	GetRetrievalFilter                           uintptr
	GetRetrievalFilterStackSize                  uintptr
	GetStorageFilter                             uintptr
	GetStorageFilterStackSize                    uintptr
	PopRetrievalFilter                           uintptr
	PopStorageFilter                             uintptr
	PushCopyOfRetrievalFilter                    uintptr
	PushCopyOfStorageFilter                      uintptr
	PushEmptyRetrievalFilter                     uintptr
	PushEmptyStorageFilter                       uintptr
	PushRetrievalFilter                          uintptr
	PushStorageFilter                            uintptr
	SetBreakOnCategory                           uintptr
	SetBreakOnID                                 uintptr
	SetBreakOnSeverity                           uintptr
	SetMessageCountLimit                         uintptr
	SetMuteDebugOutput                           uintptr
}

type iD3D11ResourceVtbl struct {
	iD3D11DeviceChildVtbl

	GetType             uintptr
	SetEvictionPriority uintptr
	GetEvictionPriority uintptr
}

type iD3D11Resource struct {
	ole.IUnknown
}

type iD3D11Texture2DVtbl struct {
	iD3D11ResourceVtbl

	GetDesc uintptr
}

type ID3D11Texture2D struct {
	vtbl *iD3D11Texture2DVtbl
}

func (obj *ID3D11Texture2D) GetDesc(desc *D3D11_TEXTURE2D_DESC) int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.GetDesc,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(desc)),
		0,
	)
	return int32(ret)
}
func (obj *ID3D11Texture2D) Release() int32 {
	ret, _, _ := syscall.SyscallN(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return int32(ret)
}
func (obj *ID3D11Texture2D) QueryInterface(iid w32.GUID, pp interface{}) int32 {
	return reflectQueryInterface(obj, obj.vtbl.QueryInterface, &iid, pp)
}

func reflectQueryInterface(self interface{}, method uintptr, interfaceID *w32.GUID, obj interface{}) int32 {
	selfValue := reflect.ValueOf(self).Elem()
	objValue := reflect.ValueOf(obj).Elem()

	hr, _, _ := syscall.SyscallN(
		method,
		3,
		selfValue.UnsafeAddr(),
		uintptr(unsafe.Pointer(interfaceID)),
		objValue.Addr().Pointer())

	return int32(hr)
}

func (v *iD3D11Resource) VTable() *iD3D11ResourceVtbl {
	return (*iD3D11ResourceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ID3D11Device) VTable() *ID3D11DeviceVtbl {
	return (*ID3D11DeviceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (obj *ID3D11Device) CreateTexture2D(desc *D3D11_TEXTURE2D_DESC, ppTexture2D **ID3D11Texture2D) int32 {
	ret, _, _ := syscall.SyscallN(
		obj.VTable().CreateTexture2D,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(desc)),
		0,
		uintptr(unsafe.Pointer(ppTexture2D)),
		0,
		0,
	)
	return int32(ret)
}

func (v *ID3D11Device) GetImmediateContext() (pImmediateContext *ID3D11DeviceContext) {
	syscall.SyscallN(v.VTable().GetImmediateContext, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&pImmediateContext)))
	return pImmediateContext
}

var ID3D11DeviceContextID = ole.NewGUID("{c0bfa96c-e089-44fb-8eaf-26f8796190da}")

type ID3D11DeviceContext struct {
	ole.IUnknown
}

type ID3D11DeviceContextVtbl struct {
	ole.IUnknownVtbl
	GetDevice                                 uintptr
	GetPrivateData                            uintptr
	SetPrivateData                            uintptr
	SetPrivateDataInterface                   uintptr
	VSSetConstantBuffers                      uintptr
	PSSetShaderResources                      uintptr
	PSSetShader                               uintptr
	PSSetSamplers                             uintptr
	VSSetShader                               uintptr
	DrawIndexed                               uintptr
	Draw                                      uintptr
	Map                                       uintptr
	Unmap                                     uintptr
	PSSetConstantBuffers                      uintptr
	IASetInputLayout                          uintptr
	IASetVertexBuffers                        uintptr
	IASetIndexBuffer                          uintptr
	DrawIndexedInstanced                      uintptr
	DrawInstanced                             uintptr
	GSSetConstantBuffers                      uintptr
	GSSetShader                               uintptr
	IASetPrimitiveTopology                    uintptr
	VSSetShaderResources                      uintptr
	VSSetSamplers                             uintptr
	Begin                                     uintptr
	End                                       uintptr
	GetData                                   uintptr
	SetPredication                            uintptr
	GSSetShaderResources                      uintptr
	GSSetSamplers                             uintptr
	OMSetRenderTargets                        uintptr
	OMSetRenderTargetsAndUnorderedAccessViews uintptr
	OMSetBlendState                           uintptr
	OMSetDepthStencilState                    uintptr
	SOSetTargets                              uintptr
	DrawAuto                                  uintptr
	DrawIndexedInstancedIndirect              uintptr
	DrawInstancedIndirect                     uintptr
	Dispatch                                  uintptr
	DispatchIndirect                          uintptr
	RSSetState                                uintptr
	RSSetViewports                            uintptr
	RSSetScissorRects                         uintptr
	CopySubresourceRegion                     uintptr
	CopyResource                              uintptr
	UpdateSubresource                         uintptr
	CopyStructureCount                        uintptr
	ClearRenderTargetView                     uintptr
	ClearUnorderedAccessViewUint              uintptr
	ClearUnorderedAccessViewFloat             uintptr
	ClearDepthStencilView                     uintptr
	GenerateMips                              uintptr
	SetResourceMinLOD                         uintptr
	GetResourceMinLOD                         uintptr
	ResolveSubresource                        uintptr
	ExecuteCommandList                        uintptr
	HSSetShaderResources                      uintptr
	HSSetShader                               uintptr
	HSSetSamplers                             uintptr
	HSSetConstantBuffers                      uintptr
	DSSetShaderResources                      uintptr
	DSSetShader                               uintptr
	DSSetSamplers                             uintptr
	DSSetConstantBuffers                      uintptr
	CSSetShaderResources                      uintptr
	CSSetUnorderedAccessViews                 uintptr
	CSSetShader                               uintptr
	CSSetSamplers                             uintptr
	CSSetConstantBuffers                      uintptr
	VSGetConstantBuffers                      uintptr
	PSGetShaderResources                      uintptr
	PSGetShader                               uintptr
	PSGetSamplers                             uintptr
	VSGetShader                               uintptr
	PSGetConstantBuffers                      uintptr
	IAGetInputLayout                          uintptr
	IAGetVertexBuffers                        uintptr
	IAGetIndexBuffer                          uintptr
	GSGetConstantBuffers                      uintptr
	GSGetShader                               uintptr
	IAGetPrimitiveTopology                    uintptr
	VSGetShaderResources                      uintptr
	VSGetSamplers                             uintptr
	GetPredication                            uintptr
	GSGetShaderResources                      uintptr
	GSGetSamplers                             uintptr
	OMGetRenderTargets                        uintptr
	OMGetRenderTargetsAndUnorderedAccessViews uintptr
	OMGetBlendState                           uintptr
	OMGetDepthStencilState                    uintptr
	SOGetTargets                              uintptr
	RSGetState                                uintptr
	RSGetViewports                            uintptr
	RSGetScissorRects                         uintptr
	HSGetShaderResources                      uintptr
	HSGetShader                               uintptr
	HSGetSamplers                             uintptr
	HSGetConstantBuffers                      uintptr
	DSGetShaderResources                      uintptr
	DSGetShader                               uintptr
	DSGetSamplers                             uintptr
	DSGetConstantBuffers                      uintptr
	CSGetShaderResources                      uintptr
	CSGetUnorderedAccessViews                 uintptr
	CSGetShader                               uintptr
	CSGetSamplers                             uintptr
	CSGetConstantBuffers                      uintptr
	ClearState                                uintptr
	Flush                                     uintptr
	GetType                                   uintptr
	GetContextFlags                           uintptr
	FinishCommandList                         uintptr
}

type D3D11_MAPPED_SUBRESOURCE struct {
	PData      uintptr
	RowPitch   uint32
	DepthPitch uint32
}

type D3D11_MAP uint32

const (
	D3D11_MAP_READ               D3D11_MAP = 1
	D3D11_MAP_WRITE              D3D11_MAP = 2
	D3D11_MAP_READ_WRITE         D3D11_MAP = 3
	D3D11_MAP_WRITE_DISCARD      D3D11_MAP = 4
	D3D11_MAP_WRITE_NO_OVERWRITE D3D11_MAP = 5
)

func (v *ID3D11DeviceContext) VTable() *ID3D11DeviceContextVtbl {
	return (*ID3D11DeviceContextVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ID3D11DeviceContextVtbl) CopyResourceF(in *ID3D11Texture2D, out *ID3D11Texture2D) error {
	r1, _, err := syscall.SyscallN(uintptr(v.CopyResource), uintptr(unsafe.Pointer(in)), uintptr(unsafe.Pointer(out)))
	if r1 != win.S_OK {
		return err
	}
	return nil
}

func (v *ID3D11DeviceContextVtbl) MapF(pResource *ID3D11Texture2D, subresource uintptr, mapType D3D11_MAP, mapFlags uintptr, pMappedResource *D3D11_MAPPED_SUBRESOURCE) error {
	_, _, _ = syscall.SyscallN(
		uintptr(v.Map),
		uintptr(unsafe.Pointer(pResource)),
		uintptr(subresource),
		uintptr(mapType),
		uintptr(mapFlags),
		uintptr(unsafe.Pointer(pMappedResource)),
	)
	return nil
}

func (v *ID3D11DeviceContextVtbl) UnmapF(pResource *ID3D11Texture2D, subresource uintptr) error {
	_, _, _ = syscall.SyscallN(
		uintptr(v.Unmap),
		uintptr(unsafe.Pointer(pResource)),
		uintptr(subresource),
	)
	return nil
}

var pD3DCreateDevice = d3d11DLL.NewProc("D3D11CreateDevice")

// CreateDevice
// https://learn.microsoft.com/en-us/windows/win32/api/d3d11/nf-d3d11-d3d11createdevice
func D3D11CreateDevice(
	pAdapter *IDXGIAdapter,
	DriverType D3D_DRIVER_TYPE,
	Software win.HMODULE,
	Flags D3D11_CREATE_DEVICE_FLAG,
	pFeatureLevels *D3D_FEATURE_LEVEL,
	FeatureLevels int,
	SDKVersion uint32,
	ppDevice **ID3D11Device,
	pFeatureLevel *D3D_FEATURE_LEVEL,
	ppImmediateContext **ID3D11DeviceContext,
) error {
	r1, _, _ := pD3DCreateDevice.Call(
		uintptr(unsafe.Pointer(pAdapter)),
		uintptr(DriverType),
		uintptr(Software),
		uintptr(Flags),
		uintptr(unsafe.Pointer(pFeatureLevels)),
		uintptr(FeatureLevels),
		uintptr(SDKVersion),
		uintptr(unsafe.Pointer(ppDevice)),
		uintptr(unsafe.Pointer(pFeatureLevel)),
		uintptr(unsafe.Pointer(ppImmediateContext)),
	)
	if r1 != win.S_OK {
		return ole.NewError(r1)
	}
	return nil
}

var pCreateDirect3D11DeviceFromDXGIDevice = d3d11DLL.NewProc("CreateDirect3D11DeviceFromDXGIDevice")

func CreateDirect3D11DeviceFromDXGIDevice(dxgiDevice *IDXGIDevice, graphicsDevice **ole.IInspectable) error {
	r1, _, err := pCreateDirect3D11DeviceFromDXGIDevice.Call(uintptr(unsafe.Pointer(dxgiDevice)), uintptr(unsafe.Pointer(graphicsDevice)))
	if r1 != win.S_OK {
		return err
	}
	return nil
}
