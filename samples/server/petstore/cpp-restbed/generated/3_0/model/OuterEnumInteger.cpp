/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI-Generator 7.3.0-SNAPSHOT.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */



#include "OuterEnumInteger.h"

#include <string>
#include <vector>
#include <map>
#include <sstream>
#include <stdexcept>
#include <regex>
#include <boost/lexical_cast.hpp>
#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include "helpers.h"

using boost::property_tree::ptree;
using boost::property_tree::read_json;
using boost::property_tree::write_json;

namespace org {
namespace openapitools {
namespace server {
namespace model {

OuterEnumInteger::OuterEnumInteger(boost::property_tree::ptree const& pt)
{
        fromPropertyTree(pt);
}


std::string OuterEnumInteger::toJsonString(bool prettyJson /* = false */) const
{
	std::stringstream ss;
	write_json(ss, this->toPropertyTree(), prettyJson);
    // workaround inspired by: https://stackoverflow.com/a/56395440
    std::regex reg("\\\"([0-9]+\\.{0,1}[0-9]*)\\\"");
    std::string result = std::regex_replace(ss.str(), reg, "$1");
    return result;
}

void OuterEnumInteger::fromJsonString(std::string const& jsonString)
{
	std::stringstream ss(jsonString);
	ptree pt;
	read_json(ss,pt);
	this->fromPropertyTree(pt);
}

ptree OuterEnumInteger::toPropertyTree() const
{
	ptree pt;
	ptree tmp_node;
	return pt;
}

void OuterEnumInteger::fromPropertyTree(ptree const &pt)
{
	ptree tmp_node;
}

std::string OuterEnumInteger::toString() const {
    return boost::lexical_cast<std::string>(getEnumValue());
}

void OuterEnumInteger::fromString(const std::string& str) {
    setEnumValue(boost::lexical_cast<int32_t>(str));
}

int32_t OuterEnumInteger::getEnumValue() const {
    return m_OuterEnumIntegerEnumValue;
}

void OuterEnumInteger::setEnumValue(const int32_t& val) {
    static const std::array<int32_t, 3> allowedValues = {
        0, 1, 2
    };
    if (std::find(allowedValues.begin(), allowedValues.end(), val) != allowedValues.end()) {
        m_OuterEnumIntegerEnumValue = val;
    } else {
        throw std::runtime_error("Value " + boost::lexical_cast<std::string>(val) + " not allowed");
    }
}

std::vector<OuterEnumInteger> createOuterEnumIntegerVectorFromJsonString(const std::string& json)
{
    std::stringstream sstream(json);
    boost::property_tree::ptree pt;
    boost::property_tree::json_parser::read_json(sstream,pt);

    auto vec = std::vector<OuterEnumInteger>();
    for (const auto& child: pt) {
        vec.emplace_back(OuterEnumInteger(child.second));
    }

    return vec;
}

}
}
}
}

