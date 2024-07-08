package winapi

import (
	"unsafe"

	"github.com/lxn/win"
	"github.com/moutend/go-wca/pkg/wca"
	"github.com/whiteboxsolutions/go-ole"
)

type IID ole.GUID
type REFIID IID

type IActivateAudioInterfaceCompletionHandlerInterface interface {
	ole.UnknownLike
	ActivateCompleted(activateOperation IActivateAudioInterfaceAsyncOperation)
}

type IActivateAudioInterfaceAsyncOperation interface {
	ole.UnknownLike
	GetActivateResult(activateResult win.HRESULT, activateInterface ole.IUnknown)
}

func ActivateAudioInterfaceAsync(deviceInterfacePath *uint16, riid REFIID, activationParams wca.PROPVARIANT, completionHandler IActivateAudioInterfaceCompletionHandlerInterface, activationOperation IActivateAudioInterfaceAsyncOperation) win.HRESULT {
	return win.HRESULT(activateAudioInterfaceAsync(
		deviceInterfacePath,
		uintptr(unsafe.Pointer(&riid.Data1)),
		uintptr(unsafe.Pointer(&activationParams.VT)),
		uintptr(unsafe.Pointer(&completionHandler)),
		uintptr(unsafe.Pointer(&activationOperation)),
	))
}
