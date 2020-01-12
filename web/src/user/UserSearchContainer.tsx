import React from "react";
import { gql } from "apollo-boost";
import { useQuery } from "@apollo/react-hooks";
import {
  GetQiitaUser,
  GetQiitaUserVariables,
  GetQiitaUser_getQiitaUser
} from "./__generated__/GetQiitaUser";
import { User } from "./__generated__/User";

const QUERY = gql`
  fragment User on qiita_User {
    id
    name
    profileImageUrl
    githubAccount {
      login
      avatarUrl
    }
  }

  query GetQiitaUser($userId: String) {
    getQiitaUser(input: { userId: $userId }) {
      ...User
      followees {
        ...User
      }
      followers {
        ...User
      }
    }
  }
`;

const QiitaUserList = ({ users }: { users: User[] }) => {
  const items = users.map(u => (
    <li>
      <ul>
        <li>
          Qiita:{" "}
          {u.profileImageUrl && <img width="24px" src={u.profileImageUrl} />}{" "}
          {u.id} {u.name}
        </li>
        {u.githubAccount && (
          <li>
            GitHub:{" "}
            {u.githubAccount.avatarUrl && (
              <img width="24px" src={u.githubAccount.avatarUrl} />
            )}{" "}
            {u.githubAccount.login}
          </li>
        )}
      </ul>
    </li>
  ));

  return <ul>{items}</ul>;
};

const QiitaUser = ({ user }: { user: GetQiitaUser_getQiitaUser }) => {
  return (
    <dl>
      <dt>ID</dt>
      <dd>{user.id}</dd>
      <dt>Name</dt>
      <dd>{user.name}</dd>
      <dt>Followees</dt>
      <dd>
        <QiitaUserList users={user.followees || []} />
      </dd>
      <dt>Followers</dt>
      <dd>
        <QiitaUserList users={user.followers || []} />
      </dd>
    </dl>
  );
};

const UserSearchContainer: React.FC = () => {
  const { data } = useQuery<GetQiitaUser, GetQiitaUserVariables>(QUERY, {
    variables: { userId: "izumin5210" }
  });

  if (data?.getQiitaUser == null) {
    return null;
  }

  console.log(data.getQiitaUser);

  return (
    <div>
      <input type="text" />
      <QiitaUser user={data.getQiitaUser} />
    </div>
  );
};
export default UserSearchContainer;
