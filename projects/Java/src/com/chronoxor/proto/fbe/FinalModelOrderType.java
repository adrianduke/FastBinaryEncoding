//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.proto.fbe;

// Fast Binary Encoding OrderType final model
public final class FinalModelOrderType extends com.chronoxor.fbe.FinalModel
{
    public FinalModelOrderType(com.chronoxor.fbe.Buffer buffer, long offset) { super(buffer, offset); }

    // Get the allocation size
    public long fbeAllocationSize(com.chronoxor.proto.OrderType value) { return fbeSize(); }

    // Get the final size
    @Override
    public long fbeSize() { return 1; }

    // Check if the value is valid
    @Override
    public long verify()
    {
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return Long.MAX_VALUE;

        return fbeSize();
    }

    // Get the value
    public com.chronoxor.proto.OrderType get(com.chronoxor.fbe.Size size)
    {
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return new com.chronoxor.proto.OrderType();

        size.value = fbeSize();
        return new com.chronoxor.proto.OrderType(readByte(fbeOffset()));
    }

    // Set the value
    public long set(com.chronoxor.proto.OrderType value)
    {
        assert ((_buffer.getOffset() + fbeOffset() + fbeSize()) <= _buffer.getSize()) : "Model is broken!";
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return 0;

        write(fbeOffset(), value.getRaw());
        return fbeSize();
    }
}
