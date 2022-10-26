//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.fbe;

// Fast Binary Encoding date field model
public final class FieldModelDate extends FieldModel
{
    public FieldModelDate(Buffer buffer, long offset) { super(buffer, offset); }

    // Get the field size
    @Override
    public long fbeSize() { return 8; }

    // Get the date value
    public java.util.Date get() { return get(new java.util.Date(0)); }
    public java.util.Date get(java.util.Date defaults)
    {
        assert (defaults != null) : "Invalid default date value!";
        if (defaults == null)
            throw new IllegalArgumentException("Invalid default date value!");

        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return defaults;

        long nanoseconds = readInt64(fbeOffset());
        return new java.util.Date(nanoseconds / 1000000);
    }

    // Set the date value
    public void set(java.util.Date value)
    {
        assert (value != null) : "Invalid date value!";
        if (value == null)
            throw new IllegalArgumentException("Invalid date value!");

        assert ((_buffer.getOffset() + fbeOffset() + fbeSize()) <= _buffer.getSize()) : "Model is broken!";
        if ((_buffer.getOffset() + fbeOffset() + fbeSize()) > _buffer.getSize())
            return;

        long nanoseconds = value.getTime() * 1000;
        write(fbeOffset(), nanoseconds);
    }
}
