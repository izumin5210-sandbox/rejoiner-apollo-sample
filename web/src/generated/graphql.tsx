import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/react-common';
import * as ApolloReactHooks from '@apollo/react-hooks';
export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string,
  String: string,
  Boolean: boolean,
  Int: number,
  Float: number,
  Long: any,
};

export type Github_GetUserRequest = {
   __typename?: 'github_GetUserRequest',
  login?: Maybe<Scalars['String']>,
};

export type Github_ListUsersRequest = {
   __typename?: 'github_ListUsersRequest',
  logins?: Maybe<Array<Maybe<Scalars['String']>>>,
};

export type Github_ListUsersResponse = {
   __typename?: 'github_ListUsersResponse',
  users?: Maybe<Array<Maybe<Github_User>>>,
};

export type Github_Plan = {
   __typename?: 'github_Plan',
  collaborators?: Maybe<Scalars['Int']>,
  filledSeats?: Maybe<Scalars['Int']>,
  name?: Maybe<Scalars['String']>,
  privateRepos?: Maybe<Scalars['Int']>,
  seats?: Maybe<Scalars['Int']>,
  space?: Maybe<Scalars['Int']>,
};

export type Github_User = {
   __typename?: 'github_User',
  avatarUrl?: Maybe<Scalars['String']>,
  bio?: Maybe<Scalars['String']>,
  blog?: Maybe<Scalars['String']>,
  collaborators?: Maybe<Scalars['Int']>,
  company?: Maybe<Scalars['String']>,
  createdAt?: Maybe<Google_Protobuf_Timestamp>,
  diskUsage?: Maybe<Scalars['Int']>,
  email?: Maybe<Scalars['String']>,
  followers?: Maybe<Scalars['Int']>,
  following?: Maybe<Scalars['Int']>,
  gravatarId?: Maybe<Scalars['String']>,
  hireable?: Maybe<Scalars['Boolean']>,
  htmlUrl?: Maybe<Scalars['String']>,
  id?: Maybe<Scalars['Long']>,
  ldapDn?: Maybe<Scalars['String']>,
  location?: Maybe<Scalars['String']>,
  login?: Maybe<Scalars['String']>,
  name?: Maybe<Scalars['String']>,
  nodeId?: Maybe<Scalars['String']>,
  ownedPrivateRepos?: Maybe<Scalars['Int']>,
  plan?: Maybe<Github_Plan>,
  privateGists?: Maybe<Scalars['Int']>,
  publicGists?: Maybe<Scalars['Int']>,
  publicRepos?: Maybe<Scalars['Int']>,
  siteAdmin?: Maybe<Scalars['Boolean']>,
  suspendedAt?: Maybe<Google_Protobuf_Timestamp>,
  totalPrivateRepos?: Maybe<Scalars['Int']>,
  twoFactorAuthentication?: Maybe<Scalars['Boolean']>,
  type?: Maybe<Scalars['String']>,
  updatedAt?: Maybe<Google_Protobuf_Timestamp>,
};

export type Google_Protobuf_Timestamp = {
   __typename?: 'google_protobuf_Timestamp',
  nanos?: Maybe<Scalars['Int']>,
  seconds?: Maybe<Scalars['Long']>,
};

export type Input_Github_GetUserRequest = {
  login?: Maybe<Scalars['String']>,
};

export type Input_Github_ListUsersRequest = {
  logins?: Maybe<Array<Maybe<Scalars['String']>>>,
};

export type Input_Github_ListUsersResponse = {
  users?: Maybe<Array<Maybe<Input_Github_User>>>,
};

export type Input_Github_Plan = {
  collaborators?: Maybe<Scalars['Int']>,
  filledSeats?: Maybe<Scalars['Int']>,
  name?: Maybe<Scalars['String']>,
  privateRepos?: Maybe<Scalars['Int']>,
  seats?: Maybe<Scalars['Int']>,
  space?: Maybe<Scalars['Int']>,
};

export type Input_Github_User = {
  avatarUrl?: Maybe<Scalars['String']>,
  bio?: Maybe<Scalars['String']>,
  blog?: Maybe<Scalars['String']>,
  collaborators?: Maybe<Scalars['Int']>,
  company?: Maybe<Scalars['String']>,
  createdAt?: Maybe<Input_Google_Protobuf_Timestamp>,
  diskUsage?: Maybe<Scalars['Int']>,
  email?: Maybe<Scalars['String']>,
  followers?: Maybe<Scalars['Int']>,
  following?: Maybe<Scalars['Int']>,
  gravatarId?: Maybe<Scalars['String']>,
  hireable?: Maybe<Scalars['Boolean']>,
  htmlUrl?: Maybe<Scalars['String']>,
  id?: Maybe<Scalars['Long']>,
  ldapDn?: Maybe<Scalars['String']>,
  location?: Maybe<Scalars['String']>,
  login?: Maybe<Scalars['String']>,
  name?: Maybe<Scalars['String']>,
  nodeId?: Maybe<Scalars['String']>,
  ownedPrivateRepos?: Maybe<Scalars['Int']>,
  plan?: Maybe<Input_Github_Plan>,
  privateGists?: Maybe<Scalars['Int']>,
  publicGists?: Maybe<Scalars['Int']>,
  publicRepos?: Maybe<Scalars['Int']>,
  siteAdmin?: Maybe<Scalars['Boolean']>,
  suspendedAt?: Maybe<Input_Google_Protobuf_Timestamp>,
  totalPrivateRepos?: Maybe<Scalars['Int']>,
  twoFactorAuthentication?: Maybe<Scalars['Boolean']>,
  type?: Maybe<Scalars['String']>,
  updatedAt?: Maybe<Input_Google_Protobuf_Timestamp>,
};

export type Input_Google_Protobuf_Timestamp = {
  nanos?: Maybe<Scalars['Int']>,
  seconds?: Maybe<Scalars['Long']>,
};

export type Input_Qiita_GetUserRequest = {
  userId?: Maybe<Scalars['String']>,
};

export type Input_Qiita_Item = {
  body?: Maybe<Scalars['String']>,
  coediting?: Maybe<Scalars['Boolean']>,
  commentsCount?: Maybe<Scalars['Int']>,
  createdAt?: Maybe<Input_Google_Protobuf_Timestamp>,
  group?: Maybe<Scalars['String']>,
  id?: Maybe<Scalars['String']>,
  likesCount?: Maybe<Scalars['Int']>,
  pageViewsCount?: Maybe<Scalars['Int']>,
  private?: Maybe<Scalars['Boolean']>,
  reactionsCount?: Maybe<Scalars['Int']>,
  renderedBody?: Maybe<Scalars['String']>,
  tags?: Maybe<Array<Maybe<Input_Qiita_Item_Tag>>>,
  title?: Maybe<Scalars['String']>,
  updatedAt?: Maybe<Input_Google_Protobuf_Timestamp>,
  url?: Maybe<Scalars['String']>,
  user?: Maybe<Input_Qiita_User>,
};

export type Input_Qiita_Item_Tag = {
  name?: Maybe<Scalars['String']>,
  versions?: Maybe<Array<Maybe<Scalars['String']>>>,
};

export type Input_Qiita_ListFolloweesRequest = {
  userId?: Maybe<Scalars['String']>,
};

export type Input_Qiita_ListFolloweesResponse = {
  users?: Maybe<Array<Maybe<Input_Qiita_User>>>,
};

export type Input_Qiita_ListFollowersRequest = {
  userId?: Maybe<Scalars['String']>,
};

export type Input_Qiita_ListFollowersResponse = {
  users?: Maybe<Array<Maybe<Input_Qiita_User>>>,
};

export type Input_Qiita_ListItemsRequest = {
  userId?: Maybe<Scalars['String']>,
};

export type Input_Qiita_ListItemsResponse = {
  items?: Maybe<Array<Maybe<Input_Qiita_Item>>>,
};

export type Input_Qiita_ListStockersRequest = {
  itemId?: Maybe<Scalars['String']>,
};

export type Input_Qiita_ListStockersResponse = {
  users?: Maybe<Array<Maybe<Input_Qiita_User>>>,
};

export type Input_Qiita_ListStocksRequest = {
  userId?: Maybe<Scalars['String']>,
};

export type Input_Qiita_ListStocksResponse = {
  items?: Maybe<Array<Maybe<Input_Qiita_Item>>>,
};

export type Input_Qiita_User = {
  description?: Maybe<Scalars['String']>,
  facebookId?: Maybe<Scalars['String']>,
  followeesCount?: Maybe<Scalars['Int']>,
  followersCount?: Maybe<Scalars['Int']>,
  githubLoginName?: Maybe<Scalars['String']>,
  id?: Maybe<Scalars['String']>,
  itemsCount?: Maybe<Scalars['Int']>,
  linkedinId?: Maybe<Scalars['String']>,
  location?: Maybe<Scalars['String']>,
  name?: Maybe<Scalars['String']>,
  organization?: Maybe<Scalars['String']>,
  permanentId?: Maybe<Scalars['Int']>,
  profileImageUrl?: Maybe<Scalars['String']>,
  teamOnly?: Maybe<Scalars['Boolean']>,
  twitterScreenName?: Maybe<Scalars['String']>,
  websiteUrl?: Maybe<Scalars['String']>,
};


export type Qiita_GetUserRequest = {
   __typename?: 'qiita_GetUserRequest',
  userId?: Maybe<Scalars['String']>,
};

export type Qiita_Item = {
   __typename?: 'qiita_Item',
  body?: Maybe<Scalars['String']>,
  coediting?: Maybe<Scalars['Boolean']>,
  commentsCount?: Maybe<Scalars['Int']>,
  createdAt?: Maybe<Google_Protobuf_Timestamp>,
  followers?: Maybe<Array<Qiita_User>>,
  group?: Maybe<Scalars['String']>,
  id?: Maybe<Scalars['String']>,
  likesCount?: Maybe<Scalars['Int']>,
  pageViewsCount?: Maybe<Scalars['Int']>,
  private?: Maybe<Scalars['Boolean']>,
  reactionsCount?: Maybe<Scalars['Int']>,
  renderedBody?: Maybe<Scalars['String']>,
  tags?: Maybe<Array<Maybe<Qiita_Item_Tag>>>,
  title?: Maybe<Scalars['String']>,
  updatedAt?: Maybe<Google_Protobuf_Timestamp>,
  url?: Maybe<Scalars['String']>,
  user?: Maybe<Qiita_User>,
};

export type Qiita_Item_Tag = {
   __typename?: 'qiita_Item_Tag',
  name?: Maybe<Scalars['String']>,
  versions?: Maybe<Array<Maybe<Scalars['String']>>>,
};

export type Qiita_ListFolloweesRequest = {
   __typename?: 'qiita_ListFolloweesRequest',
  userId?: Maybe<Scalars['String']>,
};

export type Qiita_ListFolloweesResponse = {
   __typename?: 'qiita_ListFolloweesResponse',
  users?: Maybe<Array<Maybe<Qiita_User>>>,
};

export type Qiita_ListFollowersRequest = {
   __typename?: 'qiita_ListFollowersRequest',
  userId?: Maybe<Scalars['String']>,
};

export type Qiita_ListFollowersResponse = {
   __typename?: 'qiita_ListFollowersResponse',
  users?: Maybe<Array<Maybe<Qiita_User>>>,
};

export type Qiita_ListItemsRequest = {
   __typename?: 'qiita_ListItemsRequest',
  userId?: Maybe<Scalars['String']>,
};

export type Qiita_ListItemsResponse = {
   __typename?: 'qiita_ListItemsResponse',
  items?: Maybe<Array<Maybe<Qiita_Item>>>,
};

export type Qiita_ListStockersRequest = {
   __typename?: 'qiita_ListStockersRequest',
  itemId?: Maybe<Scalars['String']>,
};

export type Qiita_ListStockersResponse = {
   __typename?: 'qiita_ListStockersResponse',
  users?: Maybe<Array<Maybe<Qiita_User>>>,
};

export type Qiita_ListStocksRequest = {
   __typename?: 'qiita_ListStocksRequest',
  userId?: Maybe<Scalars['String']>,
};

export type Qiita_ListStocksResponse = {
   __typename?: 'qiita_ListStocksResponse',
  items?: Maybe<Array<Maybe<Qiita_Item>>>,
};

export type Qiita_User = {
   __typename?: 'qiita_User',
  description?: Maybe<Scalars['String']>,
  facebookId?: Maybe<Scalars['String']>,
  followees?: Maybe<Array<Qiita_User>>,
  followeesCount?: Maybe<Scalars['Int']>,
  followers?: Maybe<Array<Qiita_User>>,
  followersCount?: Maybe<Scalars['Int']>,
  githubAccount?: Maybe<Github_User>,
  githubLoginName?: Maybe<Scalars['String']>,
  id?: Maybe<Scalars['String']>,
  items?: Maybe<Array<Qiita_Item>>,
  itemsCount?: Maybe<Scalars['Int']>,
  linkedinId?: Maybe<Scalars['String']>,
  location?: Maybe<Scalars['String']>,
  name?: Maybe<Scalars['String']>,
  organization?: Maybe<Scalars['String']>,
  permanentId?: Maybe<Scalars['Int']>,
  profileImageUrl?: Maybe<Scalars['String']>,
  teamOnly?: Maybe<Scalars['Boolean']>,
  twitterScreenName?: Maybe<Scalars['String']>,
  websiteUrl?: Maybe<Scalars['String']>,
};

export type QueryType = {
   __typename?: 'QueryType',
  getGithubUser?: Maybe<Github_User>,
  getQiitaUser?: Maybe<Qiita_User>,
  listGithubUsers?: Maybe<Array<Github_User>>,
  listQiitaItems?: Maybe<Array<Qiita_Item>>,
};


export type QueryTypeGetGithubUserArgs = {
  input?: Maybe<Input_Github_GetUserRequest>
};


export type QueryTypeGetQiitaUserArgs = {
  input?: Maybe<Input_Qiita_GetUserRequest>
};


export type QueryTypeListGithubUsersArgs = {
  input?: Maybe<Input_Github_ListUsersRequest>
};


export type QueryTypeListQiitaItemsArgs = {
  input?: Maybe<Input_Qiita_ListItemsRequest>
};

export type UserFragment = (
  { __typename?: 'qiita_User' }
  & Pick<Qiita_User, 'id' | 'name' | 'profileImageUrl'>
  & { githubAccount: Maybe<(
    { __typename?: 'github_User' }
    & Pick<Github_User, 'login' | 'avatarUrl'>
  )> }
);

export type GetQiitaUserQueryVariables = {
  userId?: Maybe<Scalars['String']>
};


export type GetQiitaUserQuery = (
  { __typename?: 'QueryType' }
  & { getQiitaUser: Maybe<(
    { __typename?: 'qiita_User' }
    & { followees: Maybe<Array<(
      { __typename?: 'qiita_User' }
      & UserFragment
    )>>, followers: Maybe<Array<(
      { __typename?: 'qiita_User' }
      & UserFragment
    )>> }
    & UserFragment
  )> }
);

export const UserFragmentDoc = gql`
    fragment User on qiita_User {
  id
  name
  profileImageUrl
  githubAccount {
    login
    avatarUrl
  }
}
    `;
export const GetQiitaUserDocument = gql`
    query GetQiitaUser($userId: String) {
  getQiitaUser(input: {userId: $userId}) {
    ...User
    followees {
      ...User
    }
    followers {
      ...User
    }
  }
}
    ${UserFragmentDoc}`;

/**
 * __useGetQiitaUserQuery__
 *
 * To run a query within a React component, call `useGetQiitaUserQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetQiitaUserQuery` returns an object from Apollo Client that contains loading, error, and data properties 
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetQiitaUserQuery({
 *   variables: {
 *      userId: // value for 'userId'
 *   },
 * });
 */
export function useGetQiitaUserQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<GetQiitaUserQuery, GetQiitaUserQueryVariables>) {
        return ApolloReactHooks.useQuery<GetQiitaUserQuery, GetQiitaUserQueryVariables>(GetQiitaUserDocument, baseOptions);
      }
export function useGetQiitaUserLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<GetQiitaUserQuery, GetQiitaUserQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<GetQiitaUserQuery, GetQiitaUserQueryVariables>(GetQiitaUserDocument, baseOptions);
        }
export type GetQiitaUserQueryHookResult = ReturnType<typeof useGetQiitaUserQuery>;
export type GetQiitaUserLazyQueryHookResult = ReturnType<typeof useGetQiitaUserLazyQuery>;
export type GetQiitaUserQueryResult = ApolloReactCommon.QueryResult<GetQiitaUserQuery, GetQiitaUserQueryVariables>;