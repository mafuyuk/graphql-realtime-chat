const { createApolloFetch } = require('apollo-fetch');

class GraphQL {
  constructor(uri) {
    this.apolloFetch = createApolloFetch({ uri });
  }

  postMessage(user, text) {
    const query = `
      mutation {
        postMessage(user:"${user}",text:"${text}") {
          user,
          id,
          text,
        }
      }`;

    return this.apolloFetch({
      query,
      // variables: { TODO
      //   user,
      //   text
      // }
    });
  }

  messages() {
    const query =`
      query {
        messages {
          user,
          id,
          text,
        }
      }`;
    return this.apolloFetch({ query });
  }

  users() {
    const query =`
      query {
        users 
      }`;
    return this.apolloFetch({ query });
  }
}

module.exports = GraphQL;