const GraphQL = require('./requestGraphQL');

const g = new GraphQL("http://localhost:8080/graphql");
g.postMessage("mafuyuk1", "aaaaa")
  .then(res => console.log(res.data)).catch(err => console.log(err));

g.messages()
  .then(res => console.log(res.data)).catch(err => console.log(err));

g.users()
  .then(res => console.log(res.data)).catch(err => console.log(err));