//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package fbe

import "errors"

// Fast Binary Encoding UInt32 field model
type FieldModelUInt32 struct {
    // Field model buffer
    buffer *Buffer
    // Field model buffer offset
    offset int
}

// Create a new UInt32 field model
func NewFieldModelUInt32(buffer *Buffer, offset int) *FieldModelUInt32 {
    return &FieldModelUInt32{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelUInt32) FBESize() int { return 4 }
// Get the field extra size
func (fm *FieldModelUInt32) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelUInt32) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelUInt32) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelUInt32) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelUInt32) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelUInt32) Verify() bool { return true }

// Get the value
func (fm *FieldModelUInt32) Get() (uint32, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm *FieldModelUInt32) GetDefault(defaults uint32) (uint32, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return defaults, nil
    }

    return ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), nil
}

// Set the value
func (fm *FieldModelUInt32) Set(value uint32) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return nil
}
