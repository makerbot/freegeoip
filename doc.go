// Copyright 2009-2014 The freegeoip authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package freegeoip provides an API for searching the geolocation of IP
// addresses. It uses a database that can be either a local file or a
// remote resource from a URL.
//
// Local databases are monitored by fsnotify and reloaded when the file is
// either updated or overwritten.
//
// Remote databases are automatically downloaded and updated in background
// so you can focus on using the API and not managing the database.
//
// Also, the freegeoip package provides http handlers that any Go http
// server (net/http) can use. These handlers can process IP geolocation
// lookup requests and return data in multiple formats like CSV, XML,
// JSON and JSONP. It has also an API for supporting custom formats.
package freegeoip
