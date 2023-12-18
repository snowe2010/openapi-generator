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
 * ClassModel.h
 *
 * Model for testing model with \&quot;_class\&quot; property
 */

#ifndef ClassModel_H_
#define ClassModel_H_



#include <string>
#include <memory>
#include <vector>
#include <boost/property_tree/ptree.hpp>
#include "helpers.h"

namespace org {
namespace openapitools {
namespace server {
namespace model {

/// <summary>
/// Model for testing model with \&quot;_class\&quot; property
/// </summary>
class  ClassModel 
{
public:
    ClassModel() = default;
    explicit ClassModel(boost::property_tree::ptree const& pt);
    virtual ~ClassModel() = default;

    ClassModel(const ClassModel& other) = default; // copy constructor
    ClassModel(ClassModel&& other) noexcept = default; // move constructor

    ClassModel& operator=(const ClassModel& other) = default; // copy assignment
    ClassModel& operator=(ClassModel&& other) noexcept = default; // move assignment

    std::string toJsonString(bool prettyJson = false) const;
    void fromJsonString(std::string const& jsonString);
    boost::property_tree::ptree toPropertyTree() const;
    void fromPropertyTree(boost::property_tree::ptree const& pt);


    /////////////////////////////////////////////
    /// ClassModel members

    /// <summary>
    /// 
    /// </summary>
    std::string get_Class() const;
    void set_Class(std::string value);

protected:
    std::string m__class = "";
};

std::vector<ClassModel> createClassModelVectorFromJsonString(const std::string& json);

template<>
inline boost::property_tree::ptree toPt<ClassModel>(const ClassModel& val) {
    return val.toPropertyTree();
}

template<>
inline ClassModel fromPt<ClassModel>(const boost::property_tree::ptree& pt) {
    ClassModel ret;
    ret.fromPropertyTree(pt);
    return ret;
}

}
}
}
}

#endif /* ClassModel_H_ */
