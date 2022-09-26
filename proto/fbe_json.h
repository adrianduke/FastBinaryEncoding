//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// FBE version: 1.12.0.0
//------------------------------------------------------------------------------

#pragma once

#if defined(__clang__)
#pragma clang system_header
#elif defined(__GNUC__)
#pragma GCC system_header
#elif defined(_MSC_VER)
#pragma system_header
#endif

#include "fbe.h"

#define RAPIDJSON_HAS_STDSTRING 1
#include <rapidjson/document.h>
#include <rapidjson/prettywriter.h>

namespace FBE {

namespace JSON {

template <class TWriter, typename T>
struct KeyWriter
{
    static bool to_json_key(TWriter& writer, const T& key)
    {
        return writer.Key(std::to_string(key));
    }
};

template <class TWriter, typename T>
bool to_json_key(TWriter& writer, const T& key)
{
    return KeyWriter<TWriter, T>::to_json_key(writer, key);
}

template <class TWriter>
struct KeyWriter<TWriter, FBE::decimal_t>
{
    static bool to_json_key(TWriter& writer, const FBE::decimal_t& key)
    {
        return writer.Key(key.string());
    }
};

template <class TWriter>
struct KeyWriter<TWriter, FBE::uuid_t>
{
    static bool to_json_key(TWriter& writer, const FBE::uuid_t& key)
    {
        return writer.Key(key.string());
    }
};

template <class TWriter>
struct KeyWriter<TWriter, std::string>
{
    static bool to_json_key(TWriter& writer, const std::string& key)
    {
        return writer.Key(key);
    }
};

template <class TWriter, size_t N>
struct KeyWriter<TWriter, char[N]>
{
    static bool to_json_key(TWriter& writer, const char (&key)[N])
    {
        return writer.Key(key, N - 1);
    }
};

template <class TWriter, typename T>
struct ValueWriter
{
    static bool to_json(TWriter& writer, const T& value, bool scope = true)
    {
        throw std::logic_error("Not implemented!");
    }
};

template <class TWriter, typename T>
bool to_json(TWriter& writer, const T& value, bool scope = true)
{
    return ValueWriter<TWriter, T>::to_json(writer, value, scope);
}

template <class TWriter>
struct ValueWriter<TWriter, bool>
{
    static bool to_json(TWriter& writer, const bool& value, bool scope = true)
    {
        return writer.Bool(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, char>
{
    static bool to_json(TWriter& writer, const char& value, bool scope = true)
    {
        return writer.Uint(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, wchar_t>
{
    static bool to_json(TWriter& writer, const wchar_t& value, bool scope = true)
    {
        return writer.Uint(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, int8_t>
{
    static bool to_json(TWriter& writer, const int8_t& value, bool scope = true)
    {
        return writer.Int(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, uint8_t>
{
    static bool to_json(TWriter& writer, const uint8_t& value, bool scope = true)
    {
        return writer.Uint(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, int16_t>
{
    static bool to_json(TWriter& writer, const int16_t& value, bool scope = true)
    {
        return writer.Int(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, uint16_t>
{
    static bool to_json(TWriter& writer, const uint16_t& value, bool scope = true)
    {
        return writer.Uint(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, int32_t>
{
    static bool to_json(TWriter& writer, const int32_t& value, bool scope = true)
    {
        return writer.Int(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, uint32_t>
{
    static bool to_json(TWriter& writer, const uint32_t& value, bool scope = true)
    {
        return writer.Uint(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, int64_t>
{
    static bool to_json(TWriter& writer, const int64_t& value, bool scope = true)
    {
        return writer.Int64(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, uint64_t>
{
    static bool to_json(TWriter& writer, const uint64_t& value, bool scope = true)
    {
        return writer.Uint64(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, float>
{
    static bool to_json(TWriter& writer, const float& value, bool scope = true)
    {
        return writer.Double(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, double>
{
    static bool to_json(TWriter& writer, const double& value, bool scope = true)
    {
        return writer.Double(value);
    }
};

template <class TWriter>
struct ValueWriter<TWriter, FBE::decimal_t>
{
    static bool to_json(TWriter& writer, const FBE::decimal_t& value, bool scope = true)
    {
        return writer.String(value.string());
    }
};

template <class TWriter>
struct ValueWriter<TWriter, FBE::uuid_t>
{
    static bool to_json(TWriter& writer, const FBE::uuid_t& value, bool scope = true)
    {
        return writer.String(value.string());
    }
};

template <class TWriter>
struct ValueWriter<TWriter, std::string>
{
    static bool to_json(TWriter& writer, const std::string& value, bool scope = true)
    {
        return writer.String(value);
    }
};

template <class TWriter, std::size_t N>
struct ValueWriter<TWriter, char[N]>
{
    static bool to_json(TWriter& writer, const char (&value)[N], bool scope = true)
    {
        return writer.String(value, N);
    }
};

template <class TWriter, typename T>
struct ValueWriter<TWriter, std::optional<T>>
{
    static bool to_json(TWriter& writer, const std::optional<T>& value, bool scope = true)
    {
        if (value)
            return ValueWriter<TWriter, T>::to_json(writer, value.value(), true);
        else
            return writer.Null();
    }
};

template <class TWriter>
struct ValueWriter<TWriter, buffer_t>
{
    static bool to_json(TWriter& writer, const buffer_t& values, bool scope = true)
    {
        return writer.String(values.base64encode());
    }
};

template <class TWriter, typename T, size_t N>
struct ValueWriter<TWriter, std::array<T, N>>
{
    static bool to_json(TWriter& writer, const std::array<T, N>& values, bool scope = true)
    {
        writer.StartArray();
        for (const auto& value : values)
            if (!ValueWriter<TWriter, T>::to_json(writer, value, true))
                return false;
        writer.EndArray();
        return true;
    }
};

template <class TWriter, typename T>
struct ValueWriter<TWriter, std::vector<T>>
{
    static bool to_json(TWriter& writer, const std::vector<T>& values, bool scope = true)
    {
        writer.StartArray();
        for (const auto& value : values)
            if (!FBE::JSON::to_json(writer, value, true))
                return false;
        writer.EndArray();
        return true;
    }
};

template <class TWriter, typename T>
struct ValueWriter<TWriter, std::list<T>>
{
    static bool to_json(TWriter& writer, const std::list<T>& values, bool scope = true)
    {
        writer.StartArray();
        for (const auto& value : values)
            if (!FBE::JSON::to_json(writer, value, true))
                return false;
        writer.EndArray();
        return true;
    }
};

template <class TWriter, typename T>
struct ValueWriter<TWriter, std::set<T>>
{
    static bool to_json(TWriter& writer, const std::set<T>& values, bool scope = true)
    {
        writer.StartArray();
        for (const auto& value : values)
            if (!FBE::JSON::to_json(writer, value, true))
                return false;
        writer.EndArray();
        return true;
    }
};

template <class TWriter, typename TKey, typename TValue>
struct ValueWriter<TWriter, std::map<TKey, TValue>>
{
    static bool to_json(TWriter& writer, const std::map<TKey, TValue>& values, bool scope = true)
    {
        writer.StartObject();
        for (const auto& value : values)
        {
            if (!FBE::JSON::to_json_key(writer, value.first))
                return false;
            if (!FBE::JSON::to_json(writer, value.second, true))
                return false;
        }
        writer.EndObject();
        return true;
    }
};

template <class TWriter, typename TKey, typename TValue>
struct ValueWriter<TWriter, std::unordered_map<TKey, TValue>>
{
    static bool to_json(TWriter& writer, const std::unordered_map<TKey, TValue>& values, bool scope = true)
    {
        writer.StartObject();
        for (const auto& value : values)
        {
            if (!FBE::JSON::to_json_key(writer, value.first))
                return false;
            if (!FBE::JSON::to_json(writer, value.second, true))
                return false;
        }
        writer.EndObject();
        return true;
    }
};

template <class TJson, typename T>
struct ValueReader
{
    static bool from_json(const TJson& json, T& value)
    {
        throw std::logic_error("Not implemented!");
    }
};

template <class TJson, typename T>
bool from_json(const TJson& json, T& value)
{
    return ValueReader<TJson, T>::from_json(json, value);
}

template <class TJson, typename T>
struct KeyReader
{
    static bool from_json_key(const TJson& json, T& key)
    {
        std::string str;
        if (!FBE::JSON::from_json(json, str))
            return false;

        std::istringstream(str) >> key;
        return true;
    }
};

template <class TJson, typename T>
bool from_json_key(const TJson& json, T& key)
{
    return KeyReader<TJson, T>::from_json_key(json, key);
}

template <class TJson>
struct KeyReader<TJson, std::string>
{
    static bool from_json_key(const TJson& json, std::string& key)
    {
        return FBE::JSON::from_json(json, key);
    }
};

template <class TJson>
struct KeyReader<TJson, FBE::decimal_t>
{
    static bool from_json_key(const TJson& json, FBE::decimal_t& key)
    {
        return FBE::JSON::from_json(json, key);
    }
};

template <class TJson>
struct KeyReader<TJson, FBE::uuid_t>
{
    static bool from_json_key(const TJson& json, FBE::uuid_t& key)
    {
        return FBE::JSON::from_json(json, key);
    }
};

template <class TJson, size_t N>
struct KeyReader<TJson, char[N]>
{
    static bool from_json_key(const TJson& json, char (&key)[N])
    {
        return FBE::JSON::from_json(json, key);
    }
};

template <class TJson>
struct ValueReader<TJson, bool>
{
    static bool from_json(const TJson& json, bool& value)
    {
        value = false;

        // Schema validation
        if (json.IsNull() || !json.IsBool())
            return false;

        // Save the value
        value = json.GetBool();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, char>
{
    static bool from_json(const TJson& json, char& value)
    {
        value = '\0';

        // Schema validation
        if (json.IsNull() || !json.IsUint())
            return false;

        // Save the value
        value = (char)json.GetUint();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, wchar_t>
{
    static bool from_json(const TJson& json, wchar_t& value)
    {
        value = L'\0';

        // Schema validation
        if (json.IsNull() || !json.IsUint())
            return false;

        // Save the value
        value = (wchar_t)json.GetUint();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, int8_t>
{
    static bool from_json(const TJson& json, int8_t& value)
    {
        value = 0;

        // Schema validation
        if (json.IsNull() || !json.IsInt())
            return false;

        // Save the value
        value = (int8_t)json.GetInt();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, uint8_t>
{
    static bool from_json(const TJson& json, uint8_t& value)
    {
        value = 0;

        // Schema validation
        if (json.IsNull() || !json.IsUint())
            return false;

        // Save the value
        value = (uint8_t)json.GetUint();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, int16_t>
{
    static bool from_json(const TJson& json, int16_t& value)
    {
        value = 0;

        // Schema validation
        if (json.IsNull() || !json.IsInt())
            return false;

        // Save the value
        value = (int16_t)json.GetInt();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, uint16_t>
{
    static bool from_json(const TJson& json, uint16_t& value)
    {
        value = 0;

        // Schema validation
        if (json.IsNull() || !json.IsUint())
            return false;

        // Save the value
        value = (uint16_t)json.GetUint();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, int32_t>
{
    static bool from_json(const TJson& json, int32_t& value)
    {
        value = 0;

        // Schema validation
        if (json.IsNull() || !json.IsInt())
            return false;

        // Save the value
        value = (int32_t)json.GetInt();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, uint32_t>
{
    static bool from_json(const TJson& json, uint32_t& value)
    {
        value = 0;

        // Schema validation
        if (json.IsNull() || !json.IsUint())
            return false;

        // Save the value
        value = (uint32_t)json.GetUint();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, int64_t>
{
    static bool from_json(const TJson& json, int64_t& value)
    {
        value = 0;

        // Schema validation
        if (json.IsNull() || !json.IsInt64())
            return false;

        // Save the value
        value = (int64_t)json.GetInt64();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, uint64_t>
{
    static bool from_json(const TJson& json, uint64_t& value)
    {
        value = 0;

        // Schema validation
        if (json.IsNull() || !json.IsUint64())
            return false;

        // Save the value
        value = (uint64_t)json.GetUint64();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, float>
{
    static bool from_json(const TJson& json, float& value)
    {
        value = 0.0f;

        // Schema validation
        if (json.IsNull() || !(json.IsInt() || json.IsFloat()))
            return false;

        // Save the value
        value = json.GetFloat();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, double>
{
    static bool from_json(const TJson& json, double& value)
    {
        value = 0.0;

        // Schema validation
        if (json.IsNull() || !(json.IsInt() || json.IsDouble()))
            return false;

        // Save the value
        value = json.GetDouble();
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, FBE::decimal_t>
{
    static bool from_json(const TJson& json, FBE::decimal_t& value)
    {
        value = 0.0;

        // Schema validation
        if (json.IsNull() || !json.IsString())
            return false;

        // Save the value
        try
        {
            std::string str(json.GetString(), (size_t)json.GetStringLength());
            value = std::stod(str);
            return true;
        }
        catch (...)
        {
            return false;
        }
    }
};

template <class TJson>
struct ValueReader<TJson, FBE::uuid_t>
{
    static bool from_json(const TJson& json, FBE::uuid_t& value)
    {
        value = uuid_t::nil();

        // Schema validation
        if (json.IsNull() || !json.IsString())
            return false;

        // Save the value
        try
        {
            std::string str(json.GetString(), (size_t)json.GetStringLength());
            value = str;
            return true;
        }
        catch (...)
        {
            return false;
        }
    }
};

template <class TJson>
struct ValueReader<TJson, std::string>
{
    static bool from_json(const TJson& json, std::string& value)
    {
        value = "";

        // Schema validation
        if (json.IsNull() || !json.IsString())
            return false;

        // Save the value
        value.assign(json.GetString(), (size_t)json.GetStringLength());
        return true;
    }
};

template <class TJson, size_t N>
struct ValueReader<TJson, char[N]>
{
    static bool from_json(const TJson& json, char (&value)[N])
    {
        // Schema validation
        if (json.IsNull() || !json.IsString())
            return false;

        // Save the value
        size_t length = std::min((size_t)json.GetStringLength(), N);
        std::memcpy(value, json.GetString(), length);
        // Write the end of string character if possible
        if (length < N)
            value[length] = '\0';
        return true;
    }
};

template <class TJson, typename T>
struct ValueReader<TJson, std::optional<T>>
{
    static bool from_json(const TJson& json, std::optional<T>& value)
    {
        value = std::nullopt;

        // Empty optional value
        if (json.IsNull())
            return true;

        // Try to get the value
        T temp = T();
        if (!FBE::JSON::from_json(json, temp))
            return false;

        // Save the value
        value.emplace(temp);
        return true;
    }
};

template <class TJson>
struct ValueReader<TJson, buffer_t>
{
    static bool from_json(const TJson& json, buffer_t& value)
    {
        // Schema validation
        if (json.IsNull() || !json.IsString())
            return false;

        std::string str(json.GetString(), (size_t)json.GetStringLength());
        value = buffer_t::base64decode(str);
        return true;
    }
};

template <class TJson, typename T, size_t N>
struct ValueReader<TJson, std::array<T, N>>
{
    static bool from_json(const TJson& json, std::array<T, N>& values)
    {
        // Schema validation
        if (json.IsNull() || !json.IsArray())
            return false;

        // Collect array items
        size_t length = json.GetArray().Size();
        for (size_t i = 0; (i < length) && (i < N); ++i)
            if (!FBE::JSON::from_json(json.GetArray()[(rapidjson::SizeType)i], values[i]))
                return false;
        return true;
    }
};

template <class TJson, typename T>
struct ValueReader<TJson, std::vector<T>>
{
    static bool from_json(const TJson& json, std::vector<T>& values)
    {
        values.clear();

        // Schema validation
        if (json.IsNull() || !json.IsArray())
            return false;

        // Collect vector items
        values.reserve(json.GetArray().Size());
        for (const auto& item : json.GetArray())
        {
            T temp = T();
            if (!FBE::JSON::from_json(item, temp))
                return false;
            values.emplace_back(temp);
        }
        return true;
    }
};

template <class TJson, typename T>
struct ValueReader<TJson, std::list<T>>
{
    static bool from_json(const TJson& json, std::list<T>& values)
    {
        values.clear();

        // Schema validation
        if (json.IsNull() || !json.IsArray())
            return false;

        // Collect list items
        for (const auto& item : json.GetArray())
        {
            T temp = T();
            if (!FBE::JSON::from_json(item, temp))
                return false;
            values.emplace_back(temp);
        }
        return true;
    }
};

template <class TJson, typename T>
struct ValueReader<TJson, std::set<T>>
{
    static bool from_json(const TJson& json, std::set<T>& values)
    {
        values.clear();

        // Schema validation
        if (json.IsNull() || !json.IsArray())
            return false;

        // Collect set items
        for (const auto& item : json.GetArray())
        {
            T temp = T();
            if (!FBE::JSON::from_json(item, temp))
                return false;
            values.emplace(temp);
        }
        return true;
    }
};

template <class TJson, typename TKey, typename TValue>
struct ValueReader<TJson, std::map<TKey, TValue>>
{
    static bool from_json(const TJson& json, std::map<TKey, TValue>& values)
    {
        values.clear();

        // Schema validation
        if (json.IsNull() || !json.IsObject())
            return false;

        // Collect map items
        for (auto it = json.MemberBegin(); it != json.MemberEnd(); ++it)
        {
            TKey key = TKey();
            TValue value = TValue();
            if (!FBE::JSON::from_json_key(it->name, key))
                return false;
            if (!FBE::JSON::from_json(it->value, value))
                return false;
            values.emplace(key, value);
        }
        return true;
    }
};

template <class TJson, typename TKey, typename TValue>
struct ValueReader<TJson, std::unordered_map<TKey, TValue>>
{
    static bool from_json(const TJson& json, std::unordered_map<TKey, TValue>& values)
    {
        values.clear();

        // Schema validation
        if (json.IsNull() || !json.IsObject())
            return false;

        // Collect hash items
        for (auto it = json.MemberBegin(); it != json.MemberEnd(); ++it)
        {
            TKey key = TKey();
            TValue value = TValue();
            if (!FBE::JSON::from_json_key(it->name, key))
                return false;
            if (!FBE::JSON::from_json(it->value, value))
                return false;
            values.emplace(key, value);
        }
        return true;
    }
};

template <class TJson, typename T>
struct NodeReader
{
    static bool from_json(const TJson& json, T& value, const char* key)
    {
        if (key == nullptr)
            return false;

        // Try to find a member with the given key
        rapidjson::Value::ConstMemberIterator member = json.FindMember(key);
        if (member == json.MemberEnd())
            return false;

        return FBE::JSON::from_json(member->value, value);
    }
};

template <class TJson, typename T>
bool from_json(const TJson& json, T& value, const char* key)
{
    return NodeReader<TJson, T>::from_json(json, value, key);
}

template <class TJson, typename T>
struct ChildNodeReader
{
    static bool from_json_child(const TJson& json, T& value, const char* key)
    {
        if (key == nullptr)
            return false;

        // Try to find a member with the given key
        rapidjson::Value::ConstMemberIterator member = json.FindMember(key);
        if (member == json.MemberEnd())
            return false;

        // Schema validation
        if (member->value.IsNull() || !member->value.IsObject())
            return false;

        // Deserialize the child object
        return FBE::JSON::from_json(member->value.GetObj(), value);
    }
};

template <class TJson, typename T>
bool from_json_child(const TJson& json, T& value, const char* key)
{
    return ChildNodeReader<TJson, T>::from_json_child(json, value, key);
}

} // namespace JSON

} // namespace FBE
