//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.proto.fbe;

// Fast Binary Encoding Balance final model
public final class FinalModelBalance extends com.chronoxor.fbe.FinalModel
{
    public final com.chronoxor.fbe.FinalModelString currency;
    public final com.chronoxor.fbe.FinalModelDouble amount;

    public FinalModelBalance(com.chronoxor.fbe.Buffer buffer, long offset)
    {
        super(buffer, offset);
        currency = new com.chronoxor.fbe.FinalModelString(buffer, 0);
        amount = new com.chronoxor.fbe.FinalModelDouble(buffer, 0);
    }

    // Get the allocation size
    public long fbeAllocationSize(com.chronoxor.proto.Balance fbeValue)
    {
        long fbeResult = 0
            + currency.fbeAllocationSize(fbeValue.currency)
            + amount.fbeAllocationSize(fbeValue.amount)
            ;
        return fbeResult;
    }

    // Get the final type
    public static final long fbeTypeConst = 2;
    public long fbeType() { return fbeTypeConst; }

    // Check if the struct value is valid
    @Override
    public long verify()
    {
        _buffer.shift(fbeOffset());
        long fbeResult = verifyFields();
        _buffer.unshift(fbeOffset());
        return fbeResult;
    }

    // Check if the struct fields are valid
    public long verifyFields()
    {
        long fbeCurrentOffset = 0;
        long fbeFieldSize = 0;

        currency.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = currency.verify();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        amount.fbeOffset(fbeCurrentOffset);
        fbeFieldSize = amount.verify();
        if (fbeFieldSize == Long.MAX_VALUE)
            return Long.MAX_VALUE;
        fbeCurrentOffset += fbeFieldSize;

        return fbeCurrentOffset;
    }

    // Get the struct value
    public com.chronoxor.proto.Balance get(com.chronoxor.fbe.Size fbeSize) { return get(fbeSize, new com.chronoxor.proto.Balance()); }
    public com.chronoxor.proto.Balance get(com.chronoxor.fbe.Size fbeSize, com.chronoxor.proto.Balance fbeValue)
    {
        _buffer.shift(fbeOffset());
        fbeSize.value = getFields(fbeValue);
        _buffer.unshift(fbeOffset());
        return fbeValue;
    }

    // Get the struct fields values
    public long getFields(com.chronoxor.proto.Balance fbeValue)
    {
        long fbeCurrentOffset = 0;
        long fbeCurrentSize = 0;
        var fbeFieldSize = new com.chronoxor.fbe.Size(0);

        currency.fbeOffset(fbeCurrentOffset);
        fbeValue.currency = currency.get(fbeFieldSize);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        amount.fbeOffset(fbeCurrentOffset);
        fbeValue.amount = amount.get(fbeFieldSize);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        return fbeCurrentSize;
    }

    // Set the struct value
    public long set(com.chronoxor.proto.Balance fbeValue)
    {
        _buffer.shift(fbeOffset());
        long fbeSize = setFields(fbeValue);
        _buffer.unshift(fbeOffset());
        return fbeSize;
    }

    // Set the struct fields values
    public long setFields(com.chronoxor.proto.Balance fbeValue)
    {
        long fbeCurrentOffset = 0;
        long fbeCurrentSize = 0;
        var fbeFieldSize = new com.chronoxor.fbe.Size();

        currency.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = currency.set(fbeValue.currency);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        amount.fbeOffset(fbeCurrentOffset);
        fbeFieldSize.value = amount.set(fbeValue.amount);
        fbeCurrentOffset += fbeFieldSize.value;
        fbeCurrentSize += fbeFieldSize.value;

        return fbeCurrentSize;
    }
}
