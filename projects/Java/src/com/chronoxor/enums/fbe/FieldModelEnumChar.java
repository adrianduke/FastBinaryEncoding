//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.enums.fbe;

// Fast Binary Encoding EnumChar field model
public final class FieldModelEnumChar extends com.chronoxor.fbe.FieldModel
{
    public FieldModelEnumChar(com.chronoxor.fbe.Buffer buffer, long offset) { super(buffer, offset); }

    // Get the field size
    @Override
    public long fbeSize() { return 1; }

    // Get the value
    public com.chronoxor.enums.EnumChar get() { return get(new com.chronoxor.enums.EnumChar()); }
    public com.chronoxor.enums.EnumChar get(com.chronoxor.enums.EnumChar defaults)
    {
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return defaults;

        return new com.chronoxor.enums.EnumChar(readInt8(fbeOffset()));
    }

    // Set the value
    public void set(com.chronoxor.enums.EnumChar value)
    {
        assert ((_buffer.getOffset() + fbeOffset() + fbeSize()) <= _buffer.getSize()) : "Model is broken!";
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return;

        write(fbeOffset(), value.getRaw());
    }
}
