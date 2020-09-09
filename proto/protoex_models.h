// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: protoex.fbe
// Version: 1.3.0.0

#pragma once

#if defined(__clang__)
#pragma clang system_header
#elif defined(__GNUC__)
#pragma GCC system_header
#elif defined(_MSC_VER)
#pragma system_header
#endif

#include "fbe_models.h"

#include "protoex.h"

#include "proto_models.h"

namespace FBE {

// Fast Binary Encoding ::protoex::OrderSide field model
template <>
class FieldModel<::protoex::OrderSide> : public FieldModelBase<::protoex::OrderSide, uint8_t>
{
public:
    using FieldModelBase<::protoex::OrderSide, uint8_t>::FieldModelBase;
};

// Fast Binary Encoding ::protoex::OrderType field model
template <>
class FieldModel<::protoex::OrderType> : public FieldModelBase<::protoex::OrderType, uint8_t>
{
public:
    using FieldModelBase<::protoex::OrderType, uint8_t>::FieldModelBase;
};

// Fast Binary Encoding ::protoex::StateEx field model
template <>
class FieldModel<::protoex::StateEx> : public FieldModelBase<::protoex::StateEx, uint8_t>
{
public:
    using FieldModelBase<::protoex::StateEx, uint8_t>::FieldModelBase;
};

// Fast Binary Encoding ::protoex::Order field model
template <>
class FieldModel<::protoex::Order>
{
public:
    FieldModel(FBEBuffer& buffer, size_t offset) noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept { return 4; }
    // Get the field body size
    size_t fbe_body() const noexcept;
    // Get the field extra size
    size_t fbe_extra() const noexcept;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return 1; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the struct value is valid
    bool verify(bool fbe_verify_type = true) const noexcept;
    // Check if the struct fields are valid
    bool verify_fields(size_t fbe_struct_size) const noexcept;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept;

    // Get the struct value
    void get(::protoex::Order& fbe_value) const noexcept;
    // Get the struct fields values
    void get_fields(::protoex::Order& fbe_value, size_t fbe_struct_size) const noexcept;

    // Set the struct value (begin phase)
    size_t set_begin();
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin);

    // Set the struct value
    void set(const ::protoex::Order& fbe_value) noexcept;
    // Set the struct fields values
    void set_fields(const ::protoex::Order& fbe_value) noexcept;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    FieldModel<int32_t> id;
    FieldModel<std::string> symbol;
    FieldModel<::protoex::OrderSide> side;
    FieldModel<::protoex::OrderType> type;
    FieldModel<double> price;
    FieldModel<double> volume;
    FieldModel<double> tp;
    FieldModel<double> sl;
};

namespace protoex {

// Fast Binary Encoding Order model
class OrderModel : public FBE::Model
{
public:
    OrderModel() : model(this->buffer(), 4) {}
    OrderModel(const std::shared_ptr<FBEBuffer>& buffer) : FBE::Model(buffer), model(this->buffer(), 4) {}

    // Get the model size
    size_t fbe_size() const noexcept { return model.fbe_size() + model.fbe_extra(); }
    // Get the model type
    static constexpr size_t fbe_type() noexcept { return FieldModel<::protoex::Order>::fbe_type(); }

    // Check if the struct value is valid
    bool verify();

    // Create a new model (begin phase)
    size_t create_begin();
    // Create a new model (end phase)
    size_t create_end(size_t fbe_begin);

    // Serialize the struct value
    size_t serialize(const ::protoex::Order& value);
    // Deserialize the struct value
    size_t deserialize(::protoex::Order& value) const noexcept;

    // Move to the next struct value
    void next(size_t prev) noexcept { model.fbe_shift(prev); }

public:
    FieldModel<::protoex::Order> model;
};

} // namespace protoex

// Fast Binary Encoding ::protoex::Balance field model
template <>
class FieldModel<::protoex::Balance>
{
public:
    FieldModel(FBEBuffer& buffer, size_t offset) noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept { return 4; }
    // Get the field body size
    size_t fbe_body() const noexcept;
    // Get the field extra size
    size_t fbe_extra() const noexcept;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return FieldModel<::proto::Balance>::fbe_type(); }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the struct value is valid
    bool verify(bool fbe_verify_type = true) const noexcept;
    // Check if the struct fields are valid
    bool verify_fields(size_t fbe_struct_size) const noexcept;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept;

    // Get the struct value
    void get(::protoex::Balance& fbe_value) const noexcept;
    // Get the struct fields values
    void get_fields(::protoex::Balance& fbe_value, size_t fbe_struct_size) const noexcept;

    // Set the struct value (begin phase)
    size_t set_begin();
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin);

    // Set the struct value
    void set(const ::protoex::Balance& fbe_value) noexcept;
    // Set the struct fields values
    void set_fields(const ::protoex::Balance& fbe_value) noexcept;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    FieldModel<::proto::Balance> parent;
    FieldModel<double> locked;
};

namespace protoex {

// Fast Binary Encoding Balance model
class BalanceModel : public FBE::Model
{
public:
    BalanceModel() : model(this->buffer(), 4) {}
    BalanceModel(const std::shared_ptr<FBEBuffer>& buffer) : FBE::Model(buffer), model(this->buffer(), 4) {}

    // Get the model size
    size_t fbe_size() const noexcept { return model.fbe_size() + model.fbe_extra(); }
    // Get the model type
    static constexpr size_t fbe_type() noexcept { return FieldModel<::protoex::Balance>::fbe_type(); }

    // Check if the struct value is valid
    bool verify();

    // Create a new model (begin phase)
    size_t create_begin();
    // Create a new model (end phase)
    size_t create_end(size_t fbe_begin);

    // Serialize the struct value
    size_t serialize(const ::protoex::Balance& value);
    // Deserialize the struct value
    size_t deserialize(::protoex::Balance& value) const noexcept;

    // Move to the next struct value
    void next(size_t prev) noexcept { model.fbe_shift(prev); }

public:
    FieldModel<::protoex::Balance> model;
};

} // namespace protoex

// Fast Binary Encoding ::protoex::Account field model
template <>
class FieldModel<::protoex::Account>
{
public:
    FieldModel(FBEBuffer& buffer, size_t offset) noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept { return 4; }
    // Get the field body size
    size_t fbe_body() const noexcept;
    // Get the field extra size
    size_t fbe_extra() const noexcept;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return 3; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the struct value is valid
    bool verify(bool fbe_verify_type = true) const noexcept;
    // Check if the struct fields are valid
    bool verify_fields(size_t fbe_struct_size) const noexcept;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept;

    // Get the struct value
    void get(::protoex::Account& fbe_value) const noexcept;
    // Get the struct fields values
    void get_fields(::protoex::Account& fbe_value, size_t fbe_struct_size) const noexcept;

    // Set the struct value (begin phase)
    size_t set_begin();
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin);

    // Set the struct value
    void set(const ::protoex::Account& fbe_value) noexcept;
    // Set the struct fields values
    void set_fields(const ::protoex::Account& fbe_value) noexcept;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    FieldModel<int32_t> id;
    FieldModel<std::string> name;
    FieldModel<::protoex::StateEx> state;
    FieldModel<::protoex::Balance> wallet;
    FieldModel<std::optional<::protoex::Balance>> asset;
    FieldModelVector<::protoex::Order> orders;
};

namespace protoex {

// Fast Binary Encoding Account model
class AccountModel : public FBE::Model
{
public:
    AccountModel() : model(this->buffer(), 4) {}
    AccountModel(const std::shared_ptr<FBEBuffer>& buffer) : FBE::Model(buffer), model(this->buffer(), 4) {}

    // Get the model size
    size_t fbe_size() const noexcept { return model.fbe_size() + model.fbe_extra(); }
    // Get the model type
    static constexpr size_t fbe_type() noexcept { return FieldModel<::protoex::Account>::fbe_type(); }

    // Check if the struct value is valid
    bool verify();

    // Create a new model (begin phase)
    size_t create_begin();
    // Create a new model (end phase)
    size_t create_end(size_t fbe_begin);

    // Serialize the struct value
    size_t serialize(const ::protoex::Account& value);
    // Deserialize the struct value
    size_t deserialize(::protoex::Account& value) const noexcept;

    // Move to the next struct value
    void next(size_t prev) noexcept { model.fbe_shift(prev); }

public:
    FieldModel<::protoex::Account> model;
};

} // namespace protoex

// Fast Binary Encoding ::protoex::OrderMessage field model
template <>
class FieldModel<::protoex::OrderMessage>
{
public:
    FieldModel(FBEBuffer& buffer, size_t offset) noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept { return 4; }
    // Get the field body size
    size_t fbe_body() const noexcept;
    // Get the field extra size
    size_t fbe_extra() const noexcept;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return 11; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the struct value is valid
    bool verify(bool fbe_verify_type = true) const noexcept;
    // Check if the struct fields are valid
    bool verify_fields(size_t fbe_struct_size) const noexcept;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept;

    // Get the struct value
    void get(::protoex::OrderMessage& fbe_value) const noexcept;
    // Get the struct fields values
    void get_fields(::protoex::OrderMessage& fbe_value, size_t fbe_struct_size) const noexcept;

    // Set the struct value (begin phase)
    size_t set_begin();
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin);

    // Set the struct value
    void set(const ::protoex::OrderMessage& fbe_value) noexcept;
    // Set the struct fields values
    void set_fields(const ::protoex::OrderMessage& fbe_value) noexcept;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    FieldModel<::protoex::Order> body;
};

namespace protoex {

// Fast Binary Encoding OrderMessage model
class OrderMessageModel : public FBE::Model
{
public:
    OrderMessageModel() : model(this->buffer(), 4) {}
    OrderMessageModel(const std::shared_ptr<FBEBuffer>& buffer) : FBE::Model(buffer), model(this->buffer(), 4) {}

    // Get the model size
    size_t fbe_size() const noexcept { return model.fbe_size() + model.fbe_extra(); }
    // Get the model type
    static constexpr size_t fbe_type() noexcept { return FieldModel<::protoex::OrderMessage>::fbe_type(); }

    // Check if the struct value is valid
    bool verify();

    // Create a new model (begin phase)
    size_t create_begin();
    // Create a new model (end phase)
    size_t create_end(size_t fbe_begin);

    // Serialize the struct value
    size_t serialize(const ::protoex::OrderMessage& value);
    // Deserialize the struct value
    size_t deserialize(::protoex::OrderMessage& value) const noexcept;

    // Move to the next struct value
    void next(size_t prev) noexcept { model.fbe_shift(prev); }

public:
    FieldModel<::protoex::OrderMessage> model;
};

} // namespace protoex

// Fast Binary Encoding ::protoex::BalanceMessage field model
template <>
class FieldModel<::protoex::BalanceMessage>
{
public:
    FieldModel(FBEBuffer& buffer, size_t offset) noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept { return 4; }
    // Get the field body size
    size_t fbe_body() const noexcept;
    // Get the field extra size
    size_t fbe_extra() const noexcept;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return 12; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the struct value is valid
    bool verify(bool fbe_verify_type = true) const noexcept;
    // Check if the struct fields are valid
    bool verify_fields(size_t fbe_struct_size) const noexcept;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept;

    // Get the struct value
    void get(::protoex::BalanceMessage& fbe_value) const noexcept;
    // Get the struct fields values
    void get_fields(::protoex::BalanceMessage& fbe_value, size_t fbe_struct_size) const noexcept;

    // Set the struct value (begin phase)
    size_t set_begin();
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin);

    // Set the struct value
    void set(const ::protoex::BalanceMessage& fbe_value) noexcept;
    // Set the struct fields values
    void set_fields(const ::protoex::BalanceMessage& fbe_value) noexcept;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    FieldModel<::protoex::Balance> body;
};

namespace protoex {

// Fast Binary Encoding BalanceMessage model
class BalanceMessageModel : public FBE::Model
{
public:
    BalanceMessageModel() : model(this->buffer(), 4) {}
    BalanceMessageModel(const std::shared_ptr<FBEBuffer>& buffer) : FBE::Model(buffer), model(this->buffer(), 4) {}

    // Get the model size
    size_t fbe_size() const noexcept { return model.fbe_size() + model.fbe_extra(); }
    // Get the model type
    static constexpr size_t fbe_type() noexcept { return FieldModel<::protoex::BalanceMessage>::fbe_type(); }

    // Check if the struct value is valid
    bool verify();

    // Create a new model (begin phase)
    size_t create_begin();
    // Create a new model (end phase)
    size_t create_end(size_t fbe_begin);

    // Serialize the struct value
    size_t serialize(const ::protoex::BalanceMessage& value);
    // Deserialize the struct value
    size_t deserialize(::protoex::BalanceMessage& value) const noexcept;

    // Move to the next struct value
    void next(size_t prev) noexcept { model.fbe_shift(prev); }

public:
    FieldModel<::protoex::BalanceMessage> model;
};

} // namespace protoex

// Fast Binary Encoding ::protoex::AccountMessage field model
template <>
class FieldModel<::protoex::AccountMessage>
{
public:
    FieldModel(FBEBuffer& buffer, size_t offset) noexcept;

    // Get the field offset
    size_t fbe_offset() const noexcept { return _offset; }
    // Get the field size
    size_t fbe_size() const noexcept { return 4; }
    // Get the field body size
    size_t fbe_body() const noexcept;
    // Get the field extra size
    size_t fbe_extra() const noexcept;
    // Get the field type
    static constexpr size_t fbe_type() noexcept { return 13; }

    // Shift the current field offset
    void fbe_shift(size_t size) noexcept { _offset += size; }
    // Unshift the current field offset
    void fbe_unshift(size_t size) noexcept { _offset -= size; }

    // Check if the struct value is valid
    bool verify(bool fbe_verify_type = true) const noexcept;
    // Check if the struct fields are valid
    bool verify_fields(size_t fbe_struct_size) const noexcept;

    // Get the struct value (begin phase)
    size_t get_begin() const noexcept;
    // Get the struct value (end phase)
    void get_end(size_t fbe_begin) const noexcept;

    // Get the struct value
    void get(::protoex::AccountMessage& fbe_value) const noexcept;
    // Get the struct fields values
    void get_fields(::protoex::AccountMessage& fbe_value, size_t fbe_struct_size) const noexcept;

    // Set the struct value (begin phase)
    size_t set_begin();
    // Set the struct value (end phase)
    void set_end(size_t fbe_begin);

    // Set the struct value
    void set(const ::protoex::AccountMessage& fbe_value) noexcept;
    // Set the struct fields values
    void set_fields(const ::protoex::AccountMessage& fbe_value) noexcept;

private:
    FBEBuffer& _buffer;
    size_t _offset;

public:
    FieldModel<::protoex::Account> body;
};

namespace protoex {

// Fast Binary Encoding AccountMessage model
class AccountMessageModel : public FBE::Model
{
public:
    AccountMessageModel() : model(this->buffer(), 4) {}
    AccountMessageModel(const std::shared_ptr<FBEBuffer>& buffer) : FBE::Model(buffer), model(this->buffer(), 4) {}

    // Get the model size
    size_t fbe_size() const noexcept { return model.fbe_size() + model.fbe_extra(); }
    // Get the model type
    static constexpr size_t fbe_type() noexcept { return FieldModel<::protoex::AccountMessage>::fbe_type(); }

    // Check if the struct value is valid
    bool verify();

    // Create a new model (begin phase)
    size_t create_begin();
    // Create a new model (end phase)
    size_t create_end(size_t fbe_begin);

    // Serialize the struct value
    size_t serialize(const ::protoex::AccountMessage& value);
    // Deserialize the struct value
    size_t deserialize(::protoex::AccountMessage& value) const noexcept;

    // Move to the next struct value
    void next(size_t prev) noexcept { model.fbe_shift(prev); }

public:
    FieldModel<::protoex::AccountMessage> model;
};

} // namespace protoex

} // namespace FBE
