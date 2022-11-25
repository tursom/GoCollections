/*
 * Copyright (c) 2022 tursom. All rights reserved.
 * Use of this source code is governed by a GPL-3
 * license that can be found in the LICENSE file.
 */

package exceptions

var ElementFound = NewRuntimeException("", DefaultExceptionConfig().SetGetStackTrace(false).SetExceptionName("github.com.tursom.GoCollections.exceptions.ElementFound"))
var ElementNotFound = NewElementNotFoundException("", nil)
var CollectionLoopFinished = NewRuntimeException("", DefaultExceptionConfig().SetGetStackTrace(false).SetExceptionName("github.com.tursom.GoCollections.exceptions.CollectionLoopFinished"))
