/* tslint:disable */
/* eslint-disable */
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: GetQiitaUser
// ====================================================

export interface GetQiitaUser_getQiitaUser_githubAccount {
  __typename: "github_User";
  login: string | null;
  avatarUrl: string | null;
}

export interface GetQiitaUser_getQiitaUser_followees_githubAccount {
  __typename: "github_User";
  login: string | null;
  avatarUrl: string | null;
}

export interface GetQiitaUser_getQiitaUser_followees {
  __typename: "qiita_User";
  id: string | null;
  name: string | null;
  profileImageUrl: string | null;
  githubAccount: GetQiitaUser_getQiitaUser_followees_githubAccount | null;
}

export interface GetQiitaUser_getQiitaUser_followers_githubAccount {
  __typename: "github_User";
  login: string | null;
  avatarUrl: string | null;
}

export interface GetQiitaUser_getQiitaUser_followers {
  __typename: "qiita_User";
  id: string | null;
  name: string | null;
  profileImageUrl: string | null;
  githubAccount: GetQiitaUser_getQiitaUser_followers_githubAccount | null;
}

export interface GetQiitaUser_getQiitaUser {
  __typename: "qiita_User";
  id: string | null;
  name: string | null;
  profileImageUrl: string | null;
  githubAccount: GetQiitaUser_getQiitaUser_githubAccount | null;
  followees: GetQiitaUser_getQiitaUser_followees[] | null;
  followers: GetQiitaUser_getQiitaUser_followers[] | null;
}

export interface GetQiitaUser {
  getQiitaUser: GetQiitaUser_getQiitaUser | null;
}

export interface GetQiitaUserVariables {
  userId?: string | null;
}
