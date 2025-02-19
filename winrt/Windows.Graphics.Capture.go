package winrt

import (
	"errors"
	"fmt"
	"runtime"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
	"github.com/whiteboxsolutions/go-ole"
)

// GraphicsCaptureItemClass
// https://learn.microsoft.com/en-us/uwp/api/windows.graphics.capture.graphicscaptureitem?view=winrt-22621

var GraphicsCaptureItemClass = "Windows.Graphics.Capture.GraphicsCaptureItem"

// IGraphicsCaptureItem

var IGraphicsCaptureItemID = ole.NewGUID("{79c3f95b-31f7-4ec2-a464-632ef5d30760}")
var IGraphicsCaptureItemClass = "Windows.Graphics.Capture.IGraphicsCaptureItem"

type TimeSpan struct {
	Duration int64
}

type IGraphicsCaptureItem struct {
	ole.IInspectable
}

type IGraphicsCaptureItemVtbl struct {
	ole.IInspectableVtbl
	DisplayName   uintptr
	Size          uintptr
	add_Closed    uintptr
	remove_Closed uintptr
}

func (v *IGraphicsCaptureItem) VTable() *IGraphicsCaptureItemVtbl {
	return (*IGraphicsCaptureItemVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IGraphicsCaptureItem) DisplayName() (string, error) {
	var hRet ole.HString

	r1, _, _ := syscall.SyscallN(v.VTable().DisplayName, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&hRet)))
	if r1 != 0 {
		return "", ole.NewError(r1)
	}

	var ret = hRet.String()
	ole.DeleteHString(hRet)

	return ret, nil
}

func (v *IGraphicsCaptureItem) Size() (*SizeInt32, error) {
	var size SizeInt32
	r1, _, _ := syscall.SyscallN(v.VTable().Size, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&size.Width)))
	if r1 != 0 {
		return nil, ole.NewError(r1)
	}
	return &size, nil
}

// IGraphicsCaptureItemStatics

var IGraphicsCaptureItemStaticsID = ole.NewGUID("{a87ebea5-457c-5788-ab47-0cf1d3637e74}")
var IGraphicsCaptureItemStaticsClass = "Windows.Graphics.Capture.IGraphicsCaptureItemStatics"

type IGraphicsCaptureItemStatics struct {
	ole.IInspectable
}

type IGraphicsCaptureItemStaticsVtbl struct {
	ole.IInspectableVtbl
	CreateFromVisual uintptr
}

func (v *IGraphicsCaptureItemStatics) VTable() *IGraphicsCaptureItemStaticsVtbl {
	return (*IGraphicsCaptureItemStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

// IGraphicsCaptureItemStatics2

var IGraphicsCaptureItemStatics2ID = ole.NewGUID("{3b92acc9-e584-5862-bf5c-9c316c6d2dbb}")
var IGraphicsCaptureItemStatics2Class = "Windows.Graphics.Capture.IGraphicsCaptureItemStatics2"

type IGraphicsCaptureItemStatics2 struct {
	ole.IInspectable
}

type IGraphicsCaptureItemStatics2Vtbl struct {
	ole.IInspectableVtbl
	TryCreateFromWindowId  uintptr
	TryCreateFromDisplayId uintptr
}

func (v *IGraphicsCaptureItemStatics2) VTable() *IGraphicsCaptureItemStatics2Vtbl {
	return (*IGraphicsCaptureItemStatics2Vtbl)(unsafe.Pointer(v.RawVTable))
}

// Direct3D11CaptureFramePool
// https://learn.microsoft.com/en-us/uwp/api/windows.graphics.capture.direct3d11captureframepool?view=winrt-22621
var Direct3D11CaptureFramePoolClass = "Windows.Graphics.Capture.Direct3D11CaptureFramePool"

// IDirect3D11CaptureFramePool

var IDirect3D11CaptureFramePoolID = ole.NewGUID("{24EB6D22-1975-422E-82E7-780DBD8DDF24}")

type IDirect3D11CaptureFramePool struct {
	ole.IUnknown
}

type IDirect3D11CaptureFramePoolVtbl struct {
	ole.IInspectableVtbl
	Recreate             uintptr
	TryGetNextFrame      uintptr
	add_FrameArrived     uintptr
	remove_FrameArrived  uintptr
	CreateCaptureSession uintptr
	get_DispatcherQueue  uintptr
}

var generatedDirect3D11CaptureFramePool = map[uintptr]*Direct3D11CaptureFramePoolVtbl{}

type Direct3D11CaptureFramePool struct {
	ole.IUnknown
}

type Direct3D11CaptureFramePoolVtbl struct {
	ole.IInspectableVtbl
	Invoke  uintptr
	counter *int
}

func (v *IDirect3D11CaptureFramePool) VTable() *IDirect3D11CaptureFramePoolVtbl {
	return (*IDirect3D11CaptureFramePoolVtbl)(unsafe.Pointer(v.RawVTable))
}

// type Direct3D11CaptureFramePoolFrameArrivedProcType func(this *Direct3D11CaptureFramePool, sender *IDirect3D11CaptureFramePool, args *ole.IInspectable) uintptr
type Direct3D11CaptureFramePoolFrameArrivedProcType func(this *uintptr, sender *IDirect3D11CaptureFramePool, args *ole.IInspectable) uintptr

/*
eventHandler:

	interface {
		IUnknown
		Invoke(sender *IDirect3D11CaptureFramePool, args *ole.IInspectable) uintptr
	}
*/
func (v *IDirect3D11CaptureFramePool) AddFrameArrived(eventHandler unsafe.Pointer) (*EventRegistrationToken, error) {
	var token EventRegistrationToken
	r1, _, _ := syscall.SyscallN(v.VTable().add_FrameArrived, uintptr(unsafe.Pointer(v)), uintptr(eventHandler), uintptr(unsafe.Pointer(&token.value)))
	if r1 != win.S_OK {
		fmt.Println("Not S_OK in AddFrameArrived")
		return nil, ole.NewError(r1)
	}
	fmt.Println("token:", token)
	return &token, nil
}

func (v *IDirect3D11CaptureFramePool) RemoveFrameArrived(token *EventRegistrationToken) error {
	r1, _, _ := syscall.SyscallN(v.VTable().remove_FrameArrived, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&token.value)))
	if r1 != win.S_OK {
		return ole.NewError(r1)
	}
	return nil
}

func (v *IDirect3D11CaptureFramePool) CreateCaptureSession(item *IGraphicsCaptureItem) (*IGraphicsCaptureSession, error) {
	var session *IGraphicsCaptureSession
	r1, _, _ := syscall.SyscallN(v.VTable().CreateCaptureSession, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(item)), uintptr(unsafe.Pointer(&session)))
	if r1 != win.S_OK {
		return nil, ole.NewError(r1)
	}

	return session, nil
}
func NewDirect3D11CaptureFramePool(invoke Direct3D11CaptureFramePoolFrameArrivedProcType) *Direct3D11CaptureFramePool {
	var counter = 1
	var v = &Direct3D11CaptureFramePoolVtbl{
		Invoke:  syscall.NewCallback(invoke),
		counter: &counter,
	}

	var newV = new(Direct3D11CaptureFramePool)
	newV.RawVTable = (*interface{})(unsafe.Pointer(v))

	v.QueryInterface = syscall.NewCallback(newV.queryInterface)
	v.AddRef = syscall.NewCallback(newV.addRef)
	v.Release = syscall.NewCallback(newV.release)

	generatedDirect3D11CaptureFramePool[uintptr(unsafe.Pointer(newV))] = v

	return newV
}

func (v *IDirect3D11CaptureFramePool) TryGetNextFrame() (*IDirect3D11CaptureFrame, error) {
	var frame *IDirect3D11CaptureFrame
	r1, _, _ := syscall.SyscallN(v.VTable().TryGetNextFrame, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&frame)))
	if r1 != win.S_OK {
		return nil, ole.NewError(r1)
	}

	return frame, nil
}

func (v *Direct3D11CaptureFramePool) VTable() *Direct3D11CaptureFramePoolVtbl {
	return (*Direct3D11CaptureFramePoolVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *Direct3D11CaptureFramePool) Invoke(sender *IDirect3D11CaptureFramePool, args *ole.IInspectable) error {
	r1, _, _ := syscall.SyscallN(v.VTable().Invoke, uintptr(unsafe.Pointer(sender)), uintptr(unsafe.Pointer(args)))
	return ole.NewError(r1)
}

// QueryInterface(vp *Direct3D11CaptureFramePool, riid ole.GUID, lppvObj **ole.Inspectable)
func (v *Direct3D11CaptureFramePool) queryInterface(lpMyObj *uintptr, riid *uintptr, lppvObj **uintptr) uintptr {
	// Validate input
	if lpMyObj == nil {
		return win.E_INVALIDARG
	}

	var V = new(Direct3D11CaptureFramePool)

	var err error
	// Check dereferencability
	func() {
		defer func() {
			if recover() != nil {
				err = errors.New("InvalidObject")
			}
		}()
		// if object cannot be dereferenced, then panic occurs
		*V = *(*Direct3D11CaptureFramePool)(unsafe.Pointer(lpMyObj))
		V.VTable()
	}()
	if err != nil {
		return win.E_INVALIDARG
	}

	*lppvObj = nil
	var id = new(ole.GUID)
	*id = *(*ole.GUID)(unsafe.Pointer(riid))

	// Convert
	switch id.String() {
	case ole.IID_IUnknown.String(), ITypedEventHandlerID.String(), IAgileObjectID.String():
		V.AddRef()
		*lppvObj = (*uintptr)(unsafe.Pointer(V))

		return win.S_OK
	default:
		return win.E_NOINTERFACE
	}
}

func (v *Direct3D11CaptureFramePool) addRef(lpMyObj *uintptr) uintptr {
	// Validate input
	if lpMyObj == nil {
		return 0
	}

	var V = (*Direct3D11CaptureFramePool)(unsafe.Pointer(lpMyObj))
	*V.VTable().counter++

	return uintptr(*V.VTable().counter)
}

func (v *Direct3D11CaptureFramePool) release(lpMyObj *uintptr) uintptr {
	// Validate input
	if lpMyObj == nil {
		return 0
	}

	var V = (*Direct3D11CaptureFramePool)(unsafe.Pointer(lpMyObj))
	*V.VTable().counter--

	if *V.VTable().counter == 0 {
		V.RawVTable = nil
		_, ok := generatedDirect3D11CaptureFramePool[uintptr(unsafe.Pointer(lpMyObj))]
		if ok {
			delete(generatedDirect3D11CaptureFramePool, uintptr(unsafe.Pointer(lpMyObj)))
			runtime.GC()
		}
		return 0
	}

	return uintptr(*V.VTable().counter)
}

// IDirect3D11CaptureFramePoolStatics

var IDirect3D11CaptureFramePoolStaticsID = ole.NewGUID("{7784056A-67AA-4D53-AE54-1088D5A8CA21}")

type IDirect3D11CaptureFramePoolStatics struct {
	ole.IInspectable
}

type IDirect3D11CaptureFramePoolStaticsVtbl struct {
	ole.IInspectableVtbl
	Create uintptr
}

func (v *IDirect3D11CaptureFramePoolStatics) VTable() *IDirect3D11CaptureFramePoolStaticsVtbl {
	return (*IDirect3D11CaptureFramePoolStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IDirect3D11CaptureFramePoolStatics) Create(device *IDirect3DDevice, pixelFormat DirectXPixelFormat, numberOfBuffers int32, size *SizeInt32) (*IDirect3D11CaptureFramePool, error) {
	var ret *IDirect3D11CaptureFramePool
	r1, _, _ := syscall.SyscallN(
		v.VTable().Create, uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(device)), uintptr(pixelFormat), uintptr(numberOfBuffers), uintptr(size.Width)<<32+uintptr(size.Height), uintptr(unsafe.Pointer(&ret)),
	)
	if r1 != win.S_OK {
		return nil, ole.NewError(r1)
	}

	return ret, nil
}

// IDirect3D11CaptureFramePoolStatics2

var IDirect3D11CaptureFramePoolStatics2ID = ole.NewGUID("{589B103F-6BBC-5DF5-A991-02E28B3B66D5}")

type IDirect3D11CaptureFramePoolStatics2 struct {
	ole.IInspectable
}

type IDirect3D11CaptureFramePoolStatics2Vtbl struct {
	ole.IInspectableVtbl
	CreateFreeThreaded uintptr
}

func (v *IDirect3D11CaptureFramePoolStatics2) VTable() *IDirect3D11CaptureFramePoolStatics2Vtbl {
	return (*IDirect3D11CaptureFramePoolStatics2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IDirect3D11CaptureFramePoolStatics2) CreateFreeThreaded(device *IDirect3DDevice, pixelFormat DirectXPixelFormat, numberOfBuffers int32, size *SizeInt32) (*IDirect3D11CaptureFramePool, error) {
	var ret *IDirect3D11CaptureFramePool
	r1, _, _ := syscall.SyscallN(
		v.VTable().CreateFreeThreaded, uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(device)), uintptr(pixelFormat), uintptr(numberOfBuffers), uintptr(size.Width)<<32+uintptr(size.Height), uintptr(unsafe.Pointer(&ret)),
	)
	if r1 != win.S_OK {
		return nil, ole.NewError(r1)
	}

	return ret, nil
}

type IGraphicsCaptureSession struct {
	ole.IInspectable
}

type IGraphicsCaptureSessionVtbl struct {
	ole.IInspectableVtbl
	StartCapture uintptr
}

var IGraphicsCaptureSessionID = ole.NewGUID("{814E42A9-F70F-4AD7-939B-FDDCC6EB880D}")

func (v *IGraphicsCaptureSession) VTable() *IGraphicsCaptureSessionVtbl {
	return (*IGraphicsCaptureSessionVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IGraphicsCaptureSession) StartCapture() error {
	r1, _, _ := syscall.SyscallN(v.VTable().StartCapture, uintptr(unsafe.Pointer(v)))
	if r1 != win.S_OK {
		return ole.NewError(r1)
	}

	return nil
}

type IGraphicsCaptureSession2 struct {
	ole.IInspectable
}

type IGraphicsCaptureSession2Vtbl struct {
	ole.IInspectableVtbl
	get_IsCursorCaptureEnabled uintptr
	put_IsCursorCaptureEnabled uintptr
}

var IGraphicsCaptureSession2ID = ole.NewGUID("{2C39AE40-7D2E-5044-804E-8B6799D4CF9E}")

func (v *IGraphicsCaptureSession2) VTable() *IGraphicsCaptureSession2Vtbl {
	return (*IGraphicsCaptureSession2Vtbl)(unsafe.Pointer(v.RawVTable))
}

type IGraphicsCaptureSession3 struct {
	ole.IInspectable
}

type IGraphicsCaptureSession3Vtbl struct {
	ole.IInspectableVtbl
	get_IsBorderRequired uintptr
	put_IsBorderRequired uintptr
}

var IGraphicsCaptureSession3ID = ole.NewGUID("{F2CDD966-22AE-5EA1-9596-3A289344C3BE}")

func (v *IGraphicsCaptureSession3) VTable() *IGraphicsCaptureSession3Vtbl {
	return (*IGraphicsCaptureSession3Vtbl)(unsafe.Pointer(v.RawVTable))
}

type IGraphicsCaptureSessionStatics struct {
	ole.IInspectable
}

type IGraphicsCaptureSessionStaticsVtbl struct {
	ole.IInspectableVtbl
	IsSupported uintptr
}

var IGraphicsCaptureSessionStaticsID = ole.NewGUID("{2224A540-5974-49AA-B232-0882536F4CB5}")

func (v *IGraphicsCaptureSessionStatics) VTable() *IGraphicsCaptureSessionStaticsVtbl {
	return (*IGraphicsCaptureSessionStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IGraphicsCaptureSessionStatics) IsSupported() (ok bool) {
	syscall.SyscallN(v.VTable().IsSupported, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&ok)))
	return ok
}

var IDirect3D11CaptureFrameClass = "Windows.Graphics.Capture.IDirect3D11CaptureFrame"
var IDirect3D11CaptureFrameID = "{FA50C623-38DA-4B32-ACF3-FA9734AD800E}"

type IDirect3D11CaptureFrame struct {
	ole.IInspectable
}

type IDirect3D11CaptureFrameVtbl struct {
	ole.IInspectableVtbl
	Surface                uintptr
	SystemRelativeTime     uintptr
	ContentSize            uintptr
	Get_Surface            uintptr
	Get_SystemRelativeTime uintptr
	Get_ContentSize        uintptr
}

type IDirect3D11Surface struct {
	ole.IUnknown
}

type IDirect3D11SurfaceVtbl struct {
	ole.IUnknownVtbl
	Description uintptr
	Dispose     uintptr
}

type IDirect3DSurfaceDescription struct {
	Format                 DirectXPixelFormat
	Height                 int
	MultisampleDescription Direct3DMultisampleDescription
	Width                  int
}

type Direct3DMultisampleDescription struct {
	Count   int
	Quality int
}

func (v *IDirect3D11CaptureFrame) VTable() *IDirect3D11CaptureFrameVtbl {
	return (*IDirect3D11CaptureFrameVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IDirect3D11CaptureFrame) Get_Surface() IDirect3D11Surface {
	var _result IDirect3D11Surface
	_hr, _, _ := syscall.SyscallN(v.VTable().Get_Surface, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (v *IDirect3D11CaptureFrame) Get_SystemRelativeTime() TimeSpan {
	var _result TimeSpan
	_hr, _, _ := syscall.SyscallN(v.VTable().Get_SystemRelativeTime, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (v *IDirect3D11CaptureFrame) Get_ContentSize() unsafe.Pointer {
	var _result unsafe.Pointer
	_hr, _, _ := syscall.SyscallN(v.VTable().Get_ContentSize, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}

func (v *IDirect3D11Surface) VTable() *IDirect3D11SurfaceVtbl {
	return (*IDirect3D11SurfaceVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IDirect3D11Surface) Get_Description() IDirect3DSurfaceDescription {
	var _result IDirect3DSurfaceDescription
	_hr, _, _ := syscall.SyscallN(v.VTable().Description, uintptr(unsafe.Pointer(v)), uintptr(unsafe.Pointer(&_result)))
	_ = _hr
	return _result
}
