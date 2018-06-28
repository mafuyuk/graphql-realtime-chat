require('isomorphic-fetch');

class GraphQL {
  constructor(url) {
    this.url = url;
    this.config = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
    };
  }

  postMessage(user, text) {
    return fetch(this.url, Object.assign({
      body: JSON.stringify({
        query: `mutation postMessage { postMessage(user:"${user}", text:"${text}") { user id text }}`
      })
    }, this.config))
      .then(res => res.json());
  }

  messages() {
    return fetch(this.url, Object.assign({
      body: JSON.stringify({
        query: "query messages { messages { user id text } }"
      })
    }, this.config))
      .then(res => res.json());
  }
}

module.exports = GraphQL;