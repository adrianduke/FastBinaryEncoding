//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumUInt8 final model
type FinalModelEnumUInt8 struct {
    // Final model buffer
    buffer *fbe.Buffer
    // Final model buffer offset
    offset int
}

// Create a new EnumUInt8 final model
func NewFinalModelEnumUInt8(buffer *fbe.Buffer, offset int) *FinalModelEnumUInt8 {
    return &FinalModelEnumUInt8{buffer: buffer, offset: offset}
}

// Get the allocation size
func (fm *FinalModelEnumUInt8) FBEAllocationSize(value *EnumUInt8) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelEnumUInt8) FBESize() int { return 1 }

// Get the final offset
func (fm *FinalModelEnumUInt8) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumUInt8) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumUInt8) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumUInt8) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FinalModelEnumUInt8) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return fbe.MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm *FinalModelEnumUInt8) Get() (*EnumUInt8, int, error) {
    var value EnumUInt8
    size, err := fm.GetValueDefault(&value, EnumUInt8(0))
    return &value, size, err
}

// Get the value with provided default value
func (fm *FinalModelEnumUInt8) GetDefault(defaults EnumUInt8) (*EnumUInt8, int, error) {
    var value EnumUInt8
    size, err := fm.GetValueDefault(&value, defaults)
    return &value, size, err
}

// Get the value by the given pointer
func (fm *FinalModelEnumUInt8) GetValue(value *EnumUInt8) (int, error) {
    return fm.GetValueDefault(value, EnumUInt8(0))
}

// Get the value by the given pointer with provided default value
func (fm *FinalModelEnumUInt8) GetValueDefault(value *EnumUInt8, defaults EnumUInt8) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return 0, errors.New("model is broken")
    }

    *value = EnumUInt8(fbe.ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return fm.FBESize(), nil
}

// Set the value by the given pointer
func (fm *FinalModelEnumUInt8) Set(value *EnumUInt8) (int, error) {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FinalModelEnumUInt8) SetValue(value EnumUInt8) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint8(value))
    return fm.FBESize(), nil
}
