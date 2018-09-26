import json
import re

type_map = {
    "Boolean": {
        "type": "boolean",
    },
    "Date": {
        "type": "string",
        "format": "date",
    },
    "Integer": {
        "type": "integer",
    },
    "String": {
        "type": "string",
    },
    "Time": {
        "type": "string",
        "format": "date-time",
    },
}

"""
Array
{
    "type": "array",
    "items": {
        "$ref": "#/definitions/pet"
    },
}

List
{
    "type": "array",
    "items": {
        "type": "...",
        "format": "...",
    },
    "collectionFormat": "csv",
}
"""

def read_json(fname):
    with open(fname, "r", encoding="utf8") as f:
        return json.load(f)

def write_json(fname, data):
    with open(fname, "w") as f:
        json.dump(data, f)

swagger_v2 = read_json("swagger.v2.json")

paths = [api["path"].replace("{format}", "json").split("/")[-1] for api in read_json("definitions\\definitions.json")["apis"]]

path_var_match = re.compile(r"\{[a-z_]*\}")

def find_new_param_name(new_path):
    found = path_var_match.findall(new_path)
    if found:
        return found[-1][1:-1]
    else:
        return None

def process_operation(old_path, old_operation, new_path, new_operation):
    # update 200 response class
    responseClass = operation["responseClass"] if "responseClass" in operation else None
    if responseClass:
        if "responses" in new_operation:
            responses = new_operation["responses"]
            if "200" in responses:
                response = responses["200"]
                response["description"] = "OK"
                if "schema" not in response:
                    response["schema"] = {}
                response["schema"]["$ref"] = f"#/definitions/{responseClass}"
                if "type" in response["schema"]:
                    del response["schema"]["type"]
            else:
                print("MISSING 200")
        else:
            print("MISSING RESPONSES")
    else:
        print("NO RESPONSE CLASS")
    
    # update parameter types
    if "parameters" in new_operation:    
        if "parameters" in old_operation:
            for old_parameter in old_operation["parameters"]:
                old_name = old_parameter["name"]
                old_description = old_parameter["description"]
                old_paramType = old_parameter["paramType"]
                old_dataType = old_parameter["dataType"]
                old_required = old_parameter["required"]
                
                new_parameter_found = False
                for new_parameter in new_operation["parameters"]:
                    new_name = new_parameter["name"]
                    if new_name == old_name:
                        new_parameter_found = True
                if not new_parameter_found:
                    if old_name == "id":
                        found_name = find_new_param_name(new_path)
                        if found_name:
                            old_name = found_name
                        else:
                            print("COULD NOT FIND NEW PARAM NAME")

                new_parameter_found = False
                for new_parameter in new_operation["parameters"]:
                    new_name = new_parameter["name"]
                    if new_name == old_name:
                        new_parameter_found = True
                        new_description = new_parameter["description"]
                        new_in = new_parameter["in"]
                        new_required = new_parameter["required"]
                        if old_dataType in type_map:
                            new_type = type_map[old_dataType]
                            new_parameter["type"] = new_type["type"]
                            if "format" in new_type:
                                new_parameter["format"] = new_type["format"]
                        elif old_dataType == "List":
                            new_parameter["type"] = "array"
                            new_parameter["collectionFormat"] = "csv"
                            # known integer parameters
                            if new_name in ("filter[ids]", "channel_ids"):
                                new_parameter["items"] = {
                                    "type": "integer",
                                }
                            else:
                                # default to strings
                                new_parameter["items"] = {
                                    "type": "string",
                                }
                        elif old_dataType == "Array":
                            if old_paramType == "query":
                                new_parameter["type"] = "array"
                                new_parameter["collectionFormat"] = "csv"
                                # known integer parameters
                                if new_name in ("policy_ids",):
                                    new_parameter["items"] = {
                                        "type": "integer",
                                    }
                                else:
                                    # default to strings
                                    new_parameter["items"] = {
                                        "type": "string",
                                    }
                            else:
                                print("UNRECOGNIZED ARRAY TYPE")
                        else:
                            # data type is a complex type
                            if new_in == "body":
                                new_parameter["schema"] = {
                                    "$ref": f"#/definitions/{old_dataType}",
                                }
                            else:
                                print("COMPLEX TYPE NOT IN BODY")
                        if new_in in ("query", "path"):
                            new_type = new_parameter["type"]
                            new_format = new_parameter["format"] if "format" in new_parameter else None
                            new_allowEmptyValue = new_parameter["allowEmptyValue"] if "allowEmptyValue" in new_parameter else None
                            new_default = new_parameter["default"] if "default" in new_parameter else None
                            if new_type == "array":
                                new_items = new_parameter["items"]
                                new_collectionFormat = new_parameter["collectionFormat"] if "collectionFormat" in new_parameter else None
                        elif new_in == "header":
                            print("IN HEADER")
                        elif new_in == "formData":
                            print("IN FORMDATA")
                        elif new_in == "body":
                            new_schema = new_parameter["schema"]
                        else:
                            print("UNRECOGNIZED \"IN\"")
                        # new_type = new_parameter["type"]
                if not new_parameter_found:
                    print(f"PARAMETER {old_name} NOT FOUND")
        else:
            print("NO OLD PARAMETERS")
    else:
        print("NO NEW PARAMETERS")


for fname in paths:
    data = read_json(f"definitions\\{fname}")
    for api in data["apis"]:
        path = api["path"].replace("{format}", "json")
        for operation in api["operations"]:
            httpMethod = operation["httpMethod"]
            # print(f"{httpMethod} {path}")
            path_operation_found = False
            for path_v2 in swagger_v2["paths"]:
                if path_v2.lower() == path.lower() or path_var_match.sub("{id}", path_v2.lower()) == path_var_match.sub("{id}", path.lower()):
                    operations = swagger_v2["paths"][path_v2]
                    for operation_v2 in operations:
                        if operation_v2.lower() == httpMethod.lower():
                            path_operation_found = True
                            process_operation(path, operation, path_v2, operations[operation_v2])
            if not path_operation_found:
                print("PATH + OPERATION NOT FOUND")

write_json("swagger_2.v2.json", swagger_v2)
