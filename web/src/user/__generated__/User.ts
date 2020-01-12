/* tslint:disable */
/* eslint-disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL fragment: User
// ====================================================

export interface User_githubAccount {
  __typename: "github_User";
  login: string | null;
  avatarUrl: string | null;
}

export interface User {
  __typename: "qiita_User";
  id: string | null;
  name: string | null;
  profileImageUrl: string | null;
  githubAccount: User_githubAccount | null;
}
