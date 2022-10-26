//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package test

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = errors.New
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding StructSimple final model
type FinalModelStructSimple struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset

    Id *fbe.FinalModelInt32
    F1 *fbe.FinalModelBool
    F2 *fbe.FinalModelBool
    F3 *fbe.FinalModelByte
    F4 *fbe.FinalModelByte
    F5 *fbe.FinalModelChar
    F6 *fbe.FinalModelChar
    F7 *fbe.FinalModelWChar
    F8 *fbe.FinalModelWChar
    F9 *fbe.FinalModelInt8
    F10 *fbe.FinalModelInt8
    F11 *fbe.FinalModelUInt8
    F12 *fbe.FinalModelUInt8
    F13 *fbe.FinalModelInt16
    F14 *fbe.FinalModelInt16
    F15 *fbe.FinalModelUInt16
    F16 *fbe.FinalModelUInt16
    F17 *fbe.FinalModelInt32
    F18 *fbe.FinalModelInt32
    F19 *fbe.FinalModelUInt32
    F20 *fbe.FinalModelUInt32
    F21 *fbe.FinalModelInt64
    F22 *fbe.FinalModelInt64
    F23 *fbe.FinalModelUInt64
    F24 *fbe.FinalModelUInt64
    F25 *fbe.FinalModelFloat
    F26 *fbe.FinalModelFloat
    F27 *fbe.FinalModelDouble
    F28 *fbe.FinalModelDouble
    F29 *fbe.FinalModelDecimal
    F30 *fbe.FinalModelDecimal
    F31 *fbe.FinalModelString
    F32 *fbe.FinalModelString
    F33 *fbe.FinalModelTimestamp
    F34 *fbe.FinalModelTimestamp
    F35 *fbe.FinalModelTimestamp
    F36 *fbe.FinalModelUUID
    F37 *fbe.FinalModelUUID
    F38 *fbe.FinalModelUUID
    F39 *proto.FinalModelOrderSide
    F40 *proto.FinalModelOrderType
    F41 *proto.FinalModelOrder
    F42 *proto.FinalModelBalance
    F43 *proto.FinalModelState
    F44 *proto.FinalModelAccount
}

// Create a new StructSimple final model
func NewFinalModelStructSimple(buffer *fbe.Buffer, offset int) *FinalModelStructSimple {
    fbeResult := FinalModelStructSimple{buffer: buffer, offset: offset}
    fbeResult.Id = fbe.NewFinalModelInt32(buffer, 0)
    fbeResult.F1 = fbe.NewFinalModelBool(buffer, 0)
    fbeResult.F2 = fbe.NewFinalModelBool(buffer, 0)
    fbeResult.F3 = fbe.NewFinalModelByte(buffer, 0)
    fbeResult.F4 = fbe.NewFinalModelByte(buffer, 0)
    fbeResult.F5 = fbe.NewFinalModelChar(buffer, 0)
    fbeResult.F6 = fbe.NewFinalModelChar(buffer, 0)
    fbeResult.F7 = fbe.NewFinalModelWChar(buffer, 0)
    fbeResult.F8 = fbe.NewFinalModelWChar(buffer, 0)
    fbeResult.F9 = fbe.NewFinalModelInt8(buffer, 0)
    fbeResult.F10 = fbe.NewFinalModelInt8(buffer, 0)
    fbeResult.F11 = fbe.NewFinalModelUInt8(buffer, 0)
    fbeResult.F12 = fbe.NewFinalModelUInt8(buffer, 0)
    fbeResult.F13 = fbe.NewFinalModelInt16(buffer, 0)
    fbeResult.F14 = fbe.NewFinalModelInt16(buffer, 0)
    fbeResult.F15 = fbe.NewFinalModelUInt16(buffer, 0)
    fbeResult.F16 = fbe.NewFinalModelUInt16(buffer, 0)
    fbeResult.F17 = fbe.NewFinalModelInt32(buffer, 0)
    fbeResult.F18 = fbe.NewFinalModelInt32(buffer, 0)
    fbeResult.F19 = fbe.NewFinalModelUInt32(buffer, 0)
    fbeResult.F20 = fbe.NewFinalModelUInt32(buffer, 0)
    fbeResult.F21 = fbe.NewFinalModelInt64(buffer, 0)
    fbeResult.F22 = fbe.NewFinalModelInt64(buffer, 0)
    fbeResult.F23 = fbe.NewFinalModelUInt64(buffer, 0)
    fbeResult.F24 = fbe.NewFinalModelUInt64(buffer, 0)
    fbeResult.F25 = fbe.NewFinalModelFloat(buffer, 0)
    fbeResult.F26 = fbe.NewFinalModelFloat(buffer, 0)
    fbeResult.F27 = fbe.NewFinalModelDouble(buffer, 0)
    fbeResult.F28 = fbe.NewFinalModelDouble(buffer, 0)
    fbeResult.F29 = fbe.NewFinalModelDecimal(buffer, 0)
    fbeResult.F30 = fbe.NewFinalModelDecimal(buffer, 0)
    fbeResult.F31 = fbe.NewFinalModelString(buffer, 0)
    fbeResult.F32 = fbe.NewFinalModelString(buffer, 0)
    fbeResult.F33 = fbe.NewFinalModelTimestamp(buffer, 0)
    fbeResult.F34 = fbe.NewFinalModelTimestamp(buffer, 0)
    fbeResult.F35 = fbe.NewFinalModelTimestamp(buffer, 0)
    fbeResult.F36 = fbe.NewFinalModelUUID(buffer, 0)
    fbeResult.F37 = fbe.NewFinalModelUUID(buffer, 0)
    fbeResult.F38 = fbe.NewFinalModelUUID(buffer, 0)
    fbeResult.F39 = proto.NewFinalModelOrderSide(buffer, 0)
    fbeResult.F40 = proto.NewFinalModelOrderType(buffer, 0)
    fbeResult.F41 = proto.NewFinalModelOrder(buffer, 0)
    fbeResult.F42 = proto.NewFinalModelBalance(buffer, 0)
    fbeResult.F43 = proto.NewFinalModelState(buffer, 0)
    fbeResult.F44 = proto.NewFinalModelAccount(buffer, 0)
    return &fbeResult
}

// Get the allocation size
func (fm *FinalModelStructSimple) FBEAllocationSize(fbeValue *StructSimple) int {
    fbeResult := 0 +
        fm.Id.FBEAllocationSize(fbeValue.Id) +
        fm.F1.FBEAllocationSize(fbeValue.F1) +
        fm.F2.FBEAllocationSize(fbeValue.F2) +
        fm.F3.FBEAllocationSize(fbeValue.F3) +
        fm.F4.FBEAllocationSize(fbeValue.F4) +
        fm.F5.FBEAllocationSize(fbeValue.F5) +
        fm.F6.FBEAllocationSize(fbeValue.F6) +
        fm.F7.FBEAllocationSize(fbeValue.F7) +
        fm.F8.FBEAllocationSize(fbeValue.F8) +
        fm.F9.FBEAllocationSize(fbeValue.F9) +
        fm.F10.FBEAllocationSize(fbeValue.F10) +
        fm.F11.FBEAllocationSize(fbeValue.F11) +
        fm.F12.FBEAllocationSize(fbeValue.F12) +
        fm.F13.FBEAllocationSize(fbeValue.F13) +
        fm.F14.FBEAllocationSize(fbeValue.F14) +
        fm.F15.FBEAllocationSize(fbeValue.F15) +
        fm.F16.FBEAllocationSize(fbeValue.F16) +
        fm.F17.FBEAllocationSize(fbeValue.F17) +
        fm.F18.FBEAllocationSize(fbeValue.F18) +
        fm.F19.FBEAllocationSize(fbeValue.F19) +
        fm.F20.FBEAllocationSize(fbeValue.F20) +
        fm.F21.FBEAllocationSize(fbeValue.F21) +
        fm.F22.FBEAllocationSize(fbeValue.F22) +
        fm.F23.FBEAllocationSize(fbeValue.F23) +
        fm.F24.FBEAllocationSize(fbeValue.F24) +
        fm.F25.FBEAllocationSize(fbeValue.F25) +
        fm.F26.FBEAllocationSize(fbeValue.F26) +
        fm.F27.FBEAllocationSize(fbeValue.F27) +
        fm.F28.FBEAllocationSize(fbeValue.F28) +
        fm.F29.FBEAllocationSize(fbeValue.F29) +
        fm.F30.FBEAllocationSize(fbeValue.F30) +
        fm.F31.FBEAllocationSize(fbeValue.F31) +
        fm.F32.FBEAllocationSize(fbeValue.F32) +
        fm.F33.FBEAllocationSize(fbeValue.F33) +
        fm.F34.FBEAllocationSize(fbeValue.F34) +
        fm.F35.FBEAllocationSize(fbeValue.F35) +
        fm.F36.FBEAllocationSize(fbeValue.F36) +
        fm.F37.FBEAllocationSize(fbeValue.F37) +
        fm.F38.FBEAllocationSize(fbeValue.F38) +
        fm.F39.FBEAllocationSize(&fbeValue.F39) +
        fm.F40.FBEAllocationSize(&fbeValue.F40) +
        fm.F41.FBEAllocationSize(&fbeValue.F41) +
        fm.F42.FBEAllocationSize(&fbeValue.F42) +
        fm.F43.FBEAllocationSize(&fbeValue.F43) +
        fm.F44.FBEAllocationSize(&fbeValue.F44) +
        0
    return fbeResult
}

// Get the final size
func (fm *FinalModelStructSimple) FBESize() int { return 0 }

// Get the final extra size
func (fm *FinalModelStructSimple) FBEExtra() int { return 0 }

// Get the final type
func (fm *FinalModelStructSimple) FBEType() int { return 110 }

// Get the final offset
func (fm *FinalModelStructSimple) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelStructSimple) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelStructSimple) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelStructSimple) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FinalModelStructSimple) Verify() int {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult := fm.VerifyFields()
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult
}

// Check if the struct fields are valid
func (fm *FinalModelStructSimple) VerifyFields() int {
    fbeCurrentOffset := 0
    fbeFieldSize := 0


    fm.Id.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.Id.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F1.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F2.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F3.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F3.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F4.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F4.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F5.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F5.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F6.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F6.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F7.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F7.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F8.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F8.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F9.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F9.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F10.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F10.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F11.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F11.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F12.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F12.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F13.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F13.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F14.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F14.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F15.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F15.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F16.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F16.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F17.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F17.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F18.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F18.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F19.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F19.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F20.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F20.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F21.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F21.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F22.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F22.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F23.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F23.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F24.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F24.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F25.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F25.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F26.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F26.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F27.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F27.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F28.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F28.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F29.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F29.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F30.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F30.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F31.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F31.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F32.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F32.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F33.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F33.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F34.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F34.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F35.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F35.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F36.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F36.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F37.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F37.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F38.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F38.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F39.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F39.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F40.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F40.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F41.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F41.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F42.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F42.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F43.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F43.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    fm.F44.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize = fm.F44.Verify(); fbeFieldSize == fbe.MaxInt {
        return fbe.MaxInt
    }
    fbeCurrentOffset += fbeFieldSize

    return fbeCurrentOffset
}

// Get the struct value
func (fm *FinalModelStructSimple) Get() (*StructSimple, int, error) {
    fbeResult := NewStructSimple()
    fbeSize, err := fm.GetValue(fbeResult)
    return fbeResult, fbeSize, err
}

// Get the struct value by the given pointer
func (fm *FinalModelStructSimple) GetValue(fbeValue *StructSimple) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeSize, err := fm.GetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeSize, err
}

// Get the struct fields values
func (fm *FinalModelStructSimple) GetFields(fbeValue *StructSimple) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.Id.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.Id, fbeFieldSize, err = fm.Id.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F1, fbeFieldSize, err = fm.F1.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F2, fbeFieldSize, err = fm.F2.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F3.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F3, fbeFieldSize, err = fm.F3.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F4.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F4, fbeFieldSize, err = fm.F4.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F5.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F5, fbeFieldSize, err = fm.F5.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F6.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F6, fbeFieldSize, err = fm.F6.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F7.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F7, fbeFieldSize, err = fm.F7.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F8.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F8, fbeFieldSize, err = fm.F8.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F9.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F9, fbeFieldSize, err = fm.F9.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F10.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F10, fbeFieldSize, err = fm.F10.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F11.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F11, fbeFieldSize, err = fm.F11.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F12.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F12, fbeFieldSize, err = fm.F12.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F13.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F13, fbeFieldSize, err = fm.F13.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F14.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F14, fbeFieldSize, err = fm.F14.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F15.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F15, fbeFieldSize, err = fm.F15.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F16.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F16, fbeFieldSize, err = fm.F16.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F17.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F17, fbeFieldSize, err = fm.F17.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F18.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F18, fbeFieldSize, err = fm.F18.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F19.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F19, fbeFieldSize, err = fm.F19.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F20.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F20, fbeFieldSize, err = fm.F20.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F21.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F21, fbeFieldSize, err = fm.F21.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F22.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F22, fbeFieldSize, err = fm.F22.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F23.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F23, fbeFieldSize, err = fm.F23.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F24.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F24, fbeFieldSize, err = fm.F24.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F25.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F25, fbeFieldSize, err = fm.F25.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F26.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F26, fbeFieldSize, err = fm.F26.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F27.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F27, fbeFieldSize, err = fm.F27.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F28.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F28, fbeFieldSize, err = fm.F28.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F29.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F29, fbeFieldSize, err = fm.F29.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F30.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F30, fbeFieldSize, err = fm.F30.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F31.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F31, fbeFieldSize, err = fm.F31.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F32.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F32, fbeFieldSize, err = fm.F32.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F33.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F33, fbeFieldSize, err = fm.F33.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F34.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F34, fbeFieldSize, err = fm.F34.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F35.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F35, fbeFieldSize, err = fm.F35.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F36.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F36, fbeFieldSize, err = fm.F36.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F37.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F37, fbeFieldSize, err = fm.F37.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F38.SetFBEOffset(fbeCurrentOffset)
    if fbeValue.F38, fbeFieldSize, err = fm.F38.Get(); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F39.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F39.GetValue(&fbeValue.F39); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F40.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F40.GetValue(&fbeValue.F40); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F41.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F41.GetValue(&fbeValue.F41); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F42.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F42.GetValue(&fbeValue.F42); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F43.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F43.GetValue(&fbeValue.F43); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F44.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F44.GetValue(&fbeValue.F44); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}

// Set the struct value
func (fm *FinalModelStructSimple) Set(fbeValue *StructSimple) (int, error) {
    fm.buffer.Shift(fm.FBEOffset())
    fbeResult, err := fm.SetFields(fbeValue)
    fm.buffer.Unshift(fm.FBEOffset())
    return fbeResult, err
}

// Set the struct fields values
func (fm *FinalModelStructSimple) SetFields(fbeValue *StructSimple) (int, error) {
    var err error = nil
    fbeCurrentSize := 0
    fbeCurrentOffset := 0
    fbeFieldSize := 0

    fm.Id.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.Id.Set(fbeValue.Id); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F1.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F1.Set(fbeValue.F1); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F2.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F2.Set(fbeValue.F2); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F3.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F3.Set(fbeValue.F3); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F4.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F4.Set(fbeValue.F4); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F5.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F5.Set(fbeValue.F5); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F6.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F6.Set(fbeValue.F6); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F7.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F7.Set(fbeValue.F7); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F8.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F8.Set(fbeValue.F8); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F9.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F9.Set(fbeValue.F9); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F10.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F10.Set(fbeValue.F10); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F11.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F11.Set(fbeValue.F11); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F12.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F12.Set(fbeValue.F12); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F13.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F13.Set(fbeValue.F13); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F14.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F14.Set(fbeValue.F14); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F15.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F15.Set(fbeValue.F15); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F16.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F16.Set(fbeValue.F16); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F17.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F17.Set(fbeValue.F17); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F18.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F18.Set(fbeValue.F18); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F19.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F19.Set(fbeValue.F19); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F20.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F20.Set(fbeValue.F20); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F21.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F21.Set(fbeValue.F21); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F22.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F22.Set(fbeValue.F22); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F23.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F23.Set(fbeValue.F23); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F24.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F24.Set(fbeValue.F24); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F25.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F25.Set(fbeValue.F25); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F26.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F26.Set(fbeValue.F26); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F27.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F27.Set(fbeValue.F27); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F28.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F28.Set(fbeValue.F28); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F29.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F29.Set(fbeValue.F29); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F30.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F30.Set(fbeValue.F30); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F31.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F31.Set(fbeValue.F31); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F32.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F32.Set(fbeValue.F32); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F33.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F33.Set(fbeValue.F33); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F34.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F34.Set(fbeValue.F34); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F35.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F35.Set(fbeValue.F35); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F36.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F36.Set(fbeValue.F36); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F37.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F37.Set(fbeValue.F37); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F38.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F38.Set(fbeValue.F38); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F39.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F39.Set(&fbeValue.F39); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F40.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F40.Set(&fbeValue.F40); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F41.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F41.Set(&fbeValue.F41); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F42.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F42.Set(&fbeValue.F42); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F43.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F43.Set(&fbeValue.F43); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    fm.F44.SetFBEOffset(fbeCurrentOffset)
    if fbeFieldSize, err = fm.F44.Set(&fbeValue.F44); err != nil {
        return fbeCurrentSize, err
    }
    fbeCurrentOffset += fbeFieldSize
    fbeCurrentSize += fbeFieldSize

    return fbeCurrentSize, err
}
