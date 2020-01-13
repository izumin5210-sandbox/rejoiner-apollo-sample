import React from "react";

import {
  useGetQiitaUserQuery,
  UserFragment,
  GetQiitaUserQuery
} from "../generated/graphql";

const QiitaUserList = ({ users }: { users: UserFragment[] }) => {
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

type User = NonNullable<GetQiitaUserQuery["getQiitaUser"]>;
const QiitaUser = ({ user }: { user: User }) => {
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
  const { data } = useGetQiitaUserQuery({
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
