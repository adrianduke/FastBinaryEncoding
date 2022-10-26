//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package test

import "errors"
import "../fbe"

// Fast Binary Encoding EnumEmpty field model
type FieldModelEnumEmpty struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int
}

// Create a new EnumEmpty field model
func NewFieldModelEnumEmpty(buffer *fbe.Buffer, offset int) *FieldModelEnumEmpty {
    return &FieldModelEnumEmpty{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelEnumEmpty) FBESize() int { return 4 }
// Get the field extra size
func (fm *FieldModelEnumEmpty) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelEnumEmpty) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumEmpty) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumEmpty) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumEmpty) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelEnumEmpty) Verify() bool { return true }

// Get the value
func (fm *FieldModelEnumEmpty) Get() (*EnumEmpty, error) {
    var value EnumEmpty
    return &value, fm.GetValueDefault(&value, EnumEmpty(0))
}

// Get the value with provided default value
func (fm *FieldModelEnumEmpty) GetDefault(defaults EnumEmpty) (*EnumEmpty, error) {
    var value EnumEmpty
    err := fm.GetValueDefault(&value, defaults)
    return &value, err
}

// Get the value by the given pointer
func (fm *FieldModelEnumEmpty) GetValue(value *EnumEmpty) error {
    return fm.GetValueDefault(value, EnumEmpty(0))
}

// Get the value by the given pointer with provided default value
func (fm *FieldModelEnumEmpty) GetValueDefault(value *EnumEmpty, defaults EnumEmpty) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return nil
    }

    *value = EnumEmpty(fbe.ReadInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return nil
}

// Set the value by the given pointer
func (fm *FieldModelEnumEmpty) Set(value *EnumEmpty) error {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FieldModelEnumEmpty) SetValue(value EnumEmpty) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int32(value))
    return nil
}
