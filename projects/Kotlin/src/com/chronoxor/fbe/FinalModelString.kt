//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.fbe

// Fast Binary Encoding string final model
class FinalModelString(buffer: Buffer, offset: Long) : FinalModel(buffer, offset)
{
    // Get the allocation size
    fun fbeAllocationSize(value: String): Long = 4 + 3 * (value.length.toLong() + 1)

    // Check if the string value is valid
    override fun verify(): Long
    {
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
            return Long.MAX_VALUE

        val fbeStringSize = readUInt32(fbeOffset).toLong()
        if ((_buffer.offset + fbeOffset + 4 + fbeStringSize) > _buffer.size)
            return Long.MAX_VALUE

        return 4 + fbeStringSize
    }

    // Get the string value
    fun get(size: Size): String
    {
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
        {
            size.value = 0
            return ""
        }

        val fbeStringSize = readUInt32(fbeOffset).toLong()
        assert((_buffer.offset + fbeOffset + 4 + fbeStringSize) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + 4 + fbeStringSize) > _buffer.size)
        {
            size.value = 4
            return ""
        }

        size.value = 4 + fbeStringSize
        return readString(fbeOffset + 4, fbeStringSize)
    }

    // Set the string value
    fun set(value: String): Long
    {
        assert((_buffer.offset + fbeOffset + 4) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + 4) > _buffer.size)
            return 0

        val bytes = value.toByteArray(java.nio.charset.StandardCharsets.UTF_8)

        val fbeStringSize = bytes.size.toLong()
        assert((_buffer.offset + fbeOffset + 4 + fbeStringSize) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + 4 + fbeStringSize) > _buffer.size)
            return 4

        write(fbeOffset, fbeStringSize.toUInt())
        write(fbeOffset + 4, bytes)
        return 4 + fbeStringSize
    }
}
