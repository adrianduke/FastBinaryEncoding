//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.fbe

// Fast Binary Encoding buffer based on dynamic byte array
class Buffer
{
    // Is the buffer empty?
    val empty: Boolean
        get() = size == 0L
    // Get bytes memory buffer
    var data = ByteArray(0)
        private set
    // Get bytes memory buffer capacity
    val capacity: Long
        get() = data.size.toLong()
    // Get bytes memory buffer size
    var size: Long = 0
        private set
    // Get bytes memory buffer offset
    var offset: Long = 0
        private set

    // Initialize a new expandable buffer with zero capacity
    constructor()
    // Initialize a new expandable buffer with the given capacity
    constructor(capacity: Long) { attach(capacity) }
    // Initialize a new buffer based on the specified byte array
    constructor(buffer: ByteArray) { attach(buffer) }
    // Initialize a new buffer based on the specified region (offset) of a byte array
    constructor(buffer: ByteArray, offset: Long) { attach(buffer, offset) }
    // Initialize a new buffer based on the specified region (size and offset) of a byte array
    constructor(buffer: ByteArray, size: Long, offset: Long) { attach(buffer, size, offset) }

    // Attach memory buffer methods
    fun attach() { data = ByteArray(0); size = 0; offset = 0 }
    fun attach(capacity: Long) { data = ByteArray(capacity.toInt()); size = 0; offset = 0 }
    fun attach(buffer: ByteArray) { data = buffer; size = buffer.size.toLong(); offset = 0 }
    fun attach(buffer: ByteArray, offset: Long) { data = buffer; size = buffer.size.toLong(); this.offset = offset }
    fun attach(buffer: ByteArray, size: Long, offset: Long) { data = buffer; this.size = size; this.offset = offset }

    // Allocate memory in the current buffer and return offset to the allocated memory block
    fun allocate(size: Long): Long
    {
        assert(size >= 0) { "Invalid allocation size!" }
        if (size < 0)
            throw IllegalArgumentException("Invalid allocation size!")

        val offset = this.size

        // Calculate a new buffer size
        val total = this.size + size

        if (total <= data.size)
        {
            this.size = total
            return offset
        }

        val data = ByteArray(kotlin.math.max(total, 2L * this.data.size).toInt())
        System.arraycopy(this.data, 0, data, 0, this.size.toInt())
        this.data = data
        this.size = total
        return offset
    }

    // Remove some memory of the given size from the current buffer
    fun remove(offset: Long, size: Long)
    {
        assert((offset + size) <= this.size) { "Invalid offset & size!" }
        if ((offset + size) > this.size)
            throw IllegalArgumentException("Invalid offset & size!")

        System.arraycopy(data, (offset + size).toInt(), data, offset.toInt(), (this.size - size - offset).toInt())
        this.size -= size
        if (this.offset >= offset + size)
            this.offset -= size
        else if (this.offset >= offset)
        {
            this.offset -= this.offset - offset
            if (this.offset > this.size)
                this.offset = this.size
        }
    }

    // Reserve memory of the given capacity in the current buffer
    fun reserve(capacity: Long)
    {
        assert(capacity >= 0) { "Invalid reserve capacity!" }
        if (capacity < 0)
            throw IllegalArgumentException("Invalid reserve capacity!")

        if (capacity > data.size)
        {
            val data = ByteArray(kotlin.math.max(capacity, 2L * this.data.size).toInt())
            System.arraycopy(this.data, 0, data, 0, size.toInt())
            this.data = data
        }
    }

    // Resize the current buffer
    fun resize(size: Long)
    {
        reserve(size)
        this.size = size
        if (offset > this.size)
            offset = this.size
    }

    // Reset the current buffer and its offset
    fun reset()
    {
        size = 0
        offset = 0
    }

    // Shift the current buffer offset
    fun shift(offset: Long) { this.offset += offset }
    // Unshift the current buffer offset
    fun unshift(offset: Long) { this.offset -= offset }

    companion object
    {
        // Buffer I/O methods

        fun readBoolean(buffer: ByteArray, offset: Long): Boolean
        {
            val index = offset.toInt()
            return buffer[index].toInt() != 0
        }

        fun readByte(buffer: ByteArray, offset: Long): Byte
        {
            val index = offset.toInt()
            return buffer[index]
        }

        fun readChar(buffer: ByteArray, offset: Long): Char
        {
            return readInt8(buffer, offset).toInt().toChar()
        }

        fun readWChar(buffer: ByteArray, offset: Long): Char
        {
            return readInt32(buffer, offset).toChar()
        }

        fun readInt8(buffer: ByteArray, offset: Long): Byte
        {
            val index = offset.toInt()
            return buffer[index]
        }

        fun readUInt8(buffer: ByteArray, offset: Long): UByte
        {
            val index = offset.toInt()
            return buffer[index].toUByte()
        }

        fun readInt16(buffer: ByteArray, offset: Long): Short
        {
            val index = offset.toInt()
            return (((buffer[index + 0].toInt() and 0xFF) shl 0) or
                    ((buffer[index + 1].toInt() and 0xFF) shl 8)).toShort()
        }

        fun readUInt16(buffer: ByteArray, offset: Long): UShort
        {
            val index = offset.toInt()
            return (((buffer[index + 0].toUInt() and 0xFFu) shl 0) or
                    ((buffer[index + 1].toUInt() and 0xFFu) shl 8)).toUShort()
        }

        fun readInt32(buffer: ByteArray, offset: Long): Int
        {
            val index = offset.toInt()
            return ((buffer[index + 0].toInt() and 0xFF) shl  0) or
                   ((buffer[index + 1].toInt() and 0xFF) shl  8) or
                   ((buffer[index + 2].toInt() and 0xFF) shl 16) or
                   ((buffer[index + 3].toInt() and 0xFF) shl 24)
        }

        fun readUInt32(buffer: ByteArray, offset: Long): UInt
        {
            val index = offset.toInt()
            return ((buffer[index + 0].toUInt() and 0xFFu) shl  0) or
                   ((buffer[index + 1].toUInt() and 0xFFu) shl  8) or
                   ((buffer[index + 2].toUInt() and 0xFFu) shl 16) or
                   ((buffer[index + 3].toUInt() and 0xFFu) shl 24)
        }

        fun readInt64(buffer: ByteArray, offset: Long): Long
        {
            val index = offset.toInt()
            return ((buffer[index + 0].toLong() and 0xFF) shl  0) or
                   ((buffer[index + 1].toLong() and 0xFF) shl  8) or
                   ((buffer[index + 2].toLong() and 0xFF) shl 16) or
                   ((buffer[index + 3].toLong() and 0xFF) shl 24) or
                   ((buffer[index + 4].toLong() and 0xFF) shl 32) or
                   ((buffer[index + 5].toLong() and 0xFF) shl 40) or
                   ((buffer[index + 6].toLong() and 0xFF) shl 48) or
                   ((buffer[index + 7].toLong() and 0xFF) shl 56)
        }

        fun readUInt64(buffer: ByteArray, offset: Long): ULong
        {
            val index = offset.toInt()
            return ((buffer[index + 0].toULong() and 0xFFu) shl  0) or
                   ((buffer[index + 1].toULong() and 0xFFu) shl  8) or
                   ((buffer[index + 2].toULong() and 0xFFu) shl 16) or
                   ((buffer[index + 3].toULong() and 0xFFu) shl 24) or
                   ((buffer[index + 4].toULong() and 0xFFu) shl 32) or
                   ((buffer[index + 5].toULong() and 0xFFu) shl 40) or
                   ((buffer[index + 6].toULong() and 0xFFu) shl 48) or
                   ((buffer[index + 7].toULong() and 0xFFu) shl 56)
        }

        private fun readInt64BE(buffer: ByteArray, offset: Long): Long
        {
            val index = offset.toInt()
            return ((buffer[index + 0].toLong() and 0xFF) shl 56) or
                   ((buffer[index + 1].toLong() and 0xFF) shl 48) or
                   ((buffer[index + 2].toLong() and 0xFF) shl 40) or
                   ((buffer[index + 3].toLong() and 0xFF) shl 32) or
                   ((buffer[index + 4].toLong() and 0xFF) shl 24) or
                   ((buffer[index + 5].toLong() and 0xFF) shl 16) or
                   ((buffer[index + 6].toLong() and 0xFF) shl  8) or
                   ((buffer[index + 7].toLong() and 0xFF) shl  0)
        }

        fun readFloat(buffer: ByteArray, offset: Long): Float
        {
            val bits = readInt32(buffer, offset)
            return java.lang.Float.intBitsToFloat(bits)
        }

        fun readDouble(buffer: ByteArray, offset: Long): Double
        {
            val bits = readInt64(buffer, offset)
            return java.lang.Double.longBitsToDouble(bits)
        }

        fun readBytes(buffer: ByteArray, offset: Long, size: Long): ByteArray
        {
            val result = ByteArray(size.toInt())
            System.arraycopy(buffer, offset.toInt(), result, 0, size.toInt())
            return result
        }

        fun readString(buffer: ByteArray, offset: Long, size: Long): String
        {
            return String(buffer, offset.toInt(), size.toInt(), java.nio.charset.StandardCharsets.UTF_8)
        }

        fun readUUID(buffer: ByteArray, offset: Long): java.util.UUID
        {
            return java.util.UUID(readInt64BE(buffer, offset), readInt64BE(buffer, offset + 8))
        }

        fun write(buffer: ByteArray, offset: Long, value: Boolean)
        {
            buffer[offset.toInt()] = (if (value) 1 else 0).toByte()
        }

        fun write(buffer: ByteArray, offset: Long, value: Byte)
        {
            buffer[offset.toInt()] = value
        }

        fun write(buffer: ByteArray, offset: Long, value: UByte)
        {
            buffer[offset.toInt()] = value.toByte()
        }

        fun write(buffer: ByteArray, offset: Long, value: Short)
        {
            buffer[offset.toInt() + 0] = (value.toInt() shr 0).toByte()
            buffer[offset.toInt() + 1] = (value.toInt() shr 8).toByte()
        }

        fun write(buffer: ByteArray, offset: Long, value: UShort)
        {
            buffer[offset.toInt() + 0] = (value.toUInt() shr 0).toByte()
            buffer[offset.toInt() + 1] = (value.toUInt() shr 8).toByte()
        }

        fun write(buffer: ByteArray, offset: Long, value: Int)
        {
            buffer[offset.toInt() + 0] = (value shr  0).toByte()
            buffer[offset.toInt() + 1] = (value shr  8).toByte()
            buffer[offset.toInt() + 2] = (value shr 16).toByte()
            buffer[offset.toInt() + 3] = (value shr 24).toByte()
        }

        fun write(buffer: ByteArray, offset: Long, value: UInt)
        {
            buffer[offset.toInt() + 0] = (value shr  0).toByte()
            buffer[offset.toInt() + 1] = (value shr  8).toByte()
            buffer[offset.toInt() + 2] = (value shr 16).toByte()
            buffer[offset.toInt() + 3] = (value shr 24).toByte()
        }

        fun write(buffer: ByteArray, offset: Long, value: Long)
        {
            buffer[offset.toInt() + 0] = (value shr  0).toByte()
            buffer[offset.toInt() + 1] = (value shr  8).toByte()
            buffer[offset.toInt() + 2] = (value shr 16).toByte()
            buffer[offset.toInt() + 3] = (value shr 24).toByte()
            buffer[offset.toInt() + 4] = (value shr 32).toByte()
            buffer[offset.toInt() + 5] = (value shr 40).toByte()
            buffer[offset.toInt() + 6] = (value shr 48).toByte()
            buffer[offset.toInt() + 7] = (value shr 56).toByte()
        }

        fun write(buffer: ByteArray, offset: Long, value: ULong)
        {
            buffer[offset.toInt() + 0] = (value shr  0).toByte()
            buffer[offset.toInt() + 1] = (value shr  8).toByte()
            buffer[offset.toInt() + 2] = (value shr 16).toByte()
            buffer[offset.toInt() + 3] = (value shr 24).toByte()
            buffer[offset.toInt() + 4] = (value shr 32).toByte()
            buffer[offset.toInt() + 5] = (value shr 40).toByte()
            buffer[offset.toInt() + 6] = (value shr 48).toByte()
            buffer[offset.toInt() + 7] = (value shr 56).toByte()
        }

        private fun writeBE(buffer: ByteArray, offset: Long, value: Long)
        {
            buffer[offset.toInt() + 0] = (value shr 56).toByte()
            buffer[offset.toInt() + 1] = (value shr 48).toByte()
            buffer[offset.toInt() + 2] = (value shr 40).toByte()
            buffer[offset.toInt() + 3] = (value shr 32).toByte()
            buffer[offset.toInt() + 4] = (value shr 24).toByte()
            buffer[offset.toInt() + 5] = (value shr 16).toByte()
            buffer[offset.toInt() + 6] = (value shr  8).toByte()
            buffer[offset.toInt() + 7] = (value shr  0).toByte()
        }

        fun write(buffer: ByteArray, offset: Long, value: Float)
        {
            val bits = java.lang.Float.floatToIntBits(value)
            write(buffer, offset, bits)
        }

        fun write(buffer: ByteArray, offset: Long, value: Double)
        {
            val bits = java.lang.Double.doubleToLongBits(value)
            write(buffer, offset, bits)
        }

        fun write(buffer: ByteArray, offset: Long, value: ByteArray)
        {
            System.arraycopy(value, 0, buffer, offset.toInt(), value.size)
        }

        fun write(buffer: ByteArray, offset: Long, value: ByteArray, valueOffset: Long, valueSize: Long)
        {
            System.arraycopy(value, valueOffset.toInt(), buffer, offset.toInt(), valueSize.toInt())
        }

        fun write(buffer: ByteArray, offset: Long, value: Byte, valueCount: Long)
        {
            for (i in 0 until valueCount)
                buffer[(offset + i).toInt()] = value
        }

        fun write(buffer: ByteArray, offset: Long, value: java.util.UUID)
        {
            writeBE(buffer, offset, value.mostSignificantBits)
            writeBE(buffer, offset + 8, value.leastSignificantBits)
        }
    }
}
