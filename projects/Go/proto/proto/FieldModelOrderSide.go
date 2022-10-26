//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package proto

import "errors"
import "../fbe"

// Fast Binary Encoding OrderSide field model
type FieldModelOrderSide struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int
}

// Create a new OrderSide field model
func NewFieldModelOrderSide(buffer *fbe.Buffer, offset int) *FieldModelOrderSide {
    return &FieldModelOrderSide{buffer: buffer, offset: offset}
}

// Get the field size
func (fm *FieldModelOrderSide) FBESize() int { return 1 }
// Get the field extra size
func (fm *FieldModelOrderSide) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelOrderSide) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelOrderSide) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelOrderSide) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelOrderSide) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FieldModelOrderSide) Verify() bool { return true }

// Get the value
func (fm *FieldModelOrderSide) Get() (*OrderSide, error) {
    var value OrderSide
    return &value, fm.GetValueDefault(&value, OrderSide(0))
}

// Get the value with provided default value
func (fm *FieldModelOrderSide) GetDefault(defaults OrderSide) (*OrderSide, error) {
    var value OrderSide
    err := fm.GetValueDefault(&value, defaults)
    return &value, err
}

// Get the value by the given pointer
func (fm *FieldModelOrderSide) GetValue(value *OrderSide) error {
    return fm.GetValueDefault(value, OrderSide(0))
}

// Get the value by the given pointer with provided default value
func (fm *FieldModelOrderSide) GetValueDefault(value *OrderSide, defaults OrderSide) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return nil
    }

    *value = OrderSide(fbe.ReadByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return nil
}

// Set the value by the given pointer
func (fm *FieldModelOrderSide) Set(value *OrderSide) error {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FieldModelOrderSide) SetValue(value OrderSide) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), byte(value))
    return nil
}
