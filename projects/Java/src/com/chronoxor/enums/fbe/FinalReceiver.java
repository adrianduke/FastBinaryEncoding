//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.enums.fbe;

// Fast Binary Encoding com.chronoxor.enums final receiver
public class FinalReceiver extends com.chronoxor.fbe.Receiver
{
    // Receiver values accessors

    // Receiver models accessors

    public FinalReceiver()
    {
        super(true);
    }
    public FinalReceiver(com.chronoxor.fbe.Buffer buffer)
    {
        super(buffer, true);
    }

    // Receive handlers

    @Override
    public boolean onReceive(long type, byte[] buffer, long offset, long size)
    {
        switch ((int)type)
        {
            default: break;
        }

        return false;
    }
}
