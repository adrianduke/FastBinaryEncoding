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

// Fast Binary Encoding StructEmpty field model
type FieldModelStructEmpty struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int
}

// Create a new StructEmpty field model
func NewFieldModelStructEmpty(buffer *fbe.Buffer, offset int) *FieldModelStructEmpty {
    fbeResult := FieldModelStructEmpty{buffer: buffer, offset: offset}
    return &fbeResult
}

// Get the field size
func (fm *FieldModelStructEmpty) FBESize() int { return 4 }

// Get the field body size
func (fm *FieldModelStructEmpty) FBEBody() int {
    fbeResult := 4 + 4 +
        0
    return fbeResult
}

// Get the field extra size
func (fm *FieldModelStructEmpty) FBEExtra() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4) > fm.buffer.Size()) {
        return 0
    }

    fm.buffer.Shift(fbeStructOffset)

    fbeResult := fm.FBEBody() +
        0

    fm.buffer.Unshift(fbeStructOffset)

    return fbeResult
}

// Get the field type
func (fm *FieldModelStructEmpty) FBEType() int { return 143 }

// Get the field offset
func (fm *FieldModelStructEmpty) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelStructEmpty) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelStructEmpty) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelStructEmpty) FBEUnshift(size int) { fm.offset -= size }

// Check if the struct value is valid
func (fm *FieldModelStructEmpty) Verify() bool { return fm.VerifyType(true) }

// Check if the struct value and its type are valid
func (fm *FieldModelStructEmpty) VerifyType(fbeVerifyType bool) bool {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return true
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return false
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return false
    }

    fbeStructType := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4))
    if fbeVerifyType && (fbeStructType != fm.FBEType()) {
        return false
    }

    fm.buffer.Shift(fbeStructOffset)
    fbeResult := fm.VerifyFields(fbeStructSize)
    fm.buffer.Unshift(fbeStructOffset)
    return fbeResult
}

// // Check if the struct value fields are valid
func (fm *FieldModelStructEmpty) VerifyFields(fbeStructSize int) bool {
    return true
}

// Get the struct value (begin phase)
func (fm *FieldModelStructEmpty) GetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, nil
    }

    fbeStructOffset := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    if (fbeStructOffset == 0) || ((fm.buffer.Offset() + fbeStructOffset + 4 + 4) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset))
    if fbeStructSize < (4 + 4) {
        return 0, errors.New("model is broken")
    }

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Get the struct value (end phase)
func (fm *FieldModelStructEmpty) GetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Get the struct value
func (fm *FieldModelStructEmpty) Get() (*StructEmpty, error) {
    fbeResult := NewStructEmpty()
    return fbeResult, fm.GetValue(fbeResult)
}

// Get the struct value by the given pointer
func (fm *FieldModelStructEmpty) GetValue(fbeValue *StructEmpty) error {
    fbeBegin, err := fm.GetBegin()
    if fbeBegin == 0 {
        return err
    }

    fbeStructSize := int(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset()))
    fm.GetFields(fbeValue, fbeStructSize)
    fm.GetEnd(fbeBegin)
    return nil
}

// Get the struct fields values
func (fm *FieldModelStructEmpty) GetFields(fbeValue *StructEmpty, fbeStructSize int) {
}

// Set the struct value (begin phase)
func (fm *FieldModelStructEmpty) SetBegin() (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbeStructSize := fm.FBEBody()
    fbeStructOffset := fm.buffer.Allocate(fbeStructSize) - fm.buffer.Offset()
    if (fbeStructOffset <= 0) || ((fm.buffer.Offset() + fbeStructOffset + fbeStructSize) > fm.buffer.Size()) {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(fbeStructOffset))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset, uint32(fbeStructSize))
    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fbeStructOffset + 4, uint32(fm.FBEType()))

    fm.buffer.Shift(fbeStructOffset)
    return fbeStructOffset, nil
}

// Set the struct value (end phase)
func (fm *FieldModelStructEmpty) SetEnd(fbeBegin int) {
    fm.buffer.Unshift(fbeBegin)
}

// Set the struct value
func (fm *FieldModelStructEmpty) Set(fbeValue *StructEmpty) error {
    fbeBegin, err := fm.SetBegin()
    if fbeBegin == 0 {
        return err
    }

    err = fm.SetFields(fbeValue)
    fm.SetEnd(fbeBegin)
    return err
}

// Set the struct fields values
func (fm *FieldModelStructEmpty) SetFields(fbeValue *StructEmpty) error {
    return nil
}
