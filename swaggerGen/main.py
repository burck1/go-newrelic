import json

def read_json(fname):
    with open(fname, "r", encoding="utf8") as f:
        return json.load(f)

def write_json(fname, data):
    with open(fname, "w") as f:
        json.dump(data, f)

swagger_v2 = read_json("swagger.v2.json")

paths = [api["path"].replace("{format}", "json").split("/")[-1] for api in read_json("definitions\\definitions.json")["apis"]]

for fname in paths:
    data = read_json(f"definitions\\{fname}")
    for api in data["apis"]:
        path = api["path"].replace("{format}", "json")
        for operation in api["operations"]:
            httpMethod = operation["httpMethod"]
            responseClass = operation["responseClass"] if "responseClass" in operation else None
            print(f"{httpMethod} {path} > {responseClass}")
            if responseClass:
                for path_v2 in swagger_v2["paths"]:
                    if path_v2.lower() == path.lower() or path_v2.lower() == path.lower().replace("{id}", "{condition_id}") or path_v2.lower().replace("{id}", "{policy_id}") == path.lower():
                        operations = swagger_v2["paths"][path_v2]
                        for operation_v2 in operations:
                            if operation_v2.lower() == httpMethod.lower():
                                method_v2 = operations[operation_v2]
                                if "responses" in method_v2:
                                    responses = method_v2["responses"]
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
                print("NO RESPONSECLASS")

write_json("swagger_2.v2.json", swagger_v2)
