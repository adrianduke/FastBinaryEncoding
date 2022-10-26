//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.fbe

// Fast Binary Encoding base model
open class Model
{
    // Get bytes buffer
    var buffer = Buffer()
        private set

    // Initialize a new model
    protected constructor()
    protected constructor(buffer: Buffer) { this.buffer = buffer }

    // Attach memory buffer methods
    fun attach() { buffer.attach() }
    fun attach(capacity: Long) { buffer.attach(capacity) }
    fun attach(buffer: ByteArray) { this.buffer.attach(buffer) }
    fun attach(buffer: ByteArray, offset: Long) { this.buffer.attach(buffer, offset) }
    fun attach(buffer: ByteArray, size: Long, offset: Long) { this.buffer.attach(buffer, size, offset) }
    fun attach(buffer: Buffer) { this.buffer.attach(buffer.data, buffer.size, buffer.offset) }
    fun attach(buffer: Buffer, offset: Long) { this.buffer.attach(buffer.data, buffer.size, offset) }

    // Model buffer operations
    fun allocate(size: Long): Long { return buffer.allocate(size) }
    fun remove(offset: Long, size: Long) { buffer.remove(offset, size) }
    fun reserve(capacity: Long) { buffer.reserve(capacity) }
    fun resize(size: Long) { buffer.resize(size) }
    fun reset() { buffer.reset() }
    fun shift(offset: Long) { buffer.shift(offset) }
    fun unshift(offset: Long) { buffer.unshift(offset) }

    // Buffer I/O methods
    protected fun readUInt32(offset: Long): UInt { return Buffer.readUInt32(buffer.data, buffer.offset + offset) }
    protected fun write(offset: Long, value: UInt) { Buffer.write(buffer.data, buffer.offset + offset, value) }
}
