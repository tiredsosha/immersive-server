# immersive-server

A simple background http server that controls Mad Mapper (or any osc) via webpage.

## Requirements

If you use binary file:

- `Windows 10 or greater`

If you use source code:

- `Windows 10 or greater`
- `Go 1.18 or greater`

### config.yaml

App will look for this file in the current working directory (directory from where you launched Warden). If there is no conf file app will create default file. Change this file or create **config.yaml** file and put desired parameters into it. Or just copy an example of this file from config folder in the repo.

Example file:

    httpPort: 80
    oscIp: 127.0.0.1
    oscPort: 8011

By the way, App will validate you config before starting and notice you whether you forget something!

## Logging

App starts logging immediately after launch. It makes **server.log** file in the current working directory and overwrite file every 7 days

## Building

You can build it by yourself.

    go build -o bin/warden.exe -ldflags "-H windowsgui"

## License

Copyright 2023 Alexandra Chichko &lt;tiredsosha@gmail.com&gt;

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this software except in compliance with the License.
You may obtain a copy of the License at

> http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
