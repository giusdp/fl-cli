<!--
  ~ Licensed to the Apache Software Foundation (ASF) under one
  ~ or more contributor license agreements.  See the NOTICE file
  ~ distributed with this work for additional information
  ~ regarding copyright ownership.  The ASF licenses this file
  ~ to you under the Apache License, Version 2.0 (the
  ~ "License"); you may not use this file except in compliance
  ~ with the License.  You may obtain a copy of the License at
  ~
  ~   http://www.apache.org/licenses/LICENSE-2.0
  ~
  ~ Unless required by applicable law or agreed to in writing,
  ~ software distributed under the License is distributed on an
  ~ "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  ~ KIND, either express or implied.  See the License for the
  ~ specific language governing permissions and limitations
  ~ under the License.
  ~
-->

## 0.1.0 (2022-04-22)

### Refactor

- **client**: refactor client.go and reduce clientAPI to just get request

### Feat

- create client and fn service in main and bind it to kong
- **fn**: use FnService.Invoke when cmd fn is used
- **FnService**: add FnService with simple Invoke
- **client**: add send method and remove interface
- **client**: add an initial client package
- setup kong library with sample cli main
- create go project
