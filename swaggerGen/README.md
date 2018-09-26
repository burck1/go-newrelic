# New Relic's Swagger Docs

New Relic uses Swagger v1.1 internally to generate their [api docs](https://rpm.newrelic.com/api/explore/api/explore) on page load. The swagger json is not officialy published by New Relic, so the intention of this tool is to "scrape" their api docs and create a single `swagger.{json|yaml}`. This swagger doc can then be used to generate a client library for the New Relic apis.

## New Relic's API Explorer

The [API Explorer](https://rpm.newrelic.com/api/explore/api/explore) provides documentation for most of New Relic's APIs. Using Chrome's dev tools to inspect the network requests on page load reveals that the page is using the [swagger.js](https://rpm.newrelic.com/explorer/lib/swagger.js) and [swagger-ui.js](https://rpm.newrelic.com/explorer/swagger-ui.js) javascript libraries to dynamically generate their API docs. Further inspection of the XHR requests being made reveals that the swagger 1.1 spec splits the swagger definitions into multiple files. So, in order for us to "scrape" a single `swagger.{json|yaml}`, we will need to read their root [definitions.json](https://api.newrelic.com/v2/definitions.json) document and then gather the rest of the swagger definitions based on the `apis` objects.

## Converting from Swagger 1.1 to 1.2

These docs use swagger 1.1 which was released on 2012-08-22. In for the docs to be useful for generating a go client, we will need to first convert the swagger format to 2.0.

Unfortunately, I cannot find any swagger 1.1 to swagger 2.0 converters, and not even 1.1 to 1.2 converters. So we will need to do the initial 1.1 to 1.2 conversion ourselves.

### diff based on [1.2 spec](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/1.2.md)

[Resource Listing](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/1.2.md#51-resource-listing)

* `swaggerVersion` should be updated from "1.1" to "1.2"
* `apiVersion` should stay the same
* `basePath` should be removed
* `operations` should be removed
* `apis` paths should be updated to the https://raw.githubusercontent.com url for the corresponding definition
    * Note: can optionally add a `description` to each api object

Note: May want to inject an [Info Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/1.2.md#513-info-object) and an [Authorizations Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/1.2.md#514-authorizations-object).

[API Declarations](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/1.2.md#52-api-declaration)

* `swaggerVersion` should be updated from "1.1" to "1.2"
* `apiVersion` should stay the same
* `basePath` should stay the same
* `resourcePath` should be removed if empty or else stay the same
* `notes` should be moved from the API Declaration to the `description` property for each of the `apis` in the Resource Listing
* `nrExtensions` should be removed
* `apis` may need to combine operations for apis with the same path
    * `path` should replace "{format}" with "json"
    * `description` should stay the same if it exists
    * `operations`
        * `responseClass` should be changed to `type`, `type` can be "void", "array", or a primative, or a model's id
        * `httpMethod` should be changed to `method`
        * `summary` should stay the same
        * `notes` should stay the same
        * `nickname` should stay the same
        * `produces` should stay the same
        * `consumes` should stay the same
        * `nrExtensions` should be removed
        * `parameters`
            * `dataType` should be changed to `type` & `format`
            * `paramType` should stay the same
            * `name` should stay the same
            * `description` should stay the same
            * `required` should stay the same
        * `errorResponses` should be changed to `responseMessages`
            * `code` should stay the same
            * `reason` should be changed to `message`
            * `responseModel` can be added
* `models`
    * `id` should be added, same as key for model
    * `properties`
        * `type` should stay the same
        * `desc` should be changed to `description`
        * `update` should be removed???

Note: Can add `produces`, `consumes`, and `authorizations` to the root object.

## Converting from 1.2 to 2.0

We will be using [APIMATIC](https://www.apimatic.io) to convert our 1.2 swagger doc to a 2.0 swagger doc. To run the conversion you will need an api key from apimatic. This can be created by creating a free account then going to https://www.apimatic.io/account/manage#/authkeys to generate an api key. We will then use this api key to call the [transform](https://www.apimatic.io/apidocs/cgaas-api/v/1_0#/http/api-endpoints/api-transformer/transform-using-file) api endpoint.

## Enriching the 2.0 docs

TODO: add "notes" from original 1.1 swagger docs to 2.0 docs

## [definitions.json](https://api.newrelic.com/v2/definitions.json) : 09/16/2018
The [API Explorer](https://rpm.newrelic.com/api/explore/api/explore) provides documentation for most of New Relic's APIs. Using Chrome's dev tools to inspect the network requests on page load reveals that the page is using the [swagger.js](https://rpm.newrelic.com/explorer/lib/swagger.js) and [swagger-ui.js](https://rpm.newrelic.com/explorer/swagger-ui.js) javascript libraries to dynamically generate their API docs. Furthermore, inspection of the XHR requests being made reveals that the swagger docs are templatized and compiled before being passed to the `swagger.js` library. So, in order for us to "scrape" a single `swagger.{json|yaml}`, we will need to read their root [definitions.json](https://api.newrelic.com/v2/definitions.json) document and then gather the rest of the swagger definitions based on the `apis` objects.

### [definitions.json](https://api.newrelic.com/v2/definitions.json) : 09/16/2018

The `path` property of each of the `apis` objects contains a `{format}` variable. Based on the network requests seen in Chrome, this variable should always be `json`. After some testing, it looks like `xml` is supported as well, but we'll just be using `json` for this project.

```json
{
    "apiVersion": "v2",
    "swaggerVersion": "1.1",
    "basePath": "https://api.newrelic.com",
    "operations": [],
    "apis": [
        { "path": "/v2/definitions/applications.{format}" },
        { "path": "/v2/definitions/application_hosts.{format}" },
        { "path": "/v2/definitions/application_instances.{format}" },
        { "path": "/v2/definitions/application_deployments.{format}" },
        { "path": "/v2/definitions/mobile_applications.{format}" },
        { "path": "/v2/definitions/browser_applications.{format}" },
        { "path": "/v2/definitions/key_transactions.{format}" },
        { "path": "/v2/definitions/servers.{format}" },
        { "path": "/v2/definitions/usages.{format}" },
        { "path": "/v2/definitions/users.{format}" },
        { "path": "/v2/definitions/alerts_events.{format}" },
        { "path": "/v2/definitions/alerts_conditions.{format}" },
        { "path": "/v2/definitions/alerts_plugins_conditions.{format}" },
        { "path": "/v2/definitions/alerts_external_service_conditions.{format}" },
        { "path": "/v2/definitions/alerts_synthetics_conditions.{format}" },
        { "path": "/v2/definitions/alerts_nrql_conditions.{format}" },
        { "path": "/v2/definitions/alerts_policies.{format}" },
        { "path": "/v2/definitions/alerts_channels.{format}" },
        { "path": "/v2/definitions/alerts_policy_channels.{format}" },
        { "path": "/v2/definitions/alerts_violations.{format}" },
        { "path": "/v2/definitions/alerts_incidents.{format}" },
        { "path": "/v2/definitions/alerts_entity_conditions.{format}" },
        { "path": "/v2/definitions/dashboards.{format}" },
        { "path": "/v2/definitions/plugins.{format}" },
        { "path": "/v2/definitions/components.{format}" },
        { "path": "/v2/definitions/labels.{format}" }
    ]
}
```

## [applications.json](https://api.newrelic.com/v2/definitions/applications.json) : 09/16/2018

```json
{
    "apiVersion": "v2",
    "swaggerVersion": "1.1",
    "basePath": "https://api.newrelic.com",
    "resourcePath": "",
    "notes": "<p>The applications resource provides general information about applications monitored by New Relic, including\nresponse time, apdex, throughput and settings. It allows for querying and filtering on multiple parameters\nas well as retrieving information for a single application.</p>\n",
    "apis": [
        {
            "path": "/v2/applications.{format}",
            "operations": [
                {
                    "notes": "<p>This API endpoint returns a <a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/requirements/pagination-api-output\" target=\"_blank\">paginated</a>\nlist of the Applications associated with your New Relic account. The time range for summary data is the last 3-4 minutes.</p>\n\n<p>Applications can be filtered by their name, hosts, the list of application IDs or the application language as\nreported by the agents.</p>\n\n<p>See our documentation for a discussion and examples of\nusing <a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/application-examples-v2/list-application-specific-server-host-instance-ids\" target=\"_blank\"> filters </a>\nand <a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/application-examples-v2/summary-data-examples-v2\" target=\"_blank\">summary data output</a>.</p>\n\n",
                    "summary": "List",
                    "nickname": "list",
                    "httpMethod": "GET",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "filter[name]",
                            "description": "Filter by application name",
                            "dataType": "String",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "filter[host]",
                            "description": "Filter by application host",
                            "dataType": "String",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "filter[ids]",
                            "description": "Filter by application ids",
                            "dataType": "List",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "filter[language]",
                            "description": "Filter by application language",
                            "dataType": "String",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "exclude_links",
                            "description": "Exclude links section from the response",
                            "dataType": "Boolean",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "page",
                            "description": "Pagination index",
                            "dataType": "Integer",
                            "required": false
                        }
                    ],
                    "consumes": [
                        "application/json",
                        "application/xml"
                    ],
                    "produces": [
                        "application/json",
                        "application/xml"
                    ],
                    "nrExtensions": {},
                    "responseClass": "ApplicationResponse",
                    "errorResponses": [
                        {
                            "code": 401,
                            "reason": "Invalid API key"
                        },
                        {
                            "code": 401,
                            "reason": "Invalid request, API key required"
                        },
                        {
                            "code": 403,
                            "reason": "New Relic API access has not been enabled"
                        },
                        {
                            "code": 500,
                            "reason": "A server error occurred, please contact <a href=\"http://support.newrelic.com\">New Relic support</a>"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/v2/applications/{id}.{format}",
            "operations": [
                {
                    "notes": "<p>This API endpoint returns a single Application, identified by ID. The time range for summary data is the last 3-4 minutes.</p>\n\n<p>See our documentation for a discussion of the\n<a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/application-examples-v2/summary-data-examples-v2\" target=\"_blank\"> summary data output</a>.</p>\n\n",
                    "summary": "Show",
                    "nickname": "show",
                    "httpMethod": "GET",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Application ID",
                            "dataType": "Integer",
                            "required": true
                        }
                    ],
                    "consumes": [
                        "application/json",
                        "application/xml"
                    ],
                    "produces": [
                        "application/json",
                        "application/xml"
                    ],
                    "nrExtensions": {},
                    "responseClass": "ApplicationResponse",
                    "errorResponses": [
                        {
                            "code": 401,
                            "reason": "Invalid API key"
                        },
                        {
                            "code": 401,
                            "reason": "Invalid request, API key required"
                        },
                        {
                            "code": 403,
                            "reason": "New Relic API access has not been enabled"
                        },
                        {
                            "code": 500,
                            "reason": "A server error occurred, please contact <a href=\"http://support.newrelic.com\">New Relic support</a>"
                        },
                        {
                            "code": 404,
                            "reason": "No Application found with the given ID"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/v2/applications/{id}.{format}",
            "operations": [
                {
                    "notes": "<p>This API endpoint allows you to update certain parameters of your application.</p>\n\n<p>The input is expected to be in <strong>JSON or XML</strong> format in the body parameter of the PUT request. The exact\nschema is defined below. Any extra parameters passed in the body <strong>will be ignored</strong>.</p>\n\n<p>See our documentation for a discussion and simple example of\n<a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/application-examples-v2/changing-alias-your-application-v2\" target=\"_blank\"> updating</a>\nan application.</p>\n\n",
                    "summary": "Update",
                    "nickname": "update",
                    "httpMethod": "PUT",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Application ID",
                            "dataType": "Integer",
                            "required": true
                        },
                        {
                            "paramType": "body",
                            "name": "application",
                            "description": "Application schema",
                            "dataType": "Application",
                            "required": true
                        }
                    ],
                    "consumes": [
                        "application/json",
                        "application/xml"
                    ],
                    "produces": [
                        "application/json",
                        "application/xml"
                    ],
                    "nrExtensions": {},
                    "responseClass": "ApplicationResponse",
                    "errorResponses": [
                        {
                            "code": 401,
                            "reason": "Invalid API key"
                        },
                        {
                            "code": 401,
                            "reason": "Invalid request, API key required"
                        },
                        {
                            "code": 403,
                            "reason": "New Relic API access has not been enabled"
                        },
                        {
                            "code": 500,
                            "reason": "A server error occurred, please contact <a href=\"http://support.newrelic.com\">New Relic support</a>"
                        },
                        {
                            "code": 422,
                            "reason": "Validation error occurred when updating the application"
                        },
                        {
                            "code": 404,
                            "reason": "No Application found with the given ID"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/v2/applications/{id}.{format}",
            "operations": [
                {
                    "notes": "<p>This API endpoint deletes an application and all of its reported data.</p>\n\n<p><strong>WARNING</strong>: Only applications that have stopped reporting can be deleted. This is an irreversible process\nwhich will delete all reported data for this application.</p>\n",
                    "summary": "Delete",
                    "nickname": "delete",
                    "httpMethod": "DELETE",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "id",
                            "description": "Application ID",
                            "dataType": "Integer",
                            "required": true
                        }
                    ],
                    "consumes": [
                        "application/json",
                        "application/xml"
                    ],
                    "produces": [
                        "application/json",
                        "application/xml"
                    ],
                    "nrExtensions": {},
                    "responseClass": "ApplicationResponse",
                    "errorResponses": [
                        {
                            "code": 401,
                            "reason": "Invalid API key"
                        },
                        {
                            "code": 401,
                            "reason": "Invalid request, API key required"
                        },
                        {
                            "code": 403,
                            "reason": "New Relic API access has not been enabled"
                        },
                        {
                            "code": 500,
                            "reason": "A server error occurred, please contact <a href=\"http://support.newrelic.com\">New Relic support</a>"
                        },
                        {
                            "code": 404,
                            "reason": "No Application found with the given ID"
                        },
                        {
                            "code": 409,
                            "reason": "Cannot delete an application that is still reporting data."
                        }
                    ]
                }
            ]
        },
        {
            "path": "/v2/applications/{application_id}/metrics.{format}",
            "operations": [
                {
                    "notes": "<p>Return a list of known metrics and their value names for the given resource.</p>\n\n<p>See our documentation for a discussion\non <a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/requirements/pagination-api-output\" target=\"_blank\"> output pagination</a>\nand for examples of <a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/requirements/new-relic-rest-api-v2-getting-started#examples\" target=\"_blank\">requesting and using metric values</a>.</p>\n",
                    "summary": "Metric Names",
                    "nickname": "metric_names",
                    "httpMethod": "GET",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "application_id",
                            "description": "Application ID",
                            "dataType": "Integer",
                            "required": true
                        },
                        {
                            "paramType": "query",
                            "name": "name",
                            "description": "Filter metrics by name",
                            "dataType": "String",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "page",
                            "description": "Pagination index (will be deprecated)",
                            "dataType": "Integer",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "cursor",
                            "description": "Cursor for next page (replacing page param)",
                            "dataType": "String",
                            "required": false
                        }
                    ],
                    "consumes": [
                        "application/json",
                        "application/xml"
                    ],
                    "produces": [
                        "application/json",
                        "application/xml"
                    ],
                    "nrExtensions": {},
                    "responseClass": "MetricParserResponse",
                    "errorResponses": [
                        {
                            "code": 401,
                            "reason": "Invalid API key"
                        },
                        {
                            "code": 401,
                            "reason": "Invalid request, API key required"
                        },
                        {
                            "code": 403,
                            "reason": "New Relic API access has not been enabled"
                        },
                        {
                            "code": 500,
                            "reason": "A server error occurred, please contact <a href=\"http://support.newrelic.com\">New Relic support</a>"
                        }
                    ]
                }
            ]
        },
        {
            "path": "/v2/applications/{application_id}/metrics/data.{format}",
            "operations": [
                {
                    "notes": "<p>This API endpoint returns a list of values for each of the requested metrics. The list of available metrics\ncan be returned using the Metric Name API endpoint.</p>\n\n<p>Metric data can be filtered by a number of parameters, including multiple names and values, and by time range.\nMetric names and values will be matched intelligently in the background.</p>\n\n<p>You can also retrieve a summarized data point across the entire time range selected by using the summarize\nparameter.</p>\n\n<p>See our documentation for a discussion on <a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/requirements/pagination-api-output\" target=\"_blank\">\noutput pagination</a>, <a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/requirements/extracting-metric-data\" target=\"_blank\"> time range</a> \nrelated considerations, and for examples of <a href=\"https://docs.newrelic.com/docs/apis/rest-api-v2/requirements/new-relic-rest-api-v2-getting-started#examples\" target=\"_blank\">requesting and using metric values</a>.</p>\n\n",
                    "summary": "Metric Data",
                    "nickname": "metric_data",
                    "httpMethod": "GET",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "application_id",
                            "description": "Application ID",
                            "dataType": "Integer",
                            "required": true
                        },
                        {
                            "paramType": "query",
                            "name": "names",
                            "description": "Retrieve specific metrics by name",
                            "dataType": "Array",
                            "required": true
                        },
                        {
                            "paramType": "query",
                            "name": "values",
                            "description": "Retrieve specific metric values",
                            "dataType": "Array",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "from",
                            "description": "Retrieve metrics after this time",
                            "dataType": "Time",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "to",
                            "description": "Retrieve metrics before this time",
                            "dataType": "Time",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "period",
                            "description": "Period of timeslices in seconds",
                            "dataType": "Integer",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "summarize",
                            "description": "Summarize the data",
                            "dataType": "Boolean",
                            "required": false
                        },
                        {
                            "paramType": "query",
                            "name": "raw",
                            "description": "Return unformatted raw values",
                            "dataType": "Boolean",
                            "required": false
                        }
                    ],
                    "consumes": [
                        "application/json",
                        "application/xml"
                    ],
                    "produces": [
                        "application/json",
                        "application/xml"
                    ],
                    "nrExtensions": {},
                    "responseClass": "MetricDataResponse",
                    "errorResponses": [
                        {
                            "code": 401,
                            "reason": "Invalid API key"
                        },
                        {
                            "code": 401,
                            "reason": "Invalid request, API key required"
                        },
                        {
                            "code": 403,
                            "reason": "New Relic API access has not been enabled"
                        },
                        {
                            "code": 500,
                            "reason": "A server error occurred, please contact <a href=\"http://support.newrelic.com\">New Relic support</a>"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "ApplicationResponse": {
            "properties": {
                "application": {
                    "type": "ApplicationResponseType"
                }
            }
        },
        "ApplicationResponseType": {
            "properties": {
                "id": {
                    "type": "integer",
                    "desc": "Application unique identifier"
                },
                "name": {
                    "type": "string",
                    "desc": "Application name",
                    "update": true
                },
                "language": {
                    "type": "string",
                    "desc": "Application language as reported by agents"
                },
                "health_status": {
                    "type": "string",
                    "desc": "Application code status"
                },
                "reporting": {
                    "type": "boolean",
                    "desc": "Whether application has reported data in the last 10 minutes"
                },
                "last_reported_at": {
                    "type": "time",
                    "desc": "The last time the application reported data (accurate to within a few minutes)"
                },
                "application_summary": {
                    "type": "AppSummaryResponse"
                },
                "end_user_summary": {
                    "type": "EndUserSummaryResponse"
                },
                "settings": {
                    "type": "AppSettingsResponse"
                },
                "links": {
                    "type": "ApplicationLinksResponse"
                }
            }
        },
        "Application": {
            "properties": {
                "application": {
                    "type": "ApplicationBody"
                }
            }
        },
        "ApplicationBody": {
            "properties": {
                "name": {
                    "type": "string",
                    "desc": "Application name",
                    "update": true
                },
                "settings": {
                    "type": "AppSettingsBody"
                }
            }
        },
        "AppSummaryResponse": {
            "properties": {
                "response_time": {
                    "type": "float",
                    "desc": "Application average response time"
                },
                "throughput": {
                    "type": "float",
                    "desc": "Application throughput in the last 5 minutes"
                },
                "error_rate": {
                    "type": "float",
                    "desc": "Application error rate in the last 5 minutes"
                },
                "apdex_target": {
                    "type": "float",
                    "desc": "Application Apdex target time"
                },
                "apdex_score": {
                    "type": "float",
                    "desc": "Application Apdex score in the last 5 minutes"
                },
                "host_count": {
                    "type": "integer",
                    "desc": "Number of application hosts"
                },
                "instance_count": {
                    "type": "integer",
                    "desc": "Number of application instances"
                },
                "concurrent_instance_count": {
                    "type": "integer",
                    "desc": "Number of concurrent application instances for background applications"
                }
            }
        },
        "EndUserSummaryResponse": {
            "properties": {
                "response_time": {
                    "type": "float",
                    "desc": "End user average response time (ms)"
                },
                "throughput": {
                    "type": "float",
                    "desc": "End user throughput in the last 5 minutes"
                },
                "apdex_target": {
                    "type": "float",
                    "desc": "End user Apdex target time"
                },
                "apdex_score": {
                    "type": "float",
                    "desc": "End user Apdex score in the last 5 minutes"
                }
            }
        },
        "AppSettingsResponse": {
            "properties": {
                "app_apdex_threshold": {
                    "type": "float",
                    "desc": "Application Apdex threshold",
                    "update": true
                },
                "end_user_apdex_threshold": {
                    "type": "float",
                    "desc": "End user Apdex threshold",
                    "update": true
                },
                "enable_real_user_monitoring": {
                    "type": "boolean",
                    "desc": "Whether end user monitoring is enabled",
                    "update": true
                },
                "use_server_side_config": {
                    "type": "boolean",
                    "desc": "Whether configurations are server-side or at newrelic.com"
                }
            }
        },
        "AppSettingsBody": {
            "properties": {
                "app_apdex_threshold": {
                    "type": "float",
                    "desc": "Application Apdex threshold",
                    "update": true
                },
                "end_user_apdex_threshold": {
                    "type": "float",
                    "desc": "End user Apdex threshold",
                    "update": true
                },
                "enable_real_user_monitoring": {
                    "type": "boolean",
                    "desc": "Whether end user monitoring is enabled",
                    "update": true
                }
            }
        },
        "ApplicationLinksResponse": {
            "properties": {
                "servers": {
                    "type": "array",
                    "items": {
                        "$ref": "integer"
                    },
                    "desc": "List of servers associated with this application"
                },
                "application_hosts": {
                    "type": "array",
                    "items": {
                        "$ref": "integer"
                    },
                    "desc": "List of application hosts"
                },
                "application_instances": {
                    "type": "array",
                    "items": {
                        "$ref": "integer"
                    },
                    "desc": "List of application instances"
                },
                "alert_policy": {
                    "type": "integer",
                    "desc": "Alert policy in which this application is a member"
                }
            }
        },
        "MetricParserResponse": {
            "properties": {
                "metric": {
                    "type": "MetricParserResponseType"
                }
            }
        },
        "MetricParserResponseType": {
            "properties": {
                "name": {
                    "type": "string",
                    "desc": "Name of metric collected"
                },
                "values": {
                    "type": "array",
                    "items": {
                        "$ref": "string"
                    },
                    "desc": "The name of metric values available"
                }
            }
        },
        "MetricDataResponse": {
            "properties": {
                "metric_data": {
                    "type": "MetricDataResponseType"
                }
            }
        },
        "MetricDataResponseType": {
            "properties": {
                "from": {
                    "type": "time",
                    "desc": "Find metric data that was collected after this time."
                },
                "to": {
                    "type": "time",
                    "desc": "Find metric data that was collected before this time."
                },
                "metrics_not_found": {
                    "type": "string",
                    "desc": "Metrics that were requested but not found for this time period."
                },
                "metrics_found": {
                    "type": "string",
                    "desc": "Requested metrics found for this time period."
                },
                "metrics": {
                    "type": "array",
                    "items": {
                        "$ref": "MetricResponse"
                    }
                }
            }
        },
        "MetricResponse": {
            "properties": {
                "name": {
                    "type": "string",
                    "desc": "Name of the metric."
                },
                "timeslices": {
                    "type": "array",
                    "items": {
                        "$ref": "TimesliceResponse"
                    }
                }
            }
        },
        "TimesliceResponse": {
            "properties": {
                "from": {
                    "type": "time",
                    "desc": "Find data that was collected after this time."
                },
                "to": {
                    "type": "time",
                    "desc": "Find data that was collected before this time."
                },
                "values": {
                    "type": "hash",
                    "desc": "List of values to return for each data point."
                }
            }
        },
        "EventResponse": {
            "properties": {
                "event": {
                    "type": "EventResponseType"
                }
            }
        },
        "EventResponseType": {
            "properties": {
                "id": {
                    "type": "integer",
                    "desc": "Event unique identifier"
                },
                "type": {
                    "type": "string",
                    "desc": "Event type: event,problem,note,settings_changed,application_apdex_problem,application_error_rate_problem,application_throughput_problem,application_downtime,browser_apdex_problem,server_cpu_problem,server_disk_io_problem,server_disk_usage_problem,server_memory_problem,metric_problem,server_not_reporting_problem,custom_event,custom_metric_problem,custom_problem,mobile_failed_call_rate_problem,mobile_response_time_problem,mobile_status_code_error_rate_problem",
                    "update": true
                },
                "level": {
                    "type": "string",
                    "desc": "Severity level: {\"unknown\"=>-1, \"ok\"=>0, \"info\"=>1, \"warning\"=>2, \"error\"=>3, \"downtime\"=>4}",
                    "update": true
                },
                "message": {
                    "type": "string",
                    "desc": "Description",
                    "update": true
                },
                "created_at": {
                    "type": "datetime",
                    "desc": "Created at",
                    "update": true
                },
                "closed_at": {
                    "type": "datetime",
                    "desc": "Closed at",
                    "update": true
                },
                "payload": {
                    "type": "hash",
                    "desc": "JSON payload",
                    "update": true
                },
                "links": {
                    "type": "EventLinksResponse"
                }
            }
        },
        "Event": {
            "properties": {
                "event": {
                    "type": "EventBody"
                }
            }
        },
        "EventBody": {
            "properties": {
                "type": {
                    "type": "string",
                    "desc": "Event type: event,problem,note,settings_changed,application_apdex_problem,application_error_rate_problem,application_throughput_problem,application_downtime,browser_apdex_problem,server_cpu_problem,server_disk_io_problem,server_disk_usage_problem,server_memory_problem,metric_problem,server_not_reporting_problem,custom_event,custom_metric_problem,custom_problem,mobile_failed_call_rate_problem,mobile_response_time_problem,mobile_status_code_error_rate_problem",
                    "update": true
                },
                "level": {
                    "type": "string",
                    "desc": "Severity level: {\"unknown\"=>-1, \"ok\"=>0, \"info\"=>1, \"warning\"=>2, \"error\"=>3, \"downtime\"=>4}",
                    "update": true
                },
                "message": {
                    "type": "string",
                    "desc": "Description",
                    "update": true
                },
                "created_at": {
                    "type": "datetime",
                    "desc": "Created at",
                    "update": true
                },
                "closed_at": {
                    "type": "datetime",
                    "desc": "Closed at",
                    "update": true
                },
                "payload": {
                    "type": "hash",
                    "desc": "JSON payload",
                    "update": true
                }
            }
        },
        "EventLinksResponse": {
            "properties": {
                "incidents": {
                    "type": "array",
                    "items": {
                        "$ref": "integer"
                    },
                    "desc": "Event incidents"
                },
                "downtime": {
                    "type": "integer",
                    "desc": "Downtime event"
                }
            }
        },
        "TargetDowntimeEventResponse": {
            "properties": {
                "downtime": {
                    "type": "TargetDowntimeEventResponseType"
                }
            }
        },
        "TargetDowntimeEventResponseType": {
            "properties": {
                "id": {
                    "type": "integer",
                    "desc": "Downtime unique identifier"
                },
                "opened_at": {
                    "type": "datetime",
                    "desc": "Created at"
                },
                "last_failure": {
                    "type": "datetime",
                    "desc": "Last failure"
                },
                "duration": {
                    "type": "integer",
                    "desc": "Duration in seconds"
                },
                "number_of_failures": {
                    "type": "integer",
                    "desc": "Number of ping failures"
                },
                "url": {
                    "type": "string",
                    "desc": "Ping url for this downtime event"
                },
                "ip_address": {
                    "type": "string",
                    "desc": "Ping target IP address"
                }
            }
        }
    },
    "nrExtensions": {}
}
```

## Converting New Relic's Docs to the Seagger 2.0 spec

