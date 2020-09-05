# That！s Me 
That is me is an anonymous gaming social platform that pays tribute to Turing.

## Key Features
* **COMPLETELY** anonymous
* Coop in limited time
* Random Theme
* Limited time chatting as winners' reward
* Reinforcement learning AI

## Tech Stacks
* vue.js for static webapp
* protobuf for frontend/backend protocol
* redis for AI data storage
* gin for web framework
* websocket for frontend/backend communication
* docker and docker-compose for deployment and hot-reload development environment
* github action for CI/CD

## Usage
> WARNING: the developers have *never ever* run this project on **`Windows`**
### Requirements
* `docker` 
* `docker-compose`
* `bash`
### Configurations
* `tools/release/docker-compose.yml` for port mapping/redis data persistence/etc... 
* `tools/release/config/config.json` for app configuration
### Build and Run Release
```shell script
tools/release/service.sh up --build
```
Also, a docker image named `hariolate/that-is-me` is available on dockerhub, which means that you can:
```shell script
docker pull hariolate/that-is-me:latest
```

## Credits 
* Members of Team `十里山路不换肩`(aka `hariolate`) 
    * Mentor - 谢函瀚 ([@xhhzuikeaiya](https://github.com/xhhzuikeaiya)) <br>
    * Mentor - 黄嘉杰 ([@huangjj979](https://github.com/huangjj979)) <br>
    * Lead Developer - 黄扬 ([@huang825172](https://github.com/huang825172)) <br>
    * USELESS BRAINLESS Coder - 方泓睿 ([@chfanghr](https://github.com/chfanghr)) <br>
* For third party credits, See [CREDITS.md](CREDITS.md)

## License
MIT License

Copyright (c) 2020 hariolate

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

 