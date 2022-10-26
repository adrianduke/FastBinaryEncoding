//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.fbe;

// Fast Binary Encoding java.util.UUID field model
public final class FieldModelUUID extends FieldModel
{
    public FieldModelUUID(Buffer buffer, long offset) { super(buffer, offset); }

    // Get the field size
    @Override
    public long fbeSize() { return 16; }

    // Get the value
    public java.util.UUID get() { return get(UUIDGenerator.nil()); }
    public java.util.UUID get(java.util.UUID defaults)
    {
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return defaults;

        return readUUID(fbeOffset());
    }

    // Set the value
    public void set(java.util.UUID value)
    {
        assert ((_buffer.getOffset() + fbeOffset() + fbeSize()) <= _buffer.getSize()) : "Model is broken!";
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return;

        write(fbeOffset(), value);
    }
}
