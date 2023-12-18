/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI-Generator 7.2.0-SNAPSHOT.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/*
 * ChildWithNullable.h
 *
 * 
 */

#ifndef ChildWithNullable_H_
#define ChildWithNullable_H_



#include "ParentWithNullable.h"
#include <string>
#include <memory>
#include <vector>
#include <array>
#include <boost/property_tree/ptree.hpp>
#include "ParentWithNullable.h"
#include "helpers.h"

namespace org {
namespace openapitools {
namespace server {
namespace model {

/// <summary>
/// 
/// </summary>
class  ChildWithNullable : public ParentWithNullable
{
public:
    ChildWithNullable() = default;
    explicit ChildWithNullable(boost::property_tree::ptree const& pt);
    virtual ~ChildWithNullable() = default;

    ChildWithNullable(const ChildWithNullable& other) = default; // copy constructor
    ChildWithNullable(ChildWithNullable&& other) noexcept = default; // move constructor

    ChildWithNullable& operator=(const ChildWithNullable& other) = default; // copy assignment
    ChildWithNullable& operator=(ChildWithNullable&& other) noexcept = default; // move assignment

    std::string toJsonString(bool prettyJson = false) const;
    void fromJsonString(std::string const& jsonString);
    boost::property_tree::ptree toPropertyTree() const;
    void fromPropertyTree(boost::property_tree::ptree const& pt);


    /////////////////////////////////////////////
    /// ChildWithNullable members

    /// <summary>
    /// 
    /// </summary>
    std::string getType() const;
    void setType(std::string value);

    /// <summary>
    /// 
    /// </summary>
    std::string getNullableProperty() const;
    void setNullableProperty(std::string value);

    /// <summary>
    /// 
    /// </summary>
    std::string getOtherProperty() const;
    void setOtherProperty(std::string value);

protected:
    std::string m_Type = "";
    std::string m_NullableProperty = "";
    std::string m_OtherProperty = "";
};

std::vector<ChildWithNullable> createChildWithNullableVectorFromJsonString(const std::string& json);

template<>
inline boost::property_tree::ptree toPt<ChildWithNullable>(const ChildWithNullable& val) {
    return val.toPropertyTree();
}

template<>
inline ChildWithNullable fromPt<ChildWithNullable>(const boost::property_tree::ptree& pt) {
    ChildWithNullable ret;
    ret.fromPropertyTree(pt);
    return ret;
}

}
}
}
}

#endif /* ChildWithNullable_H_ */
