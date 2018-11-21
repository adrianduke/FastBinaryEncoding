// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

// Fast Binary Encoding rune final model class
type FinalModelChar struct {
    buffer Buffer // Final model buffer
    offset int    // Final model buffer offset
}

// Get the allocation size
func (fm FinalModelChar) FBEAllocationSize() int { return fm.FBESize() }

// Get the final size
func (fm FinalModelChar) FBESize() int { return 1 }
// Get the final extra size
func (fm FinalModelChar) FBEExtra() int { return 0 }

// Get the final offset
func (fm FinalModelChar) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelChar) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelChar) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelChar) FBEUnshift(size int) { fm.offset -= size }

func NewFinalModelChar(buffer Buffer, offset int) *FinalModelChar {
    return &FinalModelChar{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm FinalModelChar) Verify() int {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm FinalModelChar) Get() (rune, int) {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return '\000', 0
    }

    return ReadChar(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), fm.FBESize()
}

// Set the value
func (fm *FinalModelChar) Set(value rune) int {
    if fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize() > fm.buffer.Size() {
        return 0
    }

    WriteChar(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return fm.FBESize()
}
