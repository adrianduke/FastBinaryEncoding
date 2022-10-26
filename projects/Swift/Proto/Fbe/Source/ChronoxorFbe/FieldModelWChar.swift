//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

import Foundation

// Fast Binary Encoding Character field model
public class FieldModelWChar: FieldModel {
    public var _buffer = Buffer()
    public var _offset: Int = 0

    // Field size
    public let fbeSize: Int = 4

    public required init() {
        _buffer = Buffer()
        _offset = 0
    }

    public required init(buffer: Buffer, offset: Int) {
        _buffer = buffer
        _offset = offset
    }

    // Get the value
    public func get(defaults: Character = Character("0")) -> Character {
        if (_buffer.offset + fbeOffset + fbeSize) > _buffer.size {
            return defaults
        }

        return readWChar(offset: fbeOffset)
    }

    // Set the value
    public func set(value: Character) throws {
        if (_buffer.offset + fbeOffset + fbeSize) > _buffer.size {
            assertionFailure("Model is broken!")
            return
        }

        write(offset: fbeOffset, value: value.utf16.map { UInt32($0) }[0])
    }
}
