// Copyright 2017 Santhosh Kumar Tekuri. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package jsonschema provides json-schema compilation and validation.

This implementation of JSON Schema, supports draft4 and draft6.
Passes all tests(including optional) in https://github.com/json-schema/JSON-Schema-Test-Suite

An example of using this package:

	schema, err := jsonschema.Compile("schemas/purchaseOrder.json")
	if err != nil {
		return err
	}
	f, err := os.Open("purchaseOrder.json")
	if err != nil {
		return err
	}
	defer f.Close()
	if err = schema.Validate(f); err != nil {
		return err
	}

The schema is compiled against the version specified in `$schema` property.
If `$schema` property is missing, it uses latest draft which currently is draft6.
You can force to use draft4 when `$schema` is missing, as follows:

	compiler := jsonschema.NewCompiler()
	compler.Draft = jsonschema.Draft4

This package supports loading json-schema from filePath and fileURL.

To load json-schema from HTTPURL, add following import:

	import _ "github.com/santhosh-tekuri/jsonschema/httploader"

Loading from urls for other schemes (such as ftp), can be plugged in. see package jsonschema/httploader
for an example

To load json-schema from in-memory:

	data := `{"type": "string"}`
	url := "sch.json"
	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource(url, strings.NewReader(data)); err != nil {
		return err
	}
	schema, err := compiler.Compile(url)
	if err != nil {
		return err
	}
	f, err := os.Open("doc.json")
	if err != nil {
		return err
	}
	defer f.Close()
	if err = schema.Validate(f); err != nil {
		return err
	}

This package supports json string formats: date-time, hostname, email, ip-address, ipv4, ipv6, uri, uriref, regex,
format, json-pointer, uri-template (limited validation).

Developers can define their own formats using package jsonschema/formats.

The ValidationError returned by Validate method contains detailed context to understand why and where the error is.

*/
package jsonschema
