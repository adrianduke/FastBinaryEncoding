//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

package com.chronoxor.test.fbe;

// Fast Binary Encoding com.chronoxor.test proxy
public class Proxy extends com.chronoxor.fbe.Receiver
{
    // Imported proxy
    public com.chronoxor.proto.fbe.Proxy protoProxy;

    // Proxy models accessors

    public Proxy()
    {
        super(false);
        protoProxy = new com.chronoxor.proto.fbe.Proxy(getBuffer());
    }
    public Proxy(com.chronoxor.fbe.Buffer buffer)
    {
        super(buffer, false);
        protoProxy = new com.chronoxor.proto.fbe.Proxy(getBuffer());
    }

    // Proxy handlers

    @Override
    public boolean onReceive(long type, byte[] buffer, long offset, long size)
    {
        switch ((int)type)
        {
            default: break;
        }

        if ((protoProxy != null) && protoProxy.onReceive(type, buffer, offset, size))
            return true;

        return false;
    }
}
