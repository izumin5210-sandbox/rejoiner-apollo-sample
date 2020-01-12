import React from "react";
import "./App.css";

import ApolloClient from "apollo-boost";
import { ApolloProvider } from "@apollo/react-hooks";
import UserSearchContainer from "./user/UserSearchContainer";

const apolloClient = new ApolloClient({
  uri: "http://localhost:8080/graphql"
});

const App: React.FC = () => (
  <ApolloProvider client={apolloClient}>
    <h1>App</h1>
    <UserSearchContainer />
  </ApolloProvider>
);

export default App;
