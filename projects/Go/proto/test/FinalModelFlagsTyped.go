//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package test

import "errors"
import "../fbe"

// Fast Binary Encoding FlagsTyped final model
type FinalModelFlagsTyped struct {
    // Final model buffer
    buffer *fbe.Buffer
    // Final model buffer offset
    offset int
}

// Create a new FlagsTyped final model
func NewFinalModelFlagsTyped(buffer *fbe.Buffer, offset int) *FinalModelFlagsTyped {
    return &FinalModelFlagsTyped{buffer: buffer, offset: offset}
}

// Get the allocation size
func (fm *FinalModelFlagsTyped) FBEAllocationSize(value *FlagsTyped) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelFlagsTyped) FBESize() int { return 8 }

// Get the final offset
func (fm *FinalModelFlagsTyped) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelFlagsTyped) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelFlagsTyped) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelFlagsTyped) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FinalModelFlagsTyped) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return fbe.MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm *FinalModelFlagsTyped) Get() (*FlagsTyped, int, error) {
    var value FlagsTyped
    size, err := fm.GetValueDefault(&value, FlagsTyped(0))
    return &value, size, err
}

// Get the value with provided default value
func (fm *FinalModelFlagsTyped) GetDefault(defaults FlagsTyped) (*FlagsTyped, int, error) {
    var value FlagsTyped
    size, err := fm.GetValueDefault(&value, defaults)
    return &value, size, err
}

// Get the value by the given pointer
func (fm *FinalModelFlagsTyped) GetValue(value *FlagsTyped) (int, error) {
    return fm.GetValueDefault(value, FlagsTyped(0))
}

// Get the value by the given pointer with provided default value
func (fm *FinalModelFlagsTyped) GetValueDefault(value *FlagsTyped, defaults FlagsTyped) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return 0, errors.New("model is broken")
    }

    *value = FlagsTyped(fbe.ReadUInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return fm.FBESize(), nil
}

// Set the value by the given pointer
func (fm *FinalModelFlagsTyped) Set(value *FlagsTyped) (int, error) {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FinalModelFlagsTyped) SetValue(value FlagsTyped) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint64(value))
    return fm.FBESize(), nil
}
