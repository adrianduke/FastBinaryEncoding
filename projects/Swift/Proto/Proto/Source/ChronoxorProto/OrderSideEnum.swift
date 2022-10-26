//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// FBE version: 1.14.1.0
//------------------------------------------------------------------------------

import Foundation

public enum OrderSideEnum {
    typealias RawValue = UInt8
    case buy
    case sell

    var rawValue: RawValue {
        switch self {
        case .buy: return 0 + 0
        case .sell: return 0 + 1
        }
    }

    init(value: UInt8) { self = OrderSideEnum(rawValue: NSNumber(value: value).uint8Value) }
    init(value: UInt16) { self = OrderSideEnum(rawValue: NSNumber(value: value).uint8Value) }
    init(value: UInt32) { self = OrderSideEnum(rawValue: NSNumber(value: value).uint8Value) }
    init(value: UInt64) { self = OrderSideEnum(rawValue: NSNumber(value: value).uint8Value) }
    init(value: OrderSideEnum) { self = OrderSideEnum(rawValue: value.rawValue) }
    init(rawValue: UInt8) { self = Self.mapValue(value: rawValue)! }

    var description: String {
        switch self {
        case .buy: return "buy"
        case .sell: return "sell"
        }
    }

    static let rawValuesMap: [RawValue: OrderSideEnum] = {
        var value = [RawValue: OrderSideEnum]()
        value[OrderSideEnum.buy.rawValue] = .buy
        value[OrderSideEnum.sell.rawValue] = .sell
        return value
    }()

    static func mapValue(value: UInt8) -> OrderSideEnum? {
        return rawValuesMap[value]
    }
}
