//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

import ChronoxorFbe
import ChronoxorProto

// Fast Binary Encoding Order final model
public class FinalModelOrder: FinalModel {
    public var _buffer: Buffer
    public var _offset: Int

    let id: ChronoxorFbe.FinalModelInt32
    let symbol: ChronoxorFbe.FinalModelString
    let side: FinalModelOrderSide
    let type: FinalModelOrderType
    let price: ChronoxorFbe.FinalModelDouble
    let volume: ChronoxorFbe.FinalModelDouble
    let tp: ChronoxorFbe.FinalModelDouble
    let sl: ChronoxorFbe.FinalModelDouble

    // Field type
    public var fbeType: Int = fbeTypeConst

    public static let fbeTypeConst: Int = 1

    public required init(buffer: Buffer, offset: Int) {
        _buffer = buffer
        _offset = offset

        id = FinalModelInt32(buffer: buffer, offset: 0)
        symbol = FinalModelString(buffer: buffer, offset: 0)
        side = FinalModelOrderSide(buffer: buffer, offset: 0)
        type = FinalModelOrderType(buffer: buffer, offset: 0)
        price = FinalModelDouble(buffer: buffer, offset: 0)
        volume = FinalModelDouble(buffer: buffer, offset: 0)
        tp = FinalModelDouble(buffer: buffer, offset: 0)
        sl = FinalModelDouble(buffer: buffer, offset: 0)
    }

    // Get the allocation size
    public func fbeAllocationSize(value fbeValue: Order) -> Int {
        return 0
            + id.fbeAllocationSize(value: fbeValue.id)
            + symbol.fbeAllocationSize(value: fbeValue.symbol)
            + side.fbeAllocationSize(value: fbeValue.side)
            + type.fbeAllocationSize(value: fbeValue.type)
            + price.fbeAllocationSize(value: fbeValue.price)
            + volume.fbeAllocationSize(value: fbeValue.volume)
            + tp.fbeAllocationSize(value: fbeValue.tp)
            + sl.fbeAllocationSize(value: fbeValue.sl)
    }

    // Check if the struct value is valid
    public func verify() -> Int {
        _buffer.shift(offset: fbeOffset)
        let fbeResult = verifyFields()
        _buffer.unshift(offset: fbeOffset)
        return fbeResult
    }

    // Check if the struct fields are valid
    public func verifyFields() -> Int {
        var fbeCurrentOffset: Int = 0
        var fbeFieldSize: Int = 0

        id.fbeOffset = fbeCurrentOffset
        fbeFieldSize = id.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        symbol.fbeOffset = fbeCurrentOffset
        fbeFieldSize = symbol.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        side.fbeOffset = fbeCurrentOffset
        fbeFieldSize = side.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        type.fbeOffset = fbeCurrentOffset
        fbeFieldSize = type.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        price.fbeOffset = fbeCurrentOffset
        fbeFieldSize = price.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        volume.fbeOffset = fbeCurrentOffset
        fbeFieldSize = volume.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        tp.fbeOffset = fbeCurrentOffset
        fbeFieldSize = tp.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        sl.fbeOffset = fbeCurrentOffset
        fbeFieldSize = sl.verify()
        if fbeFieldSize == Int.max {
            return Int.max
        }
        fbeCurrentOffset += fbeFieldSize

        return fbeCurrentOffset
    }

    // Get the struct value
    public func get(size: inout Size) -> Order {
        var fbeValue = Order()
        return get(size: &size, fbeValue: &fbeValue)
    }

    // Get the struct value
    public func get(size: inout Size, fbeValue: inout Order) -> Order {
        _buffer.shift(offset: fbeOffset)
        size.value = getFields(fbeValue: &fbeValue)
        _buffer.unshift(offset: fbeOffset)
        return fbeValue
    }

    // Get the struct fields values
    public func getFields(fbeValue: inout Order) -> Int {
        var fbeCurrentOffset: Int = 0
        var fbeCurrentSize: Int = 0
        var fbeFieldSize = Size()

        id.fbeOffset = fbeCurrentOffset
        fbeValue.id = id.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        symbol.fbeOffset = fbeCurrentOffset
        fbeValue.symbol = symbol.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        side.fbeOffset = fbeCurrentOffset
        fbeValue.side = side.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        type.fbeOffset = fbeCurrentOffset
        fbeValue.type = type.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        price.fbeOffset = fbeCurrentOffset
        fbeValue.price = price.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        volume.fbeOffset = fbeCurrentOffset
        fbeValue.volume = volume.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        tp.fbeOffset = fbeCurrentOffset
        fbeValue.tp = tp.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        sl.fbeOffset = fbeCurrentOffset
        fbeValue.sl = sl.get(size: &fbeFieldSize)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        return fbeCurrentSize
    }

    // Set the struct value
    public func set(value fbeValue: Order) throws -> Int {
        _buffer.shift(offset: fbeOffset)
        let fbeSize = try setFields(fbeValue: fbeValue)
        _buffer.unshift(offset: fbeOffset)
        return fbeSize
    }

    // Set the struct fields values
    public func setFields(fbeValue: Order) throws -> Int {
        var fbeCurrentOffset: Int = 0
        var fbeCurrentSize: Int = 0
        let fbeFieldSize = Size()

        id.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try id.set(value: fbeValue.id)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        symbol.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try symbol.set(value: fbeValue.symbol)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        side.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try side.set(value: fbeValue.side)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        type.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try type.set(value: fbeValue.type)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        price.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try price.set(value: fbeValue.price)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        volume.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try volume.set(value: fbeValue.volume)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        tp.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try tp.set(value: fbeValue.tp)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        sl.fbeOffset = fbeCurrentOffset
        fbeFieldSize.value = try sl.set(value: fbeValue.sl)
        fbeCurrentOffset += fbeFieldSize.value
        fbeCurrentSize += fbeFieldSize.value

        return fbeCurrentSize
    }
}
