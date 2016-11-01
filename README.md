# todo2016
A simple TODO software developed on 2016

* [server](server) contain a server which will maintain the TODO database and expose over a REST endpoint, protected by JWT
* [client](client) contains the webclient which talks with the above server. Implemented using [react](https://facebook.github.io/react/), [redux](http://redux.js.org/) and uses JSON Web Tokens, [JWT](https://jwt.io/) for authentication
